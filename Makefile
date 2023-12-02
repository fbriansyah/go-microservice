DB_URL=postgresql://root:secret@localhost:5432/db_test?sslmode=disable

postgres:
	docker run --name pg-local -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it pg-local createdb --username=root --owner=root db_test

dropdb:
	docker exec -it pg-local dropdb db_test

migrateup:
	migrate -path migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path migration -database "$(DB_URL)" -verbose down 1

new_migration:
	migrate create -ext sql -dir migration -seq $(name)

re-db: dropdb createdb migrateup

run:
	go run ./cmd/api/

.PHONY: postgres createdb migrateup migrateup1 migratedown migratedown1 new_migration re-db run
