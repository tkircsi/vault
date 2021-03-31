package services

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"
	"os"
)

var (
	key []byte
	//nonce []byte
)

// const (
// 	nonceSize = 12
// )

func init() {
	k, ok := os.LookupEnv("SECRET_KEY")
	if !ok {
		log.Fatal("secret key environment variable is missing")
	}
	key = []byte(k)
	if len(key) != 32 {
		log.Fatal("the aes secret key must be 32 byte length")
	}

}

func Encrpyt(data string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatal("error generating nonce")
	}
	ciphertext := aesgcm.Seal(nonce, nonce, []byte(data), nil)
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

func Decrypt(data string) (string, error) {
	ciphertext, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce, enc := ciphertext[:aesgcm.NonceSize()], ciphertext[aesgcm.NonceSize():]
	plaintext, err := aesgcm.Open(nil, nonce, enc, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}
