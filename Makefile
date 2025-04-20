init:
	docker-compose up --build

stop:
	docker-compose down

run:
	docker-compose up

clear:
	docker-compose down --volumes --remove-orphans

doc:
	godoc -http=:6060