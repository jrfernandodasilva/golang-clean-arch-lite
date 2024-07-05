.PHONY: help clean build run docker_build docker_run

# Binary name
BINARY_NAME=clean-arch-lite

# Colors for output help messages
GREEN := $(shell tput setaf 2)
RESET := $(shell tput sgr0)

help:
	@echo "Available commands:"
	@echo "  $(GREEN)help$(RESET): See this help"
	@echo "  $(GREEN)clean$(RESET): remove binary"
	@echo "  $(GREEN)build$(RESET): build binary"
	@echo "  $(GREEN)run$(RESET): run binary"
	@echo "  $(GREEN)docker_build$(RESET): build docker image"
	@echo "  $(GREEN)docker_run$(RESET): run docker image"

clean:
	rm -f $(BINARY_NAME)

build:
	go build -o $(BINARY_NAME) .

run: build
	./$(BINARY_NAME)

docker_build:
	docker build -t $(BINARY_NAME):unstable .

docker_run: docker_build
	docker run -it --rm $(BINARY_NAME):unstable