package main

import (
	"gopkg.in/mgo.v2"
)

// AddSchedule 함수는 DB에 Schedule 을 추가한다.
func AddSchedule(session *mgo.Session, s Schedule) error {
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(*flagDBName).C(*flagUser)
	err := c.Insert(s)
	if err != nil {
		return err
	}
	return nil
}
