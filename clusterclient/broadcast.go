////////////////////////////////////////////////////////////////////////////////
// Copyright © 2018 Privategrity Corporation                                   /
//                                                                             /
// All rights reserved.                                                        /
////////////////////////////////////////////////////////////////////////////////

// broadcast.go - comms client server functions that send to all servers in
//                the cluster.
package clusterclient

import (
	"golang.org/x/net/context"
	pb "gitlab.com/privategrity/comms/mixmessages"
	jww "github.com/spf13/jwalterweatherman"
)

func SetPublicKey(addr string, message *pb.PublicKeyMessage) (*pb.Ack, error) {
	// Attempt to connect to addr
	c := Connect(addr)
	ctx, cancel := DefaultContext()

	// Send the message
	result, err := c.SetPublicKey(ctx, message)

	// Make sure there are no errors with sending the message
	if err != nil {
		jww.ERROR.Printf("RealtimePermute: Error received: %s", err)
	}
	cancel()
	return result, err
}


func SendAskOnline(addr string, message *pb.Ping) (*pb.Pong, error) {
	// Attempt to connect to addr
	c := Connect(addr)
	ctx, cancel := DefaultContext()

	// Send the message
	result, err := c.AskOnline(ctx, message)

	// Make sure there are no errors with sending the message
	if err != nil {
		jww.ERROR.Printf("AskOnline: Error received: %s", err)
	}
	cancel()
	return result, err
}


func SendNetworkError(addr string, message *pb.ErrorMessage) (*pb.ErrorAck, error) {
	// Attempt to connect to addr
	c := Connect(addr)
	ctx, cancel := DefaultContext()

	// Send the message
	result, err := c.NetworkError(ctx, message)

	// Make sure there are no errors with sending the message
	if err != nil {
		jww.ERROR.Printf("NetworkError: Error received: %s", err)
	}
	cancel()
	return result, err
}

func SendNewRound(addr string, message *pb.InitRound) (*pb.InitRoundAck, error) {
	c := Connect(addr)

	// Send the message
	result, err := c.NewRound(context.Background(), message)

	// Make sure there are no errors with sending the message
	if err != nil {
		jww.ERROR.Printf("NewRound: Error received: %s", err)
	}
	return result, err
}