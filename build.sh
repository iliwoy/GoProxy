#!/bin/bash

if [ ! -d "src/github.com/djimenez/iconv-go/" ]; then
mkdir -p src/github.com/djimenez/iconv-go/
git clone https://github.com/djimenez/iconv-go/ src/github.com/djimenez/iconv-go/
fi
export GOPATH=`pwd`
go build main
