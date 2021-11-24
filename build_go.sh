#!/bin/bash
set -e
echo "Building go code..."
go run go_code_gen.go
go fmt errcode/constant.go