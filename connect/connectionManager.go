////////////////////////////////////////////////////////////////////////////////
// Copyright © 2018 Privategrity Corporation                                   /
//                                                                             /
// All rights reserved.                                                        /
////////////////////////////////////////////////////////////////////////////////

// Contains functionality for connecting to gateways and servers

package connect

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	jww "github.com/spf13/jwalterweatherman"
	pb "gitlab.com/elixxir/comms/mixmessages"
	"gitlab.com/elixxir/crypto/signature/rsa"
	tlsCreds "gitlab.com/elixxir/crypto/tls"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials"
	"math"
	"sort"
	"sync"
	"time"
)

// Stores information used to connect to a server
type ConnectionInfo struct {
	Address string
	// You can also get the server name from Creds if you need it
	Creds        credentials.TransportCredentials
	RsaPublicKey *rsa.PublicKey
	Connection   *grpc.ClientConn
}

type ConnectionManager struct {
	// A map of string IDs to open connections
	connections     map[string]*ConnectionInfo
	connectionsLock sync.Mutex
	privateKey      *rsa.PrivateKey
}

// Default maximum number of retries
const MAX_RETRIES = 100

// Set private key to data to a PEM block
func (m *ConnectionManager) SetPrivateKey(data []byte) error {
	key, err := rsa.LoadPrivateKeyFromPem(data)
	if err != nil {
		s := fmt.Sprintf("Failed to form private key file from data at %s: %+v", data, err)
		return errors.New(s)
	}

	m.privateKey = key
	return nil
}

// Get connection manager's private key
func (m *ConnectionManager) GetPrivateKey() *rsa.PrivateKey {
	return m.privateKey
}

func (m *ConnectionManager) GetConnectionInfo(id string) *ConnectionInfo {
	return m.connections[id]
}

// DEPRECATED - Use ConnectToRemote instead
// Connect to a certain registration server
// connectionInfo can be nil if the connection already exists for this id
func (m *ConnectionManager) ConnectToRegistration(id fmt.Stringer,
	addr string, certPEMblock []byte, hasTimeout bool) error {
	return m.ConnectToRemote(id, addr, certPEMblock, hasTimeout)
}

// ConnectToRemote connects to a remote server at address addr with the passed
// cert. The connection is stored locally at the passed id.  that ID can be
// used to identify the private keys of the sender of incoming messages so
// it must be the same as used across the network.
func (m *ConnectionManager) ConnectToRemote(id fmt.Stringer,
	addr string, certPEMblock []byte, hasTimeout bool) error {
	// Make TransportCredentials
	var creds credentials.TransportCredentials
	var pubKey *rsa.PublicKey

	if certPEMblock != nil && len(certPEMblock) != 0 {

		var err error

		//Gets the DNS name from the cert so it cna override for testing
		//fix-me: this should not run on a live deployment
		cert, err := tlsCreds.LoadCertificate(string(certPEMblock))

		if err != nil {
			s := fmt.Sprintf("Error forming transportCredentials: %+v", err)
			return errors.New(s)
		}

		jww.DEBUG.Printf("Cert: %+v", cert)

		dnsName := ""
		if len(cert.DNSNames) > 0 {
			dnsName = cert.DNSNames[0]
		}

		//create the TLS cert
		creds, err = tlsCreds.NewCredentialsFromPEM(string(certPEMblock),
			dnsName)
		if err != nil {
			s := fmt.Sprintf("Error forming transportCredentials: %+v", err)
			return errors.New(s)
		}

		pubKey, err = tlsCreds.NewPublicKeyFromPEM(certPEMblock)
		if err != nil {
			s := fmt.Sprintf("Error extracting PublicKey: %+v", err)
			return errors.New(s)
		}
	}
	// NewCredentialsFromPem, NewCredentialsFromFile, NewP
	m.connect(id.String(), addr, creds, pubKey, hasTimeout)
	return nil
}

func (m *ConnectionManager) GetRegistrationConnection(id fmt.Stringer) pb.
	RegistrationClient {
	conn := m.get(id)
	return pb.NewRegistrationClient(conn)
}

// DEPRECATED - Use ConnectToRemote instead
// Connect to a certain gateway
// connectionInfo can be nil if the connection already exists for this id
func (m *ConnectionManager) ConnectToGateway(id fmt.Stringer,
	addr string, certPEMblock []byte, hasTimeout bool) error {
	return m.ConnectToRemote(id, addr, certPEMblock, hasTimeout)
}

// DEPRECATED - Use ConnectToRemote instead
func (m *ConnectionManager) GetGatewayConnection(id fmt.Stringer) pb.
	GatewayClient {
	conn := m.get(id)
	return pb.NewGatewayClient(conn)
}

// Connect to a certain node
// connectionInfo can be nil if the connection already exists for this id
// Should this return an error if the connection doesn't exist and the
// connection info is nil?
func (m *ConnectionManager) ConnectToNode(id fmt.Stringer,
	addr string, certPEMblock []byte, hasTimeout bool) error {
	return m.ConnectToRemote(id, addr, certPEMblock, hasTimeout)
}

func (m *ConnectionManager) GetNodeConnection(id fmt.Stringer) pb.NodeClient {
	conn := m.get(id)
	return pb.NewNodeClient(conn)
}

// Returns true if the connection is non-nil and alive
func isConnectionGood(connection *grpc.ClientConn) bool {
	if connection == nil {
		return false
	}
	state := connection.GetState()
	return state == connectivity.Idle || state == connectivity.Connecting ||
		state == connectivity.Ready
}

