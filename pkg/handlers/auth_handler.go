package handlers

import "github.com/labstack/echo/v4"

// AuthAPI is a handler for the auth module
var AuthAPI = authAPI{}

// authAPI is a handler for the auth module
type authAPI struct{}

// CheckIdentifier checks an identifier
func (h authAPI) CheckIdentifier(c echo.Context) error {
	return nil
}

// GetAccountInfo returns account info
func (h authAPI) GetAccountInfo(c echo.Context) error {
	return nil
}

// RefreshToken refreshes a token
func (h authAPI) RefreshToken(c echo.Context) error {
	return nil
}

// SignMessage signs a message
func (h authAPI) SignMessage(c echo.Context) error {
	return nil
}

// VerifySignature verifies a signature
func (h authAPI) VerifySignature(c echo.Context) error {
	return nil
}

// VerifyToken verifies a token
func (h authAPI) VerifyToken(c echo.Context) error {
	return nil
}
