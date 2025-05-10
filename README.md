# sample-cart-system

## Setup
### Local Database

```
brew install sqldef/sqldef/mysqldef
```

Reflect the local database schema into structure.sql
```
make dump-schema
```

Update the local database schema based on the contents of structure.sql
```
make apply-schema
```

## Using
### start api server

```
go run cmd/apiserver/server.go
```

## dependency rule

## API tech stack
| Type                 | Value          |
|----------------------|----------------|
| Server-side Language | Go             |
| Framework            | Echo           |
| RDBMS                | MySQL          |
| Schema Management    | structure.sql  |
| ORM                  | GORM           |
| Environment Variables| Viper          |
| Authentication       | go-jose        |
| Logging              | slog           |
| Linter               | golangci-lint  |
| Test Helper          | testfixtures   |
| Environment Setup    | Docker         |
| Code Generation      | oapi-codegen   |
| Hot Reloading        | air            |


