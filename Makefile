.PHONY: clean build packing

test:
	@go test ./... -coverprofile=./coverage.out & go tool cover -html=./coverage.out
	
build:
	@GOOS=linux GOARCH=amd64
	@echo ">> Building GRPC..."
	@go build -o phonebook-service-grpc ./cmd/grpc
	@echo ">> Finished"

run:
	@./phonebook-service-grpc
	