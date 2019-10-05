package main

import (
	"testing"
	"time"
)

func Test_checkTime(t *testing.T) {
	cases := []struct {
		start         string
		end           string
		startLocation string
		endLocation   string
		want          bool
	}{{
		start:         "2019-09-13T22:04:32+09:00",
		end:           "2019-09-14T22:04:32+09:00",
		startLocation: "Asia/Seoul",
		endLocation:   "Asia/Seoul",
		want:          true,
	}, {
		start:         "2019-09-15T22:04:32+09:00",
		end:           "2019-09-14T22:04:32+09:00",
		startLocation: "Asia/Seoul",
		endLocation:   "Asia/Seoul",
		want:          false,
	}, {
		start:         "2019-09-15T21:04:31+08:00", // 중국시간
		end:           "2019-09-15T22:04:32+09:00",
		startLocation: "Asia/Seoul",
		endLocation:   "Asia/Seoul",
		want:          true,
	}, {
		start:         "2019-09-15T22:04:32+09:00",
		end:           "2019-09-15T22:04:33+09:00",
		startLocation: "Asia/Seoul",
		endLocation:   "Asia/Seoul",
		want:          true,
	}}

	for _, c := range cases {
		s, err := time.Parse("2006-01-02T15:04:05-07:00", c.start)
		if err != nil {
			t.Fatal(err)
		}
		e, err := time.Parse("2006-01-02T15:04:05-07:00", c.end)
		if err != nil {
			t.Fatal(err)
		}
		startLoc, err := time.LoadLocation(c.startLocation)
		if err != nil {
			t.Fatal(err)
		}
		endLoc, err := time.LoadLocation(c.endLocation)
		if err != nil {
			t.Fatal(err)
		}
		start := s.In(startLoc)
		end := e.In(endLoc)
		if end.After(start) != c.want {
			t.Fatalf("Test_checkTime(%s,%s): 얻은 값: %v, 원하는 값: %v\n", s.In(startLoc).UTC(), e.In(endLoc).UTC(), end.After(start), c.want)
		}
	}
}
