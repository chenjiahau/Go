N=1 # Number of migration
DES=init-schema # Migration name

create_db:
	docker exec -it server-db-mis createdb --username=mis --owner=mis mis

drop_db:
	docker exec -it server-db-mis dropdb mis -U mis

migrate_create:
	migrate create -ext sql -dir db/migration -seq ${DES}

migrate_up_all:
	migrate -path db/migration -database "postgres://mis:mis@localhost:5432/mis?sslmode=disable" up

migrate_up:
	migrate -path db/migration -database "postgres://mis:mis@localhost:5432/mis?sslmode=disable" up ${N}

migrate_down_all:
	migrate -path db/migration -database "postgres://mis:mis@localhost:5432/mis?sslmode=disable" down

migrate_down:
	migrate -path db/migration -database "postgres://mis:mis@localhost:5432/mis?sslmode=disable" down ${N}

migrate_version:
	migrate -path db/migration -database "postgres://mis:mis@localhost:5432/mis?sslmode=disable" version

migrate_force:
	migrate -path db/migration -database "postgres://mis:mis@localhost:5432/mis?sslmode=disable" force ${N}

.PHONY: postgres createdb dropdb migrateupAll migrateup migratedownAll migratedown migrateversion migrateforce

