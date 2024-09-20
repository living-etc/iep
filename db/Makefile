default: build

build:
	go build -o bin cmd/main.go

test:
	go test ./...

init:
	./bin/main init

migrate:
	./bin/main migrate

migration:
	touch "migrations/$(shell date +%Y%m%d%H%M%S)_placeholder.sql"
