export APP=anthropos
export ANTHROPOS_CONSUL_URL=127.0.0.1:8500
export ANTHROPOS_CONSUL_PATH=anthropos

.PHONY: all test coverage
all: get build install
format:
	gofmt -l -s -w .
get:
	go get ./...
build:
	go build ./...
install:
	go install ./...
build-run:
	@curl --request \
                PUT --data-binary @config.example.yml \
                http://localhost:8500/v1/kv/${ANTHROPOS_CONSUL_PATH}
	go build -v .
	./${APP} serve -v
migration-up:
	@curl --request \
                PUT --data-binary @config.example.yml \
                http://localhost:8500/v1/kv/${ANTHROPOS_CONSUL_PATH}
	go build -v .
	./${APP} migration up
migration-down:
	@curl --request \
                PUT --data-binary @config.example.yml \
                http://localhost:8500/v1/kv/${ANTHROPOS_CONSUL_PATH}
	go build -v .
	./${APP} migration down
migration-reset:
	@curl --request \
                PUT --data-binary @config.example.yml \
                http://localhost:8500/v1/kv/${ANTHROPOS_CONSUL_PATH}
	go build -v .
	./${APP} migration down
	./${APP} migration up
test:
	go test ./... -v -coverprofile .coverage.txt
	go tool cover -func .coverage.txt
coverage: test
	go tool cover -html=.coverage.txt
local-env:
	echo "Preparing local env for local debugging. please wait........"
	echo "1. Setting up config to local consul........"
	docker compose up -d
	sleep 2
	curl --request PUT --data-binary @config.example.yml http://localhost:8500/v1/kv/${ANTHROPOS_CONSUL_PATH}
	exit
stage-env:

remove-env:
	echo "Removing docker container and clearing ports, please wait..."
	docker compose down
	exit
