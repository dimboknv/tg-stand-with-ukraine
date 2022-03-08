export GO111MODULE := on

build:
	go build -o bin/tg-bot ./app

test_race:
	go test -race -v -mod=vendor -timeout=60s -count 1 ./...

lint:
	golangci-lint run --config .golangci.yml ./...

lint-fix:
	golangci-lint run --config .golangci.yml --fix ./...

fieldalignment:
	 fieldalignment ./...

fieldalignment-fix:
	 fieldalignment --fix ./...

docker-build:
	docker build -t tg-stand-with-ukraine .

.DEFAULT_GOAL := build
