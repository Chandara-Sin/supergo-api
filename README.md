# SUPERGO-API

- ### [For Development](#development)
- ### [Use Docker](#docker)

## Development

This is project use MakeFile

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
  public: <public-key-vlaue>
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
