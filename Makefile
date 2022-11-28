run:
	SERVER_ADDRESS=localhost SERVER_PORT=8080 DB_USER=hello_fastapi DB_PASSWORD=hello_fastapi DB_ADDRESS=localhost DB_PORT=5432 DB_NAME=banking go run main.go

files:
	 git ls-files | xargs wc -l

start_db:
	pwd
	docker-compose -f resources/postgres/docker-compose.yaml up

mock:
	go generate ./...

test:
	go test -v -cover ./...