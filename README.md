[![Go Report Card](https://goreportcard.com/badge/github.com/Luzifer/wasm-openssl)](https://goreportcard.com/report/github.com/Luzifer/wasm-openssl)
![](https://badges.fyi/github/license/Luzifer/wasm-openssl)
![](https://badges.fyi/github/downloads/Luzifer/wasm-openssl)
![](https://badges.fyi/github/latest-release/Luzifer/wasm-openssl)
![](https://knut.in/project-status/wasm-openssl)

# Luzifer / wasm-openssl

`wasm-openssl` is a WASM wrapper around [go-openssl](https://github.com/Luzifer/go-openssl) to be used in Javascript projects.

**A word of warning:** This relies on the **experimental** WASM implementation in Golang. It is working but most likely will not have its final state. When the Golang implementation of WASM changes this likely will change too. As long as the WASM implementation in Go is experimental this only serves as a proof-of-concept and maybe shouldn't be used in production!

## Usage

You will need to have `wasm_exec.js` installed in your project to load the binary. This file can be found in [golang/go](https://github.com/golang/go/tree/master/misc/wasm) repository. (Make sure the version of the file matches the version of Go used to compile the WASM file.

For an embedding example see the `example` folder in this repo.

If you have a top-level function `opensslLoaded()` defined, this will be called in the initialization of the `openssl.wasm`. This serves as a notification you do have now access to the library functions:

```javascript
OpenSSL.decrypt = (ciphertext, passphrase, callback, hashAlgo = OpenSSL.SHA256, usePBKDF = true, iterations = 10000) => {...}
OpenSSL.encrypt = (plaintext, passphrase, callback, hashAlgo = OpenSSL.SHA256, usePBKDF = true, iterations = 10000) => {...}
```

The functions will not return anything in the moment as in the current state Go WASM support does not have return values. Instead the callback function you've provided will be called and always have two arguments: `function callback(result, error)` - The `result` will be the plaintext on `decrypt` and the ciphertext on `encrypt`. The `error` will either be `null` or a string containing details about the error. When an error occurred the `result` is `null`.

### Examples

```javascript
// Encryption with default parameters
OpenSSL.encrypt('test', 'test', (res, err) => console.log(res, err))
// U2FsdGVkX18r5Lf94A7Ng1iO03jPtKeM1hq8cZdrrww= null

// Decryption with default parameters
OpenSSL.decrypt('U2FsdGVkX18r5Lf94A7Ng1iO03jPtKeM1hq8cZdrrww=', 'test', (res, err) => console.log(res, err))
// test null

// Encryption with custom parameters for key derivation
OpenSSL.encrypt('test', 'test', (res, err) => console.log(res, err), OpenSSL.SHA1, true, 25000)
// U2FsdGVkX19HCUnnaJefLxuljrhoDpCdzDOuvIqiB9Q= null

// Decryption with the same custom parameters
OpenSSL.decrypt('U2FsdGVkX19HCUnnaJefLxuljrhoDpCdzDOuvIqiB9Q=', 'test', (res, err) => console.log(res, err), OpenSSL.SHA1, true, 25000)
// test null

// Decryption with non-matching custom parameters (leads to an error)
OpenSSL.decrypt('U2FsdGVkX19HCUnnaJefLxuljrhoDpCdzDOuvIqiB9Q=', 'test', (res, err) => console.log(res, err), OpenSSL.SHA1, true)
// null "decrypt failed: invalid padding"
```
