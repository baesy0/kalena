package main

import (
	"gopkg.in/mgo.v2"
)

// AddCalendar : DB에 일정을 추가
func AddCalendar(session *mgo.Session, cal Calendar) error {
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("kalena").C("calendar")
	err := c.Insert(cal)
	if err != nil {
		return err
	}
	return nil
}
