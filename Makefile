GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
AGENT_NAME=agent
SERVER_NAME=server
GOARCH=amd64
GOOS=darwin

all: proto clean build-agent build-server

proto:
	protoc -I. --go_out=plugins=grpc:. ./common/proto/heimdall.proto

build-agent:
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GOBUILD) -o ./bin/$(AGENT_NAME) -v ./agent

clean-agent:
	rm -f ./bin/$(AGENT_NAME)

build-server:
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GOBUILD) -o ./bin/$(SERVER_NAME) -v ./server

clean-server:
	rm -f ./bin/$(SERVER_NAME)

clean: clean-agent clean-server
