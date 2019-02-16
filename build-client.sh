#!/usr/bin/env bash

dir=datadict-cli

#windows32
mkdir -p $dir/datadict-cli-windows-386
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o datadict.exe
sleep 1
zip $dir/datadict-cli-windows-386.zip datadict.exe
sleep 1
mv datadict.exe $dir/datadict-cli-windows-386/datadict.exe

#windows64
mkdir -p $dir/datadict-cli-windows-amd64
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o datadict.exe
sleep 1
zip $dir/datadict-cli-windows-amd64.zip datadict.exe
sleep 1
mv datadict.exe $dir/datadict-cli-windows-amd64/datadict.exe

#linux32
mkdir -p $dir/datadict-cli-linux-386
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o datadict
sleep 1
zip $dir/datadict-cli-linux-386.zip datadict
sleep 1
mv datadict $dir/datadict-cli-linux-386/datadict

#linux64
mkdir -p $dir/datadict-cli-linux-amd64
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o datadict
sleep 1
zip $dir/datadict-cli-linux-amd64.zip datadict
sleep 1
mv datadict $dir/datadict-cli-linux-amd64/datadict

#mac32
mkdir -p $dir/datadict-cli-darwin-386
CGO_ENABLED=0 GOOS=darwin GOARCH=386 go build -o datadict
sleep 1
zip $dir/datadict-cli-darwin-386.zip datadict
sleep 1
mv datadict $dir/datadict-cli-darwin-386/datadict

#mac64
mkdir -p $dir/datadict-cli-darwin-amd64
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o datadict
sleep 1
zip $dir/datadict-cli-darwin-amd64.zip datadict
sleep 1
mv datadict $dir/datadict-cli-darwin-amd64/datadict
