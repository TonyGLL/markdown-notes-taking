initschema:
	migrate create -ext sql -dir db/migration -seq init_schema

postgres:
	docker compose up -d --build db

stop-postgres:
	docker stop postgres_markdown_notes

createdb:
	docker exec -it postgres_markdown_notes createdb --username=root --owner=root markdown_notes

dropdb:
	docker exec -it postgres_markdown_notes dropdb markdown_notes

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/markdown_notes?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/markdown_notes?sslmode=disable" -verbose down

build:
	@echo "Building..."

	@CONFIG_FILE=local.env go build -o main main.go

	@echo "Build successfully!!!"

start: build
	CONFIG_FILE=local.env air

# Live Reload
watch: build
	@if command -v air > /dev/null; then \
	    CONFIG_FILE=local.env air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/air-verse/air@latest; \
	        CONFIG_FILE=local.env air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

.PHONY: initschema postgres createdb dropdb migrateup migratedown