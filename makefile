#!/bin/bash
help:
	@echo "Go Full"
	@echo "https://github.com/piovani/go_full"
	@echo "-----------------------------------------------"
	@echo "COMMANDS:                                      "
	@echo "make help    # prints usage info               "
	@echo "make migrate # execute migrates in database    "
	@echo "make rest    # start ruining API Rest          "

migrate:
	go run main.go migrate

rest:
	go run main.go rest