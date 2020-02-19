package main

import (
	"errors"
	"strconv"
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
	c := session.DB(*flagDBName).C(s.Collection + "." + s.Layer)
	err = c.Insert(s)
	if err != nil {
		return err
	}
	return nil
}

// RmSchedule 함수는 DB에서 id가 일치하는 Schedule을 삭제한다.
func RmSchedule(session *mgo.Session, Collection, Layer string, id bson.ObjectId) error {
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(*flagDBName).C(Collection + "." + Layer)
	err := c.RemoveId(id)
	if err != nil {
		return err
	}
	return nil
}

// allSchedules는 DB에서 전체 스케쥴 정보를 가져오는 함수입니다.
func allSchedules(session *mgo.Session, Collection, Layer string) ([]Schedule, error) {
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(*flagDBName).C(Collection + "." + Layer)
	var result []Schedule
	err := c.Find(bson.M{}).All(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// SearchMonth 함수는 Collection, Year, Month을 입력받아 start의 값이 일치하면 반환한다.
func SearchMonth(session *mgo.Session, Collection, Layer, year, month string) ([]Schedule, error) {
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(*flagDBName).C(Collection + "." + Layer)
	var results []Schedule

	y, err := strconv.Atoi(year)
	if err != nil {
		return nil, err
	}
	m, err := strconv.Atoi(month)
	if err != nil {
		return nil, err
	}
	start, err := BeginningOfMonth(y, m)
	if err != nil {
		return nil, err
	}
	end, err := EndOfMonth(y, m)
	if err != nil {
		return nil, err
	}
	s, err := TimeToNum(start.Format(time.RFC3339))
	if err != nil {
		return nil, err
	}
	e, err := TimeToNum(end.Format(time.RFC3339))
	if err != nil {
		return nil, err
	}

	query := []bson.M{}
	query = append(query, bson.M{"startnum": bson.M{"$gt": e}})
	query = append(query, bson.M{"endnum": bson.M{"$lt": s}})
	q := bson.M{"$nor": query}
	err = c.Find(q).All(&results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// GetLayers 함수는 DB Collection 에서 사용되는 모든 layer값을 반환한다.
func GetLayers(session *mgo.Session, Collection string) ([]Layer, error) {
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(*flagDBName).C(Collection + ".layers")
	var layers []Layer
	err := c.Find(bson.M{}).All(&layers)
	if err != nil {
		return layers, err
	}
	return layers, nil
}

// AddLayer 함수는 Collection 에 layer를 추가한다.
func AddLayer(session *mgo.Session, Collection, name, color string, order int) error {
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(*flagDBName).C(Collection + ".layers")
	num, err := c.Find(bson.M{"name": name}).Count()
	if err != nil {
		return err
	}
	if num > 0 {
		return errors.New(name + " layer가 존재합니다")
	}
	l := Layer{
		Name:   name,
		Order:  order,
		Hidden: false,
	}
	if regexWebColor.MatchString(color) {
		l.Color = color
	}
	err = c.Insert(l)
	if err != nil {
		return err
	}
	return nil
}

// RmLayer 는 이름이 name과 일치하는 layer를 삭제한다.
func RmLayer(session *mgo.Session, Collection, name string) error {
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(*flagDBName).C(Collection + ".layers")
	num, err := c.Find(bson.M{"name": name}).Count()
	if err != nil {
		return err
	}
	if num == 0 {
		return errors.New(name + " layer가 존재하지 않습니다")
	}
	err = c.Remove(bson.M{"name": name})
	if err != nil {
		return nil
	}
	return nil
}
