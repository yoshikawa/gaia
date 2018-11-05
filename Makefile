DB_CONTAINER_NAME:=fieldsensing-mysql
DBNAME:=fieldsensing
TESTDBNAME:=test_fieldsensing

migrate/init:
	mysql -u root -h localhost --protocol tcp -e "create database \`$(DBNAME)\`" -p

migrate/test-init:
	mysql -u root -h localhost --protocol tcp -e "create database \`$(TESTDBNAME)\`" -p

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

npm/install:
	docker-compose exec front npm install

npm/start:
	docker-compose exec front npm run watch

npm/build:
	docker-compose exec front npm run build