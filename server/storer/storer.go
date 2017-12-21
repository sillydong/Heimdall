package storer

import (
	"context"
	"net"

	"github.com/sillydong/heimdall/common/proto"
	"github.com/sillydong/heimdall/server/model"
	"github.com/sillydong/heimdall/server/worker"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Storer struct {
	option *StorerOption
	m      *model.Model
	w      *worker.Worker
}

type StorerOption struct {
	Listen string `yaml:"listen"`
}

func NewStorer(option *StorerOption, m *model.Model, w *worker.Worker) *Storer {
	return &Storer{
		option: option,
		m:      m,
		w:      w,
	}
}

func (s *Storer) Run() {
	lis, err := net.Listen("tcp", s.option.Listen)
	if err != nil {
		logrus.Fatal(err)
	}
	server := grpc.NewServer()
	proto.RegisterHeimdallServiceServer(server, s)
	if err := server.Serve(lis); err != nil {
		logrus.Fatal(err)
	}
}

func (s *Storer) Regist(ctx context.Context, req *proto.RegistRequest) (*proto.RegistResponse, error) {
	return nil, nil
}

func (s *Storer) UnRegist(ctx context.Context, req *proto.UnRegistRequest) (*proto.UnRegistResponse, error) {
	return nil, nil
}

func (s *Storer) Log(ctxt context.Context, req *proto.LogRequest) (*proto.LogResponse, error) {
	return nil, nil
}

func (s *Storer) Heartbeat(ctx context.Context, req *proto.HeartbeatRequest) (*proto.HeartbeatResponse, error) {
	return nil, nil
}

func (s *Storer) Rule(ctx context.Context, req *proto.RuleRequest) (*proto.RuleResponse, error) {
	return nil, nil
}
