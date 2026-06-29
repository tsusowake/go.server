# build app
# Build context is the repository root. The Go module now lives in
# packages/server, so build from there.
FROM golang:1.26.4 AS build
ARG GITHUB_ACCESS_TOKEN
COPY . /app
WORKDIR /app/packages/server

RUN echo "machine github.com login $GITHUB_ACCESS_TOKEN" > ~/.netrc \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 make build

# run go app
FROM alpine:latest
COPY --from=build /app/packages/server/.dist/app /bin/app

RUN addgroup -g 1001 yunne && adduser -D -G yunne -u 1001 sswk
USER 1001

ENTRYPOINT ["/bin/app"]