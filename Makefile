PACKAGE := github.com/kitt-technology/protos-common

.PHONE: all
all: test

.PHONY: deps
deps:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	@go install github.com/kitt-technology/protoc-gen-graphql@master
	@go mod download

.PHONY: generate
generate:
	protoc \
		--proto_path . \
		-I=. \
		--go_out="module=${PACKAGE}:./" \
		--graphql_out="module=${PACKAGE}:./" \
		common.proto
