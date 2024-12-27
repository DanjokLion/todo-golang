docker run --name=todo-db-go -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD='
qwerty' -v E:/dockerFolder/todo-db-go:/var/lib/postgresql/data -p 5433:5432 -d --rm postgres

migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5433/postgres?sslmode=disable' up - to up database schema
