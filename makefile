#!/bin/bash
help:
	@echo "Go Full"
	@echo "https://github.com/piovani/go_full"
	@echo "-----------------------------------------------"
	@echo "COMMANDS:                                      "
	@echo "make help    # prints usage info               "
	@echo "make migrate # execute migrates in database    "
	@echo "make init    # preprer the application to work "
	@echo "make rest    # start ruining API Rest          "

cop-env:
	@echo "coping envs..."
	cp .env.example .env

init-docker:
	@echo "initing docker..."
	docker-compose down
	docker-compose up -d --build

migrate:
	@echo "runing migrate..."
	go run main.go migrate


init: cop-env init-docker migrate

rest:
	~/go/bin/air rest