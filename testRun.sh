#!/bin/sh
set -e

go run assets/asset_generate.go
go install
kalena -http :80
