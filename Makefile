GO_VERSION=1.14.3

default: build_golang

build_golang:
	curl -fLo ./example/wasm_exec.js "https://raw.githubusercontent.com/golang/go/go${GO_VERSION}/misc/wasm/wasm_exec.js"
	GOOS=js GOARCH=wasm go build \
			 -o example/openssl.wasm \
			 main.go
