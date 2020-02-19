package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	// TEMPLATES 는 kalena에서 사용하는 템플릿 글로벌 변수이다.
	TEMPLATES    = template.New("")
	flagAdd      = flag.Bool("add", false, "Add mode")
	flagSearch   = flag.Bool("search", false, "Search mode")
	flagAddLayer = flag.Bool("addlayer", false, "add layer mode")
	flagRmLayer  = flag.Bool("rmlayer", false, "remove layer mode")

	flagTitle       = flag.String("title", "", "Title")
	flagLayerName   = flag.String("layername", "", "Layer name")
	flagLayerColor  = flag.String("layercolor", "", "Layer Color")
	flagLayerHidden = flag.Bool("layerhidden", false, "Layer hidden")

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
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("kalena: ")
	flag.Parse()

	if *flagAdd {
		if *flagCollection == "" {
			log.Fatal("Collection이 빈 문자열 입니다")
		}
		s := Schedule{}
		s.Collection = *flagCollection
		s.Color = *flagLayerColor
		s.Title = *flagTitle
		s.Layer = *flagLayerName

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
	} else if *flagSearch { // 해당 연도, 달의 모든 스케쥴을 가져온다.
		if *flagCollection == "" {
			log.Fatal("Collection이 빈 문자열 입니다")
		}
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
		testdata, err := SearchMonth(session, *flagCollection, *flagLayerName, *flagYear, *flagMonth)
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
	} else if *flagAddLayer {
		if *flagCollection == "" {
			log.Fatal("collection을 입력해주세요")
		}
		if *flagLayerName == "" {
			log.Fatal("layername을 입력해주세요")
		}
		if *flagLayerColor == "" {
			log.Fatal("layercolor를 입력해주세요")
		}
		if !regexWebColor.MatchString(*flagLayerColor) {
			log.Fatal("#FFFFFF형식으로 입력해주세요")
		}
		session, err := mgo.Dial(*flagDBIP)
		if err != nil {
			log.Fatal(err)
		}
		defer session.Close()
		c := session.DB(*flagDBName).C(*flagCollection + ".layers")
		num, err := c.Find(bson.M{}).Count()
		order := num + 1
		err = AddLayer(session, *flagCollection, *flagLayerName, *flagLayerColor, order)
		if err != nil {
			log.Fatal(err)
		}
	} else if *flagRmLayer {
		if *flagCollection == "" {
			log.Fatal("collection을 입력해주세요")
		}
		if *flagLayerName == "" {
			log.Fatal("layername을 입력해주세요")
		}
		session, err := mgo.Dial(*flagDBIP)
		if err != nil {
			log.Fatal(err)
		}
		defer session.Close()
		err = RmLayer(session, *flagCollection, *flagLayerName)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		flag.PrintDefaults()
		os.Exit(1)
	}
}
