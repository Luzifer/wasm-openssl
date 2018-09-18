function decryptResponse(plaintext, error) {
  console.log(["decryptResponse", plaintext, error])
}

function encryptResponse(ciphertext, error) {
  console.log(["encryptResponse", ciphertext, error])
  if (error === null) {
    decrypt(ciphertext, "password", decryptResponse)
  }
}

function opensslLoaded() {
  encrypt("Knut", "password", encryptResponse)
}

const go = new Go()
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then(async obj => await go.run(obj.instance))
