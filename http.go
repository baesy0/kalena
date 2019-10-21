package main

import (
	"net/http"
	"gopkg.in/mgo.v2"
	"log"
	"encoding/json"
)

func webserver() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/add", handleAdd)
	http.ListenAndServe(*flagHTTPPort, nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	userID := q.Get("userid")

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

func handleAdd(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("add page"))
}
