PROJECT = $(shell basename $(CURDIR))
REVISION ?= $(shell git rev-parse --short HEAD)
PORT ?= 8081
DOCKER_REPO ?= jboursiquot

.PHONY: build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/$(PROJECT) .;
	docker build --no-cache -t $(PROJECT) .

run:
	docker run -i -t --rm -p=$(PORT):$(PORT) --name="$(PROJECT)" $(PROJECT)

stop:
	docker stop $(PROJECT); docker rm $(PROJECT)

tag: tag-latest tag-version

tag-latest:
	docker tag $(PROJECT) $(DOCKER_REPO)/$(PROJECT):latest

tag-version:
	docker tag $(PROJECT) $(DOCKER_REPO)/$(PROJECT):$(REVISION)

publish: publish-latest publish-version

publish-latest: tag-latest
	docker push $(DOCKER_REPO)/$(PROJECT):latest

publish-version: tag-version
	docker push $(DOCKER_REPO)/$(PROJECT):$(REVISION)

release: build publish