DOCKER_PORT = 80
HOST_PORT = 6060

FOLDER_NAME = $(shell basename $(CURDIR))
DOCKER_VOLUME_PATH = /go/src/$(FOLDER_NAME)
DOCKER_IMAGE = golang:1.11-rc

BASE_COMMAND = docker run -it -v $(PWD):$(DOCKER_VOLUME_PATH) --rm -p $(HOST_PORT):$(DOCKER_PORT) $(DOCKER_IMAGE)

TARGET_FILE ?= playground.go

.PHONY: bash
bash:
	$(BASE_COMMAND) bash

.PHONY: go-version
go-version:
	$(BASE_COMMAND) go version

.PHONY: go-run
go-run:
	$(BASE_COMMAND) go run $(DOCKER_VOLUME_PATH)/$(TARGET_FILE)