.PHONY: iep

default: iep

test:
	go test ./...

iep:
	go build -o ./bin ./cmd/iep

db-migration:
	touch "db/migrations/$(shell date +%Y%m%d%H%M%S)_placeholder.go"
