// +build tools

// Package tools manages development tool versions through the module system.
//
// See https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
package tools

import (
	// grpc
	_ "github.com/golang/protobuf/protoc-gen-go"
	// _ "github.com/gogo/protobuf/protoc-gen-gofast"
	// _ "github.com/gogo/protobuf/protoc-gen-gogofast"
	// _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"
	// _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger"
	_ "github.com/envoyproxy/protoc-gen-validate"

	_ "github.com/axw/gocov/gocov"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/matm/gocov-html"
	_ "golang.org/x/tools/cmd/goimports"
	_ "gotest.tools/gotestsum"

	// stringer
	_ "golang.org/x/tools/cmd/stringer"

	_ "github.com/agiledragon/gomonkey"
	// 测试
	_ "github.com/golang/mock/mockgen"
	_ "github.com/stretchr/testify"
)
