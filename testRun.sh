#!/bin/sh
go run assets/asset_generate.go
go install
kalena -http :8088