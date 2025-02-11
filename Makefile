.PHONY: iep

export XDG_STATE_HOME=$(shell pwd)/.local/state
export XDG_DATA_HOME=$(shell pwd)/.local/share

default: iep

test:
	go test ./...

iep:
	go build -o ./bin ./cmd/iep

db-migration:
	touch "db/migrations/$(shell date +%Y%m%d%H%M%S)_placeholder.go"

run:
	./bin/iep
