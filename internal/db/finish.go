package db

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func (m *Mongo) SearchEvent(typeEvent string) bool {
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

	if len(events) == 0 {
		return false
	}

	for _, v := range events {
		if v.State == 0 {
			return true
		}
	}

	return false
}

func (m *Mongo) FinishEvent(typeEvent string) bool {
	if !m.SearchEvent(typeEvent) {
		return false
	} else {
		err := m.Collection.Update(bson.M{"type": typeEvent, "state": 0}, bson.M{"$set": bson.M{"state": 1,
			"finished_at": time.Now().Format("01-02-2006 15:04:05.000000")}})
		if err != nil {
			//Если не получилось с первого раза, попробуем еще несколько
			for i := 0; i < 4; i++ {
				time.Sleep(time.Second)
				err = m.Collection.Update(bson.M{"type": typeEvent, "state": 0}, bson.M{"$set": bson.M{"state": 1,
					"finished_at": time.Now().Format("01-02-2006 15:04:05.000000")}})
				if err == nil {
					break
				}
			}
			if err != nil {
				fmt.Println(err)
				panic(err)
			}
		}
	}
	return true
}
