package main

import (
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/sillydong/heimdall/agent/mirror"
	"github.com/sillydong/heimdall/agent/parser"
	"github.com/sillydong/heimdall/agent/transfer"
	"github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v1"
)

func main() {
	runtimedir := filepath.Dir(os.Args[0])
	conf := Conf{}
	conffile, err := ioutil.ReadFile(runtimedir + string(os.PathSeparator) + "conf.yaml")
	if err != nil {
		logrus.Fatal(err)
	}
	err = yaml.Unmarshal(conffile, &conf)
	if err != nil {
		logrus.Fatal(err)
	}

	//init transfer
	t := transfer.NewTransfer(conf.Transfer)

	//init parser
	p := parser.NewParser(conf.Parser, t)
	p.Run()

	//init mirror
	m := mirror.NewMirror(conf.Mirror, p)
	m.Run()

	//wait for quit signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGQUIT)
	logrus.Infof("running pid: %v", os.Getpid())
	sig := <-c
	logrus.Infof("receive sig: %v", sig)

}

type Conf struct {
	Transfer *transfer.TransferOption `yaml:"transfer"`
	Mirror   *mirror.MirrorOption     `yaml:"mirror"`
	Parser   *parser.ParserOption     `yaml:"parse"`
}
