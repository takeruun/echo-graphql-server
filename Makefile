setup:
	docker-compose build
	docker-compose run --rm echo_graphql ash -c "sql-migrate up && go run db/create_database.go"

db.create:
	docker-compose run --rm echo_graphql go run db/create_database.go

db.migrate:
	docker-compose run --rm echo_graphql sql-migrate up

db.seed:
	docker-compose run --rm echo_graphql go run db/seed/seeder.go

start:
	docker-compose up

end:
	docker-compose down

entry-server-container:
	docker-compose exec echo_graphql ash

entry-db-container:
	docker-compose exec db bash

test-cover:
	docker-compose exec echo_graphql go test -cover ./... -coverprofile=cover.out
	docker-compose exec echo_graphql go tool cover -html=cover.out -o cover.html