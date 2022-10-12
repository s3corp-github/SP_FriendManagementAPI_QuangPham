createdb:
	docker exec -it postgres12-alpine createdb --username=root --owner=root friends_management

migrateup:
	migrate -path data/migrations -database "postgresql://root:secret@localhost:5432/friends_management?sslmode=disable" -verbose up

migratedown:
	migrate -path data/migrations -database "postgresql://root:secret@localhost:5432/friends_management?sslmode=disable" -verbose down

sqlboiler:
	sqlboiler psql -c sqlboiler.toml --wipe --no-tests
.PHONY:	createdb migrateup migratedown sqlboiler