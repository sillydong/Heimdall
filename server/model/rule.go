package model

import (
	"github.com/sillydong/heimdall/common/proto"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

//collection name for rule
func rulecollection() string {
	return "rule"
}

//create rule
func (m *Model) CreateRule(rule proto.Rule) {
	if err := m.db.C(rulecollection()).Insert(rule); err != nil {
		logrus.Error(err)
	}
}

//update rule
func (m *Model) UpdateRule(id string, rule proto.Rule) {
	if err := m.db.C(rulecollection()).UpdateId(id, rule); err != nil {
		logrus.Error(err)
	}
}

//find rules
func (m *Model) FindRules(query bson.M, offset, limit int) []proto.Rule {
	var rules []proto.Rule
	if err := m.db.C(rulecollection()).Find(query).Limit(limit).Skip(offset).All(&rules); err != nil {
		logrus.Error(err)
	}
	return rules
}
