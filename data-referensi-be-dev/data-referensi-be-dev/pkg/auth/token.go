package auth

import (
	"crypto/rsa"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyTokenSpesifik(tokenStr string) (*UserClaimsSpesifikRole, error) {
	publicKey, err := loadPublicKey()
	if err != nil {
		return nil, fmt.Errorf("error loading public key: %w", err)
	}

	token, err := jwt.ParseWithClaims(tokenStr, &UserClaimsSpesifikRole{}, func(token *jwt.Token) (any, error) {
		// verify the signing method
		_, ok := token.Method.(*jwt.SigningMethodRSA)
		if !ok {
			return nil, fmt.Errorf("invalid token signing method")
		}

		return publicKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	claims, ok := token.Claims.(*UserClaimsSpesifikRole)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}

func loadPublicKey() (*rsa.PublicKey, error) {
	publicKeyData, err := os.ReadFile(os.Getenv("JWT_PUBLIC_KEY_PATH"))
	if err != nil {
		return nil, err
	}

	// Parsing public key
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyData)
	if err != nil {
		return nil, err
	}

	return publicKey, nil
}
