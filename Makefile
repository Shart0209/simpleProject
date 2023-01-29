createdb:
	docker exec -it postgresDB createdb --username=root --owner=root simple_dbcreatedb:

dropdb:
	docker exec -it postgresDB dropdb simple_db

.PHONY: createdb

.PHONY: run
run:
	go run cmd/api/main.go