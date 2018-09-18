default: build wasm_exec.js

build:
	GOOS=js GOARCH=wasm go build -o main.wasm main.go
	gzip -c main.wasm > main.wasm.gz

wasm_exec.js:
	curl -sSfLo wasm_exec.js "https://raw.githubusercontent.com/golang/go/go1.11/misc/wasm/wasm_exec.js"

publish:
	bash golang.sh
