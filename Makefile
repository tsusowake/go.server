.PHONY: build
build:
	go build -o .dist/app main.go

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
	rm -rf internal/database/generated