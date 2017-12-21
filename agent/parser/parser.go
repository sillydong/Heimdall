package parser

import (
	"bytes"
	"context"
	"net/http"
	"time"

	"github.com/sillydong/heimdall/common/proto"
	"github.com/sirupsen/logrus"
)

type Parser struct {
	option *ParserOption
	system *System
	t      proto.HeimdallServiceClient

	requestChannels []chan *WrapRequest
	handlers        []*handler
}

type System struct {
	Id          string
	Hostname    string
	Version     string
	Hostaddress string
	Deviceinfo  string
}

type ParserOption struct {
	Shard  int `yaml:"shard"`
	Buffer int `yaml:"buffer"`
}

func NewParser(option *ParserOption, t proto.HeimdallServiceClient) *Parser {
	p := &Parser{
		t:      t,
		option: option,
	}
	p.system = p.systemInfo()
	return p
}

func (p *Parser) Run() {
	for shard := 0; shard < p.option.Shard; shard++ {
		p.requestChannels[shard] = make(chan *WrapRequest, p.option.Buffer)
		go p.parseWorker(shard)
	}
	p.t.Regist(context.Background(), &proto.RegistRequest{
		Agent: &proto.Agent{
			Id:          p.system.Id,
			Hostname:    p.system.Hostname,
			Version:     p.system.Version,
			Hostaddress: p.system.Hostaddress,
			Deviceinfo:  p.system.Deviceinfo,
		},
	})
}

func (p *Parser) Parse(req *http.Request) {
	shard := time.Now().Nanosecond() % p.option.Shard
	wrapreq := &WrapRequest{
		Request: req,
		Start:   time.Now().UnixNano(),
		Stop:    0,
		Match:   false,
		RuleID:  "",
	}
	p.requestChannels[shard] <- wrapreq
}

func (p *Parser) parseWorker(shard int) {
	for {
		request := <-p.requestChannels[shard]
		if ruleid, match := p.handlers[shard].Handle(request.Request); match > 0 {
			request.Stop = time.Now().UnixNano()
			request.Match = true
			request.RuleID = ruleid
			var body bytes.Buffer
			err := request.Request.Write(&body)
			if err != nil {
				logrus.Error(err)
			}
			p.t.Log(context.Background(), &proto.LogRequest{
				Log: &proto.Log{
					AgentId:  p.system.Id,
					Hostname: p.system.Hostname,
					Version:  p.system.Version,
					Start:    request.Start,
					Stop:     request.Stop,
					Match:    match,
					Ruleid:   ruleid,
					Request:  body.Bytes(),
				},
			})
		}
	}
}

func (p *Parser) systemInfo() *System {
	s := &System{}
	return s
}

type WrapRequest struct {
	Request *http.Request
	Start   int64
	Stop    int64
	Match   bool
	RuleID  string
}
