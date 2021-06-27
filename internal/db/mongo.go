package db

import (
	"context"
	"gopkg.in/mgo.v2"
)

var MongoParam Mongo

type Mongo struct {
	Collection *mgo.Collection
	Ctx context.Context
	Session *mgo.Session
	Addr string
	DBName string
	CollName string
}

type DocStruct struct{
	Id string `bson:"_id"`
	TypeEvent string `bson:"type"`
	State int `bson:"state"`
	Start string `bson:"started_at"`
	Finish string `bson:"finished_at"`
}

func (m *Mongo) GetClient() {
	var err error
	m.Session, err = mgo.Dial(m.Addr)
	if err != nil {
		panic(err)
	}
	m.Collection = m.Session.DB(m.DBName).C(m.CollName)
}
