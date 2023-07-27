package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sonrhq/core/pkg/highway/types"
)

// The function checks if a user is authenticated.
func IsAuthenticated(c *gin.Context) bool {
	_, _, _, err := fetchAuthCookies(c)
	if err != nil {
		return false
	}
	return true
}

// The function GetAuthCookies takes a gin.Context as input and returns three strings and an error.
func fetchAuthCookies(c *gin.Context) (string, string, string, error) {
	jwtToken, err := c.Cookie("sonr-jwt")
	if err != nil {
		return "", "", "", fmt.Errorf("no jwt cookie found")
	}
	did, err := c.Cookie("sonr-did")
	if err != nil {
		return "", "", "", fmt.Errorf("no did cookie found")
	}
	alias, err := c.Cookie("sonr-alias")
	if err != nil {
		return "", "", "", fmt.Errorf("no alias cookie found")
	}
	return jwtToken, did, alias, nil
}

// The function stores authentication cookies in the context.
func StoreAuthCookies(c *gin.Context, res *types.AuthenticationResult, origin string) gin.H {
	c.SetCookie("sonr-jwt", res.JWT, 1800, "/", origin, true, true)
	c.SetCookie("sonr-did", res.DID, 1800, "/", origin, true, false)
	c.SetCookie("sonr-alias", res.Alias, 1800, "/", origin, true, false)
	return gin.H{
		"success":      true,
		"account":      res.Account,
		"did_document": res.DIDDocument,
		"token":        res.JWT,
	}
}
