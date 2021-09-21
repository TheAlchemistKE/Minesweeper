APP_PORT=8080
BUILD_DIR=$(PWD)/.build

clean:
	@rm -rf $(BUILD_DIR)/*

fmt:
	@go fmt ./...

test:
	@go test ./...

coverage: fmt
	@./coverage.sh
	@go tool cover -html=.coverage/coverage.out -o=.coverage/coverage.html

deps:
	@go mod downlaod

build-backend-api: clean
	@echo Building Local Backend API
	@mkdir -p $(BUILD_DIR)/backend-api
	@go build -o $(BUILD_DIR)/backend-api/backend-api-server cmd/httpserver/main.go

lint: fmt
	@echo Linting Project
	@revive -config revive-linter.toml -formatter stylish ./...

run-backend-api:
	@echo Running Backend API
	@$(BUILD_DIR)/backend-api/backend-api-server

build-run-backend-api: build-backend-api run-backend-api

zip-backend-api: build-backend-api
	@rm -rf backend-api.zip && cd $(BUILD_DIR)/ && zip -r ../backend-api.zip backend-api/*