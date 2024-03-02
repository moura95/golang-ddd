# Golang Base Project Using DDD


![img](doc/db_image.png)

##  Download Collection Postman
doc/collection.json

## Run

## Docker
docker is configured to up the database, create the tables, insert some records and running the api.
```bash
docker-compose up
```


## Running Development Env

## Requirements
- Docker
- Golang
- Make
- Migrate


Install Dependencias

```bash
  go mod tidy
```

Start Service

```bash
  make start
```


Running Tests
```bash
go test -cover ./...
```

