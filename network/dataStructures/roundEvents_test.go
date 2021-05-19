///////////////////////////////////////////////////////////////////////////////
// Copyright © 2020 xx network SEZC                                          //
//                                                                           //
// Use of this source code is governed by a license that can be found in the //
// LICENSE file                                                              //
///////////////////////////////////////////////////////////////////////////////

package dataStructures

import (
	pb "gitlab.com/elixxir/comms/mixmessages"
	"gitlab.com/elixxir/comms/testutils"
	"gitlab.com/elixxir/primitives/states"
	"gitlab.com/xx_network/primitives/id"
	"testing"
	"time"
)

// AddRoundEvent should increase the number of round events in the data structure
func TestRoundEvents_AddRoundEvent(t *testing.T) {
	events := NewRoundEvents()

	// Adding with no states should result in no states with callbacks
	events.AddRoundEvent(id.Round(1), func(ri *pb.RoundInfo, timedOut bool) {}, time.Minute)
	for _, callback := range events.callbacks[id.Round(1)] {
		if len(callback) != 0 {
			t.Error("Adding round event with no states shouldn't add callback")
		}
	}

	// Adding with some states should result in one round added
	events.AddRoundEvent(id.Round(1), func(ri *pb.RoundInfo, timedOut bool) {}, time.Minute, states.PENDING, states.QUEUED)
	if len(events.callbacks) != 1 {
		t.Error("Adding round event with some states should make 1 round in the map")
	}

	// Adding another with a same state should result in two events for that state
	events.AddRoundEvent(id.Round(1), func(ri *pb.RoundInfo, timedOut bool) {}, time.Minute, states.PENDING)
	if len(events.callbacks) != 1 {
		t.Error("Adding round event with some states should make 1 round in the map")
	}
	if len(events.callbacks[id.Round(1)][states.PENDING]) != 2 {
		t.Error("Pending should have two events")
	}
	if len(events.callbacks[id.Round(1)][states.QUEUED]) != 1 {
		t.Error("Queued should have one event")
	}

	// It should be possible to add events to more than one round, of course
	events.AddRoundEvent(id.Round(2), func(ri *pb.RoundInfo, timedOut bool) {}, time.Minute, states.PENDING)
	if len(events.callbacks) != 2 {
		t.Error("Should have 2 rounds' events in the map")
	}
}

// AddRoundEvent should result in round timeouts after the specified amount of time
func TestRoundEvents_AddRoundEvent_Timeout(t *testing.T) {
	events := NewRoundEvents()

	called := false
	timeout := time.Millisecond
	events.AddRoundEvent(id.Round(1), func(ri *pb.RoundInfo, timedOut bool) {
		called = true
		if !timedOut {
			t.Error("Should have called event with timedOut true")
		}
	}, timeout, states.PENDING)

	time.Sleep(timeout * 2)
	if !called {
		t.Error("Event callback should have been called to let us know that we timed out")
	}
}

// Remove should remove one event from the data structure
// If there was one add call, removing it should leave the map empty
func TestRoundEvents_Remove(t *testing.T) {
	events := NewRoundEvents()
	rid := id.Round(1)
	callback := events.AddRoundEvent(rid, func(ri *pb.RoundInfo, timedOut bool) {}, time.Minute, states.PENDING, states.QUEUED)
	events.Remove(rid, callback)
	if len(events.callbacks) != 0 {
		t.Error("callbacks map should be empty after removing")
	}
}

// Round events should be callable after being added
func TestRoundEvents_TriggerRoundEvent(t *testing.T) {
	// Normal path
	events := NewRoundEvents()
	rid := id.Round(1)
	called := false
	events.AddRoundEvent(rid, func(ri *pb.RoundInfo, timedOut bool) {
		called = true
	}, time.Minute, states.PENDING)

	// Construct a mock round object
	ri := &pb.RoundInfo{
		ID:    uint64(rid),
		State: uint32(states.PENDING),
	}

	if err := testutils.SignRoundInfoRsa(ri, t); err != nil {
		t.Errorf("Failed to sign mock round info: %v", err)
	}

	pubKey, err := testutils.LoadPublicKeyTesting(t)
	if err != nil {
		t.Errorf("Failed to load public key: %v", err)
		t.FailNow()
	}
	rnd := NewRound(ri, pubKey, nil)
	events.TriggerRoundEvent(rnd)

	// wait for calling
	time.Sleep(5 * time.Millisecond)
	if !called {
		t.Error("callback should have been called")
	}
	if len(events.callbacks) != 0 {
		t.Error("callback should have been removed after calling")
	}

	// No matching round events: nothing should happen
	// (just to cover that branch)
	called = false
	events.TriggerRoundEvent(rnd)
	time.Sleep(5 * time.Millisecond)
	if called {
		t.Error("second trigger shouldn't have resulted in a call")
	}
}

// Add a round event with a channel and make sure it can be triggered
func TestRoundEvents_AddRoundEventChan(t *testing.T) {
	// Normal path
	events := NewRoundEvents()
	rid := id.Round(1)
	eventChan := make(chan EventReturn)
	events.AddRoundEventChan(rid, eventChan, time.Minute, states.PENDING)

	// Construct a mock round object
	ri := &pb.RoundInfo{
		ID:    uint64(rid),
		State: uint32(states.PENDING),
	}
	if err := testutils.SignRoundInfoRsa(ri, t); err != nil {
		t.Errorf("Failed to sign mock round info: %v", err)
	}

	pubKey, err := testutils.LoadPublicKeyTesting(t)
	if err != nil {
		t.Errorf("Failed to load public key: %v", err)
		t.FailNow()
	}
	rnd := NewRound(ri, pubKey, nil)
	events.TriggerRoundEvent(rnd)

	// wait for calling
	time.Sleep(5 * time.Millisecond)
	select {
	case <-eventChan:
		t.Log("event called")
	default:
		t.Error("no call")
	}
	if len(events.callbacks) != 0 {
		t.Error("callback should have been removed after calling")
	}
}
