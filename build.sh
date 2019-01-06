#!/bin/sh

PKG="github.com/fsufitch/algebra-expression/cmd/calculate-expression"

GOOS=windows GOARCH=386 go build -o build/calculator-x86.exe $PKG
GOOS=darwin GOARCH=amd64 go build -o build/calculator-darwin-amd64 $PKG
GOOS=linux GOARCH=amd64 go build -o build/calculator-linux-amd64 $PKG