package main

import (
	"net/http"
	"gopkg.in/mgo.v2"
	"log"
	"encoding/json"
)

func webserver() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/search", handleSearch)
	http.HandleFunc("/add", handleAdd)
	http.ListenAndServe(*flagHTTPPort, nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello kalena"))
}

func handleAdd(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("add page"))
}

// handleSearch
func handleSearch(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	userID := q.Get("userid")
	year := q.Get("year")
	month := q.Get("month")
	day := q.Get("day")
	layer := q.Get("layer")
	
	log.Println(year, month, day, layer)

	session, err := mgo.Dial(*flagDBIP)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
	schedules, err := allSchedules(session, userID)
	if err != nil{
		w.Write([]byte(err.Error()))
		return
	}

	err = json.NewEncoder(w).Encode(schedules)
	if err != nil {
		log.Println(err)
	}
}

