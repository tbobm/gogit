SRC		=	gogit.go \
			actions.go \
			commands.go

all:
	go install $(SRC)
