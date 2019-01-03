////////////////////////////////////////////////////////////////////////////////
// Copyright © 2018 Privategrity Corporation                                   /
//                                                                             /
// All rights reserved.                                                        /
////////////////////////////////////////////////////////////////////////////////

// Contains server comms initialization functionality

package node

import (
	pb "gitlab.com/elixxir/comms/mixmessages"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	jww "github.com/spf13/jwalterweatherman"
	"gitlab.com/elixxir/comms/utils"
	"math"
	"net"
	"time"
)

// Callback interface provided by the Server repository to StartServer
var serverHandler ServerHandler

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
func StartServer(localServer string, handler ServerHandler,
	certPath, keyPath string) func() {
	var grpcServer *grpc.Server
	// Set the serverHandler
	serverHandler = handler

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
	mixmessageServer := server{gs: grpcServer}

	go func() {
		// Make the port close when the gateway dies
		defer func() {
			err := lis.Close()
			if err != nil {
				jww.WARN.Printf("Unable to close listening port: %s", err.Error())
			}
		}()

		pb.RegisterMixMessageNodeServer(mixmessageServer.gs, &mixmessageServer)

		// Register reflection service on gRPC server.
		reflection.Register(mixmessageServer.gs)
		if err := mixmessageServer.gs.Serve(lis); err != nil {
			jww.FATAL.Panicf("failed to serve: %v", err)
		}
	}()

	return mixmessageServer.ShutDown
}
