build:
	@go build -o bin/todo-auth

run: build
	@./bin/todo-auth

test:
	@go test -v ./...