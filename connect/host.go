////////////////////////////////////////////////////////////////////////////////
// Copyright © 2018 Privategrity Corporation                                   /
//                                                                             /
// All rights reserved.                                                        /
////////////////////////////////////////////////////////////////////////////////

// Contains functionality for describing and creating connections

package connect

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/pkg/errors"
	jww "github.com/spf13/jwalterweatherman"
	"gitlab.com/elixxir/crypto/signature/rsa"
	tlsCreds "gitlab.com/elixxir/crypto/tls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials"
	"time"
)

// Information used to describe a connection to a host
type Host struct {
	// address:Port being connected to
	address string

	// PEM-format TLS Certificate
	certificate []byte

	// Configure the maximum number of connection attempts
	maxRetries int

	// GRPC connection object
	connection *grpc.ClientConn

	// Credentials object used to establish the connection
	credentials credentials.TransportCredentials

	// RSA Public Key corresponding to the TLS Certificate
	rsaPublicKey *rsa.PublicKey
}

// Ensures the given Host's connection is alive
// and attempts to recover if not
func (h *Host) validateConnection() (err error) {
	// If Host connection does not exist, open the connection
	if h.connection == nil {
		err = h.connect()
	}
	if err != nil {
		return
	}

	// If Host connection is not active, attempt to reestablish
	if !h.isAlive() {
		jww.WARN.Printf("Bad host connection state, reconnecting: %v", h)
		h.disconnect()
		err = h.connect()
	}

	return
}

// Sets up or recovers the Host's connection
// Then runs the given Send function
func (h *Host) Send(f func(conn *grpc.ClientConn) (*any.Any, error)) (
	result *any.Any, err error) {

	// Ensure the connection is running
	err = h.validateConnection()
	if err != nil {
		return
	}

	// Run the send function
	return f(h.connection)
}

// Sets up or recovers the Host's connection
// Then runs the given Stream function
func (h *Host) Stream(f func(conn *grpc.ClientConn) (interface{}, error)) (
	client interface{}, err error) {

	// Ensure the connection is running
	err = h.validateConnection()
	if err != nil {
		return
	}

	// Run the stream function
	return f(h.connection)
}

// Returns the Host address
func (h *Host) GetAddress() string {
	return h.address
}

// Returns a copy of the Host certificate
func (h *Host) GetCertificate() []byte {
	cert := make([]byte, len(h.certificate))
	copy(cert, h.certificate)
	return cert
}

// Returns true if the connection is non-nil and alive
func (h *Host) isAlive() bool {
	if h.connection == nil {
		return false
	}
	state := h.connection.GetState()
	return state == connectivity.Idle || state == connectivity.Connecting ||
		state == connectivity.Ready
}

// Closes a the Host connection
func (h *Host) disconnect() {
	if h.connection != nil {
		return
	}
	err := h.connection.Close()
	if err != nil {
		jww.ERROR.Printf("Unable to close connection to %s: %+v",
			h.address, errors.New(err.Error()))
	}
	h.connection = nil
}

// Connect creates a connection
func (h *Host) connect() (err error) {

	// Configure TLS options
	var securityDial grpc.DialOption
	if h.credentials != nil {
		// Create the gRPC client with TLS
		securityDial = grpc.WithTransportCredentials(h.credentials)
	} else {
		// Create the gRPC client without TLS
		jww.WARN.Printf("Connecting to %v without TLS!", h.address)
		securityDial = grpc.WithInsecure()
	}

	// Attempt to establish a new connection
	for numRetries := 0; numRetries < h.maxRetries && !h.isAlive(); numRetries++ {

		jww.INFO.Printf("Connecting to address %+v. Attempt number %+v of %+v",
			h.address, numRetries, h.maxRetries)

		// If timeout is enabled, the max wait time becomes
		// ~14 seconds (with maxRetries=100)
		backoffTime := 2 * (numRetries/16 + 1)
		if backoffTime > 15 {
			backoffTime = 15
		}
		ctx, cancel := ConnectionContext(time.Duration(backoffTime))

		// Create the connection
		h.connection, err = grpc.DialContext(ctx, h.address, securityDial,
			grpc.WithBlock(), grpc.WithBackoffMaxDelay(time.Minute*5))
		if err != nil {
			jww.ERROR.Printf("Attempt number %+v to connect to %s failed: %+v\n",
				numRetries, h.address, errors.New(err.Error()))
		}
		cancel()
	}

	// Verify that the connection was established successfully
	if !h.isAlive() {
		return errors.New(fmt.Sprintf(
			"Last try to connect to %s failed. Giving up", h.address))
	}

	// Add the successful connection to the Manager
	jww.INFO.Printf("Successfully connected to %v", h.address)
	return
}

// Sets TransportCredentials and RSA PublicKey objects
// using a PEM-encoded TLS Certificate
func (h *Host) setCredentials() error {

	// If no TLS Certificate specified, print a warning and do nothing
	if h.certificate == nil || len(h.certificate) == 0 {
		jww.WARN.Printf("No TLS Certificate specified!")
		return nil
	}

	// Obtain the DNS name included with the certificate
	dnsName := ""
	cert, err := tlsCreds.LoadCertificate(string(h.certificate))
	if err != nil {
		s := fmt.Sprintf("Error forming transportCredentials: %+v", err)
		return errors.New(s)
	}
	if len(cert.DNSNames) > 0 {
		dnsName = cert.DNSNames[0]
	}

	// Create the TLS Credentials object
	h.credentials, err = tlsCreds.NewCredentialsFromPEM(string(h.certificate),
		dnsName)
	if err != nil {
		s := fmt.Sprintf("Error forming transportCredentials: %+v", err)
		return errors.New(s)
	}

	// Create the RSA Public Key object
	h.rsaPublicKey, err = tlsCreds.NewPublicKeyFromPEM(h.certificate)
	if err != nil {
		s := fmt.Sprintf("Error extracting PublicKey: %+v", err)
		return errors.New(s)
	}

	return nil
}

// Stringer interface for connection
func (h *Host) String() string {
	addr := h.address
	actualConnection := h.connection
	creds := h.credentials

	var state connectivity.State
	if actualConnection != nil {
		state = actualConnection.GetState()
	}

	serverName := "<nil>"
	protocolVersion := "<nil>"
	securityVersion := "<nil>"
	securityProtocol := "<nil>"
	if creds != nil {
		serverName = creds.Info().ServerName
		securityVersion = creds.Info().SecurityVersion
		protocolVersion = creds.Info().ProtocolVersion
		securityProtocol = creds.Info().SecurityProtocol
	}
	return fmt.Sprintf(
		"Addr: %v\tState: %v\tTLS ServerName: %v\t"+
			"TLS ProtocolVersion: %v\tTLS SecurityVersion: %v\t"+
			"TLS SecurityProtocol: %v\n",
		addr, state, serverName, protocolVersion,
		securityVersion, securityProtocol)
}
