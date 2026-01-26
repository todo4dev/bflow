// infrastructure/crypto/std/std.go
package std

import (
	"fmt"
	"src/port/crypto"

	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/env"
)

func Provide() {
	config := Config{
		Key: env.Get("API_CRYPTO_STD_KEY", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	}
	if err := config.Validate(); err != nil {
		panic(fmt.Errorf("crypto config validation failed: %w", err))
	}

	instance := New(&config)

	di.SingletonAs[crypto.Client](func() crypto.Client { return instance })
}
