///////////////////////////////////////////////////////////////////////////////
// Copyright © 2020 xx network SEZC                                          //
//                                                                           //
// Use of this source code is governed by a license that can be found in the //
// LICENSE file                                                              //
///////////////////////////////////////////////////////////////////////////////

package registration

import (
	"fmt"
	"git.xx.network/xx_network/primitives/id"
	"sync"
	"testing"
)

var serverPortLock sync.Mutex
var serverPort = 5900

func getNextServerAddress() string {
	serverPortLock.Lock()
	defer func() {
		serverPort++
		serverPortLock.Unlock()
	}()
	return fmt.Sprintf("localhost:%d", serverPort)
}

func TestBadCerts(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	RegAddress := getNextServerAddress()
	testId := id.NewIdFromString("test", id.Generic, t)

	_ = StartRegistrationServer(testId, RegAddress, NewImplementation(), []byte("bad cert"), []byte("bad key"), nil)
}
