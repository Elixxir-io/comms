///////////////////////////////////////////////////////////////////////////////
// Copyright © 2020 xx network SEZC                                          //
//                                                                           //
// Use of this source code is governed by a license that can be found in the //
// LICENSE file                                                              //
///////////////////////////////////////////////////////////////////////////////

package dataStructures

import (
	"gitlab.com/elixxir/comms/mixmessages"
	"gitlab.com/elixxir/comms/testutils"
	"testing"
)

func TestUpdates_AddRound(t *testing.T) {
	u := NewUpdates()
	// Construct a mock round object
	ri := &mixmessages.RoundInfo{
		ID:       0,
		UpdateID: 0,
	}
	pubKey, err := testutils.LoadPublicKeyTesting(t)
	if err != nil {
		t.Errorf("Failed to load public key: %v", err)
		t.FailNow()
	}
	ecKey, _ := testutils.LoadEllipticPublicKey()

	rnd := NewRound(ri, pubKey, ecKey.PublicKey())
	err = u.AddRound(rnd)
	if err != nil {
		t.Errorf("Failed to add round: %+v", err)
	}
}

func TestUpdates_GetUpdate(t *testing.T) {
	u := NewUpdates()
	updateID := 3
	// Construct a mock round object
	ri := &mixmessages.RoundInfo{
		ID:       0,
		UpdateID: uint64(updateID),
	}
	if err := testutils.SignRoundInfoRsa(ri, t); err != nil {
		t.Errorf("Failed to sign mock round info: %v", err)
	}
	pubKey, err := testutils.LoadPublicKeyTesting(t)
	if err != nil {
		t.Errorf("Failed to load public key: %v", err)
		t.FailNow()
	}

	ecKey, _ := testutils.LoadEllipticPublicKey()
	if err := testutils.SignRoundInfoEddsa(ri, ecKey, t); err != nil {
		t.Errorf("Failed to sign mock round info: %v", err)
	}

	rnd := NewRound(ri, pubKey, ecKey.PublicKey())
	_ = u.AddRound(rnd)
	_, err = u.GetUpdate(updateID)
	if err != nil {
		t.Errorf("Failed to get update: %+v", err)
	}
}

func TestUpdates_GetUpdates(t *testing.T) {
	u := NewUpdates()
	updateID := 3
	// Construct a mock round object
	roundInfoOne := &mixmessages.RoundInfo{
		ID:       0,
		UpdateID: uint64(updateID),
	}
	if err := testutils.SignRoundInfoRsa(roundInfoOne, t); err != nil {
		t.Errorf("Failed to sign mock round info: %v", err)
	}
	pubKey, err := testutils.LoadPublicKeyTesting(t)
	if err != nil {
		t.Errorf("Failed to load public key: %v", err)
		t.FailNow()
	}
	ecKey, _ := testutils.LoadEllipticPublicKey()

	roundOne := NewRound(roundInfoOne, pubKey, ecKey.PublicKey())

	// Construct a second eound
	roundInfoTwo := &mixmessages.RoundInfo{
		ID:       0,
		UpdateID: uint64(updateID + 1),
	}
	if err := testutils.SignRoundInfoRsa(roundInfoTwo, t); err != nil {
		t.Errorf("Failed to sign mock round info: %v", err)
	}

	roundTwo := NewRound(roundInfoTwo, pubKey, ecKey.PublicKey())

	_ = u.AddRound(roundOne)
	// Add second round twice (shouldn't duplicate)
	_ = u.AddRound(roundTwo)
	_ = u.AddRound(roundTwo)
	l := u.GetUpdates(2)
	if len(l) != 2 {
		t.Error("Something went wrong, didn't get all results")
	}
}
