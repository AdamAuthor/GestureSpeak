docker run --name=gesture -e POSTGRES_PASSWORD='helloWorld7' -p 5432:5432 -d postgres
migrate -path ./migrations -database 'postgres://postgres:helloWorld7@localhost:5432/gesture?sslmode=disable' up