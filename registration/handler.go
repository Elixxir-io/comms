////////////////////////////////////////////////////////////////////////////////
// Copyright © 2018 Privategrity Corporation                                   /
//                                                                             /
// All rights reserved.                                                        /
////////////////////////////////////////////////////////////////////////////////

// Contains callback interface for registration functionality

package registration

import (
	jww "github.com/spf13/jwalterweatherman"
	"runtime/debug"
)

type Handler interface {
	RegisterUser(registrationCode string, Y, P, Q, G []byte) (hash,
		R, S []byte, err error)
	RegisterNode(ID []byte,
		NodeCSR, GatewayTLSCert, RegistrationCode, Addr string) error
}

type implementationFunctions struct {
	RegisterUser func(registrationCode string, Y, P, Q, G []byte) (hash,
		R, S []byte, err error)
	RegisterNode func(ID []byte,
		NodeCSR, GatewayTLSCert, RegistrationCode, Addr string) error
}

// Implementation allows users of the client library to set the
// functions that implement the node functions
type Implementation struct {
	Functions implementationFunctions
}

// NewImplementation returns a Implementation struct with all of the
// function pointers returning nothing and printing an error.
func NewImplementation() Handler {
	um := "UNIMPLEMENTED FUNCTION!"
	warn := func(msg string) {
		jww.WARN.Printf(msg)
		jww.WARN.Printf("%v", debug.Stack())
	}
	return Handler(&Implementation{
		Functions: implementationFunctions{
			RegisterUser: func(registrationCode string,
				Y, P, Q, G []byte) (hash, R, S []byte, err error) {
				warn(um)
				return nil, nil, nil, nil
			},
			RegisterNode: func(ID []byte,
				NodeCsr, GatewayTLSCert, RegistrationCode, Addr string) error {
				warn(um)
				return nil
			},
		},
	})
}

// Registers a user and returns a signed public key
func (s *Implementation) RegisterUser(registrationCode string,
	Y, P, Q, G []byte) (hash, R, S []byte, err error) {
	return s.Functions.RegisterUser(registrationCode, Y, P, Q, G)
}

func (s *Implementation) RegisterNode(ID []byte,
	NodeCsr, GatewayTLSCert, RegistrationCode, Addr string) error {
	s.Functions.RegisterNode(ID, NodeCsr, GatewayTLSCert, RegistrationCode, Addr)
	return nil
}
