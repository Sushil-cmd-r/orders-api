include .env

run:
	go run main.go

# database migrations
up:
	goose -dir migrations postgres ${DB_URL} up

down:
	goose -dir migrations postgres ${DB_URL} down


