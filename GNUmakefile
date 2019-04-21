TEST?=./...
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
PKG_NAME=main

.EXPORT_ALL_VARIABLES:
CACHE_SIZE = 5
CACHE_TTL = 5
REDIS_URL = localhost:6379

up: ## build and run all services
	docker-compose up app

down: ## remove all services
	docker-compose down

unit_test: ## run tests against redis instances
	docker run -d --name redis -p 6379:6379 redis
	go test ./... -v
	docker rm redis -f

test: ## spins up the service and runs scripted tests
	DEMO=true docker-compose up

.PHONY: fmt
fmt: ## Run gofmt on all .go files
	gofmt -w $$(find . -name '*.go' | grep -v vendor)