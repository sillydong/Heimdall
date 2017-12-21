package model

import (
	"github.com/sillydong/heimdall/common/proto"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

//collection name for agent
func agentcollection() string {
	return "agent"
}

//create agent
func (m *Model) CreateAgent(agent proto.Agent) {
	if err := m.db.C(agentcollection()).Insert(agent); err != nil {
		logrus.Error(err)
	}
}

//update agent
func (m *Model) UpdateAgent(id string, agent proto.Agent) {
	if err := m.db.C(agentcollection()).UpdateId(id, agent); err != nil {
		logrus.Error(err)
	}
}

//find agent
func (m *Model) FindAgent(id string) proto.Agent {
	agent := proto.Agent{}
	if err := m.db.C(agentcollection()).FindId(id).One(&agent); err != nil {
		logrus.Error(err)
	}
	return agent
}

//find agents
func (m *Model) FindAgents(query bson.M, offset, limit int) []proto.Agent {
	var agents []proto.Agent
	if err := m.db.C(agentcollection()).Find(query).Limit(limit).Skip(offset).All(&agents); err != nil {
		logrus.Error(err)
	}
	return agents
}
