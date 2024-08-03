# Make commands for easier local deployment

# this exports the contents of .env file to make it available to our app
include .env
export

# install installs any prerequisites and does cleanup
install:
	go mod tidy

# build creates an executable build at bin/
build: install lint
	go build -o bin/

run: install lint
	go run main.go

# lint runs the linter to catch syntax and quality issues
lint:
	golangci-lint run