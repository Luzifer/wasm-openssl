# go-openssl Wrapper for WASM support

```console
$ make
GOOS=js GOARCH=wasm go build -o main.wasm main.go
gzip -c main.wasm > main.wasm.gz
curl -sSfLo wasm_exec.js "https://raw.githubusercontent.com/golang/go/go1.11/misc/wasm/wasm_exec.js"

$ ls -lh main.wasm*
-rwxr-xr-x 1 luzifer luzifer 2.8M Sep 17 16:43 main.wasm
-rw-r--r-- 1 luzifer luzifer 611K Sep 17 16:43 main.wasm.gz
```

Chrome dev console output:

```
(index):14 (2) ["encryptResponse", "U2FsdGVkX1/baq9kUCX1EmUY/XZfnz7CwqGr70vqo6g="]
(index):10 (2) ["decryptResponse", "Knut"]
```
