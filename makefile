include .env
export

MIGRATE_PATH=./cmd/migrate/migrations

migrate-up:
	migrate -path $(MIGRATE_PATH) -database '$(DB_ADDR)' up

migrate-down:
	migrate -path $(MIGRATE_PATH) -database '$(DB_ADDR)' down 1

migrate-create:
	migrate create -seq -ext sql -dir $(MIGRATE_PATH) $(name)