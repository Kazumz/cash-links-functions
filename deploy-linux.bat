@echo off
echo Compiling Golang for Linux

set GOOS=linux
set GOARCH=amd64
go build handler.go

echo Compilation complete, please deploy using Azure Core Tools in VS Code...