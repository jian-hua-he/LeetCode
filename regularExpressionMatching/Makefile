DOCKER_IMAGE ?= golang:1.12

FOLDER_NAME = $(shell basename $(CURDIR))
DOCKER_VOLUME_PATH = /go/src/$(FOLDER_NAME)
BASE_COMMAND = docker run -it -v $(PWD):$(DOCKER_VOLUME_PATH) --rm $(DOCKER_IMAGE)

check-func-name:
ifndef FUNC_NAME 
	$(error Parameter FUNC_NAME is undefined)
endif

.PHONY: bash
bash:
	$(BASE_COMMAND) bash

.PHONY: go-version
go-version:
	$(BASE_COMMAND) go version

.PHONY: go-tfunc
go-tfunc: check-func-name
	$(BASE_COMMAND) bash -c "cd $(DOCKER_VOLUME_PATH) && go test -run $(FUNC_NAME)"

.PHONY: go-test
go-test:
	$(BASE_COMMAND) bash -c "cd $(DOCKER_VOLUME_PATH) && go test ./..."
