package db

import (
	"context"
	"gopkg.in/mgo.v2"
	"time"
)

type Mongo struct {
	Collection *mgo.Collection
	Ctx        context.Context
	Session    *mgo.Session
	Addr       string
	DBName     string
	CollName   string
}

type DocStruct struct {
	Id        string `bson:"_id"`
	TypeEvent string `bson:"type"`
	State     int    `bson:"state"`
	Start     string `bson:"started_at"`
	Finish    string `bson:"finished_at"`
}

func GetMongoClient(mongo, dbname, collname string) *Mongo {
	var err error
	m := new(Mongo)

	m.Session, err = mgo.Dial(mongo)
	if err != nil {
		//Если не получилось с первого раза, попробуем еще несколько
		for i := 0; i < 4; i++ {
			time.Sleep(time.Second)
			m.Session, err = mgo.Dial(mongo)
			if err == nil {
				break
			}
		}
		if err != nil {
			panic(err)
		}
	}
	m.Collection = m.Session.DB(dbname).C(collname)

	return m
}
