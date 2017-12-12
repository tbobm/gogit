NAME	= gogit
TAG	=

all: install

build:
	go build -o $(NAME)

install:
	go install

check:
	gofmt -w $(wildcard *.go)

clean:
	rm $(NAME)

test:
	go test

release:
	goreleaser

.PHONY: all build check clean install test
