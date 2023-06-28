include .env

postgres:
	docker run --name ${DOCKER_NAME} -p ${POSTGRES_PORT}:${POSTGRES_PORT} -e POSTGRES_USER=${POSTGRES_USER} -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -d postgres:15-alpine

createdb:
	docker exec -it ${DOCKER_NAME} createdb --username=${POSTGRES_USER} --owner=${POSTGRES_USER} ${POSTGRES_DB_NAME}

dropdb:
	docker exec -it ${DOCKER_NAME} dropdb ${POSTGRES_DB_NAME}

migrateup:
	migrate -path pkg/db/migration -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB_NAME}?sslmode=disable" -verbose up

migratedown:
	migrate -path pkg/db/migration -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB_NAME}?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown

run:
	go run cmd/api/main.go

.PHONY: run

