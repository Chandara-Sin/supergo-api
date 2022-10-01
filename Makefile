run:
	go run main.go

up:
	docker compose up -d

down:
	docker compose down

remove volume:
	docker volume rm supergo-api_mongo-data