.PHONY: db

default: build

test:
	go test ./...

db:
	go build -o ./bin ./cmd/db

db-init:
	./bin/db init

db-migrate:
	./bin/db migrate

db-migration:
	touch "migrations/$(shell date +%Y%m%d%H%M%S)_placeholder.go"
