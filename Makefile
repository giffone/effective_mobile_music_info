.PHONY: build run down

build:
	docker compose build

run: build
	docker compose up -d
	# $(MAKE) clear

down:
	docker compose down

clear:
	docker compose rm -f migrate

	migrate:
	migrate -path db/migrations -database "$(DATABASE_URL)" up