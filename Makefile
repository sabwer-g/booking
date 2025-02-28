BINARY_NAME = booking

.PHONY: up
up:
	docker-compose -p ${BINARY_NAME} -f ./deployments/docker-compose.yml up -d --build

.PHONY: down
down:
	docker-compose -p ${BINARY_NAME} -f ./deployments/docker-compose.yml down --remove-orphans

.PHONY: test
test:
	@go test -v ./...




