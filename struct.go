package main

import (
	"errors"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Schedule 자료구조
type Schedule struct {
	ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Collection string        `json:"collection" bson:"collection"` // 사용자,장비명,회의실 이름이 될 수 있다
	Title      string        `json:"title" bson:"title"`           // 스케쥴의 title
	Start      string        `json:"start" bson:"start"`           // 스케쥴 시작 시간
	Startnum   int64         `json:"startnum" bson:"startnum"`     // 스케쥴 시작 날짜 int64
	End        string        `json:"end" bson:"end"`               // 스케쥴 끝나는 시간
	Endnum     int64         `json:"endnum" bson:"endnum"`         // 스케쥴 끝나는 날짜 int64
	Color      string        `json:"color" bson:"color"`           //#FF3366, #ff3366, #f36, #F36
	Layer      string        `json:"layer" bson:"layer"`           // 스케쥴이 속한 레이어의 이름
	Hidden     bool          `json:"hidden" bson:"hidden"`         // 스케쥴 숨김 속성
}

// CheckError 매소드는 Schedule 자료구조에 에러가 있는지 체크한다.
func (s Schedule) CheckError() error {
	if s.Collection == "" {
		return errors.New("Collection 이 빈 문자열 입니다")
	}
	if s.Title == "" {
		return errors.New("Title 이 빈 문자열 입니다")
	}
	if s.Layer == "" {
		return errors.New("Layer 이름이 빈 문자열 입니다")
	}

	if s.Start == "" {
		return errors.New("Start 시간이 빈 문자열 입니다")
	}
	if s.End == "" {
		return errors.New("End 시간이 빈 문자열 입니다")
	}
	if !regexRFC3339Time.MatchString(s.Start) {
		return errors.New("Start 시간이 2019-09-09T14:43:34+09:00 형식의 문자열이 아닙니다")
	}
	if !regexRFC3339Time.MatchString(s.End) {
		return errors.New("End 시간이 2019-09-09T14:43:34+09:00 형식의 문자열이 아닙니다")
	}
	startTime, err := time.Parse("2006-01-02T15:04:05-07:00", s.Start)
	if err != nil {
		return err
	}
	endTime, err := time.Parse("2006-01-02T15:04:05-07:00", s.End)
	if err != nil {
		return err
	}
	// end가 start 시간보다 큰지 체크하는 부분
	if !endTime.After(startTime) {
		return errors.New("끝시간이 시작시간보다 작습니다")
	}
	if s.Color != "" {
		if !regexWebColor.MatchString(s.Color) {
			return errors.New("#FF0011 형식의 문자열이 아닙니다")
		}
	}
	return nil
}
