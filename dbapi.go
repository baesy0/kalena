package main

import (
	"gopkg.in/mgo.v2"
)

// AddSchedule : DB에 일정을 추가
func AddSchedule(session *mgo.Session, s Schedule) error {
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("kalena").C(*flagUser)
	err := c.Insert(s)
	if err != nil {
		return err
	}
	return nil
}
