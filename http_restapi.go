package main

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2"
)

func handleAPISchedule(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Post Only", http.StatusMethodNotAllowed)
		return
	}
	s := Schedule{}
	r.ParseForm()
	for key, values := range r.PostForm {
		switch key {
		case "collection":
			if len(values) != 1 {
				http.Error(w, "collection을 설정해 주세요", http.StatusBadRequest)
				return
			}
			s.Collection = values[0]
		case "title":
			if len(values) != 1 {
				http.Error(w, "title을 설정해 주세요", http.StatusBadRequest)
				return
			}
			s.Title = values[0]
		case "start":
			if len(values) != 1 {
				http.Error(w, "start를 설정해 주세요", http.StatusBadRequest)
				return
			}
			if !regexRFC3339Time.MatchString(values[0]) {
				http.Error(w, "시간 형식이 아닙니다", http.StatusBadRequest)
				return
			}
			s.Start = values[0]
		case "end":
			if len(values) != 1 {
				http.Error(w, "end를 설정해 주세요", http.StatusBadRequest)
				return
			}
			if !regexRFC3339Time.MatchString(values[0]) {
				http.Error(w, "시간 형식이 아닙니다", http.StatusBadRequest)
				return
			}
			s.End = values[0]
		case "color":
			if len(values) != 1 {
				http.Error(w, "color를 설정해 주세요", http.StatusBadRequest)
				return
			}
			if !regexWebColor.MatchString(values[0]) {
				http.Error(w, "#FFFFFF 형식이 아닙니다", http.StatusBadRequest)
				return
			}
			s.Color = values[0]
		case "layer":
			if len(values) != 1 {
				http.Error(w, "layer를 설정해 주세요", http.StatusBadRequest)
				return
			}
			s.Layer = values[0]
		case "hidden":
			if len(values) != 1 {
				http.Error(w, "hidden를 설정해 주세요", http.StatusBadRequest)
				return
			}
			if !(values[0] == "true" || values[0] == "false") {
				http.Error(w, "true, false 값만 사용할 수 있습니다", http.StatusBadRequest)
				return
			}
			if values[0] == "true" {
				s.Hidden = true
			} else {
				s.Hidden = false
			}
		}
	}
	err := s.CheckError()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	session, err := mgo.Dial(*flagDBIP)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer session.Close()
	err = AddSchedule(session, s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// json 으로 결과 전송
	data, _ := json.Marshal(s)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
