package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	mdw "github.com/sonr-io/sonr/internal/highway/middleware"
	"github.com/sonr-io/sonr/pkg/crypto"
)

// CreateAccount creates a new account with a given coin type and name.
//
// @Summary Create an account
// @Description Creates a new account with a given coin type and name.
// @Accept  json
// @Produce  json
// @Param   coinType path string true "Coin Type Name"
// @Success 200 {object} map[string]interface{} "Account Info"
// @Router /createAccount/{coinType} [post]
func CreateAccount(c *gin.Context) {
	coinTypeName := c.Param("coinType")
	ct := crypto.CoinTypeFromName(coinTypeName)
	jwt, err := c.Cookie("sonr-jwt")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "status": "no jwt cookie found"})
		return
	}
	cont, err := mdw.UseControllerAccount(jwt)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "status": "failed to use controller account"})
		return
	}

	accInfo, err := cont.CreateWallet(ct)
	if err != nil {
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
//
// @Summary Get an account's details
// @Description Returns an account's details given its DID.
// @Accept  json
// @Produce  json
// @Param   did path string true "DID of Account"
// @Success 200 {object} map[string]interface{} "Account Info"
// @Router /getAccount/{did} [get]
func GetAccount(c *gin.Context) {
	did := c.Param("did")
	jwt, err := c.Cookie("sonr-jwt")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "status": "no jwt cookie found"})
		return
	}
	cont, err := mdw.UseControllerAccount(jwt)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "status": "failed to use controller account"})
		return
	}
	accInfo, err := cont.GetWallet(did)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error(), "where": "GetWallet"})
		return
	}
	c.JSON(200, gin.H{
		"claimed": true,
		"account": accInfo,
	})
}

// ListAccounts returns a list of wallet accounts given a coin type.
//
// @Summary List wallet accounts
// @Description Returns a list of wallet accounts given a coin type.
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{} "Accounts List"
// @Router /listAccounts [get]
func ListAccounts(c *gin.Context) {
	jwt, err := c.Cookie("sonr-jwt")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "status": "no jwt cookie found"})
		return
	}

	cont, err := mdw.UseControllerAccount(jwt)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "status": "failed to use controller account"})
		return
	}
	c.JSON(200, gin.H{
		"jwt":      jwt,
		"accounts": cont.ListWallets(),
	})
}

// SignWithAccount signs a message with an account given its DID. Requires the JWT of their Keyshare.
//
// @Summary Sign a message with an account
// @Description Signs a message with an account given its DID. Requires the JWT of their Keyshare.
// @Accept  json
// @Produce  json
// @Param   did path string true "DID of Account"
// @Param   msg query string true "Message to Sign"
// @Success 200 {object} map[string]interface{} "Signature Info"
// @Router /signWithAccount/{did} [post]
func SignWithAccount(c *gin.Context) {
	jwt, err := c.Cookie("sonr-jwt")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "status": "no jwt cookie found"})
		return
	}
	did := c.Param("did")
	msg := c.Query("msg")
	bz, err := crypto.Base64Decode(msg)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	cont, err := mdw.UseControllerAccount(jwt)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "status": "failed to use controller account"})
		return
	}
	sig, err := cont.SignWithWallet(did, bz)
	if err != nil {
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
//
// @Summary Verify a signature with an account
// @Description Verifies a signature with an account.
// @Accept  json
// @Produce  json
// @Param   did path string true "DID of Account"
// @Param   msg query string true "Message"
// @Param   sig query string true "Signature"
// @Success 200 {object} map[string]interface{} "Verification Result"
// @Router /verifyWithAccount/{did} [post]
func VerifyWithAccount(c *gin.Context) {
	jwt, err := c.Cookie("sonr-jwt")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "status": "no jwt cookie found"})
		return
	}
	did := c.Param("did")
	msgStr := c.Query("msg")
	sigStr := c.Query("sig")
	cont, err := mdw.UseControllerAccount(jwt)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "status": "failed to use controller account"})
		return
	}
	msg, err := crypto.Base64Decode(msgStr)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	sig, err := crypto.Base64Decode(sigStr)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	valid, err := cont.VerifyWithWallet(did, msg, sig)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error(), "where": "VerifyWithWallet"})
		return
	}
	c.JSON(200, gin.H{
		// "valid": valid,
		"msg":   msgStr,
		"sig":   sigStr,
		"did":   did,
		"jwt":   jwt,
		"valid": valid,
	})
}

// ExportWallet returns the encoded Sonr Wallet structure with an encrypted keyshare, which can be opened
// with the user's password, within Sonr Clients.
//
// @Summary Export Wallet
// @Description Returns the encoded Sonr Wallet structure with an encrypted keyshare.
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{} "Wallet Export Info"
// @Router /exportWallet [get]
func ExportWallet(c *gin.Context) {

}
