package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/highlight/highlight/sdk/highlight-go"
	"github.com/sonrhq/core/internal/crypto"
	mdw "github.com/sonrhq/core/pkg/highway/middleware"
)

// CreateAccount creates a new account with a given coin type and name.
func CreateAccount(c *gin.Context) {
	coinTypeName := c.Param("coinType")
	ct := crypto.CoinTypeFromName(coinTypeName)
	jwt, err := c.Cookie("sonr-jwt")
	if err != nil {
		highlight.RecordError(c.Request.Context(), err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "status": "no jwt cookie found"})
		return
	}
	cont, err := mdw.UseControllerAccount(jwt)
	if err != nil {
		highlight.RecordError(c.Request.Context(), err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "status": "failed to use controller account"})
		return
	}

	accInfo, err := cont.CreateWallet(ct)
	if err != nil {
		highlight.RecordError(c.Request.Context(), err)
		c.JSON(500, gin.H{"error": err.Error(), "where": "CreateWallet"})
		return
	}

	c.JSON(200, gin.H{
		"jwt":      jwt,
		"account":  accInfo,
		"coinType": ct,
		"success":  true,
	})
}

// GetAccount returns an account's details given its DID.
func GetAccount(c *gin.Context) {
	did := c.Param("did")
	jwt, err := c.Cookie("sonr-jwt")
	if err != nil {
		highlight.RecordError(c.Request.Context(), err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "status": "no jwt cookie found"})
		return
	}
	cont, err := mdw.UseControllerAccount(jwt)
	if err != nil {
		highlight.RecordError(c.Request.Context(), err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "status": "failed to use controller account"})
		return
	}
	accInfo, err := cont.GetWallet(did)
	if err != nil {
		highlight.RecordError(c.Request.Context(), err)
		c.JSON(500, gin.H{"error": err.Error(), "where": "GetWallet"})
		return
	}
	c.JSON(200, gin.H{
		"claimed": true,
		"account": accInfo,
	})
}

// ListAccounts returns a list of wallet accounts given a coin type.
func ListAccounts(c *gin.Context) {
	jwt, err := c.Cookie("sonr-jwt")
	if err != nil {
		highlight.RecordError(c.Request.Context(), err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "status": "no jwt cookie found"})
		return
	}

	cont, err := mdw.UseControllerAccount(jwt)
	if err != nil {
		highlight.RecordError(c.Request.Context(), err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "status": "failed to use controller account"})
		return
	}
	c.JSON(200, gin.H{
		"jwt":      jwt,
		"accounts": cont.ListWallets(),
	})
}

// SignWithAccount signs a message with an account given its DID. Requires the JWT of their Keyshare.
func SignWithAccount(c *gin.Context) {
	jwt, err := c.Cookie("sonr-jwt")
	if err != nil {
		highlight.RecordError(c.Request.Context(), err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "status": "no jwt cookie found"})
		return
	}
	did := c.Param("did")
	msg := c.Query("msg")
	bz, err := crypto.Base64Decode(msg)
	if err != nil {
		highlight.RecordError(c.Request.Context(), err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	cont, err := mdw.UseControllerAccount(jwt)
	if err != nil {
		highlight.RecordError(c.Request.Context(), err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "status": "failed to use controller account"})
		return
	}
	sig, err := cont.SignWithWallet(did, bz)
	if err != nil {
		highlight.RecordError(c.Request.Context(), err)
		c.JSON(500, gin.H{"error": err.Error(), "where": "SignWithWallet"})
		return
	}
	c.JSON(200, gin.H{
		"jwt": jwt,
		"sig": crypto.Base64Encode(sig),
		"msg": msg,
		"did": did,
	})
}

// VerifyWithAccount verifies a signature with an account.
func VerifyWithAccount(c *gin.Context) {
jwt, err := c.Cookie("sonr-jwt")
	if err != nil {
		highlight.RecordError(c.Request.Context(), err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "status": "no jwt cookie found"})
		return
	}
	did := c.Param("did")
	msgStr := c.Query("msg")
	sigStr := c.Query("sig")
	cont, err := mdw.UseControllerAccount(jwt)
	if err != nil {
		highlight.RecordError(c.Request.Context(), err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "status": "failed to use controller account"})
		return
	}
	msg, err := crypto.Base64Decode(msgStr)
	if err != nil {
		highlight.RecordError(c.Request.Context(), err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	sig, err := crypto.Base64Decode(sigStr)
	if err != nil {
		highlight.RecordError(c.Request.Context(), err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	valid, err := cont.VerifyWithWallet(did, msg, sig)
	if err != nil {
		highlight.RecordError(c.Request.Context(), err)
		c.JSON(500, gin.H{"error": err.Error(), "where": "VerifyWithWallet"})
		return
	}
	c.JSON(200, gin.H{
		// "valid": valid,
		"msg": msgStr,
		"sig": sigStr,
		"did": did,
		"jwt": jwt,
		"valid": valid,
	})
}

// ExportWallet returns the encoded Sonr Wallet structure with an encrypted keyshare, which can be opened
// with the user's password, within Sonr Clients.
func ExportWallet(c *gin.Context) {
	isAuthenticated := mdw.IsAuthenticated(c)
	if !isAuthenticated {
		c.JSON(401, gin.H{"error": "Not authenticated"})
		return
	}
}
