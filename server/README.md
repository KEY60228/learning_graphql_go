### ディレクトリ構造

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
