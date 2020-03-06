#https://github.com/helm/helm/blob/master/Makefile
GIT_COMMIT = $(shell git rev-parse HEAD)
GIT_SHA    = $(shell git rev-parse --short HEAD)
GIT_TAG    = $(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)
GIT_DIRTY  = $(shell test -n "`git status --porcelain`" && echo "dirty" || echo "clean")
VERSION    = $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || echo v0)

PKG          := ./...
COV_REPORT   := coverage.out
TRACE_REPORT := trace.out
SRC          := $(shell find . -type f -name '*.go' -print | grep -v _test.go)


.PHONY: help
help:
	@echo "*******************************************************************************************************"
	@echo "  make build            - Generate a production build"
	@echo "  make run              - Run the main package in development mode"
	@echo "  make test             - Run the all the tests"
	@echo "  make cov              - Run the all the tests with coverage statistics"
	@echo "  make cov-report       - Run the all the tests with coverage statistics and generate an HTML report"
	@echo "  make trace            - Generate a trace report"
	@echo "  make optimize-report  - Generate the optimization plan of the compiler"
	@echo "  make clean            - Delete all the generate artifacts like the binary build and reports"
	@echo "  make info             - Show the project information"
	@echo "*******************************************************************************************************"

# Build a production release of the application
.PHONY: build
build:
	go mod tidy
	go build -tags release \
		-ldflags="-s -w -X 'main.version=$(VERSION)'" 

# Run the application in development
.PHONY: run
run:
	go run main.go 

# Run all the tests with coverage contage
.PHONY: test
test:
	go test -cover $(PKG) -v

# Run all the tests with coverage summary
.PHONY: cov
cov:
	go test -coverprofile=$(COV_REPORT) $(PKG) -v
	go tool cover -func=$(COV_REPORT)

# Run all the tests with coverage summary and HTML report
.PHONY: cov-report
cov-report: cov
	go tool cover -html=$(COV_REPORT)

# Generate a trace report of the tests
.PHONY: trace
trace:
	go test -trace=$(TRACE_REPORT) ./server
	go tool trace $(TRACE_REPORT)

.PHONY: bench
bench:
	go test -benchtime=1s -count=5 -benchmem -bench . ./server

profile:
	go test -cpuprofile cpu.out ./server

# Genereate the compiler optimization plan
# make optimize-report | grep "cannot inline"
# make optimize-report | grep "escapes to heap"
# make optimize-report | grep "leaking"
opt-report:
	go build -gcflags=-m=2 $(PKG) 2>&1
	

# Clean build, test, profile artifacts
.PHONY: clean
clean:
	go clean
	rm -f ozone main *.out *.test

.PHONY: info
info:
	@echo "Version:           ${VERSION}"
	@echo "Git Tag:           ${GIT_TAG}"
	@echo "Git Commit:        ${GIT_COMMIT}"
	@echo "Git Tree State:    ${GIT_DIRTY}"