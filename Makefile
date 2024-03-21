build:
	@go build -o goo .

migrate:
	@sqlite3 goo.db < schema.sql
	@echo "Migration successful"

run:
	@go build -o goo . && ./goo

reset-db:
	@rm goo.db
	@touch goo.db
	@$(MAKE) migrate

test:
	@go test ./... -v
