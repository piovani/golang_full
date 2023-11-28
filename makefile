GOPATH = $(shell go env GOPATH)
# GOPACKAGES = $(shell go list ./... | grep -v /vendor/)

help:
	@echo "Go Full"
	@echo "https://github.com/piovani/go_full"
	@echo "-------------------------------------------------"
	@echo "COMMANDS:                                        "
	@echo "make help      # prints usage info               "
	@echo "make migrate   # execute migrates in database    "
	@echo "make init      # preprer the application to work "
	@echo "make rest      # start ruining API Rest          "
	@echo "make all-foles # print all files in bucket s3    "

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

all-files:
	docker exec go_full_s3 awslocal s3api list-objects --bucket=my-bucket

mock:
	${GOPATH}/bin/mockgen -source=./infra/storage/contract.go -destination=./infra/test/mock/infra/storage.go -package=mock
	${GOPATH}/bin/mockgen -source=./domain/entity/student.go -destination=./infra/test/mock/repositories/student.go -package=mock
	${GOPATH}/bin/mockgen -source=./infra/storage/repository.go -destination=./infra/test/mock/repositories/file.go -package=mock

cover:
	go test ./... -coverprofile=covarage.out -covermode=count
	go tool cover -func=covarage.out