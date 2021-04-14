module gitlab.com/elixxir/comms

go 1.13

require (
	github.com/aws/aws-lambda-go v1.8.1 // indirect
	github.com/golang-collections/collections v0.0.0-20130729185459-604e922904d3
	github.com/golang/protobuf v1.4.2
	github.com/katzenpost/core v0.0.14
	github.com/nyaruka/phonenumbers v1.0.60 // indirect
	github.com/pkg/errors v0.9.1
	github.com/spf13/jwalterweatherman v1.1.0
	gitlab.com/elixxir/crypto v0.0.6
	gitlab.com/elixxir/primitives v0.0.3-0.20210409190923-7bf3cd8d97e7
	gitlab.com/xx_network/comms v0.0.4-0.20210414191603-0904bc6eeda2
	gitlab.com/xx_network/crypto v0.0.5-0.20210413200952-56bd15ec9d99
	gitlab.com/xx_network/primitives v0.0.4-0.20210412170941-7ef69bce5a5c
	gitlab.com/xx_network/ring v0.0.2
	golang.org/x/lint v0.0.0-20190313153728-d0100b6bd8b3 // indirect
	golang.org/x/net v0.0.0-20201029221708-28c70e62bb1d
	golang.org/x/text v0.3.4 // indirect
	golang.org/x/tools v0.0.0-20190524140312-2c0ae7006135 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/grpc v1.31.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
	honnef.co/go/tools v0.0.0-20190523083050-ea95bdfd59fc // indirect
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.27.1
