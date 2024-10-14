.PHONY: up-db migrate app-build run clear

up-db:
	docker compose -f docker-compose.db.yml up -d
	@echo "Database is up."

migrate:
	docker compose -f docker-compose.migrate.yml run --rm migrate
	@if [ $$? -eq 0 ]; then \
		echo "Migration completed."; \
	else \
        echo "Migration failed. Not starting app."; \
        exit 1; \
    fi

app-build:
	docker compose -f docker-compose.app.yml build
	@echo "App is built."

run: up-db migrate app-build
	docker compose -f docker-compose.app.yml up -d
	@echo "App is running."

clear:
	docker compose rm -f migrate