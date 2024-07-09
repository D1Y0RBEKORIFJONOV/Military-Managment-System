run:
	go run cmd/soldiers_service/main.go



migrate-file:
	migrate create -ext sql -dir migrations/ -seq soldiers_table

DB_URL := "postgres://postgres:password@localhost:5432/solders_service?sslmode=disable"

migrate-up:
	migrate -path migrations -database $(DB_URL) -verbose up

migrate-down:
	migrate -path migrations -database $(DB_URL) -verbose down

migrate-force:
	migrate -path migrations -database $(DB_URL) -verbose force 1