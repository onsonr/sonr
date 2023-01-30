package token

import (
	"github.com/golang-jwt/jwt"
)

// `JWTOptions` is a struct that contains a `[]byte` named `secret`, a `jwt.SigningMethod` named
// `singingMethod`, and an `int64` named `ttl`.
// @property {[]byte} secret - This is the secret key used to sign the JWT.
// @property singingMethod - The signing method used to sign the token.
// @property {int64} ttl - Time to live in seconds.
type JWTOptions struct {
	secret        []byte
	singingMethod jwt.SigningMethod
	ttl           int64
}

var (
	options JWTOptions = JWTOptions{}
)

// A method that returns a `JWTOptions` struct.
func (opts *JWTOptions) DefaultTestConfig() JWTOptions {

	return JWTOptions{
		// using rsa pub key as secret
		secret:        []byte("MIICXAIBAAKBgQCGcXG10nJ5I7OCma6hoH3VMv9OmVV4f8DpSd3Z1I2ud4GpBKFRgCs3AeGLOF87qpGdPh6lCugH3IihKy+vuhISpiploBNyGXEAdHbKyNOCHKvEuOxe9MUEE5oLYvE8qv5LPq5G971PC1HGCwPCglxLXHxzAUnOYsJkFekoWeeBIQIDAQABAoGAMVUPVJiUSL9A73tVCRnLEqBT7pN1OXInZ3MjZPsJwis3+L0qNK3DLbwS9vMIfuxn4jsZI5aM7dWOjRU7uk+csZKPqlvvPyToq8eEBcXWz5R2XpIvzrvucVI4vZMKGgZi8YlzvLnTqdfaGy/AqFDqvmU7fX/aQIu+8bC+DZ1XP8ECQQDJG07O+7hPNtsAAC4lATuIx75SCn+9tHj11I3X7pSV3em+co3XmTYIJsX/G2ehArbUyg5ED8OvjbfxD5qCQT03AkEAqyPm1OhtcI1tq9JGsM1y6av0mM4atEjHvsrxaN0Lt9HyUP8xD52ARxYEilVytYsnArzBGv5sZz7rVpZ2jocgZwJAPjD2tyW7Aqw5H4/utTzjV1JF9gMPK/Bis8surkc2pf4Bagbs/G6B+hVbh5/G9VDsj3OI491oK6MM7jxgEMXyEwJACwZvEnw+wKd7zzvmrfEuW/tl8IomkkK2C4aLctP6s0blM26dPIJLB0lV1YuXrjZetwBt+E03spcNFjDvRlNSNQJBAKd1d0+nm2ruD6vmw7tmJQ/iFIuKRMtcTKd2nVhO7TCw7ORM4tH56NWj74uioXQ1YmfFHvwZY3O0woLTa91scYk="),
		singingMethod: jwt.SigningMethodPS256,
		ttl:           3600,
	}
}

// Setting the `options` variable to the `config` variable.
func (j *JWTOptions) WithConfig(config JWTOptions) {
	options = config
}
