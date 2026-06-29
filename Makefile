.PHONY: build
build:
	go build -o .dist/app apps/app/main.go

.PHONY: build
build-race:
	go build -race -o .dist/app-race main.go

.PHONY: clean
clean:
	rm -rf .dist

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: test
test:
	go test -cover ./...

.PHONY: run
run:
	docker compose -f compose.yaml up -d

.PHONY: build-nocache
build-nocache:
	docker compose -f compose.yaml build --no-cache

.phony: dump-schema
dump-schema:
	psqldef -U user -Wpassword -h localhost -p 5432 yunne --export >sqlc/schema.sql

.phony: gen-sqlc
gen-sqlc:
	docker pull sqlc/sqlc
	docker run --rm -v .:/src -w /src sqlc/sqlc generate

.phony: clean-gen-sqlc
clean-gen-sqlc:
	rm -rf database/generated

.PHONY: gen-oapi
gen-oapi:
	go tool oapi-codegen --config openapi/cfg.yaml openapi/openapi.yaml

.PHONY: lint-oapi
lint-oapi:
	@docker run --rm -v $(PWD):/spec redocly/cli lint /spec/openapi/openapi.yaml