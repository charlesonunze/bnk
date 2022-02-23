migrateup:
	migrate -path db/migrations -database ${DB_URI} up

migratedown:
	migrate -path db/migrations -database ${DB_URI} down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: migrateup migratedown sqlc test