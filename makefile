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
pgdb:
	docker-compose up -d pgdb
pull:
	git pull origin master
push:
	git push origin master
start:
	podman-compose -f produciton.yml up --build -d
stop:
	podman-compose -f produciton.yml down