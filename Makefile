GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GODOC=$(GOCMD)doc
COMPOSE=docker-compose
COMPOSEEXEC=$(COMPOSE) exec
COMPOSEBUILD=$(COMPOSE) build
COMPOSEUP=$(COMPOSE) up -d
COMPOSELOGS=$(COMPOSE) logs
COMPOSESTOP=$(COMPOSE) stop
COMPOSERM=$(COMPOSE) rm
DBNAME:=fieldsensing
TESTDBNAME:=test_fieldsensing
MYSQL:=mysql --defaults-extra-file=./docker/mysql/my.cnf -h localhost --protocol tcp

all: docker/up dep migrate/init migrate/up ## docker up & dep ensure & migrate

front: docker/up npm/install ## docker up & npm install

migrate/init: ## migrate init
	$(MYSQL) -e "create database \`$(DBNAME)\`"

migrate/test-init: ## migrate test database init
	$(MYSQL) -e "create database \`$(TESTDBNAME)\`"

migrate/up: ## migrate up
	$(COMPOSEEXEC) api goose up

migrate/down: ## migrate down
	$(COMPOSEEXEC) api goose down

docker/build: ## docker build
	$(COMPOSEBUILD)

docker/up: ## docker up
	$(COMPOSEUP)

docker/logs: ## docker logs
	$(COMPOSELOGS)

docker/stop: ## docker stop
	$(COMPOSESTOP)

docker/clean: ## docker clean
	$(COMPOSERM)

api/bash: ## api container bash
	$(COMPOSEEXEC) api bash

react/bash: ## react container bash
	$(COMPOSEEXEC) react bash

db/bash: ## db(MySQL) container bash
	$(COMPOSEEXEC) db bash

dep: ## dep ensure
	$(COMPOSEEXEC) api dep ensure

dep/update: ## dep ensure update
	$(COMPOSEEXEC) api dep ensure -update

npm/install: ## npm install
	$(COMPOSEEXEC) react npm install

npm/watch: ## npm watch
	$(COMPOSEEXEC) react npm run watch

npm/build: ## npm build
	$(COMPOSEEXEC) react npm run build

run: ## go run main.go
	$(COMPOSEEXEC) api $(GORUN) main.go

test: ## go test
	$(COMPOSEEXEC) api $(GOTEST) -v ./...

doc: ## godoc http:6060
	$(GODOC) -http=:6060

help: ## Display this help screen
	@grep -E '^[a-zA-Z/_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'