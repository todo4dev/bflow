// application/port/security/jwt.go
package security

import "time"

type JWTPayload struct {
	ID    string
	Email string
}

type JWTAlgorithmEnum string

const (
	JWTAlgorithm_HS256 JWTAlgorithmEnum = "HS256"
	JWTAlgorithm_RS256 JWTAlgorithmEnum = "RS256"
)

type JWTGenerateOptions struct {
	Algorithm   JWTAlgorithmEnum
	Issuer      *string
	Subject     *string
	Audience    []string
	JWTID       *string
	KeyID       *string
	ExpiresIn   *time.Duration
	NotBefore   *time.Duration
	NoTimestamp bool
}

type JWTVerifyOptions struct {
	Algorithms       []JWTAlgorithmEnum
	Issuer           *string
	Subject          *string
	Audience         []string
	JWTID            *string
	IgnoreExpiration bool
	IgnoreNotBefore  bool
	ClockTolerance   *time.Duration
	MaxAge           *time.Duration
}

type JWT interface {
	Generate(payload JWTPayload, options JWTGenerateOptions) (string, error)
	Verify(token string, options JWTVerifyOptions) (JWTPayload, error)
}
