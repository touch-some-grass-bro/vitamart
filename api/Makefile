migrate_up:
	go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate -source file://db/migrations -database postgresql://postgres:password@localhost:5432/vitamart?sslmode=disable up

migrate_down:
	go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate -source file://db/migrations -database postgresql://postgres:password@localhost:5432/vitamart?sslmode=disable down 1


make_migration:
	@read -p "Enter file name: " MIGRATION_NAME; \
	go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate create -ext sql -dir db/migrations -seq $$MIGRATION_NAME
