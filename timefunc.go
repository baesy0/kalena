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

// genDate는 연도와 월을 받아서 해당 달의 요일만큼 offset한 후 배열에 날짜문자를 채우는 함수이다.
func genDate(year, month int) ([42]string, error) {
	var result [42]string // 2020-01-01 형태의 숫자가 저장될 리스트
	// 달력에는 현재 월,일 뿐 아니라 전달의 마지막 주, 다음달 첫주 날짜도 출력되어야 한다. 필요한 변수를 선언한다.
	var beforeYear int
	var afterYear int
	var beforeMonth int
	var afterMonth int
	// 1월, 12월은 이전해, 다음해의 값을 가지고 와야한다.
	switch month {
	case 1:
		beforeYear = year - 1
		afterYear = year
		beforeMonth = 12
		afterMonth = month + 1
	case 12:
		beforeYear = year
		afterYear = year + 1
		beforeMonth = month - 1
		afterMonth = 1
	default:
		beforeYear = year
		afterYear = year
		beforeMonth = month - 1
		afterMonth = month + 1
	}
	// 이전월의 마지막일을 구한다.
	beforeEnd, err := EndOfMonth(beforeYear, beforeMonth)
	if err != nil {
		return result, err
	}
	// 현재 월의 시작일을 구한다.
	currentStart, err := BeginningOfMonth(year, month)
	if err != nil {
		return result, err
	}
	// 현재 월의 마지막일을 구한다.
	currentEnd, err := EndOfMonth(year, month)
	if err != nil {
		return result, err
	}
	_, _, e := currentEnd.Date()
	// 이번달이 무슨 요일로 시작하는지 구한다. 요일값을 이용해서 날짜를 offset 하기 위함이다.
	offset := int(currentStart.Weekday())
	// 위해서 달력을 그리기에 모든 준비가 완료되었다. 위 숫자를 이용해서 달력을 채운다.
	for n := 0; n < 42; n++ {
		if n < offset { // 이번달 시작일보다 낮을때는 이전달의 마지막주의 날짜를 구하고 result에 넣는다.
			_, _, e := beforeEnd.Date()
			result[n] = fmt.Sprintf("%04d-%02d-%02d", beforeYear, beforeMonth, n+1-offset+e)
		} else if n < offset+e { // 현재달의 날짜를 구하고 result에 넣는다.
			result[n] = fmt.Sprintf("%04d-%02d-%02d", year, month, n+1-offset)
		} else { // 다음달 시작주의 날짜를 구하고 result에 넣는다.
			result[n] = fmt.Sprintf("%04d-%02d-%02d", afterYear, afterMonth, n+1-offset-e)
		}
	}
	return result, nil
}

// genNumDate는 연도와 월을 받아서 해당 달의 날짜를 엑셀날짜 형태로 채운다.
func genNumDate(year, month int) ([42]int64, error) {
	var result [42]int64 // 43748 형태의 숫자가 저장될 리스트
	// 현재 월의 시작일을 구한다.
	currentStart, err := BeginningOfMonth(year, month)
	if err != nil {
		return result, err
	}
	t := currentStart.Format("2006-01-02T15:04:05Z0700")
	// 시작일을 숫자로 변환한다.
	currentStartNum, err := TimeToNum(t)
	if err != nil {
		return result, err
	}
	// 이번달이 무슨 요일로 시작하는지 구한다. 요일값을 이용해서 날짜를 offset 하기 위함이다.
	offset := int(currentStart.Weekday())
	// 시작일에서 offset만큼 뺀 숫자부터 차례대로 채운다.
	StartNum := currentStartNum - int64(offset)
	for n := 0; n < 42; n++ {
		result[n] = StartNum
		StartNum++
	}
	return result, nil
}
