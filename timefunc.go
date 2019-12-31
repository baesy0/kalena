package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/now"
)

// TimeToNum 함수는 시간문자를 받아서 1899년 12월 30일부터 몇번째 날인지 반환한다.
// Excel 에서 날수를 계산하는 방식과 같다.
// 전설의 스토리: https://social.msdn.microsoft.com/Forums/office/en-US/f1eef5fe-ef5e-4ab6-9d92-0998d3fa6e14/what-is-story-behind-december-30-1899-as-base-date?forum=accessdev
// 1. 쿼리를 고속으로 하기위해 사용한다. == 날짜의 범위를 체크할 때
// 2. 나중에 업무를 위해서 Export Excel할 상황을 대비한다.
func TimeToNum(str string) (int64, error) {
	input, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return 0, err
	}
	initTime := time.Date(1899, time.December, 30, 0, 0, 0, 0, time.UTC)
	return (int64(input.Sub(initTime) / (time.Hour * 24))), nil
}

// EndOfMonth 는 주어진 연,월을 이용해 해당 달의 마지막 날을 구한다.
func EndOfMonth(year, month int) (time.Time, error) {
	t, err := time.Parse("2006-01-02", fmt.Sprintf("%04d-%02d-15", year, month))
	if err != nil {
		return time.Now().UTC(), err
	}
	return now.With(t.UTC()).EndOfMonth(), nil
}

// BeginningOfMonth 는 주어진 연,월을 이용해 해당 달의 첫 날을 구한다.
func BeginningOfMonth(year, month int) (time.Time, error) {
	t, err := time.Parse("2006-01-02", fmt.Sprintf("%04d-%02d-15", year, month))
	if err != nil {
		return time.Now().UTC(), err
	}
	return now.With(t.UTC()).BeginningOfMonth(), nil
}

// genDate는 연도와 월을 받아서 해당 달의 요일만큼 offset한 후 배열에 날짜를 채우는 함수이다.
func genDate(year, month int) ([42]int, error) {
	var l [42]int
	start, err := BeginningOfMonth(year, month)
	if err != nil {
		return l, err
	}
	end, err := EndOfMonth(year, month)
	if err != nil {
		return l, err
	}
	offset := int(start.Weekday())
	_, _, e := end.Date()
	for i := offset; i < e+offset; i++ {
		l[i] = i + 1 - offset
	}
	return l, nil
}

// genDateStr는 연도와 월을 받아서 해당 달의 요일만큼 offset한 후 배열에 날짜문자를 채우는 함수이다.
func genDateStr(year, month int) ([42]string, error) {
	var result [42]string // 2020-01-01 형태의 숫자가 저장된 리스트
	var lastYear int
	var nextYear int
	var lastMonth int
	var nextMonth int
	switch month {
	case 1:
		lastYear = year - 1
		nextYear = year
		lastMonth = 12
		nextMonth = month + 1
	case 12:
		lastYear = year
		nextYear = year + 1
		lastMonth = month - 1
		nextMonth = 1
	default:
		lastYear = year
		nextYear = year
		lastMonth = month - 1
		nextMonth = month + 1
	}
	lastEnd, err := EndOfMonth(lastYear, lastMonth)
	if err != nil {
		return result, err
	}
	currentStart, err := BeginningOfMonth(year, month)
	if err != nil {
		return result, err
	}
	currentEnd, err := EndOfMonth(year, month)
	if err != nil {
		return result, err
	}
	offset := int(currentStart.Weekday())
	_, _, e := currentEnd.Date()
	for n := 0; n < 42; n++ {
		if n < offset {
			_, _, e := lastEnd.Date()
			result[n] = fmt.Sprintf("%04d-%02d-%02d", lastYear, lastMonth, n+1-offset+e)
		} else if n < offset+e {
			result[n] = fmt.Sprintf("%04d-%02d-%02d", year, month, n+1-offset)
		} else {
			result[n] = fmt.Sprintf("%04d-%02d-%02d", nextYear, nextMonth, n+1-offset-e)
		}
	}
	return result, nil
}
