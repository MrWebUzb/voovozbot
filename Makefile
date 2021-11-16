-include .env

POSTGRESQL_URL='postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=disable'

migrate-local-up-postgres:
	migrate -database ${POSTGRESQL_URL} -path migrations up

migrate-local-down-postgres:
	migrate -database ${POSTGRESQL_URL} -path migrations down

create-new-migration-postgres: # make create-new-migration name=file_name
	migrate create -ext sql -dir migrations -seq $(name)

run:
	go run cmd/main.go

.PHONY: run
.DEFAULT_GOAL:=run