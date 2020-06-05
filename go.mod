module gitlab.com/elixxir/comms

go 1.13

require (
	github.com/golang/protobuf v1.4.2
	github.com/pkg/errors v0.9.1
	github.com/spf13/jwalterweatherman v1.1.0
	gitlab.com/elixxir/crypto v0.0.0-20200605212708-94b815c48c34
	gitlab.com/elixxir/primitives v0.0.0-20200605200100-cc0169bdd7e3
	golang.org/x/net v0.0.0-20200513185701-a91f0712d120
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/genproto v0.0.0-20200514193133-8feb7f20f2a2 // indirect
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.23.0
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.27.1
