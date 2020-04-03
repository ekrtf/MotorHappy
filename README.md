Automobile API
==============

Open data for automobiles.

## Setup

```
docker run --name automobile-api-postgres -e POSTGRES_PASSWORD=pw -p 5454:5432 -d postgres:11
go run main.go
```

## Env

### Local

```
POSTGRES_PORT=5454
```
