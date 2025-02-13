APP_NAME=glofox
PORT=8080
BUILD_DIR=bin

.PHONY: build run test docker-build docker-run clean fmt vet lint

build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) main.go

run: build
	$(BUILD_DIR)/$(APP_NAME)

test:
	go test ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	golangci-lint run

docker-build:
	docker build -t $(APP_NAME):latest .

docker-run: docker-build
	docker run --rm -p $(PORT):$(PORT) $(APP_NAME):latest

docker-clean:
	docker rmi $(APP_NAME):latest || true

clean:
	rm -rf $(BUILD_DIR)