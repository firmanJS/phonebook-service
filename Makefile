.PHONY: clean build packing

test:
	@go test ./... -coverprofile=./coverage.out & go tool cover -html=./coverage.out
	
build:
	@GOOS=linux GOARCH=amd64
	@echo ">> Building GRPC..."
	@go build -o phonebook-service-grpc ./cmd/grpc
	@echo ">> Finished"

# Definitions
ROOT                    := $(PWD)
GO_HTML_COV             := ./coverage.html
GO_TEST_OUTFILE         := ./c.out
GOLANG_DOCKER_IMAGE     := golang:1.15
GOLANG_DOCKER_CONTAINER := goesquerydsl-container
CC_TEST_REPORTER_ID := ${CC_TEST_REPORTER_ID}
CC_PREFIX       := github.com/firmanJS/phonebook-service

#   Usage:
#       make test
coverage:
 # download CC test reported
    docker run -w /app -v ${ROOT}:/app ${GOLANG_DOCKER_IMAGE} \
        /bin/bash -c \
        "curl -L https://codeclimate.com/downloads/test-reporter/test-reporter-latest-linux-amd64 > ./cc-test-reporter"

    docker run -w /app -v ${ROOT}:/app ${GOLANG_DOCKER_IMAGE} go test ./... -coverprofile=${GO_TEST_OUTFILE}
    docker run -w /app -v ${ROOT}:/app ${GOLANG_DOCKER_IMAGE} go tool cover -html=${GO_TEST_OUTFILE} -o ${GO_HTML_COV}

	# handle custom prefix
    $(eval PREFIX=${CC_PREFIX})
	ifdef prefix
    	$(eval PREFIX=${prefix})
	endif
    # upload data to CC
    docker run -w /app -v ${ROOT}:/app \
        -e CC_TEST_REPORTER_ID=${CC_TEST_REPORTER_ID} \
        ${GOLANG_DOCKER_IMAGE} ./cc-test-reporter after-build --prefix ${PREFIX}


run:
	@./phonebook-service-grpc
	