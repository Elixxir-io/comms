////////////////////////////////////////////////////////////////////////////////
// Copyright © 2018 Privategrity Corporation                                   /
//                                                                             /
// All rights reserved.                                                        /
////////////////////////////////////////////////////////////////////////////////

package gateway

import (
	"fmt"
	"gitlab.com/elixxir/comms/mixmessages"
	"gitlab.com/elixxir/comms/node"
	"gitlab.com/elixxir/comms/testkeys"
	"sync"
	"testing"
)

var serverPortLock sync.Mutex
var serverPort = 5500

func getNextServerAddress() string {
	serverPortLock.Lock()
	defer func() {
		serverPort++
		serverPortLock.Unlock()
	}()
	return fmt.Sprintf("localhost:%d", serverPort)
}

var gatewayPortLock sync.Mutex
var gatewayPort = 5600

func getNextGatewayAddress() string {
	gatewayPortLock.Lock()
	defer func() {
		gatewayPort++
		gatewayPortLock.Unlock()
	}()
	return fmt.Sprintf("localhost:%d", gatewayPort)
}

// Tests whether the gateway can be connected to and run an RPC with TLS enabled
func TestTLS(t *testing.T) {
	GatewayAddress := getNextServerAddress()
	shutdown := StartGateway(GatewayAddress, NewImplementation(),
		testkeys.GetGatewayCertPath(), testkeys.GetGatewayKeyPath())
	// Reset TLS-related global variables
	defer shutdown()
	err := node.SendReceiveBatch(GatewayAddress, testkeys.GetGatewayCertPath(),
		"", []*mixmessages.CmixBatch{})
	if err != nil {
		t.Error(err)
	}
}
