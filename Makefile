DB_CONTAINER_NAME:=fieldsensing-mysql
DBNAME:=fieldsensing
TESTDBNAME:=test_fieldsensing

migrate/init:
	mysql -u root -h localhost --protocol tcp -e "create database \`$(DBNAME)\`" -p

migrate/test-init:
	mysql -u root -h localhost --protocol tcp -e "create database \`$(TESTDBNAME)\`" -p

migrate/up:
	docker-compose exec api goose up

migrate/down:
	docker-compose exec api goose down

docker/build:
	docker-compose build

docker/start:
	docker-compose up -d

docker/logs:
	docker-compose logs

docker/stop:
	docker-compose stop

docker/clean:
	docker-compose rm

api/bash:
	docker-compose exec api bash

react/bash:
	docker-compose exec react bash

db/bash:
	docker-compose exec db bash

api/init:
	docker-compose exec api dep ensure

npm/install:
	docker-compose exec react npm install

npm/start:
	docker-compose exec react npm run watch

npm/build:
	docker-compose exec react npm run build

run:
	docker-compose exec api go run main.go
