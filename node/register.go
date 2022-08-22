///////////////////////////////////////////////////////////////////////////////
// Copyright © 2020 xx network SEZC                                          //
//                                                                           //
// Use of this source code is governed by a license that can be found in the //
// LICENSE file                                                              //
///////////////////////////////////////////////////////////////////////////////

package node

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/pkg/errors"
	jww "github.com/spf13/jwalterweatherman"
	pb "gitlab.com/elixxir/comms/mixmessages"
	"gitlab.com/xx_network/comms/connect"
	"gitlab.com/xx_network/comms/messages"
)

// Server -> Registration Send Function
func (s *Comms) SendNodeRegistration(host *connect.Host,
	message *pb.NodeRegistration) error {

	// Create the Send Function
	f := func(conn connect.Connection) (*any.Any, error) {
		// Set up the context
		ctx, cancel := host.GetMessagingContext()
		defer cancel()

		// Send the message
		var resultMsg *messages.Ack
		var err error
		if conn.IsWeb() {
			wc := conn.GetWebConn()
			err = wc.Invoke(ctx, "/mixmessages.Registration/RegisterNode",
				message, resultMsg)
			if err != nil {
				return nil, err
			}
		} else {
			_, err = pb.NewRegistrationClient(conn.GetGrpcConn()).
				RegisterNode(ctx, message)
			if err != nil {
				err = errors.New(err.Error())
			}
		}
		return nil, err
	}

	// Execute the Send function
	jww.TRACE.Printf("Sending Node Registration message: %+v", message)
	_, err := s.Send(host, f)
	return err
}

// Server -> Registration Send Function
func (s *Comms) SendPoll(host *connect.Host,
	message *pb.PermissioningPoll) (*pb.PermissionPollResponse, error) {

	// Create the Send Function
	f := func(conn connect.Connection) (*any.Any, error) {
		// Set up the context
		ctx, cancel := host.GetMessagingContext()
		defer cancel()
		// Pack the message for server
		authMsg, err := s.PackAuthenticatedMessage(message, host, false)
		if err != nil {
			return nil, errors.New(err.Error())
		}

		// Send the message
		var resultMsg *pb.PermissionPollResponse
		if conn.IsWeb() {
			wc := conn.GetWebConn()
			err = wc.Invoke(
				ctx, "/mixmessages.Registration/Poll", authMsg, resultMsg)
			if err != nil {
				return nil, err
			}
		} else {
			resultMsg, err = pb.NewRegistrationClient(conn.GetGrpcConn()).
				Poll(ctx, authMsg)
			if err != nil {
				return nil, errors.New(err.Error())
			}
		}
		return ptypes.MarshalAny(resultMsg)
	}

	// Execute the Send function
	jww.TRACE.Printf("Sending Poll message...")
	resultMsg, err := s.Send(host, f)
	if err != nil {
		return nil, err
	}

	// Marshall the result
	result := &pb.PermissionPollResponse{}
	return result, ptypes.UnmarshalAny(resultMsg, result)
}

// Server -> Registration Send Function
func (s *Comms) SendRegistrationCheck(host *connect.Host,
	message *pb.RegisteredNodeCheck) (*pb.RegisteredNodeConfirmation, error) {
	// Create the Send Function
	f := func(conn connect.Connection) (*any.Any, error) {
		// Set up the context
		ctx, cancel := host.GetMessagingContext()
		defer cancel()

		// Send the message
		var resultMsg *pb.RegisteredNodeConfirmation
		var err error
		if conn.IsWeb() {
			wc := conn.GetWebConn()
			err = wc.Invoke(ctx, "/mixmessages.Registration/CheckRegistration",
				message, resultMsg)
			if err != nil {
				return nil, err
			}
		} else {
			resultMsg, err = pb.NewRegistrationClient(conn.GetGrpcConn()).
				CheckRegistration(ctx, message)
			if err != nil {
				return nil, errors.New(err.Error())
			}
		}
		return ptypes.MarshalAny(resultMsg)

	}

	// Execute the Send function
	jww.TRACE.Printf("Sending Node Registration Check message: %+v", message)
	resultMsg, err := s.Send(host, f)
	if err != nil {
		return nil, err
	}

	// Marshall the result
	result := &pb.RegisteredNodeConfirmation{}
	return result, ptypes.UnmarshalAny(resultMsg, result)
}

// Server -> Authorizer Send Function
func (s *Comms) SendAuthorizerAuth(host *connect.Host,
	message *pb.AuthorizerAuth) (*messages.Ack, error) {
	// Create the Send Function
	f := func(conn connect.Connection) (*any.Any, error) {
		// Set up the context
		ctx, cancel := host.GetMessagingContext()
		defer cancel()

		// Send the message
		var resultMsg *messages.Ack
		var err error
		if conn.IsWeb() {
			wc := conn.GetWebConn()
			err = wc.Invoke(
				ctx, "/mixmessages.Authorizer/Authorize", message, resultMsg)
			if err != nil {
				return nil, err
			}
		} else {
			resultMsg, err = pb.NewAuthorizerClient(conn.GetGrpcConn()).
				Authorize(ctx, message)
			if err != nil {
				return nil, errors.New(err.Error())
			}
		}
		return ptypes.MarshalAny(resultMsg)

	}

	// Execute the Send function
	jww.TRACE.Printf("Sending Authorizer Authorize message: %+v", message)
	resultMsg, err := s.Send(host, f)
	if err != nil {
		return nil, err
	}

	// Marshall the result
	result := &messages.Ack{}
	return result, ptypes.UnmarshalAny(resultMsg, result)
}
