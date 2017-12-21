package model

import (
	"github.com/sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"
)

//model contains db connection
type Model struct {
	option *ModelOption
	db     *mgo.Database
}

//option for model
type ModelOption struct {
	Host string `yaml:"host"`
	Db   string `yaml:"db"`
}

//create model instance
func NewModel(option *ModelOption) *Model {
	m := &Model{}
	m.Init(option)
	return m
}

//init mongodb connection
func (m *Model) Init(option *ModelOption) {
	m.option = option
	session, err := mgo.Dial(m.option.Host)
	if err != nil {
		logrus.Error(err)
	}
	session.SetMode(mgo.Monotonic, true)
	m.db = session.DB(m.option.Db)
	m.db.C("agent").EnsureIndex(mgo.Index{
		Key:        []string{"hostname"},
		Unique:     true,
		DropDups:   true,
		Background: true,
	})
}

//close mongodb connection
func (m *Model) Close() {
	m.db.Session.Close()
}
