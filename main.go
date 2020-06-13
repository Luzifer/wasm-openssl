package main

import (
	"fmt"
	"syscall/js"

	openssl "github.com/Luzifer/go-openssl/v4"
)

var defaultCG = openssl.PBKDF2SHA256

func main() {
	js.Global().Set("opensslDecrypt", js.FuncOf(decrypt))
	js.Global().Set("opensslEncrypt", js.FuncOf(encrypt))

	// Trigger custom "event"
	if js.Global().Get("opensslLoaded").Type() == js.TypeFunction {
		js.Global().Call("opensslLoaded")
	}

	<-make(chan struct{}, 0)
}

func decrypt(this js.Value, i []js.Value) interface{} {
	if len(i) != 3 {
		println("decrypt requires 3 arguments")
		return nil
	}

	var (
		ciphertext = i[0].String()
		password   = i[1].String()
		callback   = i[2]
	)

	o := openssl.New()
	plaintext, err := o.DecryptBytes(password, []byte(ciphertext), defaultCG)
	if err != nil {
		callback.Invoke(nil, fmt.Sprintf("decrypt failed: %s", err))
		return nil
	}

	callback.Invoke(string(plaintext), nil)
	return nil
}

func encrypt(this js.Value, i []js.Value) interface{} {
	if len(i) != 3 {
		println("encrypt requires 3 arguments")
		return nil
	}

	var (
		plaintext = i[0].String()
		password  = i[1].String()
		callback  = i[2]
	)

	o := openssl.New()
	ciphertext, err := o.EncryptBytes(password, []byte(plaintext), defaultCG)
	if err != nil {
		callback.Invoke(nil, fmt.Sprintf("encrypt failed: %s", err))
		return nil
	}

	callback.Invoke(string(ciphertext), nil)
	return nil
}
