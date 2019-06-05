////////////////////////////////////////////////////////////////////////////////
// Copyright © 2018 Privategrity Corporation                                   /
//                                                                             /
// All rights reserved.                                                        /
////////////////////////////////////////////////////////////////////////////////

// Contains server -> server functionality for precomputation operations

package node

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/pkg/errors"
	jww "github.com/spf13/jwalterweatherman"
	"gitlab.com/elixxir/comms/connect"
	pb "gitlab.com/elixxir/comms/mixmessages"
)

func (s *NodeComms) SendPostPhase(id fmt.Stringer,
	message *pb.Batch) (*pb.Ack, error) {
	// Attempt to connect to addr
	c := s.GetNodeConnection(id)
	ctx, cancel := connect.DefaultContext()

	// Send the message
	result, err := c.PostPhase(ctx, message,
		grpc_retry.WithMax(connect.MAX_RETRIES))

	// Make sure there are no errors with sending the message
	if err != nil {
		err = errors.New(err.Error())
		jww.ERROR.Printf("PostPhase: Error received: %+v", err)
	}

	cancel()
	return result, err
}

// GetPostPhaseStream uses an id and streaming context to retrieve
// a Node_StreamPostPhaseClient object otherwise it returns
// an error if the connection is unavailable
func (nodeComms *NodeComms) GetPostPhaseStream(id fmt.Stringer, ctx context.Context) (
	pb.Node_StreamPostPhaseClient, error) {

	// Attempt to connect to addr
	c := nodeComms.GetNodeConnection(id)

	// Get the stream client using streaming context
	streamClient, err := c.StreamPostPhase(ctx,
		grpc_retry.WithMax(connect.MAX_RETRIES))

	// Make sure there are no errors with getting the stream client
	if err != nil {
		err = errors.New(err.Error())
		jww.ERROR.Printf("GetPostPhaseStream: Error received: %+v", err)
		return nil, err
	}

	return streamClient, nil
}
