postgres:
	docker run --name postgres16 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -p 5432:5432 -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root go_bank

dropdb:
	docker exec -it postgres16 dropdb go_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/go_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/go_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...


.PHONY: postgres createdb dropdb migrateup migratedown