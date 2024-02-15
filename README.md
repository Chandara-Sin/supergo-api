# SUPERGO-API

[![run time: Makefile](https://img.shields.io/badge/Run_Time-Makefile-e63946.svg?style=flat-square)](https://makefiletutorial.com/#makefile-cookbook)

[![Go: Version](https://img.shields.io/badge/Go_Version_1.12.1-007d9c.svg?style=flat-square)](https://makefiletutorial.com/#makefile-cookbook)

- ### [For Development](#development)
- ### [Use Docker](#docker)

## Development

Go version 1.21.1

### create `config.yml`

```yaml
app:
  host: localhost
  port: "8080"
mongo:
  uri: mongodb://localhost:27017
  user:
  password:
  db:
auth:
  sign: <signagure-value>
api:
  key: x-api-key
  public: <public-key-value>
```

### Use Docker + MongoDB

```sh
make up
```

### Down Docker + MongoDB

```sh
make down
```

### Remove Docker Volumns

```sh
make remove volume
```

### Run server

serve on `http://localhost:8080`

```sh
make run
```

## Docker

### create `config-local.yml`

```yaml
app:
  host: localhost
  port: "8080"
mongo:
  uri: mongodb://mongodb:27017
  user:
  password:
  db:
auth:
  sign:
api:
  key:
  public:
```

### Run Server + Docker

serve on `http://localhost:8080`

```sh
docker-compose up -d
```
