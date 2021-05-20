sqlc:
	cd database; sqlc generate; cd ../..;
run:
	go run main.go;
migrateup:
	migrate -path database/postgres/migration -database "postgresql://postgres:secret@localhost:5432/postgres?sslmode=disable" -verbose up\

migratedown:
	migrate -path database/postgres/migration -database "postgresql://postgres:secret@localhost:5432/postgres?sslmode=disable" -verbose down