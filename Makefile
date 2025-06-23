APP_NAME = taskd
BIN_DIR = bin

.PHONY: all lint build test clean run

all: build

build: 
	@echo "Building..."
	@env CGO_ENABLED=0 GOOS=linux go build -o $(BIN_DIR)/$(APP_NAME) ./main.go

test:
	@echo "Running unit tests..."
	go test -v ./...

run: lint build
	@echo "Running $(APP_NAME)..."
	./$(BIN_DIR)/$(APP_NAME)

clean:
	@echo "Cleaning..."
	rm -rf $(BIN_DIR)

lint:
	golangci-lint run -v --color=always --out-format=colored-line-number

docker-build:
	@echo "Building Docker image..."
	docker build -t task-api-gateway .

docker-run: docker-build
	@echo "Running Docker container..."
	docker run -d --name task-api \
	  -p 8080:8080 \
	  task-api-gateway --log-level=debug
