module gitlab.com/elixxir/comms

go 1.13

require (
	github.com/golang-collections/collections v0.0.0-20130729185459-604e922904d3
	github.com/golang/protobuf v1.4.3
	github.com/google/go-cmp v0.5.2 // indirect
	github.com/pkg/errors v0.9.1
	github.com/spf13/jwalterweatherman v1.1.0
	gitlab.com/elixxir/crypto v0.0.7-0.20210326171146-c137bd7b0c6e
	gitlab.com/elixxir/primitives v0.0.3-0.20210326022836-1143187bd2fe
	gitlab.com/xx_network/comms v0.0.4-0.20210326005744-5e73cbf0f525
	gitlab.com/xx_network/crypto v0.0.5-0.20210319231335-249c6b1aa323
	gitlab.com/xx_network/primitives v0.0.4-0.20210309173740-eb8cd411334a
	gitlab.com/xx_network/ring v0.0.2
	golang.org/x/net v0.0.0-20201029221708-28c70e62bb1d
	golang.org/x/text v0.3.4 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/genproto v0.0.0-20201030142918-24207fddd1c3 // indirect
	google.golang.org/grpc v1.33.1
	google.golang.org/protobuf v1.25.0
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.27.1
