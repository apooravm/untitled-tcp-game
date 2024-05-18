APP_NAME := untitled_game.exe

tidy:
	@echo "tidying up..."
	@go mod tidy
	@go mod vendor

build:
	go build -o ./bin/${APP_NAME} ./src/main.go

run: tidy build
	./bin/${APP_NAME}

dev: build
	./bin/${APP_NAME}
