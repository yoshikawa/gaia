GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GODOC=$(GOCMD)doc
COMPOSE=docker-compose
EXEC=$(COMPOSE) exec
BUILD=$(COMPOSE) build
UP=$(COMPOSE) up -d
LOGS=$(COMPOSE) logs
STOP=$(COMPOSE) stop
RM=$(COMPOSE) rm
DOWN=$(COMPOSE) down
API=$(EXEC) api
DB=$(EXEC) db
REACT=$(EXEC) react
DBNAME:=fieldsensing
TESTDBNAME:=test_fieldsensing
MYSQL:=mysql --defaults-extra-file=/home/access.cnf

all: docker/up dep migrate/init migrate/up ## docker up & dep ensure & migrate

front: docker/up npm/install ## docker up & npm install

migrate/init: ## migrate init
	$(DB) $(MYSQL) -e "create database if not exists \`$(DBNAME)\`"

migrate/test-init: ## migrate test database init
	$(DB) $(MYSQL) -e "create database if not exists \`$(TESTDBNAME)\`"

migrate/up: ## migrate up
	$(API) goose up

migrate/down: ## migrate down
	$(API) goose down

docker/build: ## docker build
	$(BUILD)

docker/up: ## docker up
	$(UP)

docker/logs: ## docker logs
	$(LOGS)

docker/stop: ## docker stop
	$(STOP)

docker/clean: ## docker clean
	$(RM)

docker/down: ## docker down
	$(DOWN)

api/bash: ## api container bash
	$(API) bash

react/bash: ## react container bash
	$(REACT) bash

db/bash: ## db(MySQL) container bash
	$(DB) bash

dep: ## dep ensure
	$(API) dep ensure

dep/update: ## dep ensure update
	$(API) dep ensure -update

npm/install: ## npm install
	$(REACT) npm install

npm/watch: ## npm watch
	$(REACT) npm run watch

npm/build: ## npm build
	$(REACT) npm run build

run: ## go run main.go
	$(API) $(GORUN) main.go

test: ## go test
	$(API) $(GOTEST) -v ./...

doc: ## godoc http:6060
	$(GODOC) -http=:6060

help: ## Display this help screen
	@grep -E '^[a-zA-Z/_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
