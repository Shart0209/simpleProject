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

.PHONY: up start down restart ps destroy

