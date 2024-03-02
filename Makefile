include .env

migrate-up:
	migrate -database ${DB_SOURCE} -path internal/infra/database/migration up

migrate-down:
	migrate -database ${DB_SOURCE} -path internal/infra/database/migration down --all

migrate-create:
	@read -p "name of migration: " name; \
	migrate create -dir db/migration -ext sql -seq $$name


down:
	docker-compose down --volumes && docker volume prune -f

up:
	docker-compose up -d


run:
	go run cmd/main.go

start:
	make up
	sleep 10
	make migrate-up
	go run cmd/main.go

restart:
	make down
	make up
	sleep 10
	make migrate-up
	go run cmd/main.go

.PHONY: migrate-up migrate-down migrate-create down up sqlc start run restart
