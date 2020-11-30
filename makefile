up:
	docker-compose up
up-d:
	docker-compose up -d
build:
	docker-compose build
down:
	docker-compose down
ps:
	docker ps -a
server:
	docker-compose up -d app
mydb:
	docker-compose up -d mydb
pgdb:
	docker-compose up -d pgdb