package tizoriCrypto

import (
	"encoding/base64"
	"log"
)

var AESKey []byte

func InitializeAESKey(aesKey64 string) {
	aesKey, err := base64.StdEncoding.DecodeString(aesKey64)
	if err != nil {
		log.Fatalf("error decoding AES key: %v\n", err)
	}

	AESKey = aesKey
}
