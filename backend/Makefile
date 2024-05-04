N=1
# make migrateup N (the first one is 1, the second one is 2, and so on)
# make migratedown N (the last one is 1, the last two is 2, and so on)
# make migrateforce N (the first one is 1, the second one is 2, and so on)

postgres:
	docker run --name postgres12 -p 5432:5432 -v ./data:/var/lib/postgresql/data -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root mis

dropdb:
	docker exec -it postgres12 dropdb mis

migratecreate:
	migrate create -ext sql -dir db/migration -seq init-schema

migrateupAll:
	migrate -path db/migration -database "postgres://root:password@localhost:5432/mis?sslmode=disable" -verbose up

migrateup:
	migrate -path db/migration -database "postgres://root:password@localhost:5432/mis?sslmode=disable" -verbose up ${N}

migratedownAll:
	migrate -path db/migration -database "postgres://root:password@localhost:5432/mis?sslmode=disable" -verbose down

migratedown:
	migrate -path db/migration -database "postgres://root:password@localhost:5432/mis?sslmode=disable" -verbose down ${N}

migrateforce:
	migrate -path db/migration -database "postgres://root:password@localhost:5432/mis?sslmode=disable" force ${N}

.PHONY: postgres createdb dropdb migrateupAll migrateup migratedownAll migratedown migrateforce