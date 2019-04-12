# GraphQL golang template

## **Setup**

module管理にはgo modを使っている。
```
go build
```

## **Schema**
./schema.graphql

変更がある場合には以下のコマンドで ./infrastructure/graphql/exec.go を再生成

```
go run github.com/99designs/gqlgen
```

## **Prisma**

インストール
```
npm install -g prisma 
```

DBのMigrationとDAOの生成
```
prisma deploy
prisma generate
```

## **起動コマンド**
```
go run server.go
```

