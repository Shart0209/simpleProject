include .env.dev

.PHONY: up
up:
	docker compose --env-file ./.env.dev up -d

.PHONY: conf
conf:
	docker compose --env-file ./.env.dev config

.PHONY: start
start:
	docker compose --env-file ./.env.dev start

.PHONY: down
down:
	docker compose --env-file ./.env.dev down

.PHONY: restart
restart:
	docker compose --env-file ./.env.dev stop
	docker compose --env-file ./.env.dev up -d

.PHONY: ps
ps:
	docker compose --env-file ./.env.dev ps

.PHONY: destroy
destroy:
	docker compose --env-file ./.env.dev down -v

.PHONY: shell-db
shell-db:
	docker compose --env-file ./.env.dev exec db psql -U ${POSTGRES_USER} -d ${POSTGRES_DB_NAME}

.PHONY: logs
logs:
	docker compose  --env-file ./.env.dev logs app

.PHONY: imgd
imgd:
	docker image prune -a

.PHONY: vold
vold:
	docker volume prune -a


.PHONY: createdb
createdb:
	docker compose  --env-file ./.env.dev exec db createdb -U ${POSTGRES_USER} -O ${POSTGRES_USER} ${POSTGRES_DB_NAME}

.PHONY: dropdb
dropdb:
	docker compose --env-file ./.env.dev exec db dropdb -U ${POSTGRES_USER} ${POSTGRES_DB_NAME}

.PHONY: migrateup
migrateup:
	migrate -path ./migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB_NAME}?sslmode=disable" -verbose up 2

.PHONY: migratedown
migratedown:
	migrate -path ./migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB_NAME}?sslmode=disable" -verbose down
