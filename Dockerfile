# syntax=docker/dockerfile:1.4

FROM golang:1.21.1-bullseye AS build
WORKDIR /app
COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download

COPY . .
RUN CGO_ENABLED=0 go build -v -o /supergo-api

FROM scratch
ENV ENVIRONMENT=local
WORKDIR /
COPY --from=build /supergo-api /supergo-api
COPY config-${ENVIRONMENT}.yml /config.yml

EXPOSE 8080
ENTRYPOINT [ "/supergo-api" ]
