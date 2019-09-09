package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	flagAdd   = flag.Bool("add", false, "Add mode")
	flagTitle = flag.String("title", "", "Title")
	flagLayerTitle = flag.String("layerTitle", "", "Layer Title")
	flagLayerColor = flag.String("layerColor", "", "Layer Color")
	flagStart = flag.String("start", "", "Start time")
	flagEnd   = flag.String("end", "", "End time")
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
	if *flagLayerColor != ""{
		if !regexWebColor.MatchString(*flagLayerColor){
			log.Fatal("#FF0011 형식의 문자열이 아닙니다.")
		}
	}
	
	//사용자에게 입력받은 데이터값이 유효한지 체크
	if !regexTime.MatchString(*flagStart) {
		log.Fatal("2019-09-09 형식의 문자열이 아닙니다")
	}
	if !regexTime.MatchString(*flagEnd) {
		log.Fatal("2019-09-09 형식의 문자열이 아닙니다")
	}
	// declare variables for struct
	c := Calendar{}
	l := Layer{}
	s := Schedule{}

	l.Title = *flagLayerTitle
	l.Color = *flagLayerColor
	s.Title = *flagTitle
	s.Start = *flagStart
	s.End = *flagEnd

	l.Schedules = append(l.Schedules, s)
	c.Layers = append(c.Layers, l)

	fmt.Println(c)
}
