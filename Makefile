up:
	docker compose up -d
start:
	docker compose start
down:
	docker compose down
restart:
	docker compose stop
	docker compose up -d
ps:
	docker compose ps
destroy:
	docker compose down -v
imgd:
	docker image prune -a
shell-db:
	docker compose exec db psql -U postgres -d postgres

.PHONY: up start down restart ps destroy imgd shell-db

