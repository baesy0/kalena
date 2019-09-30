package main

import "net/http"



func webserver(port string){
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/add", handleAdd)
	http.ListenAndServe(port, nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("calendar page"))
}

func handleAdd(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("add page"))
}