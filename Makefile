.PHONY: build test lint generate clean docker run-local run stop help gendocs

.EXPORT_ALL_VARIABLES:
HOST=127.0.0.1
PORT=8080
TIMEOUT=30
CACHE_TTL=600
DATABASE_HOST=127.0.0.1 
DATABASE_PORT=5432
DATABASE_USERNAME=zippoapps             
DATABASE_PASSWORD=zippoapps
DATABASE_NAME=zippoapps

build: # build the app
	go build -o  bin/zippo ./cmd/main.go

lint: # lint the project
	golangci-lint run

test: # test whole project
	go test -race -cover ./pkg/...

generate: # re-generate boilerplate code
	go generate ./...

clean:
	rm -rf bin

docker:
	docker build -t zippo-app .

run-local:
	go build -o bin/zippo ./cmd/main.go
	./bin/zippo

run:
	docker-compose up --build  -d

stop:
	docker-compose down
	
help: # prints all possible targets
	@grep '^[^#[:space:].].*:' Makefile

gendocs:
	@swag init --parseDependency --parseInternal --parseDepth 4 -g cmd/main.go