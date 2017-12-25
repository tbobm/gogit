SRC		=	gogit.go \
			actions.go \
			commands.go
NAME	= gogit

all: install

build:
	go build -o $(NAME)

install:
	go install

check:
	gofmt -w -l $(SRC)

clean:
	rm $(NAME)

test:
	go test

release:
	goreleaser

.PHONY: all build check clean install test
