# go.server

OpenAPI スキーマを単一ソースとするモノレポ。

## 構成

```
go.server/
├── openapi/                  # 単一スキーマ (source of truth) + 生成設定
│   ├── openapi.yaml
│   ├── cfg.yaml              # oapi-codegen (Go) 設定
│   └── orval.config.ts       # orval (TS クライアント) 設定
├── packages/
│   ├── server/               # Go API サーバー (module: github.com/tsusowake/go.server)
│   ├── web/                  # 一般ユーザー向け Web (React Router v7 framework / SSR, :3000)
│   └── admin-console/        # 管理者向け Web (React Router v7 framework / SSR, :3333)
├── package.json              # bun workspaces ルート
├── Makefile                  # ルートタスク (gen / dev / build)
└── compose.yaml / .docker/   # infra
```

## セットアップ

```sh
bun install            # フロントエンド依存をインストール
make gen               # OpenAPI から Go 型 + TS クライアントを生成
```

## コード生成 (schema-first)

```sh
make gen-oapi          # openapi.yaml -> Go 型 (packages/server)
make gen-api           # openapi.yaml -> TS クライアント (web / admin-console)
make gen               # 両方
```

## 開発

```sh
make run               # DB (docker compose)
make build-server      # Go サーバーをビルド (出力: packages/server/.dist/app)
make dev-web           # web 開発サーバー (http://localhost:3000)
make dev-admin         # admin-console 開発サーバー (http://localhost:3333)
```

dev 中、ブラウザからの `/api/*` は vite proxy 経由で Go サーバー (`localhost:8080`) に転送される。

Go サーバー固有のタスクは `packages/server/Makefile` を参照 (`make -C packages/server <target>`)。

## redis

```sh
docker run -it --rm -p 6379:6379 local/redis
```

## tools

```sh
brew install golangci-lint
```

```shell
go install github.com/golang/mock/mockgen@latest
```

```shell
brew install golang-migrate
```

## db

### docker networks

```shell
docker network create -d bridge schemaspy-network

```

### schemaspy

```shell
docker build -f .docker/schemaspy/Dockerfile -t schemaspy:dev .
cd .docker/schemaspy
docker run --rm -v $(pwd)/.output:/output schemaspy:dev \
  -t pgsql \
  -host host.docker.internal \
  -port 5432 \
  -db yunne \
  -u user \
  -p password \
  -s public \
  -dp /usr/local/bin/postgresql.jar \
  -o /output

```

### dev

```shell
docker build -t go/server:dev -f ./.docker/db/dev/Dockerfile .
```

```shell
docker run -p 5432:5432 -d go/server:dev
```

### unittest

```shell
docker build -t go/server:unittest -f ./.docker/db/unittest/Dockerfile .
```

```shell
docker run -p 5433:5432 -d go/server:unittest
```

### create migration file

```shell
migrate create -dir packages/server/migrations/ -ext .sql ${sql_names}
```
