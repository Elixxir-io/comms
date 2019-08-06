////////////////////////////////////////////////////////////////////////////////
// Copyright © 2018 Privategrity Corporation                                   /
//                                                                             /
// All rights reserved.                                                        /
////////////////////////////////////////////////////////////////////////////////

package gateway

import (
	pb "gitlab.com/elixxir/comms/mixmessages"
	"gitlab.com/elixxir/comms/node"
	"testing"
)

// Smoke test SendRequestNonceMessage
func TestSendRequestNonceMessage(t *testing.T) {
	GatewayAddress := getNextGatewayAddress()
	ServerAddress := getNextServerAddress()
	gateway := StartGateway(GatewayAddress, NewImplementation(), nil, nil)
	server := node.StartNode(ServerAddress, node.NewImplementation(),
		nil, nil)
	defer gateway.Shutdown()
	defer server.Shutdown()
	connID := MockID("gatewayToServer")
	gateway.ConnectToNode(connID, ServerAddress, nil)

	_, err := gateway.SendRequestNonceMessage(connID, &pb.NonceRequest{})
	if err != nil {
		t.Errorf("SendRequestNonceMessage: Error received: %s", err)
	}
}

// Smoke test SendConfirmNonceMessage
func TestSendConfirmNonceMessage(t *testing.T) {
	GatewayAddress := getNextGatewayAddress()
	ServerAddress := getNextServerAddress()
	gateway := StartGateway(GatewayAddress, NewImplementation(), nil, nil)
	server := node.StartNode(ServerAddress, node.NewImplementation(),
		nil, nil)
	defer gateway.Shutdown()
	defer server.Shutdown()
	connID := MockID("gatewayToServer")
	gateway.ConnectToNode(connID, ServerAddress, nil)

	_, err := gateway.SendConfirmNonceMessage(connID, &pb.DSASignature{})
	if err != nil {
		t.Errorf("SendConfirmNonceMessage: Error received: %s", err)
	}
}

func TestSendGetSignedCertMessage(t *testing.T) {
	GatewayAddress := getNextGatewayAddress()
	ServerAddress := getNextServerAddress()

	gateway := StartGateway(GatewayAddress, NewImplementation(), nil, nil)
	server := node.StartNode(ServerAddress, node.NewImplementation(), nil, nil)
	defer gateway.Shutdown()
	defer server.Shutdown()
	connID := MockID("gatewayToServer")
	gateway.ConnectToNode(connID, ServerAddress, nil)

	_, err := gateway.SendGetSignedCertMessage(connID, &pb.Ping{})
	if err != nil {
		t.Errorf("SendGetSignedCertMessage: Error received: %s", err)
	}
}
