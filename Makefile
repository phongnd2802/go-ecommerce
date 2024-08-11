
MIGRATE_CMD = migrate -path internal/database/migrations/ -database "mysql://root:12345@tcp(localhost:3306)/GoDEV?multiStatements=true"

run:
	go run cmd/server/main.go

mysql:
	docker run --name mysql-db -p 3306:3306 -e MYSQL_ROOT_PASSWORD=12345 -d mysql

createdb:
	docker exec -it mysql-db mysql -uroot -p12345 --execute='CREATE DATABASE GoDEV'
	@echo "Create database successfully"

dropdb:
	docker exec -it mysql-db mysql -uroot -p12345 --execute='DROP DATABASE GoDEV'
	@echo "Drop database successfully"

migrate-up:
	$(MIGRATE_CMD) -verbose up

migrate-down:
	$(MIGRATE_CMD) -verbose down

migare-force:
	$(MIGRATE_CMD) force $(version)

sqlc:
	sqlc generate

.PHONY: run mysql createdb dropdb migrate-up migrate-down sqlc