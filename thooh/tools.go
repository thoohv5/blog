// +build tools

// Package tools manages development tool versions through the module system.
//
// See https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
package tools

import (
	// grpc
	_ "github.com/golang/protobuf/protoc-gen-go@v1.4.3"
	// _ "github.com/gogo/protobuf/protoc-gen-gofast"
	// _ "github.com/gogo/protobuf/protoc-gen-gogofast"
	// _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"
	// _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger"
	_ "github.com/envoyproxy/protoc-gen-validate@v0.4.0"

	_ "github.com/axw/gocov/gocov@v1.0.0"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint@v1.37.1"
	_ "github.com/matm/gocov-html@v1.1"
	_ "golang.org/x/tools/cmd/goimports@v0.1.5"
	_ "gotest.tools/gotestsum@v1.7.0"

	// stringer
	_ "golang.org/x/tools/cmd/stringer@v0.1.5"

	_ "github.com/agiledragon/gomonkey@v2.0.1"
	// 测试
	_ "github.com/golang/mock/mockgen"
	_ "github.com/stretchr/testify"
)
