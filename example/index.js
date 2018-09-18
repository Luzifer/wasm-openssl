function decryptResponse(plaintext, error) {
  console.log(["decryptResponse", plaintext, error])
}

function encryptResponse(ciphertext, error) {
  console.log(["encryptResponse", ciphertext, error])
  if (error === null) {
    opensslDecrypt(ciphertext, "password", decryptResponse)
  }
}

function opensslLoaded() {
  opensslEncrypt("Knut", "password", encryptResponse)
}

const go = new Go()
WebAssembly.instantiateStreaming(fetch("openssl.wasm"), go.importObject).then(async obj => await go.run(obj.instance))
