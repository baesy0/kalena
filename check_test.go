package main

import (
	"testing"
	"time"
)

// TestUserArgvValue 테스트 코드는 사용자로부터 터미널에서 start, end 인수를 받았다고 가정하고 테스트해보는 함수이다.
// 프로그래머가 머릿속으로 시간을 다룰 때 헷갈릴경우 자신이 생각하는 값을 입력하고 시뮬레이션 할 때 사용한다.
func TestUserArgvValue(t *testing.T) {
	cases := []struct {
		start string
		end   string
		want  bool
	}{{
		start: "2019-09-13T22:04:32+09:00",
		end:   "2019-09-14T22:04:32+09:00",
		want:  true,
	}, {
		start: "2019-09-15T22:04:32+09:00",
		end:   "2019-09-14T22:04:32+09:00",
		want:  false,
	}, {
		start: "2019-09-15T21:04:31+08:00", // 한국과 중국은 1시간 차이난다. 중국시간으로 시작하고 1시간 시간을 offset후 1초 차이를 두었을 때 잘 체크되는지 체크하는 값
		end:   "2019-09-15T22:04:32+09:00",
		want:  true,
	}, {
		start: "2019-09-15T22:04:32+09:00",
		end:   "2019-09-15T22:04:33+09:00",
		want:  true,
	}}

	for _, c := range cases {
		s, err := time.Parse(time.RFC3339, c.start)
		if err != nil {
			t.Fatal(err)
		}
		e, err := time.Parse(time.RFC3339, c.end)
		if err != nil {
			t.Fatal(err)
		}
		if e.UTC().After(s.UTC()) != c.want {
			t.Fatalf("Test_checkTime(%s,%s): 얻은 값: %v, 원하는 값: %v\n", c.start, c.end, e.UTC().After(s.UTC()), c.want)
		}
	}
}
