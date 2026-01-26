// infrastructure/crypto/std/client.go
package std

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"src/port/crypto"
)

type Client struct {
	config *Config
}

var _ crypto.Client = (*Client)(nil)

func New(config *Config) *Client {
	return &Client{config: config}
}

func (c *Client) Encrypt(plainText string, optionalIV ...string) (string, error) {
	key := []byte(c.config.Key)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	var iv []byte
	if len(optionalIV) > 0 && optionalIV[0] != "" {
		// Assume input IV is base64url string
		decodedIV, err := base64.RawURLEncoding.DecodeString(optionalIV[0])
		if err != nil {
			return "", fmt.Errorf("invalid iv: %w", err)
		}
		iv = decodedIV
	} else {
		iv = make([]byte, 12)
		if _, err := rand.Read(iv); err != nil {
			return "", err
		}
	}

	// Encrypt
	sealed := gcm.Seal(nil, iv, []byte(plainText), nil)

	tagSize := gcm.Overhead()
	if len(sealed) < tagSize {
		return "", errors.New("ciphertext too short")
	}

	ciphertext := sealed[:len(sealed)-tagSize]
	tag := sealed[len(sealed)-tagSize:]

	encodedIV := base64.RawURLEncoding.EncodeToString(iv)
	encodedCipher := base64.RawURLEncoding.EncodeToString(ciphertext)
	encodedTag := base64.RawURLEncoding.EncodeToString(tag)

	return fmt.Sprintf("%s.%s.%s", encodedIV, encodedCipher, encodedTag), nil
}

func (c *Client) Decrypt(cipherText string) (string, error) {
	parts := strings.Split(cipherText, ".")
	if len(parts) != 3 {
		return "", errors.New("Invalid cipherText format")
	}

	iv, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return "", err
	}
	encrypted, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return "", err
	}
	tag, err := base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil {
		return "", err
	}

	key := []byte(c.config.Key)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	sealed := append(encrypted, tag...)

	plaintext, err := gcm.Open(nil, iv, sealed, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
