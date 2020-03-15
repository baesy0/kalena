package main

import (
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
	// RestAPI
	http.HandleFunc("/api/schedule", handleAPISchedule)
	http.HandleFunc("/api/layer", handleAPILayer)
	// 웹서버 실행
	err = http.ListenAndServe(*flagHTTPPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	// Today 자료구조는 오늘 날짜를 하이라이트 하기위해서 사용하는 자료구조이다.
	type Today struct {
		Year  int `bson:"year" json:"year"`
		Month int `bson:"month" json:"month"`
		Date  int `bson:"date" json:"date"`
	}
	today := Today{}
	y, m, d := time.Now().Date()
	today.Year = y
	today.Month = int(m)
	today.Date = d

	session, err := mgo.Dial(*flagDBIP)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer session.Close()

	// 모든 collection을 가지고 온다
	collections, err := GetCollections(session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(collections) == 0 {
		http.Error(w, "no collection exists", http.StatusInternalServerError)
		return
	}

	q := r.URL.Query()
	// collection 가지고오기.
	currentcollection := q.Get("collection")
	if currentcollection == "" {
		currentcollection = collections[0]
	}
	// 연도를 가지고 온다.
	year, err := strconv.Atoi(q.Get("year"))
	if err != nil {
		year = today.Year
	}
	// 월을 가지고 온다.
	month, err := strconv.Atoi(q.Get("month"))
	if err != nil {
		month = today.Month
	}

	// currentlayer를 가지고 온다.
	currentLayer := q.Get("currentlayer")
	if currentLayer == "" {
		// currentLayer가 빈 문자열이면 currentcollection의 레이어들중 첫번째 레이어를 currentLayer로 설정한다.
		layers, err := GetLayers(session, currentcollection)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if len(layers) == 0 {
			currentLayer = ""
		} else {
			currentLayer = layers[0].Name
		}
	}
	// 아래부터는 달력을 렌더링 하기 위해서 생성하는 코드이다.

	type recipe struct {
		Collections       []string   `bson:"collections" json:"collections"`
		CurrentCollection string     `bson:"currentcollection" json:"currentcollection"`
		QueryYear         int        `bson:"queryyear" json:"queryyear"`
		QueryMonth        int        `bson:"querymonth" json:"querymonth"`
		Layers            []Layer    `bson:"layers" json:"layers"`
		CurrentLayer      string     `bson:"currentlayer" json:"currentlayer"`
		Theme             string     `bson:"theme" json:"theme"`
		Dates             [42]string `bson:"dates" json:"dates"`
		Today             `bson:"today" json:"today"`
	}
	rcp := recipe{
		Theme: "default.css",
	}
	rcp.Today = today
	rcp.Collections = collections
	rcp.CurrentCollection = currentcollection
	rcp.QueryYear = year
	rcp.QueryMonth = month
	rcp.Layers, err = GetLayers(session, currentcollection)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rcp.CurrentLayer = currentLayer
	rcp.Dates, err = genDate(rcp.QueryYear, rcp.QueryMonth)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// 사용자로부터 받은 데이터를 이용해서 스캐줄을 가지고와야 한다.
	// 미래에 구현한다.

	// 템플릿으로 렌더링한다.
	w.Header().Set("Content-Type", "text/html")
	err = TEMPLATES.ExecuteTemplate(w, "index", rcp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
