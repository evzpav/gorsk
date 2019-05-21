POSTGRES_URL=postgres://gorsk:gorskpass@localhost:5433/gorskdb?sslmode=disable

.PHONY: build-docker run-local migrations build-postgres-docker run-postgres gorsk-initial-migration start-db

build-docker:
	docker build -t gorsk .


run-local:
	POSTGRES_URL=postgres://gorsk:gorskpass@localhost:5433/gorskdb?sslmode=disable \
	go run cmd/api/main.go

build-postgres-docker:
	docker build -t postgres_gorsk -f ./pkg/utl/postgres/postgres-docker/Dockerfile ./pkg/utl/postgres/postgres-docker/

run-postgres: build-postgres-docker
	docker run --name postgres_gorsk -d -p 5433:5432 postgres_gorsk

gorsk-initial-migration:
	go run cmd/migration/main.go

migrations: 
	docker run --rm -v ${PWD}/cmd/migration/:/migrations --network host migrate/migrate -path=/migrations/ -database ${POSTGRES_URL} up 1

check-table:
	docker exec -it postgres_gorsk psql -U postgres

start-db: build-postgres-docker run-postgres gorsk-initial-migration migrations
