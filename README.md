# SUPERGO-API

## For development

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
  sign:
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
