GOPACKAGES = $(shell go list ./...  | grep -v /vendor/)

default: build

build:
		protoc -I. --go_out=plugins=gprc:./ pb/logsvc.proto

test: test-unit

test-unit:
		@go test $(GOPACKAGES)

test-integration:
		@go test -tags=integration
