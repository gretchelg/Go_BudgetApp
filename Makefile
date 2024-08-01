# Make commands for easier local deployment

# this exports the contents of .env file to make it available to our app
include .env
export

install:
	go mod tidy

build: install
	go build -o bin/

run: install
	go run main.go