package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sonrhq/core/internal/sfs"
	mdw "github.com/sonrhq/core/pkg/highway/middleware"
	"github.com/sonrhq/core/types/crypto"
	vtt "github.com/sonrhq/core/x/vault/types"
)

// CreateAccount creates a new account with a given coin type and name.
func CreateAccount(c *gin.Context) {
	coinTypeName := c.Param("coinType")
	ct := crypto.CoinTypeFromName(coinTypeName)
	jwtToken, err := c.Cookie("sonr-jwt")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to read JWT cookie"})
		return
	}
	did, err := c.Cookie("sonr-did")
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	cred, err := sfs.RetreiveCredential(jwtToken)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	acc, err := sfs.DeriveWithKeyshares(c.Request.Context(), did, cred, ct)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"account": acc,
		"success": true,
	})
}

// GetAccount returns an account's details given its DID.
func GetAccount(c *gin.Context) {
	did := c.Param("did")
	acc, err := sfs.GetAccountInfo(c.Request.Context(), did)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error(), "claimed": false})
		return
	}
	c.JSON(200, gin.H{
		"account": acc,
		"claimed": true,
	})
}

// ListAccounts returns a list of wallet accounts given a coin type.
func ListAccounts(c *gin.Context) {
	alias := c.Param("alias")
	didDoc, err := mdw.GetDID(alias)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	accVms := didDoc.ListWalletVerificationMethods()
	accInfs := make([]*vtt.AccountInfo, len(accVms))
	for i, vm := range accVms {
		acc, err := sfs.GetAccountInfo(c.Request.Context(), vm.Id)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		accInfs[i] = acc
	}
	c.JSON(200, gin.H{
		"accounts": accInfs,
	})

}

// SignWithAccount signs a message with an account given its DID. Requires the JWT of their Keyshare.
func SignWithAccount(c *gin.Context) {
	did := c.Param("did")
	msg := c.Query("msg")
	bz, err := crypto.Base64UrlToBytes(msg)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	jwtToken, err := c.Cookie("sonr-jwt")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to read JWT cookie"})
		return
	}
	cred, err := sfs.RetreiveCredential(jwtToken)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	sfs.SignWithKeyshares(c.Request.Context(), did, cred, bz)
}

// VerifyWithAccount verifies a signature with an account.
func VerifyWithAccount(c *gin.Context) {
	did := c.Param("did")
	msgStr := c.Query("msg")
	sigStr := c.Query("sig")
	msg, err := crypto.Base64UrlToBytes(msgStr)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	sig, err := crypto.Base64UrlToBytes(sigStr)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	valid, err := sfs.VerifyWithPublicKeyshare(c.Request.Context(), did, msg, sig)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"valid": valid,
		"msg":   msgStr,
		"sig":   sigStr,
		"did":   did,
	})
}
