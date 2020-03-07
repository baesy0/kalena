# kalena
![travisCI](https://secure.travis-ci.org/lazypic/kalena.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazypic/kalena)](https://goreportcard.com/report/github.com/lazypic/kalena)

Kalena is calendar web application based on opensource.<br>
켈레나는 오픈소스 달력 웹어플리케이션입니다.

## Goal
We are making calender service for enterprise.<br>
기업을 위한 달력서비스를 만듭니다.

## Features
- collaboration-oriented
  - 협업을 위한 달력을 만듭니다.
- Support REST API
  - 켈레나는 REST API를 지원합니다.
- Developer-friendly
- Kalena can be installed on cloud and intranet.
  - 클라우드와 인트라넷에 설치가 가능합니다.
- Layer function


### 임시데이터 확인하기
#### 스케쥴 확인
```
$ mongo
> use kalena
> db.bae.find()
```
#### 레이어 확인
```
$ mongo
> use kalena
> db.bae.layers.find()
```
### 전체 데이터 삭제하기
```
> db.bae.drop()
```

### url로 스케쥴 검색하기
```
http://localhost/search?userid=bae&year=2019&month=11&day=21&layer=se
```

### License
- BSD 3-Clause License
