package db

import "gopkg.in/mgo.v2/bson"

func (m *Mongo) SearchOpenEvent(typeEvent string) bool{
	query := bson.M{
		"type" : typeEvent,
	}
	events := []DocStruct{}
	m.Collection.Find(query).All(&events)

	for _, v := range events{
		if v.State == 0{
			return true
		}
	}

	return false
}

func (m *Mongo) WriteEvent(data DocStruct) string{
	if !m.SearchOpenEvent(data.TypeEvent){
		m.Collection.Insert(data)
	}
	return "OK"
}
