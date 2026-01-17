// infrastructure/jwt/golang_jwt/utils.go
package golang_jwt

import (
	"crypto/rsa"
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

func parseKeys(config GolangJWTConfig) (signKey any, verifyKey any, err error) {
	privateBytes := []byte(config.PrivateKey)
	publicBytes := []byte(config.PublicKey)

	if config.Algorithm == "RS256" {
		var rsaPrivate *rsa.PrivateKey
		var rsaPublic *rsa.PublicKey

		rsaPrivate, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
		if err != nil {
			return nil, nil, errors.New("invalid RSA private key")
		}

		rsaPublic, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
		if err != nil {
			return nil, nil, errors.New("invalid RSA public key")
		}

		return rsaPrivate, rsaPublic, nil
	}

	return privateBytes, publicBytes, nil
}
