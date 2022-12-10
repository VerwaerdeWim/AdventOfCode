REPOROOT = $(shell git rev-parse --show-toplevel)
DAY = 3
YEAR = 2022

run:input.txt
	go run main.go
input.txt:
	curl 'https://adventofcode.com/$(YEAR)/day/$(DAY)/input' --cookie session=$(shell cat $(REPOROOT)/.sessionID)>input.txt
