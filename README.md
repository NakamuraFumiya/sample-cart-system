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

## start api server

```
go run cmd/apiserver/main.go
```

## dependency rule
