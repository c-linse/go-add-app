BINARY_NAME=main

DOCKER_IMAGE_NAME=go-add-app
DOCKER_TAG=latest

build:
	CGO_ENABLED=0 GOOS=linux go build -o ${BINARY_NAME} .

docker-build: build
	docker build -t ${DOCKER_IMAGE_NAME}:${DOCKER_TAG} .

# Clean up
clean:
	rm -f "${BINARY_NAME}"

.PHONY: build docker-build clean
