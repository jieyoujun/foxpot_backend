VERSION := v1.0.0-pre
NAME := foxpot
BUILDSTRING := $(shell git log --pretty=format:'%h' -n 1)
VERSIONSTRING := $(NAME) version $(VERSION)+$(BUILDSTRING)
BUILDDATE := $(shell date -u -Iseconds)

LDFLAGS := "-X \"main.VERSION=$(VERSIONSTRING)\" -X \"main.BUILDDATE=$(BUILDDATE)\""

.PHONY: all test clean build

default: build

build:
	go build -ldflags=$(LDFLAGS) -o bin/foxpot main.go

static:
	go build --ldflags '-extldflags "-static"' -o bin/foxpot main.go
	upx -1 bin/foxpot

clean:
	rm -rf bin/

run: build
	sudo bin/foxpot

docker:
	docker build -t foxpot .
	docker run -it foxpot