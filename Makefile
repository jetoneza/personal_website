#!make
include .env

run:
	go run ./cmd/build run

run_web:
	go run ./cmd/build run:web

run_backend_prod:
	go run ./cmd/build run:backend:prod

run_frontend_prod:
	go run ./cmd/build run:frontend:prod

build:
	go run ./cmd/build build

build_web:
	go run ./cmd/build build:web

build_prod:
	go run ./cmd/build build:prod

deploy:
	git pull
	cp .env ./web/.env
	make build_backup
	make build_prod
	make run_apps

run_apps:
	pm2 delete -s go_app || :
	pm2 start ./bin/api --name go_app
	pm2 delete -s svelte_app || :
	PORT=${SVELTE_PORT} pm2 start ./web/build/index.js --name svelte_app

build_backup:
	go build -o bin/db_backup -v ./cmd/backup
