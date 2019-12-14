GO11MODULES=on
APP?=rocketlauncher
REGISTRY?=gcr.io/images
COMMIT_SHA=$(shell git rev-parse --short HEAD)

.PHONY: build
## build: build the application
build: clean
	@echo "Building..."
	@go build -o ${APP} .

.PHONY: run
## run: runs go run main.go
run:
	go run -race .

.PHONY: clean
## clean: cleans the binary
clean:
	@echo "Cleaning"
	@rm -rf ${APP}

.PHONY: test
## test: runs go test with default values
test:
	go test -v -count=1 -race ./...


.PHONY: build-tokenizer
## build-tokenizer: build the tokenizer application
build-tokenizer:
	${MAKE} -c tokenizer build

.PHONY: setup
## setup: setup go modules
setup:
	@go mod init \
		&& go mod tidy \
		&& go mod vendor

.PHONY: help
## help: Prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ m
