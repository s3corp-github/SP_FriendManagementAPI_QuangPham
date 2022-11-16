.PHONY: db db-migrate setup down db-down run migrateup migratedown

MOUNT_VOLUME ='$(realpath .)/data/migrations:/migrations'

setup: db db-migration

run:
	docker-compose up app

db:
	docker-compose up db -d

down:
	docker-compose down -v

db-migration:
	docker-compose run --rm -v $(MOUNT_VOLUME) db-migrate \
	sh -c 'migrate -path ./migrations -database postgres://friends-management:@db:5432/friends-management?sslmode=disable up'

db-down:
	docker-compose run --rm -v $(MOUNT_VOLUME) db-migrate \
	sh -c 'migrate -path ./migrations -database postgres://friends-management:@db:5432/friends-management?sslmode=disable down -all'

migrateup:
	migrate -path data/migrations -database 'postgres://friends-management:@localhost:5432/friends-management?sslmode=disable' -verbose up

migratedown:
	migrate -path data/migrations -database 'postgres://friends-management:@localhost:5432/friends-management?sslmode=disable' -verbose down
