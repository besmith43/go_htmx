
build:
	go build -o bin/main cmd/go_htmx/main.go

run:
	go run cmd/go_htmx/main.go

all: run