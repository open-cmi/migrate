mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
ROOT_DIR := $(shell dirname $(mkfile_path))

.PHONY: build
build:
	cd src && go build -ldflags "-s -w" -o $(ROOT_DIR)/migrate main.go
