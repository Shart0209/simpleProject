postgres:
	docker run --name postgresDB -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=1234qwER -d postgres:15-alpine

createdb:
	docker exec -it postgresDB createdb --username=root --owner=root simpleDB

dropdb:
	docker exec -it postgresDB dropdb simpleDB

migrateup:
	migrate -path pkg/db/migration -database "postgresql://root:1234qwER@localhost:5432/simpleDB?sslmode=disable" -verbose up

migratedown:
	migrate -path pkg/db/migration -database "postgresql://root:1234qwER@localhost:5432/simpleDB?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown

run:
	go run cmd/api/main.go

.PHONY: run

