REPOROOT = $(shell git rev-parse --show-toplevel)
DAY = 9
YEAR = 2022
REPOURL=github.com/VerwaerdeWim/AdventOfCode
EMAIL=$(shell git config user.email)

run:input.txt
	go run main.go
input.txt:
	curl 'https://adventofcode.com/$(YEAR)/day/$(DAY)/input' --cookie session=$(shell cat $(REPOROOT)/.sessionID) --user-agent "$(REPOURL) by $(EMAIL)" > input.txt
