GOPACKAGES = $(shell go list ./...  | grep -v /vendor/)
BIN_DIR := $(GOPATH)/bin
DEP := $(BIN_DIR)/dep

default: server

server: deps
		@go run cmd/server/main.go

client: deps
		@go run cmd/client/main.go

$(DEP):
		@go get -u github.com/golang/dep/cmd/dep

deps: $(DEP)
		@dep ensure

test: deps test-unit test-integration

test-unit: deps
		@go test $(GOPACKAGES)

test-integration: deps
		@go test -tags=integration

bench: deps
		@go test -tags=integration -bench=.

.PHONY:	\
				server \
				client \
				test \
				test-unit \
				test-integration \
				bench
