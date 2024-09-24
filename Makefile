postgres:
	sudo docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	sudo docker exec -it postgres12 createdb --username=root --owner=root go_bank

dropdb:
	sudo docker exec -it postgres12 dropdb go_bank

migrateup:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/go_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/go_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/go_bank?sslmode=disable" -verbose down    

migratedown1:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/go_bank?sslmode=disable" -verbose down 1    

sqlc: 
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -build_flags=--mod=mod -package mockdb -destination db/mock/store.go github.com/jvsena42/go_bank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server mock