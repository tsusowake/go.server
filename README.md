# go.server

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
cd .docker/schemaspy
docker build -t schemaspy:dev .

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
migrate create -dir migrations/ -ext .sql ${sql_names}
```
