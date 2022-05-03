### ディレクトリ構造

#### gqlgen(ほぼ)デフォルト
```
server/
├─ graph/
│   ├─ generated/
│   │   └─ generated.go     (自動生成)
│   ├─ model/
│   │   └─ models_gen.go    型定義(自動生成)
│   ├─ resolver.go          状態を集約(？？)
│   ├─ schema.graphqls      スキーマ定義
│   └─ schema.resolvers.go  バックエンド実装
├─ generate.go              go generate用ファイル
├─ go.mod                   go modules
├─ go.sum                   go modules
├─ gqlgen.yml               gqlgenの設定ファイル
├─ README.md                本ファイル
├─ server.go                実行ファイル
└─ tools.go                 セットアップファイル
```

#### ちょっとアレンジ

schema.graphqlはserver/の外でgraphql/として持つ

```
server/
├─ cmd/
│   └─ server.go            gqlgenが初期生成したGraphQLサーバー実行ファイル
├─ domain/
│   ├─ model/               DDDの文脈で言うドメインモデルのイメージ
│   └─ service/             DDDの文脈で言うドメインサービスのイメージ
├─ graph/
│   ├─ generated/
│   │   └─ generated.go     自動生成ファイル
│   ├─ model/
│   │   ├─ xxx.go           GraphQLモデルの拡張
│   │   └─ models_gen.go    自動生成された型定義
│   ├─ resolver.go          状態を集約(？？)
│   └─ schema.resolvers.go  DDDの文脈で言うアプリケーションサービスのイメージ
├─ infrastructure/          Layered Architectureの文脈で言うインフラ層のイメージ
├─ middleware/              GraphQLリクエストに対するミドルウェア (ここじゃないかも)
├─ support/                 外部APIを叩く関数や定数など便利機能 (アンチパターンか？)
├─ generate.go              generateファイル
├─ go.mod                   go modules
├─ go.sum                   go modules
├─ gqlgen.yml               gqlgenの設定ファイル
├─ README.md                本ファイル
└─ tools.go                 セットアップファイル
```
