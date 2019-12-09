package main

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// AddSchedule 함수는 DB에 Schedule 을 추가한다.
func AddSchedule(session *mgo.Session, s Schedule) error {
	// DB에 추가되는 모든 시간은 UTC 시간을 가진다.
	err := s.ToUTC()
	if err != nil {
		return err
	}
	// 데이터를 넣기전 Excel시간도 셋팅한다.
	err = s.SetTimeNum()
	if err != nil {
		return err
	}
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(*flagDBName).C(s.Collection)
	err = c.Insert(s)
	if err != nil {
		return err
	}
	return nil
}

// RmSchedule 함수는 DB에서 id가 일치하는 Schedule을 삭제한다.
func RmSchedule(session *mgo.Session, Collection string, id bson.ObjectId) error {
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(*flagDBName).C(Collection)
	err := c.RemoveId(id)
	if err != nil {
		return err
	}
	return nil
}

// allSchedules는 DB에서 전체 스케쥴 정보를 가져오는 함수입니다.
func allSchedules(session *mgo.Session, Collection string) ([]Schedule, error) {
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(*flagDBName).C(Collection)
	var result []Schedule
	err := c.Find(bson.M{}).All(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// SearchMonth 함수는 Collection, Year, Month을 입력받아 start의 값이 일치하면 반환한다.
func SearchMonth(session *mgo.Session, Collection, year, month string) ([]Schedule, error) {
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(*flagDBName).C(Collection)
	var all []Schedule
	var results []Schedule
	err := c.Find(bson.M{}).All(&all)
	if err != nil {
		return nil, err
	}
	for _, s := range all {
		startTime, err := time.Parse("2006-01-02T15:04:05-07:00", s.Start)
		if err != nil {
			return []Schedule{}, err
		}
		endTime, err := time.Parse("2006-01-02T15:04:05-07:00", s.End)
		if err != nil {
			return []Schedule{}, err
		}
		// 현재는 한국 시간으로 한다. 사용자별로 시간대를 설정할 수 있는 기능은 나중에 만들겠다.
		monthStart, err := time.Parse("2006-01-02T15:04:05-07:00", fmt.Sprintf("%s-%s-01T00:00:00+09:00", year, month))
		if err != nil {
			return []Schedule{}, err
		}
		monthEnd, err := time.Parse("2006-01-02T15:04:05-07:00", fmt.Sprintf("%s-%s-30T23:59:59+09:00", year, month))
		if err != nil {
			return []Schedule{}, err
		}
		// endTime 이 monthStart 보다 작을 때 제외한다.
		if endTime.Before(monthStart) {
			continue
		}
		// startTime 이 monthEnd 보다 클 때 제외한다.
		if startTime.After(monthEnd) {
			continue
		}
		results = append(results, s)
	}
	return results, nil
}
