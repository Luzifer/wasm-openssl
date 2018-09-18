[![Go Report Card](https://goreportcard.com/badge/github.com/Luzifer/wasm-openssl)](https://goreportcard.com/report/github.com/Luzifer/wasm-openssl)
![](https://badges.fyi/github/license/Luzifer/wasm-openssl)
![](https://badges.fyi/github/downloads/Luzifer/wasm-openssl)
![](https://badges.fyi/github/latest-release/Luzifer/wasm-openssl)

# Luzifer / wasm-openssl

`wasm-openssl` is a WASM wrapper around [go-openssl](https://github.com/Luzifer/go-openssl) to be used in Javascript projects.

## Usage

You will need to have `wasm_exec.js` installed in your project to load the binary:

```console
$ curl -sSfLo wasm_exec.js "https://raw.githubusercontent.com/golang/go/go1.11/misc/wasm/wasm_exec.js"
```

Afterwards in your HTML you can include the `wasm_exec.js` and load the binary:

```html
<!doctype html>
<html>
  <head>
    <meta charset="utf-8">
  </head>
  <body>
    <script src="wasm_exec.js"></script>
    <script>
      function opensslLoaded() { console.log("openssl.wasm loaded") }

      const go = new Go()
      WebAssembly.instantiateStreaming(fetch("openssl.wasm"), go.importObject).then(async obj => await go.run(obj.instance))
    </script>
  </body>
</html>
```

If you have a top-level function `opensslLoaded()` defined, this will be called in the initialization of the `openssl.wasm`. This serves as a notification you do have now access to the top-level functions `encrypt` and `decrypt`:

```javascript
function decrypt(ciphertext, passphrase, callback) {}
function encrypt(plaintext, passphrase, callback) {}
```

The functions will not return anything in the moment as in the current state Go WASM support does not have return values. Instead the callback function you've provided will be called and always have two arguments: `function callback(result, error)` - The `result` will be the plaintext on `decrypt` and the ciphertext on `encrypt`. The `error` will either be `null` or a string containing details about the error. When an error occurred the `result` is `null`.
