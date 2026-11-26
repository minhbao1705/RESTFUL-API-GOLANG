include .env
export

CONN_STRING = postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)

MIGRATIONS_DIRS = internal/db/migrations


# Import database
import-db: 
	docker exec -i postgres-db psql -U root -d master-golang < backupdb-master-golang.sql

# Export database
export-db:
	docker exec postgres-db pg_dump -U root -d master-golang --encoding=UTF8 > ./backupdb-master-golang.sql

# Run server
server:
	go run .

# Create a new migration (make migrate-create NAME=profiles)
migrate-create:
	migrate create -ext sql -dir $(MIGRATIONS_DIRS) -seq $(NAME)

# Run all pending migration (make migrate-up)
migrate-up:
	migrate -path $(MIGRATIONS_DIRS) -database "$(CONN_STRING)" up 

# Rollback N migration (make migrate-down-n N=2)
migrate-down-n:
	migrate -path $(MIGRATIONS_DIRS) -database "$(CONN_STRING)" down $(N)

# Rollback the last migration (make migrate-down)
migrate-down:
	migrate -path $(MIGRATIONS_DIRS) -database "$(CONN_STRING)" down 1

# Force migration version (use with caution) (make migrate-force VERSION=1)
migrate-force:
	migrate -path $(MIGRATIONS_DIRS) -database "$(CONN_STRING)" force $(VERSION)

# Drop everything (include schema-migration)
migrate-drop:
	migrate -path $(MIGRATIONS_DIRS) -database "$(CONN_STRING)" drop

# Apply specific migration version (make migrate-goto VERSION=1)
migrate-goto:
	migrate -path $(MIGRATIONS_DIRS) -database "$(CONN_STRING)" goto $(VERSION)

.PHONY: import-db export-db server migrate-create migrate-up migrate-down migrate-force migrate-drop migrate-goto migrate-down-n