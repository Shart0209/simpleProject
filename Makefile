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
	docker compose  --env-file ./.env.dev logs app -f

.PHONY: imgd
imgd:
	docker image prune -a

.PHONY: vold
vold:
	docker volume prune -a

### dev
postgres:
	docker run --name db-test -p ${POSTGRES_PORT}:${POSTGRES_PORT} -e POSTGRES_USER=${POSTGRES_USER} -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -d postgres:15-alpine

.PHONY: createdb

createdb:
	docker exec db-test createdb -U ${POSTGRES_USER} -O ${POSTGRES_USER} ${POSTGRES_DB_NAME}

.PHONY: dropdb

dropdb:
	docker exec db-test dropdb -U ${POSTGRES_USER} ${POSTGRES_DB_NAME}
###
