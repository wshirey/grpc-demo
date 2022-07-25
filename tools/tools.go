//go:build tools
// +build tools

package tools

// The import statements allow the go command to precisely record the version information for your tools in
// your module's go.mod, while the //go:build tools build constraint prevents your normal builds from
// actually importing your tools.
import (
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
