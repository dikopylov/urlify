.PHONY: up build shell migrate rollback

up:
	docker-compose -f docker-compose.yml -f config/dev/docker-compose.yml up -d

build:
	docker-compose -f docker-compose.yml -f config/dev/docker-compose.yml build

shell:
	docker-compose exec app sh

migrate:
	docker-compose run migrator -path /opt/app/urlify/db/migrations/ -database clickhouse://clickhouse:9000/default up

rollback:
	docker-compose run migrator -path /opt/app/urlify/db/migrations/ -database clickhouse://clickhouse:9000/default down

