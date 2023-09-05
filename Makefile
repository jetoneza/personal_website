run:
	go run ./cmd/build run

run_web:
	go run ./cmd/build run:web

build:
	go run ./cmd/build build

build_web:
	go run ./cmd/build build:web

build_prod:
	go run ./cmd/build build:prod
