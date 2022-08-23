///////////////////////////////////////////////////////////////////////////////
// Copyright © 2020 xx network SEZC                                          //
//                                                                           //
// Use of this source code is governed by a license that can be found in the //
// LICENSE file                                                              //
///////////////////////////////////////////////////////////////////////////////

package testkeys

import (
	_ "embed"
)

// This file embeds the node and gateway certs and keys in this directory so
// that they do not need to be read from file at runtime.
//
// This avoids using the [os] package so that these files can be used by tests
// when compiling for WebAssembly.

var (
	//go:embed cmix.rip.crt
	CmixCrt []byte

	//go:embed cmix.rip.key
	CmixKey []byte

	//go:embed gateway.cmix.rip.crt
	GatewayCrt []byte

	//go:embed gateway.cmix.rip.key
	GatewayKey []byte
)
