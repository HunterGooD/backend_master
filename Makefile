postgres:
	docker run --name pstgr12 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=gfhjkm -p 5432:5432 -d postgres:12-alpine

postgresdel:
	docker stop pstgr12
	docker rm pstgr12

createdb:
	docker exec -it pstgr12 createdb --username=root --owner=root bank

dropdb:
	docker exec -it pstgr12 dropdb bank

migrateup:
	migrate -path db/migration -database "postgresql://root:gfhjkm@localhost:5432/bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:gfhjkm@localhost:5432/bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc