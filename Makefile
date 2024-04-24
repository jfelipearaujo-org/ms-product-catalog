# type 'make help' to visualize the help information

##@ Utility
help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ CI/CD
build: ## Build the application to the output folder (default: ./buil/main)
	@echo "Building..."	
	@go build -o build/main cmd/api/main.go

build-docker: ## Build a container image and add the version and latest tag
	@if command -v docker > /dev/null 2>&1 && docker-buildx version > /dev/null 2>&1; then \
		echo "Building..."; \
		docker-buildx build -t jose.araujo/api-transactions:latest -t jose.araujo/api-transactions:$$(git rev-parse --short HEAD) .; \
	else \
		read -p "Docker Buildx is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        brew install docker-buildx; \
	        echo "Building..."; \
			docker-buildx build -t jose.araujo/api-transactions:latest -t jose.araujo/api-transactions:$$(git rev-parse --short HEAD) .; \
	    else \
	        echo "You chose not to install Docker Buildx. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

clean: ## Clean the binary
	@echo "Cleaning..."
	@rm -f build/main

sec: ## Security checker
	@if command -v gosec > /dev/null; then \
		echo "Analyzing..."; \
		gosec ./...; \
	else \
		read -p "gosec is not installed on your machine. Do you want to install it? [Y/n] " choice; \
		if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
			go install github.com/securego/gosec/v2/cmd/gosec@latest; \
			echo "Analyzing..."; \
			gosec ./...; \
		else \
			echo "You chose not to intall gosec. Exiting..."; \
			exit 1; \
		fi; \
	fi

lint: ## Go Linter
	@if command -v golangci-lint > /dev/null; then \
		echo "Analyzing..."; \
		golangci-lint run ./...; \
	else \
		read -p "golangci-lint is not installed on your machine. Do you want to install it? [Y/n] " choice; \
		if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
			go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.56.2; \
			echo "Analyzing..."; \
			golangci-lint run ./...; \
		else \
			echo "You chose not to intall golangci-lint. Exiting..."; \
			exit 1; \
		fi; \
	fi

##@ Runner
run: ## Run the application
	make build
	@if test ! -f .env; then \
		make env; \
	fi
	@./build/main;

##@ Testing
test: ## Test the application
	@echo "Testing..."
	@go test -race -count=1 ./internal/... -coverprofile=coverage.out

cover: ## View the coverage
	@echo "Analyzing coverage..."
	@go tool cover -func=coverage.out

test-cover: ## Run the tests and view the coverage
	make test && make cover

bench: ## Run benchmarks
	@echo "Benchmarking..."
	@go test -cpu=1,2,4,6,8,16 -benchmem -bench=. ./internal/features/transactions/handlers/create_transaction -run=^Benchmark

load: ## Run the load test using k6
	@if command -v k6 > /dev/null; then \
		echo "Generating..."; \
		k6 run tests/index.js; \
	else \
		read -p "k6 is not installed on your machine. Do you want to install it? [Y/n] " choice; \
		if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
			brew install k6; \
			echo "Generating..."; \
			k6 run tests/index.js; \
		else \
			echo "You chose not to intall k6. Exiting..."; \
			exit 1; \
		fi; \
	fi

##@ Developing
env: ## Create the .env file based on example
	@echo "Generating..."
	@cp .env.example .env

docker-up: ## Run the containers
	@if command -v docker compose > /dev/null; then \
		docker compose up -d; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up -d; \
	fi

docker-down: ## Shutdown containers
	@if command -v docker compose > /dev/null; then \
		docker compose down; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi

watch: ## Live reload using air
	@if test ! -f .env; then \
		make env; \
	fi

	@if command -v air > /dev/null; then \
	    air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/cosmtrek/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

##@ Auto generated files
gen-mocks: ## Gen mock files using mockery
	@if command -v mockery > /dev/null; then \
		echo "Generating..."; \
		mockery; \
	else \
		read -p "Go 'mockery' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
		if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
			go install github.com/vektra/mockery/v2@latest; \
			echo "Generating..."; \
			mockery; \
		else \
			echo "You chose not to intall mockery. Exiting..."; \
			exit 1; \
		fi; \
	fi

gen-docs: ## Gen Swagger docs using swag
	@if command -v swag > /dev/null; then \
		swag init --parseDependency -g main.go -d cmd/api,internal; \
	else \
		read -p "Go 'swag' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
		if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
			go install github.com/swaggo/swag/cmd/swag@latest; \
			swag init --parseDependency -g main.go -d cmd/api,internal; \
		else \
			echo "You chose not to intall swag. Exiting..."; \
			exit 1; \
		fi; \
	fi

gen-pkg-docs: ## Gen Package docs using gomarkdoc
	@if command -v gomarkdoc > /dev/null; then \
		gomarkdoc --output '{{.Dir}}/README.md' ./...; \
	else \
		read -p "Go 'gomarkdoc' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
		if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
			go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest; \
			gomarkdoc --output '{{.Dir}}/README.md' ./...; \
		else \
			echo "You chose not to intall gomarkdoc. Exiting..."; \
			exit 1; \
		fi; \
	fi

fmt-docs: ## Format generated Swagger docs using swag
	@if command -v swag > /dev/null; then \
		swag fmt -d internal -g cmd/api/main.go; \
	else \
		read -p "Go 'swag' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
		if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
			go install github.com/swaggo/swag/cmd/swag@latest; \
			swag fmt -d internal -g cmd/api/main.go; \
		else \
			echo "You chose not to intall swag. Exiting..."; \
			exit 1; \
		fi; \
	fi

.PHONY: build run test clean
