package main

import (
	"regexp"
	"time"
)

var (
	regexRFC3339Time = regexp.MustCompile(`^\d{4}-(0?[1-9]|1[012])-(0?[1-9]|[12][0-9]|3[01])T\d{2}:\d{2}:\d{2}[-+]\d{2}:\d{2}$`) //2019-09-09T02:46:52+09:00
	regexWebColor    = regexp.MustCompile(`^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$`)                                                  //#FF0011, #ff0011, #F01, #f01
)

// checkTime 함수는 시작시간과 끝 시간을 받아서 끝시간이 시작시간보다 크면 true를 반환한다.
func checkTime(start, end time.Time) bool {
	return end.After(start)
}
