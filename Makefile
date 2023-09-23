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
