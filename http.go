package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"gopkg.in/mgo.v2"

	"github.com/shurcooL/httpfs/html/vfstemplate"
)

// LoadTemplates 함수는 템플릿을 로딩합니다.
func LoadTemplates() (*template.Template, error) {
	t := template.New("").Funcs(funcMap)
	t, err := vfstemplate.ParseGlob(assets, t, "/template/*.html")
	return t, err
}

var funcMap = template.FuncMap{
	"monthBefore": monthBefore,
	"monthAfter":  monthAfter,
	"yearBefore":  yearBefore,
	"yearAfter":   yearAfter,
	"offset":      offset,
	"onlyDate":    onlyDate,
	"checkFade":   checkFade,
}

func webserver() {
	// 템플릿 로딩을 위해서 vfs(가상파일시스템)을 로딩합니다.
	vfsTemplate, err := LoadTemplates()
	if err != nil {
		log.Fatal(err)
	}
	TEMPLATES = vfsTemplate
	// assets 설정
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(assets)))
	// 웹주소 설정
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/search", handleSearch)
	// RestAPI
	http.HandleFunc("/api/schedule", handleAPISchedule)
	// 웹서버 실행
	err = http.ListenAndServe(*flagHTTPPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	//userid가 빈 값이면 첫 번째 collection값을 넣어서 리다이렉트 해준다.
	q := r.URL.Query()
	userID := q.Get("userid")
	if userID == "" {
		session, err := mgo.Dial(*flagDBIP)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer session.Close()
		collections, err := GetCollections(session)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		userID = collections[0]
		http.Redirect(w, r, fmt.Sprintf("/?userid=%s", userID), http.StatusSeeOther)
	}

	w.Header().Set("Content-Type", "text/html")
	type Today struct {
		Year  int `bson:"year" json:"year"`
		Month int `bson:"month" json:"month"`
		Date  int `bson:"date" json:"date"`
	}
	type recipe struct {
		Theme      string     `bson:"theme" json:"theme"`
		Dates      [42]string `bson:"dates" json:"dates"`
		Today      `bson:"today" json:"today"`
		QueryYear  int     `bson:"queryyear" json:"queryyear"`
		QueryMonth int     `bson:"querymonth" json:"querymonth"`
		Layers     []Layer `bson:"layers" json:"layers"`
	}
	rcp := recipe{
		Theme: "default.css",
	}
	// 75mm studio 일때만 css 파일을 변경한다. 이 구조는 개발 초기에만 사용한다.
	if userID == "75mmstudio" {
		rcp.Theme = "75mmstudio.css"
	}
	y, m, d := time.Now().Date()
	rcp.Today.Year = y
	rcp.Today.Month = int(m)
	rcp.Today.Date = d

	month, err := strconv.Atoi(q.Get("month"))
	if err != nil {
		rcp.QueryMonth = rcp.Today.Month // 입력이 제대로 안되면 이번 달을 넣는다
	} else {
		rcp.QueryMonth = month
	}

	year, err := strconv.Atoi(q.Get("year"))
	if err != nil {
		rcp.QueryYear = rcp.Today.Year // 입력이 제대로 안되면 올해 연도를 넣는다.
	} else {
		rcp.QueryYear = year
	}
	rcp.Dates, err = genDate(rcp.QueryYear, rcp.QueryMonth)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session, err := mgo.Dial(*flagDBIP)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer session.Close()
	rcp.Layers, err = GetLayers(session, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = TEMPLATES.ExecuteTemplate(w, "index", rcp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// handleSearch
func handleSearch(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	userID := q.Get("userid")
	year := q.Get("year")
	month := q.Get("month")
	day := q.Get("day")
	layer := q.Get("layer")
	if layer == "" {
		http.Error(w, "URL에 layer를 입력해주세요", http.StatusBadRequest)
		return
	}
	sortKey := q.Get("sortkey")
	if userID == "" {
		http.Error(w, "URL에 userid를 입력해주세요", http.StatusBadRequest)
		return
	}

	log.Println(year, month, day, layer, sortKey)

	session, err := mgo.Dial(*flagDBIP)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
	schedules, err := allSchedules(session, userID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = json.NewEncoder(w).Encode(schedules)
	if err != nil {
		log.Println(err)
	}
}
