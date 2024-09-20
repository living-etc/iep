default: build

build:
	go build -o bin cmd/main.go

test:
	go test ./...

migrate:
	touch "migrations/$(shell date +%Y%m%d%H%M%S)_placeholder.sql"
