package main

import (
	"time"
)

// TimeToNum 함수는 시간문자를 받아서 1899년 12월 30일부터 몇번째 날인지 반환한다.
// Excel 에서 날수를 계산하는 방식과 같다.
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
