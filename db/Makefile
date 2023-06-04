createdb:
	docker exec -it pg12 createdb --username=postgres --owner=postgres simple_bank

#  psql --username=postgres simple_bank
dropdb:
	docker exec -it pg12 dropdb simple_bank

.PHONY: createdb