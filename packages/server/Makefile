# packages/server — Go API server tasks.
# OpenAPI スキーマはリポジトリルート (../../openapi) を参照する。

.PHONY: build
build:
	go build -o .dist/app apps/app/main.go

.PHONY: build-race
build-race:
	go build -race -o .dist/app-race apps/app/main.go

.PHONY: clean
clean:
	rm -rf .dist

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: test
test:
	go test -cover ./...

.PHONY: dump-schema
dump-schema:
	psqldef -U user -Wpassword -h localhost -p 5432 yunne --export >database/schema.sql

# 注: openapi-bundled.yaml はルートの `make bundle` (redocly) が生成する。
#     通常はルートの `make gen-oapi` が bundle → これの順で呼び出す。
.PHONY: gen-oapi
gen-oapi:
	go tool oapi-codegen --config ../../openapi/cfg.yaml ../../openapi/openapi-bundled.yaml

.PHONY: lint-oapi
lint-oapi:
	@docker run --rm -v $(PWD)/../..:/spec redocly/cli lint /spec/openapi/openapi.yaml
