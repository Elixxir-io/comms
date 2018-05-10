////////////////////////////////////////////////////////////////////////////////
// Copyright © 2018 Privategrity Corporation                                   /
//                                                                             /
// All rights reserved.                                                        /
////////////////////////////////////////////////////////////////////////////////

package node

import (
	jww "github.com/spf13/jwalterweatherman"
	pb "gitlab.com/privategrity/comms/mixmessages"
	"golang.org/x/net/context"
)

// Handle a Broadcasted Network Error event
func (s *server) NetworkError(ctx context.Context, err *pb.ErrorMessage) (
	*pb.ErrorAck, error) {
	msgLen := int32(len(err.Message))
	jww.ERROR.Println(err.Message)
	return &pb.ErrorAck{MsgLen: msgLen}, nil
}

// Handle a Broadcasted Ask Online event
func (s *server) AskOnline(ctx context.Context, msg *pb.Ping) (
	*pb.Pong, error) {
	return &pb.Pong{}, nil
}

// Handle a NewRound event
func (s *server) NewRound(ctx context.Context,
	msg *pb.InitRound) (*pb.InitRoundAck, error) {
	// Call the server handler to start a new round
	serverHandler.NewRound(msg.RoundID)
	return &pb.InitRoundAck{}, nil
}

// Handle CmixMessage from Client to Server
func (s *server) ClientSendMessageToServer(ctx context.Context,
	msg *pb.CmixMessage) (*pb.Ack, error) {
	// Call the server handler with the msg
	serverHandler.ReceiveMessageFromClient(msg)
	return &pb.Ack{}, nil
}

// Request a CmixMessage from the server for the given User
func (s *server) ClientPoll(ctx context.Context,
	msg *pb.ClientPollMessage) (*pb.CmixMessage, error) {
	return serverHandler.ClientPoll(msg), nil
}

func (s *server) RequestContactList(ctx context.Context,
	msg *pb.ContactPoll) (*pb.ContactMessage, error) {
	return serverHandler.RequestContactList(msg), nil
}

func (s *server) SetNick(ctx context.Context,
	msg *pb.Contact) (*pb.Ack, error) {
	serverHandler.SetNick(msg)
	return &pb.Ack{}, nil
}

// Handle a PrecompDecrypt event
func (s *server) PrecompDecrypt(ctx context.Context,
	msg *pb.PrecompDecryptMessage) (*pb.Ack, error) {
	// Call the server handler with the msg
	serverHandler.PrecompDecrypt(msg)
	return &pb.Ack{}, nil
}

// Handle a PrecompEncrypt event
func (s *server) PrecompEncrypt(ctx context.Context,
	msg *pb.PrecompEncryptMessage) (*pb.Ack, error) {
	// Call the server handler with the msg
	serverHandler.PrecompEncrypt(msg)
	return &pb.Ack{}, nil
}

// Handle a PrecompReveal event
func (s *server) PrecompReveal(ctx context.Context,
	msg *pb.PrecompRevealMessage) (*pb.Ack, error) {
	// Call the server handler with the msg
	serverHandler.PrecompReveal(msg)
	return &pb.Ack{}, nil
}

// Handle a PrecompPermute event
func (s *server) PrecompPermute(ctx context.Context,
	msg *pb.PrecompPermuteMessage) (*pb.Ack, error) {
	// Call the server handler with the msg
	serverHandler.PrecompPermute(msg)
	return &pb.Ack{}, nil
}

// Handle a PrecompShare event
func (s *server) PrecompShare(ctx context.Context,
	msg *pb.PrecompShareMessage) (*pb.Ack, error) {
	// Call the server handler with the msg
	serverHandler.PrecompShare(msg)
	return &pb.Ack{}, nil
}

// Handle a PrecompShareInit event
func (s *server) PrecompShareInit(ctx context.Context,
	msg *pb.PrecompShareInitMessage) (*pb.Ack, error) {
	// Call the server handler with the msg
	serverHandler.PrecompShareInit(msg)
	return &pb.Ack{}, nil
}

// Handle a PrecompShareCompare event
func (s *server) PrecompShareCompare(ctx context.Context,
	msg *pb.PrecompShareCompareMessage) (*pb.Ack, error) {
	// Call the server handler with the msg
	serverHandler.PrecompShareCompare(msg)
	return &pb.Ack{}, nil
}

// Handle a PrecompShareConfirm event
func (s *server) PrecompShareConfirm(ctx context.Context,
	msg *pb.PrecompShareConfirmMessage) (*pb.Ack, error) {
	// Call the server handler with the msg
	serverHandler.PrecompShareConfirm(msg)
	return &pb.Ack{}, nil
}

// Handle a RealtimeDecrypt event
func (s *server) RealtimeDecrypt(ctx context.Context,
	msg *pb.RealtimeDecryptMessage) (*pb.Ack, error) {
	// Call the server handler with the msg
	serverHandler.RealtimeDecrypt(msg)
	return &pb.Ack{}, nil
}

// Handle a RealtimeDecrypt event
func (s *server) RealtimeEncrypt(ctx context.Context,
	msg *pb.RealtimeEncryptMessage) (*pb.Ack, error) {
	// Call the server handler with the msg
	serverHandler.RealtimeEncrypt(msg)
	return &pb.Ack{}, nil
}

// Handle a RealtimePermute event
func (s *server) RealtimePermute(ctx context.Context,
	msg *pb.RealtimePermuteMessage) (*pb.Ack, error) {
	// Call the server handler with the msg
	serverHandler.RealtimePermute(msg)
	return &pb.Ack{}, nil
}

// Handle a SetPublicKey event
func (s *server) SetPublicKey(ctx context.Context,
	msg *pb.PublicKeyMessage) (*pb.Ack, error) {
	serverHandler.SetPublicKey(msg.RoundID, msg.PublicKey)
	return &pb.Ack{}, nil
}
