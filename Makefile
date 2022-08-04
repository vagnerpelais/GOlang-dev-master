postgres:
	docker run --name postgres-bank -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:12-alpine
createdb:
	docker exec -it postgres-bank createdb --username=root --owner=root simple_bank 

dropdb:
	docker exec -it postgres-bank dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...
	
build:
	go build -o .

server:
	./simplebank
	
mock:
	mockgen --build_flags=--mod=mod -package mockdb  -destination db/mock/store.go  github.com/vagnerpelais/simplebank/db/sqlc Store	

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test build server mock