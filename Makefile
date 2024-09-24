.PHONY: db

default: iep

test:
	go test ./...

iep:
	go build -o ./bin ./cmd/iep

db:
	go build -o ./bin ./cmd/db

db-init:
	./bin/db init

db-migrate:
	./bin/db migrate

db-migration:
	touch "migrations/$(shell date +%Y%m%d%H%M%S)_placeholder.go"
