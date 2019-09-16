package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	flagAdd            = flag.Bool("add", false, "Add mode")
	flagTitle          = flag.String("title", "", "Title")
	flagLayerTitle     = flag.String("layerTitle", "", "Layer Title")
	flagLayerColor     = flag.String("layerColor", "", "Layer Color")
	flagLayerHidden    = flag.Bool("hidden", false, "Layer hidden")
	flagLayerGreyscale = flag.Bool("greyscale", false, "Layer geyscale")
	flagStart          = flag.String("start", "", "Start time")
	flagEnd            = flag.String("end", "", "End time")
	flagLocation       = flag.String("location", "Asia/Seoul", "location name")
)

func main() {
	flag.Parse()

	if !*flagAdd {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *flagTitle == "" || *flagLayerTitle == "" || *flagStart == "" || *flagEnd == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *flagLayerColor != "" {
		if !regexWebColor.MatchString(*flagLayerColor) {
			log.Fatal("#FF0011 형식의 문자열이 아닙니다")
		}
	}

	//사용자에게 입력받은 데이터값이 유효한지 체크
	if !regexRFC3339Time.MatchString(*flagStart) {
		log.Fatal("2019-09-09T14:43:34+09:00 형식의 문자열이 아닙니다")
	}
	if !regexRFC3339Time.MatchString(*flagEnd) {
		log.Fatal("2019-09-09T14:43:34+09:00 형식의 문자열이 아닙니다")
	}
	startTime, err := time.Parse("2006-01-02T15:04:05-07:00", *flagStart)
	if err != nil {
		log.Fatal(err)
	}
	endTime, err := time.Parse("2006-01-02T15:04:05-07:00", *flagEnd)
	if err != nil{
		log.Fatal(err)
	}
	// end가 start 시간보다 큰지 체크하는 부분
	if !checkTime(startTime, endTime) {
		log.Fatal("끝시간이 시작시간보다 작습니다")
	}
	// declare variables for struct
	c := Calendar{}
	l := Layer{}
	s := Schedule{}

	l.Title = *flagLayerTitle
	l.Color = *flagLayerColor
	l.Hidden = *flagLayerHidden
	l.Greyscale = *flagLayerGreyscale
	s.Title = *flagTitle
	s.Start = *flagStart
	s.End = *flagEnd

	l.Schedules = append(l.Schedules, s)
	c.Layers = append(c.Layers, l)

	fmt.Println(c)
}
