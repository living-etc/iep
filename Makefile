.PHONY: iep

CWD=$(shell pwd)

export XDG_STATE_HOME=${CWD}/.local/state
export XDG_DATA_HOME=${CWD}/.local/share
export XDG_CONFIG_HOME=${CWD}/.config

default: iep

test:
	go test ./...

iep:
	go build -o ./bin ./cmd/iep

db-migration:
	touch "db/migrations/$(shell date +%Y%m%d%H%M%S)_placeholder.sql"

run:
	./bin/iep -config-file=.config/iep/config.json

logs:
	tail -f .local/state/iep/iep.log
