ROOT := $(shell pwd)

ifdef BUILD_DIR
TARGET=$(BUILD_DIR)/bin/
else
TARGET=$(ROOT)/build/
endif

.PHONY:build
build:
	cd src && go build -ldflags "-s -w" -o $(TARGET)/migrate main.go

clean:
	rm -r build/*

