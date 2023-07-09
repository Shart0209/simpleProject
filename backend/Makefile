include ../.env.dev

.PHONY: lint
lint:
	golangci-lint run

.PHONY: build
build:
	docker build -t app .
.PHONY: run_app
run_app:
	docker run -p 8000:8000 app

.PHONY: postgres
postgres:
	docker run --name ${DOCKER_NAME} -p ${POSTGRES_PORT}:${POSTGRES_PORT} -e POSTGRES_USER=${POSTGRES_USER} -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -d postgres:15-alpine

.PHONY: createdb
createdb:
	docker exec -it ${DOCKER_NAME} createdb --username=${POSTGRES_USER} --owner=${POSTGRES_USER} ${POSTGRES_DB_NAME}

.PHONY: dropdb
dropdb:
	docker exec -it ${DOCKER_NAME} dropdb ${POSTGRES_DB_NAME}

.PHONY: migrateup
migrateup:
	migrate -path simpleProject/backend/migration -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB_NAME}?sslmode=disable" -verbose up

.PHONY: migratedown
migratedown:
	migrate -path simpleProject/backend/migration -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB_NAME}?sslmode=disable" -verbose down

.PHONY: run
run:
	go run cmd/api/main.go


