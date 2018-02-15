// serverhandler.go - Interface for interaction between comms and server
//
// Copyright © 2018 Privategrity Corporation
// All rights reserved.

package mixserver

import "gitlab.com/privategrity/comms/mixmessages"

type ServerHandler interface {
	// Server Interface for starting New Rounds
	NewRound()
	// Server Interface for SetPublicKey
	SetPublicKey(RoundID string, PublicKey []byte)

	// Server Interface for the PrecompDecrypt Messages
	PrecompDecrypt(*mixmessages.PrecompDecryptMessage)
	// Server Interface for the PrecompEncrypt Messages
	PrecompEncrypt(*mixmessages.PrecompEncryptMessage)
	// Server Interface for the PrecompReveal Messages
	PrecompReveal(*mixmessages.PrecompRevealMessage)
	// Server Interface for the PrecompPermute Messages
	PrecompPermute(*mixmessages.PrecompPermuteMessage)
	// Server Interface for the PrecompShare Messages
	PrecompShare(*mixmessages.PrecompShareMessage)
	// Server Interface for the RealtimeDecrypt Messages
	RealtimeDecrypt(*mixmessages.RealtimeDecryptMessage)
	// Server Interface for the RealtimeEncrypt Messages
	RealtimeEncrypt(*mixmessages.RealtimeEncryptMessage)
	// Server Interface for the RealtimePermute Messages
	RealtimePermute(*mixmessages.RealtimePermuteMessage)

	// Server Interface for responding to RequestMessage
	RequestMessage(*mixmessages.RequestMessage)
	// Server interface for SendMessageToServer
	SendMessageToServer(*mixmessages.CmixMessage)
}