// Get creates an existing connection
func (m *ConnectionManager) get(id fmt.Stringer) *grpc.ClientConn {
	m.connectionsLock.Lock()
	// TODO Retry/reconnect here based on current connection state?
	//  I think this could be made more robust to handle TransientFailure
	conn, ok := m.connections[id.String()]
	if !ok {
		jww.FATAL.Panicf("No connection exists for the ID \"" + id.String() + "\"")
	}
	m.connectionsLock.Unlock()
	return conn.Connection
}

// Connect creates a connection
func (m *ConnectionManager) connect(id string, addr string,
	tls credentials.TransportCredentials, pubKey *rsa.PublicKey, hasTimeout bool) {

	// Create top level vars
	var connection *grpc.ClientConn
	var err error
	connection = nil
	err = nil

	var securityDial grpc.DialOption
	if tls != nil {
		// Create the gRPC client with TLS
		securityDial = grpc.WithTransportCredentials(tls)
	} else {
		// Create the gRPC client without TLS
		jww.WARN.Printf("Connecting to %v without TLS!", addr)
		securityDial = grpc.WithInsecure()
	}

	if m.connections == nil {
		m.connections = make(map[string]*ConnectionInfo)
	}
	jww.DEBUG.Printf("Trying to connect to %v", addr)

	//Set the max number depending on if we want to timeout or not
	var maxRetries int64
	if hasTimeout {
		maxRetries = 100
	} else {
		maxRetries = math.MaxInt64
	}

	// Create a new connection if we are not present or disconnecting/disconnected
	for numRetries := int64(0); numRetries < maxRetries && !isConnectionGood(connection); numRetries++ {

		ctx, cancel := TimeoutContext(time.Duration(2 * (numRetries/16 + 1)))

		// Create the connection
		connection, err = grpc.DialContext(ctx, addr,
			securityDial, grpc.WithBlock())
		if err != nil {
			jww.ERROR.Printf("Attempt number %+v to connect to %s failed: %+v\n", numRetries, addr,
				errors.New(err.Error()))
		}

		cancel()
	}

	if !isConnectionGood(connection) {
		jww.FATAL.Panicf("Last try to connect to %s failed. Giving up", addr)
	} else {
		// Connection succeeded, so add it to the map along with any information
		// needed for reconnection
		jww.INFO.Printf("Successfully connected to %s at %v", id, addr)
		m.connectionsLock.Lock()
		m.connections[id] = &ConnectionInfo{
			Address:      addr,
			Creds:        tls,
			Connection:   connection,
			RsaPublicKey: pubKey,
		}
		m.connectionsLock.Unlock()
	}
}

// Disconnect closes client connections and removes them from the connection map
func (m *ConnectionManager) Disconnect(id string) {
	m.connectionsLock.Lock()
	connection, present := m.connections[id]
	if present {
		err := connection.Connection.Close()
		if err != nil {
			jww.ERROR.Printf("Unable to close connection to %s: %+v", id,
				errors.New(err.Error()))
		}
		delete(m.connections, id)
	}
	m.connectionsLock.Unlock()
}

// DisconnectAll closes alld client connections and removes them from the connection map
func (m *ConnectionManager) DisconnectAll() {

	m.connectionsLock.Lock()

	for id, connection := range m.connections {
		err := connection.Connection.Close()
		if err != nil {
			jww.ERROR.Printf("Unable to close connection to %s: %+v", id,
				errors.New(err.Error()))
		}
		delete(m.connections, id)
	}

	m.connectionsLock.Unlock()
}

// implements Stringer for debug printing
func (m *ConnectionManager) String() string {
	m.connectionsLock.Lock()
	defer m.connectionsLock.Unlock()

	// Sort connection IDs to print in a consistent order
	keys := make([]string, len(m.connections))
	i := 0
	for key := range m.connections {
		keys[i] = key
		i++
	}
	sort.Strings(keys)

	// Print each connection's information
	var result bytes.Buffer
	for _, key := range keys {
		// Populate fields without ever dereferencing nil
		connection := m.connections[key]
		if connection != nil {
			addr := connection.Address
			actualConnection := connection.Connection
			creds := connection.Creds

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
			result.WriteString(fmt.Sprintf(
				"[%v] Addr: %v\tState: %v\tTLS ServerName: %v\t"+
					"TLS ProtocolVersion: %v\tTLS SecurityVersion: %v\t"+
					"TLS SecurityProtocol: %v\n",
				key, addr, state, serverName, protocolVersion,
				securityVersion, securityProtocol))
		}
	}

	return result.String()
}

//TimeoutContext is the basis for the default timeout
func TimeoutContext(seconds time.Duration) (context.Context, context.CancelFunc) {
	waitingPeriod := seconds * time.Second
	jww.DEBUG.Printf("Timing out in: %s", waitingPeriod)
	ctx, cancel := context.WithTimeout(context.Background(),
		waitingPeriod)
	return ctx, cancel

}

// DefaultContexts creates a context object with the default context
// for all client messages. This is primarily used to set the default
// timeout for all clients
func DefaultContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(),
		10000*time.Millisecond)
	return ctx, cancel
}

// StreamingContext creates a context object with the default context
// for all client streaming messages. This is primarily used to
// allow a cancel option for clients and is suitable for unary streaming.
func StreamingContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	return ctx, cancel
}
