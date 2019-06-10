DOCKER_PORT = 80
HOST_PORT = 6060

FOLDER_NAME ?= $(shell basename $(CURDIR))
DOCKER_VOLUME_PATH ?= /go/src/$(FOLDER_NAME)
DOCKER_IMAGE = golang:1.11-rc

BASE_COMMAND ?= docker run -it -v $(PWD):$(DOCKER_VOLUME_PATH) --rm -p $(HOST_PORT):$(DOCKER_PORT) $(DOCKER_IMAGE)

PACKAGES ?= $(shell go list ./... | grep -v /vendor/)

.PHONY: bash
bash:
	$(BASE_COMMAND) bash

.PHONY: go-version
go-version:
	$(BASE_COMMAND) go version

.PHONY: go-test
go-test:
	$(BASE_COMMAND) bash -c "cd $(DOCKER_VOLUME_PATH) && go test ./..."

.PHONY: go-list
go-list:
	$(BASE_COMMAND) bash -c "cd $(DOCKER_VOLUME_PATH) && go list ./..."

.PHONY: go-doc
go-doc:
	$(BASE_COMMAND) godoc -http=0.0.0.0:80