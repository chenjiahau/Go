postgres:
	docker run --name postgres12 -p 5432:5432 -v ./data:/var/lib/postgresql/data -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root mis

dropdb:
	docker exec -it postgres12 dropdb mis

migratecreate:
	migrate create -ext sql -dir db/migration -seq init-schema

migrateup:
	migrate -path db/migration -database "postgres://root:password@localhost:5432/mis?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://root:password@localhost:5432/mis?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown