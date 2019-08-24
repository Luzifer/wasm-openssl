GO_VERSION=1.12.9
TINYGO_VERSION=0.7.1

default: build_golang

build_golang:
	curl -fLo ./example/wasm_exec.js "https://raw.githubusercontent.com/golang/go/go${GO_VERSION}/misc/wasm/wasm_exec.js"
	GOOS=js GOARCH=wasm go build \
			 -o example/openssl.wasm \
			 main.go

build_tinygo:
	curl -fLo ./example/wasm_exec.js "https://github.com/tinygo-org/tinygo/blob/v${TINYGO_VERSION}/targets/wasm_exec.js"
	tinygo build \
		-o example/openssl.wasm \
		-target wasm \
		--no-debug \
		main.go
