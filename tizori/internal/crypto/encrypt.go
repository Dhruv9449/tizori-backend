package tizoriCrypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

// EncryptPassword encrypts a password using AES encryption with PKCS#7 padding
func EncryptPassword(password string, key []byte) (string, error) {
	plaintext := []byte(password)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// PKCS#7 padding to align plaintext to a multiple of block size (16 bytes for AES)
	padding := aes.BlockSize - len(plaintext)%aes.BlockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	plaintext = append(plaintext, padText...)

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// DecryptPassword decrypts an encrypted password using AES encryption with PKCS#7 padding
func DecryptPassword(encryptedPassword string, key []byte) (string, error) {
	ciphertext, err := base64.URLEncoding.DecodeString(encryptedPassword)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	// Remove padding
	padding := ciphertext[len(ciphertext)-1]
	ciphertext = ciphertext[:len(ciphertext)-int(padding)]

	return string(ciphertext), nil
}
