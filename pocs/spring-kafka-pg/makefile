.PHONY: build clean test run docker-build docker-run help

# Default environment
ENV ?= default

# Environment-specific config files
env-config:
	$(eval include env/$(ENV).env)
	@echo "Using environment: $(ENV)"

# Default target
help:
	@echo "Available targets:"
	@echo "  build              - Build the application"
	@echo "  clean              - Clean build artifacts"
	@echo "  refresh            - Refresh dependencies"
	@echo "  spotless           - Run spotless"
	@echo "  test               - Run tests"
	@echo "  run ENV=<env>      - Run the application (env: defaul,t dev, sit, uat, prod)"
	@echo "  docker-build       - Build Docker image"
	@echo "  docker-run ENV=<env> - Run application in Docker"

# Build the application
build:
	./gradlew clean build

# Clean build artifacts
clean:
	./gradlew clean

refresh:
	./gradlew clean build --refresh-dependencies

spa:
	./gradlew spotlessApply

# Run tests
test:
	./gradlew test

# Run the application
run: env-config
	java -agentlib:jdwp=transport=dt_socket,server=y,suspend=n,address=5005 -jar -Dspring.profiles.active=$(ENV) app/build/libs/app.jar

# Build Docker image
docker-build:
	docker build -t spring-kafka-pg .

# Run Docker container
docker-run: env-config
	docker run -p 8080:8080 spring-kafka-pg