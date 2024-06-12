postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root task-tracker

dropdb:
	docker exec -it postgres14 dropdb task-tracker

build:
	cd app && go build .

install:
	mv ./app/app ~/.local/bin/tt

.PHONY: postgres createdb dropdb

