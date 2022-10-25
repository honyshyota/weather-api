.PHONY: test
test:
		go test -v -race -timeout 30s ./...

.PHONY: docker
docker:
		sudo docker-compose up  --remove-orphans --build

.DEFAULT_GOAL := docker