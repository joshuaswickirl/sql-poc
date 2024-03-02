https://docs.sqlc.dev/en/stable/tutorials/getting-started-postgresql.html
https://pressly.github.io/goose/
https://sqlfluff.com/

https://pressly.github.io/goose/blog/2024/goose-sqlc/

Lint
```sh
sqlfluff lint --dialect postgres .
```

Generate
```sh
sqlc generate
```

Start pg
```sh
docker run -d -e POSTGRES_PASSWORD=password -e POSTGRES_USER=app -e POSTGRES_DB=db -p 5432:5432 postgres
```

Goose
```sh
export GOOSE_DRIVER="postgres"
export GOOSE_DBSTRING="host=localhost port=5432 user=app password=password dbname=db sslmode=disable"
export GOOSE_MIGRATION_DIR="sql/migrations"

goose up
```

App
```sh
go run main.go
```

PSQL
```sh
export PGPASSWORD=password
psql -U app -d db -h localhost -c "SELECT * FROM authors;"
```
