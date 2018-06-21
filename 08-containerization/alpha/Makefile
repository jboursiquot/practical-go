PROJECT = $(shell basename $(CURDIR))
REVISION ?= $(shell git rev-parse --short HEAD)
PORT ?= 8081
DOCKER_REPO ?= jboursiquot/$(PROJECT)

default: package

build:
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/alpha .;

run:
	docker run -i -t --rm -p=$(PORT):$(PORT) --name="$(PROJECT)" $(PROJECT)

stop:
	docker stop $(PROJECT); docker rm $(PROJECT)

tag: tag-latest tag-version

tag-latest:
	docker tag $(PROJECT) $(DOCKER_REPO)/$(PROJECT):latest

tag-version:
	docker tag $(PROJECT) $(DOCKER_REPO)/$(PROJECT):$(REVISION)

publish: repo-login publish-latest publish-version

publish-latest: tag-latest
	docker push $(DOCKER_REPO)/$(PROJECT):latest

publish-version: tag-version
	docker push $(DOCKER_REPO)/$(PROJECT):$(REVISION)