SHELL:=/usr/bin/bash

.PHONY: help make operation(s) easily

help:
	@echo "--------------- HELP -----------------"
	@echo "To run the project: make run"
	@echo "To build the project: make build"
	@echo "--------------------------------------"

run:
	go run main.go

build:
	go build -o ./build/lazypip
