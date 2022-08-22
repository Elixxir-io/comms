///////////////////////////////////////////////////////////////////////////////
// Copyright © 2020 xx network SEZC                                          //
//                                                                           //
// Use of this source code is governed by a license that can be found in the //
// LICENSE file                                                              //
///////////////////////////////////////////////////////////////////////////////

// Contains send functions used for polling

package udb

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/pkg/errors"
	jww "github.com/spf13/jwalterweatherman"
	pb "gitlab.com/elixxir/comms/mixmessages"
	"gitlab.com/xx_network/comms/connect"
)

// RequestNdf is used by User Discovery to Request a NDF from permissioning
func (u *Comms) RequestNdf(host *connect.Host) (*pb.NDF, error) {

	// Create the Send Function
	f := func(conn connect.Connection) (*any.Any, error) {
		// Set up the context
		ctx, cancel := host.GetMessagingContext()
		defer cancel()

		message := &pb.NDFHash{Hash: make([]byte, 0)}

		// Send the message
		var resultMsg *pb.NDF
		var err error
		if conn.IsWeb() {
			wc := conn.GetWebConn()
			err = wc.Invoke(
				ctx, "/mixmessages.Registration/PollNdf", message, resultMsg)
			if err != nil {
				return nil, err
			}
		} else {
			resultMsg, err = pb.NewRegistrationClient(conn.GetGrpcConn()).
				PollNdf(ctx, message)
			if err != nil {
				return nil, errors.New(err.Error())
			}
		}
		return ptypes.MarshalAny(resultMsg)
	}

	// Execute the Send function
	jww.TRACE.Printf("Sending Request Ndf message...")
	resultMsg, err := u.Send(host, f)
	if err != nil {
		return nil, err
	}

	result := &pb.NDF{}
	return result, ptypes.UnmarshalAny(resultMsg, result)
}
