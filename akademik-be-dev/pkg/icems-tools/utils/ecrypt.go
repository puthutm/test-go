package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"

	"unsia.ac.id/akademic_be/pkg/icems-tools/auth"
)

func EncryptRsa(input any, path string) (string, error) {
	publicKey, err := auth.LoadPublicKey(path)
	if err != nil {
		return "", err
	}
	inputMarshal, err := json.Marshal(input)
	if err != nil {
		return "", err
	}
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, inputMarshal)
	if err != nil {
		return "", err
	}

	encodeString := base64.StdEncoding.EncodeToString(cipherText)
	return encodeString, nil
}
