function decryptResponse(plaintext) {
  console.log(["decryptResponse", plaintext])
}

function encryptResponse(ciphertext) {
  console.log(["encryptResponse", ciphertext])
  decrypt(ciphertext, "password", decryptResponse)
}

function wasmStartSuccess() {
  encrypt("Knut", "password", encryptResponse)
}

const go = new Go()
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then(async obj => await go.run(obj.instance))
