-include .env

PROJECTNAME := $(shell basename "$(PWD)")

build:
	go build -tags=jsoniter -o bin/main .
	bin/main

run:
	go run main.go