package model

import (
	"github.com/sillydong/heimdall/common/proto"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

//collection name for heartbeat, rotate by day
func heartbeatcollection() string {
	return "heartbeat_" + timeCache.Date()
}

//create heartbeat
func (m *Model) CreateHeartbeat(heartbeat proto.Heartbeat) {
	if err := m.db.C(heartbeatcollection()).Insert(heartbeat); err != nil {
		logrus.Error(err)
	}
}

//find heartbeats
func (m *Model) FindHeartbeats(query bson.M, offset, limit int) []proto.Heartbeat {
	var heartbeats []proto.Heartbeat
	if err := m.db.C(heartbeatcollection()).Find(query).Limit(limit).Skip(offset).All(&heartbeats); err != nil {
		logrus.Error(err)
	}
	return heartbeats
}
