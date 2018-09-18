default: build example/wasm_exec.js

build:
	GOOS=js GOARCH=wasm go build -o example/openssl.wasm main.go

example/wasm_exec.js:
	curl -sSfLo example/wasm_exec.js "https://raw.githubusercontent.com/golang/go/go1.11/misc/wasm/wasm_exec.js"

publish:
	bash golang.sh
