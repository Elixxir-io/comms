////////////////////////////////////////////////////////////////////////////////
// Copyright © 2018 Privategrity Corporation                                   /
//                                                                             /
// All rights reserved.                                                        /
////////////////////////////////////////////////////////////////////////////////

package node

import (
	jww "github.com/spf13/jwalterweatherman"
	"gitlab.com/privategrity/comms/connect"
	pb "gitlab.com/privategrity/comms/mixmessages"
)

// SendReceiveBatch sends a batch to the gateway
func SendReceiveBatch(addr string, message []*pb.CmixMessage) error {
	// Attempt to connect to addr
	c := connect.ConnectToGateway(addr)
	ctx, cancel := connect.DefaultContext()

	outputMessages := pb.OutputMessages{Messages: message}

	_, err := c.ReceiveBatch(ctx, &outputMessages)

	// Make sure there are no errors with sending the message
	if err != nil {
		jww.ERROR.Printf("ReceiveBatch(): Error received: %s", err)
	}
	cancel()
	return err
}