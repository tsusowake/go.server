# go.server monorepo — root tasks.
#
# Layout:
#   openapi/            … 単一スキーマ (source of truth)
#   packages/server     … Go API サーバー
#   packages/web        … 一般ユーザー向け Web (React Router v7 framework)
#   packages/admin-console … 管理者向け Web

SERVER_DIR := packages/server

# ---- infra (docker compose) ----
.PHONY: run
run:
	docker compose -f compose.yaml up -d

.PHONY: build-nocache
build-nocache:
	docker compose -f compose.yaml build --no-cache

# ---- code generation (schema-first) ----
# openapi/ の分割ファイル (paths/ schemas/) を 1 ファイルへバンドル。
# gen-oapi / gen-api はこのバンドル済み spec からコードを生成する。
.PHONY: bundle
bundle:
	bun run bundle

# Go の型 (oapi-codegen) を生成。packages/server を cwd にする必要がある。
.PHONY: gen-oapi
gen-oapi: bundle
	$(MAKE) -C $(SERVER_DIR) gen-oapi

.PHONY: lint-oapi
lint-oapi:
	$(MAKE) -C $(SERVER_DIR) lint-oapi

# TypeScript クライアント (web / admin) を orval で生成。
.PHONY: gen-api
gen-api: bundle
	bun run gen:api

# OpenAPI から全コードを再生成 (bundle は 1 回だけ実行される)。
.PHONY: gen
gen: gen-oapi gen-api

# ---- server ----
.PHONY: build-server
build-server:
	$(MAKE) -C $(SERVER_DIR) build

.PHONY: test-server
test-server:
	$(MAKE) -C $(SERVER_DIR) test

# ---- frontend (bun workspaces) ----
.PHONY: install
install:
	bun install

.PHONY: dev-web
dev-web:
	bun run dev:web

.PHONY: dev-admin
dev-admin:
	bun run dev:admin

.PHONY: build-web
build-web:
	bun run build:web

.PHONY: build-admin
build-admin:
	bun run build:admin

.PHONY: typecheck
typecheck:
	bun run typecheck
