////////////////////////////////////////////////////////////////////////////////
// Copyright © 2018 Privategrity Corporation                                   /
//                                                                             /
// All rights reserved.                                                        /
////////////////////////////////////////////////////////////////////////////////

// Contains callback interface for gateway functionality

package gateway

import (
	jww "github.com/spf13/jwalterweatherman"
	pb "gitlab.com/elixxir/comms/mixmessages"
	"gitlab.com/elixxir/primitives/id"
	"runtime/debug"
)

// Handler interface for the Gateway
type Handler interface {
	// Return any MessageIDs in the buffer for this UserID
	CheckMessages(userID *id.User, messageID string) ([]string, bool)
	// Returns the message matching the given parameters to the client
	GetMessage(userID *id.User, msgID string) (*pb.CmixBatch, bool)
	// Upload a message to the cMix Gateway
	PutMessage(message *pb.CmixBatch) bool
	// Receives a batch of messages from a server
	ReceiveBatch(messages *pb.OutputMessages)

	// Pass-through for Registration Nonce Communication
	RequestNonce(message *pb.RequestNonceMessage) (*pb.NonceMessage, error)
	// Pass-through for Registration Nonce Confirmation
	ConfirmNonce(message *pb.ConfirmNonceMessage) (*pb.
		RegistrationConfirmation, error)
}

// Handler implementation for the Gateway
type implementationFunctions struct {
	// Return any MessageIDs in the buffer for this UserID
	CheckMessages func(userID *id.User, messageID string) ([]string, bool)
	// Returns the message matching the given parameters to the client
	GetMessage func(userID *id.User, msgID string) (*pb.CmixBatch, bool)
	// Upload a message to the cMix Gateway
	PutMessage func(message *pb.CmixBatch) bool
	// Receives a batch of messages from a server
	ReceiveBatch func(messages *pb.OutputMessages)

	// Pass-through for Registration Nonce Communication
	RequestNonce func(message *pb.RequestNonceMessage) (*pb.NonceMessage, error)
	// Pass-through for Registration Nonce Confirmation
	ConfirmNonce func(message *pb.ConfirmNonceMessage) (*pb.
			RegistrationConfirmation, error)
}

// Implementation allows users of the client library to set the
// functions that implement the node functions
type Implementation struct {
	Functions implementationFunctions
}

// Creates and returns a new Handler interface
func NewImplementation() Handler {
	um := "UNIMPLEMENTED FUNCTION!"
	warn := func(msg string) {
		jww.WARN.Printf(msg)
		jww.WARN.Printf("%v", debug.Stack())
	}
	return Handler(&Implementation{
		Functions: implementationFunctions{
			CheckMessages: func(userID *id.User, messageID string) ([]string,
				bool) {
				warn(um)
				return nil, false
			},
			GetMessage: func(userID *id.User, msgID string) (*pb.CmixBatch,
				bool) {
				warn(um)
				return &pb.CmixBatch{}, false
			},
			PutMessage: func(message *pb.CmixBatch) bool {
				warn(um)
				return false
			},
			ReceiveBatch: func(messages *pb.OutputMessages) { warn(um) },

			RequestNonce: func(message *pb.RequestNonceMessage) (*pb.
				NonceMessage, error) {
				warn(um)
				return new(pb.NonceMessage), nil
			},
			ConfirmNonce: func(message *pb.ConfirmNonceMessage) (*pb.
				RegistrationConfirmation, error) {
				warn(um)
				return new(pb.RegistrationConfirmation), nil
			},
		},
	})
}

// Return any MessageIDs in the buffer for this UserID
func (s *Implementation) CheckMessages(userID *id.User, messageID string) (
	[]string, bool) {
	return s.Functions.CheckMessages(userID, messageID)
}

// Returns the message matching the given parameters to the client
func (s *Implementation) GetMessage(userID *id.User, msgID string) (
	*pb.CmixBatch, bool) {
	return s.Functions.GetMessage(userID, msgID)
}

// Upload a message to the cMix Gateway
func (s *Implementation) PutMessage(message *pb.CmixBatch) bool {
	return s.Functions.PutMessage(message)
}

// Receives a batch of messages from a server
func (s *Implementation) ReceiveBatch(messages *pb.OutputMessages) {
	s.Functions.ReceiveBatch(messages)
}

// Pass-through for Registration Nonce Communication
func (s *Implementation) RequestNonce(message *pb.RequestNonceMessage) (
	*pb.NonceMessage, error) {
	return s.Functions.RequestNonce(message)
}

// Pass-through for Registration Nonce Confirmation
func (s *Implementation) ConfirmNonce(message *pb.ConfirmNonceMessage) (*pb.
	RegistrationConfirmation, error) {
	return s.Functions.ConfirmNonce(message)
}
