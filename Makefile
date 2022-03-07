PACKAGE := github.com/kitt-technology/protos-common

.PHONE: all
all: test

.PHONY: deps
deps:
	go get google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	GO111MODULE=off go get github.com/kitt-technology/protoc-gen-graphql
	@go mod download

.PHONY: generate
generate:
	protoc \
		--proto_path . \
		-I=. \
		-I ${GOPATH}/src \
		--go_out="module=${PACKAGE}:./" \
		--graphql_out="module=${PACKAGE}:./" \
		common.proto
