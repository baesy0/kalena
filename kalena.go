package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"

	"gopkg.in/mgo.v2"
)

var (
	// TEMPLATES 는 kalena에서 사용하는 템플릿 글로벌 변수이다.
	TEMPLATES  = template.New("")
	flagAdd    = flag.Bool("add", false, "Add mode")
	flagSearch = flag.Bool("search", false, "Search mode")

	flagTitle  = flag.String("title", "", "Title")
	flagLayer  = flag.String("layer", "", "Layer Title")
	flagColor  = flag.String("color", "", "Layer Color")
	flagHidden = flag.Bool("hidden", false, "Layer hidden")

	flagStart    = flag.String("start", "", "Start time")
	flagEnd      = flag.String("end", "", "End time")
	flagLocation = flag.String("location", "Asia/Seoul", "location name")

	flagYear  = flag.String("year", "", "year to search")
	flagMonth = flag.String("month", "", "month to search")

	flagCollection = flag.String("collection", "", "username for DB collection")

	flagDBIP     = flag.String("dbip", "", "DB IP")
	flagDBName   = flag.String("dbname", "kalena", "DB name")
	flagHTTPPort = flag.String("http", "", "Web Service Port Number")
)

func main() {
	flag.Parse()
	if *flagAdd {
		s := Schedule{}
		s.Collection = *flagCollection
		s.Color = *flagColor
		s.Hidden = *flagHidden
		s.Title = *flagTitle
		s.Layer = *flagLayer

		s.Start = *flagStart
		s.End = *flagEnd

		err := s.CheckError()
		if err != nil {
			log.Fatal(err)
		}

		session, err := mgo.Dial(*flagDBIP)
		if err != nil {
			log.Fatal(err)
		}
		defer session.Close()
		err = AddSchedule(session, s)
		if err != nil {
			log.Print(err)
		}
	} else if *flagSearch {
		if !regexInt4.MatchString(*flagYear) {
			log.Fatal("검색할 연도를 4자리 정수로 입력해주세요")
		}
		if !regexInt2.MatchString(*flagMonth) {
			log.Fatal("검색할 월을 2자리 정수로 입력해주세요(3월 -> 03, 11월 -> 11")
		}
		session, err := mgo.Dial(*flagDBIP)
		if err != nil {
			log.Fatal(err)
		}
		defer session.Close()
		testdata, err := SearchMonth(session, *flagCollection, *flagYear, *flagMonth)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(testdata)
		os.Exit(1)
	} else if *flagHTTPPort != "" {
		ip, err := serviceIP()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Service start: http://%s\n", ip)
		webserver()
	} else {
		flag.PrintDefaults()
		os.Exit(1)
	}
}
