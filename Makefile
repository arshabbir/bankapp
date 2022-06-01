postgres :
	docker run --rm  -p 5432:5432 --name postgres  -e POSTGRES_PASSWORD=password  -d postgres:12-alpine 
createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres bank
dropdb:
	docker exec -it postgres dropdb --username=postgres  bank
.PHONY: createdb dropdb
