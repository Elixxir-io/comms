////////////////////////////////////////////////////////////////////////////////
// Copyright © 2022 xx foundation                                             //
//                                                                            //
// Use of this source code is governed by a license that can be found in the  //
// LICENSE file.                                                              //
////////////////////////////////////////////////////////////////////////////////

// Contains callback interface for registration functionality

package clientregistrar

import (
	"runtime/debug"

	jww "github.com/spf13/jwalterweatherman"
	pb "gitlab.com/elixxir/comms/mixmessages"
	"gitlab.com/xx_network/comms/connect"
	"gitlab.com/xx_network/comms/messages"
	"gitlab.com/xx_network/primitives/id"
)

// Registration object used to implement
// endpoints and top-level comms functionality
type Comms struct {
	*connect.ProtoComms
	handler Handler
	*pb.UnimplementedClientRegistrarServer
	*messages.UnimplementedGenericServer
}

// Starts a new server on the address:port specified by localServer
// and a callback interface for server operations
// with given path to public and private key for TLS connection
func StartClientRegistrarServer(id *id.ID, localServer string, handler Handler,
	certPEMblock, keyPEMblock []byte) *Comms {

	pc, err := connect.StartCommServer(id, localServer,
		certPEMblock, keyPEMblock, nil)
	if err != nil {
		jww.FATAL.Panicf("Unable to start comms server: %+v", err)
	}

	clientRegistrarServer := Comms{
		ProtoComms: pc,
		handler:    handler,
	}
	pb.RegisterClientRegistrarServer(clientRegistrarServer.GetServer(), &clientRegistrarServer)
	messages.RegisterGenericServer(clientRegistrarServer.GetServer(), &clientRegistrarServer)

	pc.ServeWithWeb()
	return &clientRegistrarServer
}

type Handler interface {
	RegisterUser(msg *pb.ClientRegistration) (confirmation *pb.SignedClientRegistrationConfirmations, err error)
}

type implementationFunctions struct {
	RegisterUser func(msg *pb.ClientRegistration) (confirmation *pb.SignedClientRegistrationConfirmations, err error)
}

// Implementation allows users of the client library to set the
// functions that implement the node functions
type Implementation struct {
	Functions implementationFunctions
}

// NewImplementation returns a Implementation struct with all of the
// function pointers returning nothing and printing an error.
func NewImplementation() *Implementation {
	um := "UNIMPLEMENTED FUNCTION!"
	warn := func(msg string) {
		jww.WARN.Printf(msg)
		jww.WARN.Printf("%s", debug.Stack())
	}
	return &Implementation{
		Functions: implementationFunctions{

			RegisterUser: func(msg *pb.ClientRegistration) (confirmation *pb.SignedClientRegistrationConfirmations, err error) {
				warn(um)
				return &pb.SignedClientRegistrationConfirmations{}, nil
			},
		},
	}
}

// Registers a user and returns a signed public key
func (s *Implementation) RegisterUser(msg *pb.ClientRegistration) (confirmation *pb.SignedClientRegistrationConfirmations, err error) {
	return s.Functions.RegisterUser(msg)
}
