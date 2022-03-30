PACKAGE := github.com/kitt-technology/protos-common

.PHONE: all
all: test

.PHONY: deps
deps:
	GO111MODULE=off go get github.com/kitt-technology/protoc-gen-graphql
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26

.PHONY: generate
generate:
	protoc \
		--proto_path . \
		-I=. \
		-I ${GOPATH}/src \
		--go_out="module=${PACKAGE}:./" \
		--graphql_out="module=${PACKAGE}:./" \
		common.proto
