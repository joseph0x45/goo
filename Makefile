build:
	@go build -o goo .

migrate:
	@sqlite3 goo.db < schema.sql
	@echo "Migration successful"

reset-db:
	@rm goo.db
	@touch goo.db
	@$(MAKE) migrate
