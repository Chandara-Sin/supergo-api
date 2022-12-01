# syntax=docker/dockerfile:1

FROM golang:1.18-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -v -o /supergo-api

FROM scratch
WORKDIR /
COPY --from=build /supergo-api /supergo-api
COPY --from=build /app/config.yml /config.yml

EXPOSE 8080
ENTRYPOINT [ "/supergo-api" ]
