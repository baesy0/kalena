package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gopkg.in/mgo.v2"
)

func handleAPISchedule(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
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
		data, err := json.Marshal(s)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	} else {
		http.Error(w, "Not Supported Method", http.StatusMethodNotAllowed)
		return
	}
}

// handleAPILayer 핸들러는 Layer 정보를 설정한다.
func handleAPILayer(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var collection string
		l := Layer{}
		r.ParseForm()
		for key, values := range r.PostForm {
			switch key {
			case "collection":
				if len(values) != 1 {
					http.Error(w, "collection을 설정해 주세요", http.StatusBadRequest)
					return
				}
				collection = values[0]
			case "name":
				if len(values) != 1 {
					http.Error(w, "name을 설정해 주세요", http.StatusBadRequest)
					return
				}
				l.Name = values[0]
			case "order":
				if len(values) != 1 {
					http.Error(w, "순서를 설정해 주세요", http.StatusBadRequest)
					return
				}
				order, err := strconv.Atoi(values[0])
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				l.Order = order
			case "color":
				if len(values) != 1 {
					http.Error(w, "color를 설정해 주세요", http.StatusBadRequest)
					return
				}
				l.Color = values[0]
			}
		}
		err := l.CheckError()
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
		err = AddLayer(session, collection, l)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// json 으로 결과 전송
		data, err := json.Marshal(l)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	} else {
		http.Error(w, "Not Supported Method", http.StatusMethodNotAllowed)
		return
	}
}
