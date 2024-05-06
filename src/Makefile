MIGRATE=migrate -database 'mysql://${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}' -path ./migrations
.PHONY: migrate_up migrate_down

migrate_up:
	${MIGRATE} up

migrate_down:
	${MIGRATE} down