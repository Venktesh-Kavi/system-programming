APP_NAME := gaws
SOURCE := cmd/main.go

export AWS_ACCESS_KEY := ""
export AWS_SECRET_KEY := ""
export AWS_REGION := ap-south-1

all: build

.PHONY: build
build:
	@echo "building ${APP_NAME}"
	@go build -o ${APP_NAME} ${SOURCE}

.PHONY: run
run:
	@echo "running ${APP_NAME}"
	@go run ${SOURCE}

.PHONY: help
help:
	@echo "make build - build the application"
	@echo "make run - run the application"
	@echo "make help - display this help message"

.PHONY: clean
clean:
	@rm -rf gaws