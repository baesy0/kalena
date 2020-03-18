package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// handleAPISchedule 함수는 Schedule을 POST 한다.
func handleAPISchedule(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		s := Schedule{}
		r.ParseForm()
		collection := r.FormValue("collection")
		if collection == "" {
			http.Error(w, "collection을 설정해 주세요", http.StatusBadRequest)
			return
		}
		s.Collection = collection
		title := r.FormValue("title")
		if title == "" {
			http.Error(w, "title을 설정해 주세요", http.StatusBadRequest)
			return
		}
		s.Title = title
		start := r.FormValue("start")
		if start == "" {
			http.Error(w, "start를 설정해 주세요", http.StatusBadRequest)
			return
		}
		if !regexRFC3339Time.MatchString(start) {
			http.Error(w, "시간 형식이 아닙니다", http.StatusBadRequest)
			return
		}
		s.Start = start
		end := r.FormValue("end")
		if end == "" {
			http.Error(w, "end를 설정해 주세요", http.StatusBadRequest)
			return
		}
		if !regexRFC3339Time.MatchString(end) {
			http.Error(w, "시간 형식이 아닙니다", http.StatusBadRequest)
			return
		}
		s.End = end
		color := r.FormValue("color")
		if color == "" {
			http.Error(w, "color를 설정해 주세요", http.StatusBadRequest)
			return
		}
		if !regexWebColor.MatchString(color) {
			http.Error(w, "#FFFFFF 형식이 아닙니다", http.StatusBadRequest)
			return
		}
		s.Color = color
		layer := r.FormValue("layer")
		if layer == "" {
			http.Error(w, "layer를 설정해 주세요", http.StatusBadRequest)
			return
		}
		s.Layer = layer

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
		c := session.DB(*flagDBName).C(s.Collection + ".layers")
		// 레이어설정 컬렉션에서 이름이 일치하는 레이어가 있는지 검사한다.
		layerNum, err := c.Find(bson.M{"name": s.Layer}).Count()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// 이름이 일치하는 레이어가 없으면 에러처리
		if layerNum == 0 {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// 스케쥴 추가
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

// handleAPILayer 핸들러는 Layer를 POST 한다.
func handleAPILayer(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		l := Layer{}
		r.ParseForm()
		collection := r.FormValue("collection")
		if collection == "" {
			http.Error(w, "collection을 설정해 주세요", http.StatusBadRequest)
			return
		}
		name := r.FormValue("name")
		if name == "" {
			http.Error(w, "name 설정해 주세요", http.StatusBadRequest)
			return
		}
		l.Name = name
		order := r.FormValue("order")
		if order == "" {
			http.Error(w, "order를 설정해 주세요", http.StatusBadRequest)
			return
		}
		o, err := strconv.Atoi(order)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		l.Order = o
		color := r.FormValue("color")
		if color == "" {
			http.Error(w, "color를 설정해 주세요", http.StatusBadRequest)
			return
		}
		l.Color = color
		hidden := r.FormValue("hidden")
		if hidden == "" {
			http.Error(w, "hidden을 설정해 주세요", http.StatusBadRequest)
			return
		}
		b, err := strconv.ParseBool(hidden)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		l.Hidden = b

		err = l.CheckError()
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
