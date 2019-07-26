////////////////////////////////////////////////////////////////////////////////
// Copyright © 2018 Privategrity Corporation                                   /
//                                                                             /
// All rights reserved.                                                        /
////////////////////////////////////////////////////////////////////////////////

// Contains server gRPC endpoints

package node

// TODO: A lot of message types from gRPC are passed through, and a number of
//       errors that can occur are not accounted for.

import (
	jww "github.com/spf13/jwalterweatherman"
	pb "gitlab.com/elixxir/comms/mixmessages"
	"golang.org/x/net/context"
)

// Handle a Broadcasted Ask Online event
func (s *NodeComms) AskOnline(ctx context.Context, msg *pb.Ping) (
	*pb.Ack, error) {
	return &pb.Ack{}, nil
}

// Handle a broadcasted DownloadTopology event
func (s *NodeComms) DownloadTopology(ctx context.Context,
	msg *pb.SignedMessage) (*pb.Ack, error) {
	sender := msg.ID
	pubKey := s.ConnectionManager.GetConnectionInfo(sender).RsaPublicKey
	// VERIFY AND UNWRAP
	verified := pb.NodeTopology{}
	err := s.ConnectionManager.VerifySignature(msg, &verified, pubKey)
	if err != nil {
		jww.ERROR.Printf("Failed to verify message contents: %+v", err)
		return nil, err
	}

	s.handler.DownloadTopology(&verified)
	return &pb.Ack{}, nil
}

// Handle a NewRound event
func (s *NodeComms) CreateNewRound(ctx context.Context,
	msg *pb.RoundInfo) (*pb.Ack, error) {
	// Call the server handler to start a new round
	return &pb.Ack{}, s.handler.CreateNewRound(msg)
}

// PostNewBatch polls the first node and sends a batch when it is ready
func (s *NodeComms) PostNewBatch(ctx context.Context, msg *pb.Batch) (*pb.Ack, error) {
	// Call the server handler to post a new batch
	err := s.handler.PostNewBatch(msg)

	return &pb.Ack{}, err
}

// Handle a Phase event
func (s *NodeComms) PostPhase(ctx context.Context, msg *pb.Batch) (*pb.Ack,
	error) {
	// Call the server handler with the msg
	s.handler.PostPhase(msg)
	return &pb.Ack{}, nil
}

// Handle a phase event using a stream server
func (s *NodeComms) StreamPostPhase(server pb.Node_StreamPostPhaseServer) error {
	return s.handler.StreamPostPhase(server)
}

// Handle a PostRoundPublicKey message
func (s *NodeComms) PostRoundPublicKey(ctx context.Context,
	msg *pb.RoundPublicKey) (*pb.Ack, error) {
	// Call the server handler that receives the key share
	s.handler.PostRoundPublicKey(msg)
	return &pb.Ack{}, nil
}

// GetBufferInfo returns buffer size (number of completed precomputations)
func (s *NodeComms) GetRoundBufferInfo(ctx context.Context,
	msg *pb.RoundBufferInfo) (
	*pb.RoundBufferInfo, error) {
	bufSize, err := s.handler.GetRoundBufferInfo()
	if bufSize < 0 {
		bufSize = 0
	}
	size := uint32(bufSize)
	return &pb.RoundBufferInfo{RoundBufferSize: size}, err
}

// Handles Registration Nonce Communication
func (s *NodeComms) RequestNonce(ctx context.Context,
	msg *pb.NonceRequest) (*pb.Nonce, error) {
	pk := msg.GetClient()
	sig := msg.GetClientSignedByServer()

	// Obtain the nonce by passing to server
	nonce, err := s.handler.RequestNonce(msg.GetSalt(),
		pk.GetY(), pk.GetP(), pk.GetQ(),
		pk.GetG(), sig.GetHash(), sig.GetR(), sig.GetS())

	// Obtain the error message, if any
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}

	// Return the NonceMessage
	return &pb.Nonce{
		Nonce: nonce,
		Error: errMsg,
	}, err
}

// Handles Registration Nonce Confirmation
func (s *NodeComms) ConfirmRegistration(ctx context.Context,
	msg *pb.DSASignature) (*pb.RegistrationConfirmation, error) {

	// Obtain signed client public key by passing to server
	hash, R, S, Y, P, Q, G, err := s.handler.ConfirmRegistration(
		msg.GetHash(),
		msg.GetR(), msg.GetS())

	// Obtain the error message, if any
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}

	// Return the RegistrationConfirmation
	return &pb.RegistrationConfirmation{
		ClientSignedByServer: &pb.DSASignature{
			Hash: hash,
			R:    R,
			S:    S,
		},
		Server: &pb.DSAPublicKey{
			Y: Y,
			P: P,
			Q: Q,
			G: G,
		},
		Error: errMsg,
	}, err
}

// PostPrecompResult sends final Message and AD precomputations.
func (s *NodeComms) PostPrecompResult(ctx context.Context,
	msg *pb.Batch) (*pb.Ack, error) {
	// Call the server handler to start a new round
	err := s.handler.PostPrecompResult(msg.GetRound().GetID(),
		msg.GetSlots())
	return &pb.Ack{}, err
}

// FinishRealtime broadcasts to all nodes when the realtime is completed
func (s *NodeComms) FinishRealtime(ctx context.Context, msg *pb.RoundInfo) (*pb.Ack, error) {
	// Call the server handler to finish realtime
	err := s.handler.FinishRealtime(msg)

	return &pb.Ack{}, err
}

// GetCompletedBatch should return a completed batch that the calling gateway
// hasn't gotten before
func (s *NodeComms) GetCompletedBatch(ctx context.Context,
	msg *pb.Ping) (*pb.Batch, error) {
	return s.handler.GetCompletedBatch()
}

func (s *NodeComms) GetMeasure(ctx context.Context, msg *pb.RoundInfo) (*pb.RoundMetrics, error) {
	rm, err := s.handler.GetMeasure(msg)
	return rm, err
}
