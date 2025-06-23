APP_NAME = taskd
BIN_DIR = bin

.PHONY: all lint build test clean run

all: build

build: lint
	@echo "Building..."
	go build -o $(BIN_DIR)/$(APP_NAME) ./main.go

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
