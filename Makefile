# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
PRO_NAME=shamq
PRO_UNIX=$(PRO_NAME)_unix

all: clean test build
build:
	$(GOBUILD) -o $(PRO_NAME) -v
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(PRO_NAME)
	rm -f $(PRO_UNIX)
run:
	$(GOBUILD) -o $(PRO_NAME) -v
	./$(PRO_NAME)
deps:
	make -f ../Makefile deps