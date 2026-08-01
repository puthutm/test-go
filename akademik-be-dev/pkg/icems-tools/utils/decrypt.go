package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"

	"unsia.ac.id/akademic_be/pkg/icems-tools/auth"
)

func DecryptRsa[T any](inputEncode string, pathPrivateKey string) (T, error) {
	input, err := base64.StdEncoding.DecodeString(inputEncode)
	if err != nil {
		var zero T
		return zero, err
	}

	privateKey, err := auth.LoadPrivateKey(pathPrivateKey)
	if err != nil {
		var zero T
		return zero, err
	}

	decryptedData, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, input)
	if err != nil {
		var zero T
		return zero, err
	}

	var result T
	err = json.Unmarshal(decryptedData, &result)
	if err != nil {
		var zero T
		return zero, err
	}

	return result, nil
}
