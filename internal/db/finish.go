package db

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
)

func (m *Mongo) SearchEvent(typeEvent string) bool{
	query := bson.M{
		"type" : typeEvent,
	}
	events := []DocStruct{}
	m.Collection.Find(query).All(&events)

	if len(events) == 0{
		return false
	}

	for _, v := range events{
		if v.State == 0{
			return true
		}
	}

	return false
}

func (m *Mongo) FinishEvent(typeEvent string) bool{
	if !m.SearchEvent(typeEvent){
		return false
	} else {
		_, err := m.Collection.UpdateAll(bson.M{"type": typeEvent}, bson.M{"$set": bson.M{"state": 1, "finished_at":strconv.Itoa(int(time.Now().Unix()))}})
		if err != nil{
			fmt.Println(err)
		}
	}
	return true
}