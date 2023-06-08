postgres:
	docker run --name postgres12 -e POSTGRES_USER=postgres  -e  POSTGRES_PASSWORD=blimbeng38 -d -p 5433:5433 postgres
#docker run --name pg12 -e POSTGRES_USER=postgres  -e  POSTGRES_PASSWORD=blimbeng38 -d -p 5434:5433 postgres:12-alpine
# docker run --name postgres12 -e POSTGRES_PASSWORD=blimbeng38 -d postgres
createdb:
	docker exec -it pg12 createdb --username=postgres --owner=postgres simple_bank
#docker exec -it pg12 createdb --username=postgres --owner=postgres db_bank
#  psql --username=postgres simple_bank
# docker exec -it pg12 psql -U postgres -d simple_bank
dropdb:
	docker exec -it pg12  dropdb --username=postgres  simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://postgres:blimbeng38@localhost:5433/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://postgres:blimbeng38@localhost:5433/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY:postgres createdb dropdb migrateup migratedown sqlc