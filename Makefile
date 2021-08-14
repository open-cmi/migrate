ROOT := $(shell pwd)

ifdef BUILD_DIR
TARGET=$(BUILD_DIR)/bin/
else
TARGET=$(ROOT)/build/
endif

.PHONY:build
build:
	go build -ldflags "-s -w" -o $(TARGET)/migrate main.go

.PHONY:test
test:
	go test ./...

.PHONY:clean
clean:
	rm -r build/*

