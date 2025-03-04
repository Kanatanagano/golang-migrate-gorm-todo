.PHONY: upd
upd:
	docker compose up -d
.PHONY: up
up:
	docker compose up

.PHONY: down
down:
	docker compose down

.PHONY: build
build:
	docker compose build
. PHONY: stop
stop:
	docker compose stop
.PHONY: docker-psql
docker-psql:
	docker exec -it todo-pgsql psql -U postgres -d kanata_todo_db
.PHONY: migrate-up
migrate-up:
	migrate -path ./server/db/migrations -database "postgres://postgres:password@localhost:5432/kanata_todo_db?sslmode=disable" up
.PHONY: migrate-down
migrate-down:
	migrate -path ./server/db/migrations -database "postgres://postgres:password@localhost:5432/kanata_todo_db?sslmode=disable" down
.PHONY: migrate-version
migrate-version:
	migrate -path ./server/db/migrations -database "postgres://postgres:password@localhost:5432/kanata_todo_db?sslmode=disable" version
.PHONY: psql-schema
psql-schema:
	docker exec -it todo-pgsql psql -U postgres -d kanata_todo_db -c "\dt"
.PHONY: seed
seed:
	migrate -path ./server/db/seed -database "postgres://postgres:password@localhost:5432/kanata_todo_db?sslmode=disable&x-migrations-table=seed_migrations" up
.PHONY: seed-down
seed-down:
	migrate -path ./server/db/seed -database "postgres://postgres:password@localhost:5432/kanata_todo_db?sslmode=disable&x-migrations-table=seed_migrations" down
