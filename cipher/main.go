package main

import (
	"fmt"
	"crypto/cipher"
	"code.google.com/p/go.crypto/blowfish"
)


func main() {
	in := []byte("Hello, World!!")
	key := []byte("This is key!!")
	if block, err := blowfish.NewCipher(key); err != nil {
		// Encrypt
		enc := cipher.NewCBCEncrypter(block, make([]byte, blowfish.BlockSize))
		var encrypted []byte
		enc.CryptBlocks(encrypted, in)
		fmt.Println(encrypted)
	}
}
