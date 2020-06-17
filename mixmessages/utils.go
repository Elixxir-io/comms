///////////////////////////////////////////////////////////////////////////////
// Copyright © 2020 xx network SEZC                                          //
//                                                                           //
// Use of this source code is governed by a license that can be found in the //
// LICENSE file                                                              //
///////////////////////////////////////////////////////////////////////////////

// Contains utils functions for comms

package mixmessages

import jww "github.com/spf13/jwalterweatherman"

func DebugMode() {
	jww.SetLogThreshold(jww.LevelDebug)
	jww.SetStdoutThreshold(jww.LevelDebug)
}

func TraceMode() {
	jww.SetLogThreshold(jww.LevelTrace)
	jww.SetStdoutThreshold(jww.LevelTrace)
}
