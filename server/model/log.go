package model

import (
	"github.com/sillydong/heimdall/common/proto"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

//collection name for log, rotate by day
func logcollection() string {
	return "log_" + timeCache.Date()
}

//create log
func (m *Model) CreateLog(log proto.Log) {
	if err := m.db.C(logcollection()).Insert(log); err != nil {
		logrus.Error(err)
	}
}

//find logs
func (m *Model) FindLogs(query bson.M, offset, limit int) []proto.Log {
	var logs []proto.Log
	if err := m.db.C(logcollection()).Find(query).Limit(limit).Skip(offset).All(&logs); err != nil {
		logrus.Error(err)
	}
	return logs
}
