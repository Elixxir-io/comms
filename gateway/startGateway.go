////////////////////////////////////////////////////////////////////////////////
// Copyright © 2018 Privategrity Corporation                                   /
//                                                                             /
// All rights reserved.                                                        /
////////////////////////////////////////////////////////////////////////////////

package gateway

import (
	jww "github.com/spf13/jwalterweatherman"
	pb "gitlab.com/privategrity/comms/mixmessages"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

// Passed into StartGateway to serve as an interface
// for interacting with the gateway repo
var gatewayHandler GatewayHandler

// gateway object
type gateway struct {
	gs *grpc.Server
}

func ShutDown(s *gateway) {
	time.Sleep(time.Millisecond * 500)
	s.gs.GracefulStop()
}

// Start local comm server
func StartGateway(localServer string, handler GatewayHandler) {
	// Set the gatewayHandler
	gatewayHandler = handler

	// Listen on the given address
	lis, err := net.Listen("tcp", localServer)
	if err != nil {
		jww.FATAL.Panicf("failed to listen: %v", err)
	}
	mixmessageServer := gateway{gs: grpc.NewServer()}
	pb.RegisterMixMessageGatewayServer(mixmessageServer.gs, &mixmessageServer)

	// Register reflection service on gRPC server.
	reflection.Register(mixmessageServer.gs)
	if err := mixmessageServer.gs.Serve(lis); err != nil {
		jww.FATAL.Panicf("failed to serve: %v", err)
	}
}
