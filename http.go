package main

import (
	"encoding/json"
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
	t := template.New("")
	t, err := vfstemplate.ParseGlob(assets, t, "/template/*.html")
	return t, err
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
	http.HandleFunc("/add", handleAdd)
	// RestAPI
	http.HandleFunc("/api/schedule", handleAPISchedule)
	// 웹서버 실행
	err = http.ListenAndServe(*flagHTTPPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	type Today struct {
		Year  int        `bson:"year" json:"year"`
		Month time.Month `bson:"month" json:"month"` // time.Month로 받는게 좋을까 int로 받는게 좋을까
		Date  int        `bson:"date" json:"date"`
	}
	type recipe struct {
		Theme string  `bson:"theme" json:"theme"`
		Dates [42]int `bson:"dates" json:"dates"`
		Today `bson:"today" json:"today"`
	}
	rcp := recipe{
		Theme: "default.css",
	}
	rcp.Today.Year, rcp.Today.Month, rcp.Today.Date = time.Now().Date() // 오늘에 해당하는 year, month는 추후 다시 사용한다
	q := r.URL.Query()
	userID := q.Get("userid")
	year, err := strconv.Atoi(q.Get("year"))
	if err != nil {
		year = rcp.Today.Year // 입력이 제대로 안되면 올해 연도를 넣는다.
	}
	month, err := strconv.Atoi(q.Get("month"))
	if err != nil {
		m := rcp.Today.Month // 입력이 제대로 안되면 이번 달을 넣는다
		month = int(m)
	}

	// 75mm studio 일때만 css 파일을 변경한다. 이 구조는 개발 초기에만 사용한다.
	if userID == "75mmstudio" {
		rcp.Theme = "75mmstudio.css"
	}
	rcp.Dates, err = genDate(year, month)
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
