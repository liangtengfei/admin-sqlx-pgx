DB_URL=postgres://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable
UP:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up
UPP:
	migrate -path db/migrations -database postgresql://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable -verbose up