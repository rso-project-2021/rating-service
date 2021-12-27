postgres:
	docker run --name rating_service_postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:alpine

createdb:
	docker exec -it rating_service_postgres createdb --username=root --owner=root electric_station

dropdb:
	docker exec -it rating_service_postgres dropdb electric_station

migrateup:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/electric_station?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/electric_station?sslmode=disable" -verbose down

test: 
	go test -cover ./...

.PHONY: postgres, createdb, dropdb, migrateup, migratedown, test