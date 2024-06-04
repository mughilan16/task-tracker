postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root task-tracker

dropdb:
	docker exec -it postgres14 dropdb task-tracker

.PHONY: postgres createdb dropdb

