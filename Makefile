TARGET_FILE ?= playground.go
DOCKER_IMAGE ?= golang:1.11-rc

FOLDER_NAME = $(shell basename $(CURDIR))
DOCKER_VOLUME_PATH = /go/src/$(FOLDER_NAME)
BASE_COMMAND = docker run -it -v $(PWD):$(DOCKER_VOLUME_PATH) --rm $(DOCKER_IMAGE)

.PHONY: bash
bash:
	$(BASE_COMMAND) bash

.PHONY: go-version
go-version:
	$(BASE_COMMAND) go version

.PHONY: go-run
go-run:
	$(BASE_COMMAND) go run $(DOCKER_VOLUME_PATH)/$(TARGET_FILE)

.PHONY: go-test
go-test:
	$(BASE_COMMAND) bash -c "cd $(DOCKER_VOLUME_PATH) && go test ./..."