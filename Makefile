run:
	go run main.go

up:
	docker compose up db -d

down:
	docker compose down

remove volume:
	rm -r db-data