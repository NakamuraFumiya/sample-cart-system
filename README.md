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
|種類|値|
|--|--|
|サーバーサイド言語|[Go](https://go.dev/)|
|フレームワーク|[echo](https://github.com/labstack/echo)|
|RDBMS|[MySQL](https://dev.mysql.com/doc/)|
|スキーマ管理|structure.sql|
|ORM|[gorm](https://github.com/go-gorm/gorm)|
|環境変数|[viper](https://github.com/spf13/viper)|
|認証|[go-jose](https://github.com/go-jose/go-jose)|
|ロギング|[slog](https://pkg.go.dev/log/slog)|
|Linter|[golangci-lint](https://golangci-lint.run/)|
|テストヘルパー|[testfixtures](https://github.com/go-testfixtures/testfixtures)|
|環境構築|[Docker](https://docs.docker.com/)|
|コード自動生成|[oapi-codegen](https://github.com/oapi-codegen/oapi-codegen)|
|ホットリロード|[air](https://github.com/air-verse/air)|

