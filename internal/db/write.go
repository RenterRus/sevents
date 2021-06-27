package db

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

func (m *Mongo) SearchOpenEvent(typeEvent string) bool {
	query := bson.M{
		"type": typeEvent,
	}
	events := []DocStruct{}
	err := m.Collection.Find(query).All(&events)
	if err != nil {
		//Если не получилось с первого раза, попробуем еще несколько
		for i := 0; i < 4; i++ {
			time.Sleep(time.Second)
			err = m.Collection.Find(query).All(&events)
			if err == nil {
				break
			}
		}
		if err != nil {
			panic(err.Error())
		}
	}
	for _, v := range events {
		if v.State == 0 {
			return true
		}
	}

	return false
}

func (m *Mongo) WriteEvent(data DocStruct) string {
	if !m.SearchOpenEvent(data.TypeEvent) {
		err := m.Collection.Insert(data)
		if err != nil {
			//Если не получилось с первого раза, попробуем еще несколько
			for i := 0; i < 4; i++ {
				time.Sleep(time.Second)
				err = m.Collection.Insert(data)
				if err == nil {
					break
				}
			}
			if err != nil {
				return err.Error()
			}
		}
	}
	return "OK"
}
