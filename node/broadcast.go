////////////////////////////////////////////////////////////////////////////////
// Copyright © 2018 Privategrity Corporation                                   /
//                                                                             /
// All rights reserved.                                                        /
////////////////////////////////////////////////////////////////////////////////

// Contains server -> all servers functionality

package node

import (
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/pkg/errors"
	jww "github.com/spf13/jwalterweatherman"
	"gitlab.com/elixxir/comms/connect"
	pb "gitlab.com/elixxir/comms/mixmessages"
)

func (s *Server) SendServerMetrics(id fmt.Stringer,
	message *pb.ServerMetrics) (*pb.Ack, error) {
	// Attempt to connect to addr
	c := s.ConnectToNode(id, nil)
	ctx, cancel := connect.DefaultContext()

	// Send the message
	result, err := c.GetServerMetrics(ctx, message)

	// Make sure there are no errors with sending the message
	if err != nil {
		err = errors.New(err.Error())
		jww.ERROR.Printf("ServerMetrics: Error received: %+v", err)
	}

	cancel()
	return result, err
}

func (s *Server) SendRoundtripPing(id fmt.Stringer,
	message *pb.TimePing) (*pb.Ack, error) {
	// Attempt to connect to addr
	c := s.ConnectToNode(id, nil)
	ctx, cancel := connect.DefaultContext()

	// Send the message
	result, err := c.RoundtripPing(ctx, message)

	// Make sure there are no errors with sending the message
	if err != nil {
		err = errors.New(err.Error())
		jww.ERROR.Printf("RoundtripPing: Error received: %+v", err)
	}

	cancel()
	return result, err
}

func (s *Server) SendAskOnline(id fmt.Stringer, message *pb.Ping) (
	*pb.Ack, error) {
	// Attempt to connect to addr
	c := s.ConnectToNode(id, nil)
	ctx, cancel := connect.DefaultContext()

	// Send the message
	result, err := c.AskOnline(ctx, message,
		grpc_retry.WithMax(connect.MAX_RETRIES))

	// Make sure there are no errors with sending the message
	if err != nil {
		err = errors.New(err.Error())
		jww.ERROR.Printf("AskOnline: Error received: %+v", err)
	}

	cancel()
	return result, err
}

func (s *Server) SendNewRound(id fmt.Stringer, message *pb.RoundInfo) (
	*pb.Ack, error) {
	c := s.ConnectToNode(id, nil)
	ctx, cancel := connect.DefaultContext()

	// Send the message
	result, err := c.CreateNewRound(ctx, message,
		grpc_retry.WithMax(connect.MAX_RETRIES))

	// Make sure there are no errors with sending the message
	if err != nil {
		jww.ERROR.Printf("NewRound: Error received: %+v", err)
	}

	cancel()
	return result, err
}

func (s *Server) SendPostRoundPublicKey(id fmt.Stringer,
	message *pb.RoundPublicKey) (*pb.Ack, error) {
	c := s.ConnectToNode(id, nil)
	ctx, cancel := connect.DefaultContext()

	// Send the message
	result, err := c.PostRoundPublicKey(ctx, message,
		grpc_retry.WithMax(connect.MAX_RETRIES))

	// Make sure there are no errors with sending the message
	if err != nil {
		err = errors.New(err.Error())
		jww.ERROR.Printf("SendPostRoundPublicKey: Error received: %+v", err)
	}

	cancel()
	return result, err
}

// SendPostPrecompResult sends the final message and AD precomputations to
// other nodes.
func (s *Server) SendPostPrecompResult(id fmt.Stringer,
	roundID uint64, slots []*pb.Slot) (*pb.Ack, error) {
	c := s.ConnectToNode(id, nil)
	ctx, cancel := connect.DefaultContext()

	// Send the message
	result, err := c.PostPrecompResult(ctx,
		&pb.Batch{
			Round: &pb.RoundInfo{
				ID: roundID,
			},
			Slots: slots,
		},
		grpc_retry.WithMax(connect.MAX_RETRIES))

	// Make sure there are no errors with sending the message
	if err != nil {
		err = errors.New(err.Error())
		jww.ERROR.Printf("PostPrecompResult: Error received: %+v",
			err)
	}

	cancel()
	return result, err
}
