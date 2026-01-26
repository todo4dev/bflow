// port/crypto/client.go
package crypto

type Client interface {
	// Encrypt encrypts the plainText using the default IV
	Encrypt(plainText string, optionalIV ...string) (string, error)

	// Decrypt decrypts the cipherText using the default IV
	Decrypt(cipherText string) (string, error)
}
