package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	flagAdd   = flag.Bool("add", false, "Add mode")
	flagTitle = flag.String("title", "", "Title")
	flagStart = flag.String("start", "", "Start time")
	flagEnd   = flag.String("end", "", "End time")
)

func main() {
	flag.Parse()

	if !*flagAdd {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *flagTitle == "" || *flagStart == "" || *flagEnd == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	// declare variables for struct
	c := Calendar{}
	l := Layer{}
	s := Schedule{}

	s.Title = *flagTitle
	s.Start = *flagStart
	s.End = *flagEnd

	l.Schedules = append(l.Schedules, s)
	c.Layers = append(c.Layers, l)

	fmt.Println(c)
}
