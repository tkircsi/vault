package services

import (
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	data := `{
    "name": "Tibcsi",
    "age": 45
}`

	enc, err := Encrpyt(data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("enc: %s\n", enc)

	dec, err := Decrypt(enc)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("dec: %s\n", dec)
}
