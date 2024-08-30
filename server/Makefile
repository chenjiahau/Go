N=1
# make migrateup N
# make migratedown N
# make migrateforce N

createdb:
	docker exec -it server-db-mis createdb --username=mis --owner=mis mis

dropdb:
	docker exec -it server-db-mis dropdb mis -U mis

migratecreate:
	migrate create -ext sql -dir db/migration -seq init-schema

migrateupAll:
	migrate -path db/migration -database "postgres://mis:mis@localhost:5432/mis?sslmode=disable" -verbose up

migrateup:
	migrate -path db/migration -database "postgres://mis:mis@localhost:5432/mis?sslmode=disable" -verbose up ${N}

migratedownAll:
	migrate -path db/migration -database "postgres://mis:mis@localhost:5432/mis?sslmode=disable" -verbose down

migratedown:
	migrate -path db/migration -database "postgres://mis:mis@localhost:5432/mis?sslmode=disable" -verbose down ${N}

migrateforce:
	migrate -path db/migration -database "postgres://mis:mis@localhost:5432/mis?sslmode=disable" force ${N}

.PHONY: postgres createdb dropdb migrateupAll migrateup migratedownAll migratedown migrateforce

