# kalena
![travisCI](https://secure.travis-ci.org/lazypic/kalena.png)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazypic/kalena)](https://goreportcard.com/report/github.com/lazypic/kalena)

Kalena is calendar web application based on opensource.<br>
켈레나는 오픈소스 달력 웹어플리케이션입니다.

## Features
- Support REST API
  - 켈레나는 REST API를 지원합니다.
- collaboration-oriented
  - 협업을 위한 달력을 만듭니다.
- Layer function
- Kalena can be installed on clound and intranet.
- It can be implemented in pipeline for contents creation.
  - 콘텐츠 제작 파이프라인에 활용할 수 있는다


### 임시데이터 확인하기
```
$ mongo
> use kalena
> db.user1.find()
```
### 전체 데이터 삭제하기
```
> db.user1.drop()
```

### url로 스케쥴 검색하기
```
http://localhost/search?userid=bae&year=2019&month=11&day=21&layer=se
```
