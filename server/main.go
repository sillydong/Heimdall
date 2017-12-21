package main

import (
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/sillydong/heimdall/server/model"
	"github.com/sillydong/heimdall/server/storer"
	"github.com/sillydong/heimdall/server/web"
	"github.com/sillydong/heimdall/server/worker"
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

	//model
	m := model.NewModel(conf.Model)

	//init worker
	w := worker.NewWorker()

	//init storer
	store := storer.NewStorer(conf.Storer, m, w)
	store.Run()

	//init web
	wapi := web.NewWeb(m)
	wapi.Run()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGQUIT)
	logrus.Infof("running pid: %v", os.Getpid())
	sig := <-c
	logrus.Infof("receive sig: %v", sig)
	m.Close()
}

type Conf struct {
	Model  *model.ModelOption   `yaml:"model"`
	Storer *storer.StorerOption `yaml:"storer"`
}
