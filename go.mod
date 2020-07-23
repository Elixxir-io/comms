module gitlab.com/elixxir/comms

go 1.13

require (
	github.com/golang/protobuf v1.4.2
	github.com/pkg/errors v0.9.1
	github.com/spf13/jwalterweatherman v1.1.0
	gitlab.com/elixxir/crypto v0.0.0-20200707005343-97f868cbd930
	gitlab.com/elixxir/primitives v0.0.0-20200721220915-7250ade6f041
	gitlab.com/xx_network/collections/ring v0.0.1
	golang.org/x/net v0.0.0-20200707034311-ab3426394381
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/grpc v1.30.0
	google.golang.org/protobuf v1.25.0 // indirect
)

replace (
	gitlab.com/xx_network/collections/ring => gitlab.com/xx_network/collections/ring.git v0.0.1
	google.golang.org/grpc => github.com/grpc/grpc-go v1.27.1
)
