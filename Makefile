.PHONY: up build shell migrate rollback

up:
	docker-compose -f docker-compose.yml -f deployments/dev/docker-compose.yml up

build:
	docker-compose -f docker-compose.yml -f deployments/dev/docker-compose.yml build

shell:
	docker-compose exec app sh