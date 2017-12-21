package model

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

//collection name for alert
func alertcollection() string {
	return "alert"
}

//alert data structure
type Alert struct {
	Id string
}

//create alert
func (m *Model) CreateAlert(alert Alert) {
	if err := m.db.C(alertcollection()).Insert(alert); err != nil {
		logrus.Error(err)
	}
}

//find alert
func (m *Model) FindAlerts(query bson.M, offset, limit int) []Alert {
	var alerts []Alert
	if err := m.db.C(alertcollection()).Find(query).Limit(limit).Skip(offset).All(&alerts); err != nil {
		logrus.Error(err)
	}
	return alerts
}
