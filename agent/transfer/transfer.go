package transfer

import (
	"github.com/sillydong/heimdall/common/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

//host address of server
type TransferOption struct {
	Host string `yaml:"host"`
}

//create transfer instance
func NewTransfer(option *TransferOption) proto.HeimdallServiceClient {
	conn, err := grpc.Dial(option.Host)
	if err != nil {
		logrus.Fatal(err)
	}

	return proto.NewHeimdallServiceClient(conn)
}
