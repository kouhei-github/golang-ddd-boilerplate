MIGRATE=migrate -database 'mysql://${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}' -path ./migrations
.PHONY: migrate_up migrate_down

# make migrate_create name=create_table_user
migrate_create:
	${MIGRATE} create -ext sql -dir ./migrations -seq $(name)

migrate_up:
	${MIGRATE} up

migrate_down:
	${MIGRATE} down