package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"hash"
	"syscall/js"

	openssl "github.com/Luzifer/go-openssl/v4"
)

var (
	hashAlgo = []func() hash.Hash{
		sha256.New,
		sha1.New,
		md5.New,
	}
)

func main() {
	js.Global().Set("OpenSSL", map[string]interface{}{
		// Function definitions
		"decrypt": js.FuncOf(decrypt),
		"encrypt": js.FuncOf(encrypt),

		// Hash algorithm indices
		"SHA256": 0,
		"SHA1":   1,
		"MD5":    2,
	})

	// Trigger custom "event"
	if js.Global().Get("opensslLoaded").Type() == js.TypeFunction {
		js.Global().Call("opensslLoaded")
	}

	<-make(chan struct{}, 0)
}

func decrypt(this js.Value, i []js.Value) interface{} {
	if len(i) < 3 {
		println("decrypt requires at least 3 arguments")
		return nil
	}

	var (
		ciphertext string   = i[0].String()
		password   string   = i[1].String()
		callback   js.Value = i[2]
	)

	o := openssl.New()
	plaintext, err := o.DecryptBytes(password, []byte(ciphertext), getCredentialGenerator(i))
	if err != nil {
		callback.Invoke(nil, fmt.Sprintf("decrypt failed: %s", err))
		return nil
	}

	callback.Invoke(string(plaintext), nil)
	return nil
}

func encrypt(this js.Value, i []js.Value) interface{} {
	if len(i) < 3 {
		println("encrypt requires at least 3 arguments")
		return nil
	}

	var (
		plaintext string   = i[0].String()
		password  string   = i[1].String()
		callback  js.Value = i[2]
	)

	o := openssl.New()
	ciphertext, err := o.EncryptBytes(password, []byte(plaintext), getCredentialGenerator(i))
	if err != nil {
		callback.Invoke(nil, fmt.Sprintf("encrypt failed: %s", err))
		return nil
	}

	callback.Invoke(string(ciphertext), nil)
	return nil
}

func getCredentialGenerator(i []js.Value) openssl.CredsGenerator {
	var (
		algo       = hashAlgo[0]
		usePBKDF   = true
		iterations = 10000
	)

	switch len(i) {
	case 6:
		iterations = i[5].Int()
		fallthrough
	case 5:
		usePBKDF = i[4].Bool()
		fallthrough
	case 4:
		algo = hashAlgo[i[3].Int()]
	}

	if usePBKDF {
		return openssl.NewPBKDF2Generator(algo, iterations)
	}
	return openssl.NewBytesToKeyGenerator(algo)
}
