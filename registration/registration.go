////////////////////////////////////////////////////////////////////////////////
// Copyright © 2018 Privategrity Corporation                                   /
//                                                                             /
// All rights reserved.                                                        /
////////////////////////////////////////////////////////////////////////////////

// Contains registration server comms initialization functionality

package registration

import (
	jww "github.com/spf13/jwalterweatherman"
	pb "gitlab.com/elixxir/comms/mixmessages"
	"gitlab.com/elixxir/comms/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"math"
	"net"
	"time"
)

// Callback interface provided by the Server repository to StartServer
var registrationHandler Handler

// Server object containing a GRPC server
type server struct {
	gs *grpc.Server
}

// Performs a graceful shutdown of the server
func (s *server) ShutDown() {
	s.gs.GracefulStop()
	time.Sleep(time.Millisecond * 500)
}

// Starts a new server on the address:port specified by localServer
// and a callback interface for server operations
// with given path to public and private key for TLS connection
func StartRegistrationServer(localServer string, handler Handler,
	certPath, keyPath string) func() {
	var grpcServer *grpc.Server
	// Set the serverHandler
	registrationHandler = handler

	// Listen on the given address
	lis, err := net.Listen("tcp", localServer)
	if err != nil {
		jww.FATAL.Panicf("Failed to listen: %v", err)
	}

	// If TLS was specified
	if certPath != "" && keyPath != "" {
		// Create the TLS credentials
		certPath = utils.GetFullPath(certPath)
		keyPath = utils.GetFullPath(keyPath)
		creds, err := credentials.NewServerTLSFromFile(certPath, keyPath)
		if err != nil {
			jww.FATAL.Panicf("Could not load TLS keys: %s", err)
		}

		// Create the GRPC server with TLS
		jww.INFO.Printf("Starting server with TLS...")
		grpcServer = grpc.NewServer(grpc.Creds(creds),
			grpc.MaxConcurrentStreams(math.MaxUint32),
			grpc.MaxRecvMsgSize(math.MaxInt32))
	} else {
		// Create the GRPC server without TLS
		jww.INFO.Printf("Starting server with TLS disabled...")
		grpcServer = grpc.NewServer(grpc.MaxConcurrentStreams(math.MaxUint32),
			grpc.MaxRecvMsgSize(math.MaxInt32))
	}
	registrationServer := server{gs: grpcServer}

	go func() {
		// Make the port close when the gateway dies
		defer func() {
			err := lis.Close()
			if err != nil {
				jww.WARN.Printf("Unable to close listening port: %s", err.Error())
			}
		}()

		pb.RegisterRegistrationServer(registrationServer.gs, &registrationServer)

		// Register reflection service on gRPC server.
		reflection.Register(registrationServer.gs)
		if err := registrationServer.gs.Serve(lis); err != nil {
			jww.FATAL.Panicf("failed to serve: %v", err)
		}
	}()

	return registrationServer.ShutDown
}