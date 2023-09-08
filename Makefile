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

.PHONY: schemaspy
schemaspy:
	docker compose -f compose.schemaspy.yaml up -d
