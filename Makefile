.PHONY: iep

default: iep

test:
	go test ./...

iep:
	go build -o ./bin ./cmd/iep

db-migration:
	touch "cmd/db/migrations/$(shell date +%Y%m%d%H%M%S)_placeholder.go"
