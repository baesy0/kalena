package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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

	// 체크 구문
	err := l.CheckError()
	if err != nil {
		log.Fatal(err)
	}
	err = s.CheckError()
	if err != nil {
		log.Fatal(err)
	}

	l.Schedules = append(l.Schedules, s)
	c.Layers = append(c.Layers, l)

	fmt.Println(c)
}
