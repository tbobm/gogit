SRC		=	gogit.go \
			action.go \
			commands.go

all:
	go install $(SRC)
