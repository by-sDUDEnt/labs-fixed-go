export MIGRATIONS_PATH=migrations
export POSTGRES_URL=postgres://go-user:go-password@localhost:5432/go-db?sslmode=disable


run-wire:
	@cd internal/bootstrap && wire

install-migrate-mac:
	brew install golang-migrate

migrate-create:
	migrate create -ext sql -dir migrations -seq new

migrate-up-one:
	migrate -path ${MIGRATIONS_PATH} -database ${POSTGRES_URL} up 1

migrate-up:
	migrate -path ${MIGRATIONS_PATH} -database ${POSTGRES_URL} up

migrate-down-one:
	migrate -path ${MIGRATIONS_PATH} -database ${POSTGRES_URL} down 1


deps:
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/swaggo/swag/cmd/swag@latest
