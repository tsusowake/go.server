.PHONY: clean
clean:
	rm -rf .dist

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: test
test:
	test -cover ./...