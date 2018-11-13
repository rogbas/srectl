ifndef VERBOSE
	MAKEFLAGS += --silent
endif

GIT_SHA=$(shell git rev-parse --verify HEAD)
PWD=$(shell pwd)
GOBUILD=go build
BINFILE="bin/srectl" 

default: compile

compile:
	${GOBUILD} -o ${BINFILE} main.go
