#!/bin/sh
APP="kalena"

go run assets/asset_generate.go

GOOS=linux GOARCH=amd64 go build -o ./bin/linux/${APP} kalena.go struct.go http.go dbapi.go check.go assets_vfsdata.go
GOOS=darwin GOARCH=amd64 go build -o ./bin/darwin/${APP} kalena.go struct.go http.go dbapi.go check.go assets_vfsdata.go
GOOS=windows GOARCH=amd64 go build -o ./bin/windows/${APP} kalena.go struct.go http.go dbapi.go check.go assets_vfsdata.go

# Github Release에 업로드 하기위해 압축
cd ./bin/linux/ && mkdir thumbnail && tar -zcvf ../${APP}_linux_x86-64.tgz . && cd -
cd ./bin/darwin/ && mkdir thumbnail && tar -zcvf ../${APP}_darwin_x86-64.tgz . && cd -
cd ./bin/windows/ && mkdir thumbnail && tar -zcvf ../${APP}_windows_x86-64.tgz . && cd -

# 삭제
rm -rf ./bin/linux
rm -rf ./bin/darwin
rm -rf ./bin/windows
