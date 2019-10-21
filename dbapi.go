package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

// allSchedules는 DB에서 전체 스케쥴 정보를 가져오는 함수입니다.
func allSchedules(session *mgo.Session, userID string) ([]Schedule, error) {
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(*flagDBName).C(userID)
	var result []Schedule
	err := c.Find(bson.M{}).All(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}
