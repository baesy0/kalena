package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
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

	flagHTTPPort = flag.String("http", "", "Web Service Port Number")
)

func main() {
	flag.Parse()
	if *flagAdd {
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
	} else if *flagHTTPPort != "" {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("kalena 1,2,3...31"))
		})
		http.ListenAndServe(*flagHTTPPort, nil)
	} else {
		flag.PrintDefaults()
		os.Exit(1)
	}
}
