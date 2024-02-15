package middleware

type Matrix struct {
	HomeServer   string
	sharedSecret string
}

func NewMatrix(homeServer, sharedSecret string) *Matrix {
	return &Matrix{
		HomeServer:   homeServer,
		sharedSecret: sharedSecret,
	}
}
