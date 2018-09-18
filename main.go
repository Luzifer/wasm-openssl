package main

import (
	"fmt"
	"syscall/js"

	openssl "github.com/Luzifer/go-openssl"
)

func main() {
	js.Global().Set("opensslDecrypt", js.NewCallback(decrypt))
	js.Global().Set("opensslEncrypt", js.NewCallback(encrypt))

	// Trigger custom "event"
	if js.Global().Get("opensslLoaded").Type() == js.TypeFunction {
		js.Global().Call("opensslLoaded")
	}

	<-make(chan struct{}, 0)
}

func decrypt(i []js.Value) {
	if len(i) != 3 {
		println("decrypt requires 3 arguments")
		return
	}

	var (
		ciphertext = i[0].String()
		password   = i[1].String()
		callback   = i[2]
	)

	o := openssl.New()
	plaintext, err := o.DecryptString(password, ciphertext)
	if err != nil {
		callback.Invoke(nil, fmt.Sprintf("decrypt failed: %s", err))
		return
	}

	callback.Invoke(string(plaintext), nil)
}

func encrypt(i []js.Value) {
	if len(i) != 3 {
		println("encrypt requires 3 arguments")
		return
	}

	var (
		plaintext = i[0].String()
		password  = i[1].String()
		callback  = i[2]
	)

	o := openssl.New()
	ciphertext, err := o.EncryptString(password, plaintext)
	if err != nil {
		callback.Invoke(nil, fmt.Sprintf("encrypt failed: %s", err))
		return
	}

	callback.Invoke(string(ciphertext), nil)
}
