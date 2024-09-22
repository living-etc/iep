default: build

build:
	go build -o ./bin ./cmd/db

test:
	go test ./...

init:
	./bin/db init

migrate:
	./bin/db migrate

migration:
	touch "migrations/$(shell date +%Y%m%d%H%M%S)_placeholder.sql"
