up:
	docker-compose up -d

down:
	docker-compose down

run:
	go run cmd/server/main.go

test:
	go test -

migrate:
	migrate -path migrations -database "postgres://user:password@localhost:5432/tracker_db?sslmode=disable" up