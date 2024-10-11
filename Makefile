.PHONY: build run

build:
	docker compose build

run: build
	docker compose up -d

down:
	docker compose down