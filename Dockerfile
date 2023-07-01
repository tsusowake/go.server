# build app
FROM golang:1.20.4 AS build
ARG GITHUB_ACCESS_TOKEN
COPY . /app
WORKDIR /app

RUN echo "machine github.com login $GITHUB_ACCESS_TOKEN" > ~/.netrc \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 make build

# run go app
FROM alpine:latest
COPY --from=build /app/.dist/app /bin/app

RUN addgroup -g 1001 yunne && adduser -D -G yunne -u 1001 sswk
USER 1001

ENTRYPOINT ["/bin/app"]