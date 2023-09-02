package handler

import (
	"github.com/gin-gonic/gin"

	mdw "github.com/sonrhq/core/internal/highway/middleware"
	"github.com/sonrhq/core/internal/highway/types"
	domainproxy "github.com/sonrhq/core/x/domain/client/proxy"
	identityproxy "github.com/sonrhq/core/x/identity/client/proxy"
	serviceproxy "github.com/sonrhq/core/x/service/client/proxy"
)

// RegisterEscrowIdentity returns the credential assertion options to start account login.
//
// @Summary Register escrow identity
// @Description Returns the credential assertion options to start account login.
// @Accept  json
// @Produce  json
// @Param amount query string true "Amount"
// @Param email query string true "Email"
// @Success 200 {object} map[string]interface{} "Success response"
// @Failure 500 {object} map[string]string "Error message"
// @Router /registerEscrowIdentity [get]
func RegisterEscrowIdentity(c *gin.Context) {
	origin := c.Query("amount")
	alias := c.Query("email")

	record, err := serviceproxy.GetServiceRecord(c.Request.Context(), origin)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error(), "where": "GetServiceRecord"})
		return
	}
	assertionOpts, chal, _, err := mdw.IssueCredentialAssertionOptions(alias, record)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error(), "where": "GetCredentialAssertionOptions"})
		return
	}
	c.JSON(200, gin.H{
		"assertion_options": assertionOpts,
		"challenge":         chal.String(),
		"origin":            origin,
		"alias":             alias,
	})
}

// RegisterControllerIdentity registers a credential for a given Unclaimed Wallet.
//
// @Summary Register controller identity
// @Description Registers a credential for a given Unclaimed Wallet.
// @Accept  json
// @Produce  json
// @Param origin path string true "Origin"
// @Param alias path string true "Alias"
// @Param attestation query string true "Attestation"
// @Param challenge query string true "Challenge"
// @Success 200 {object} map[string]interface{} "Success response"
// @Failure 404 {object} map[string]string "Error message"
// @Failure 412 {object} map[string]string "Error message"
// @Failure 400 {object} map[string]string "Error message"
// @Failure 401 {object} map[string]string "Error message"
// @Router /registerControllerIdentity/{origin}/{alias} [post]
func RegisterControllerIdentity(c *gin.Context) {
	origin := c.Param("origin")
	alias := c.Param("alias")
	attestionResp := c.Query("attestation")
	challenge := c.Query("challenge")
	// Get the service record from the origin
	record, err := serviceproxy.GetServiceRecord(c.Request.Context(), origin)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error(), "where": "GetServiceRecord"})
		return
	}
	credential, err := record.VerifyCreationChallenge(attestionResp, challenge)
	if err != nil && credential == nil {
		c.JSON(412, gin.H{"error": err.Error(), "where": "VerifyCreationChallenge"})
		return
	}
	cont, resp, err := mdw.PublishControllerAccount(alias, credential, origin)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "where": "PublishControllerAccount"})
		return
	}
	token, err := types.NewSessionJWTClaims(alias, cont.Account())
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error(), "where": "NewSessionJWTClaims"})
		return
	}
	c.JSON(200, gin.H{
		"tx_hash": resp.TxHash,
		"address": cont.Account().Address,
		"token":   token,
		"origin":  origin,
		"success": true,
	})
}

// SignInWithCredential verifies a credential for a given account and returns the JWT Keyshare.
//
// @Summary Sign in with credential
// @Description Verifies a credential for a given account and returns the JWT Keyshare.
// @Accept  json
// @Produce  json
// @Param origin path string true "Origin"
// @Param alias path string true "Alias"
// @Param assertion query string true "Assertion"
// @Success 200 {object} map[string]interface{} "Success response"
// @Failure 441 {object} map[string]string "Error message"
// @Failure 442 {object} map[string]string "Error message"
// @Failure 443 {object} map[string]string "Error message"
// @Failure 444 {object} map[string]string "Error message"
// @Failure 445 {object} map[string]string "Error message"
// @Router /signInWithCredential/{origin}/{alias} [post]
func SignInWithCredential(c *gin.Context) {
	origin := c.Param("origin")
	alias := c.Param("alias")
	assertionResp := c.Query("assertion")
	record, err := serviceproxy.GetServiceRecord(c.Request.Context(), origin)
	if err != nil {
		c.JSON(441, gin.H{"error": err.Error(), "where": "GetServiceRecord"})
		return
	}
	_, err = record.VerifyAssertionChallenge(assertionResp)
	if err != nil {
		c.JSON(442, gin.H{"error": err.Error(), "where": "VerifyCreationChallenge"})
		return
	}
	addr, err := domainproxy.GetEmailRecordCreator(c.Request.Context(), alias)
	if err != nil {
		c.JSON(443, gin.H{"error": err.Error(), "where": "GetEmailRecordCreator"})
		return
	}
	contAcc, err := identityproxy.GetControllerAccount(c.Request.Context(), addr)
	if err != nil {
		c.JSON(444, gin.H{"error": err.Error(), "where": "GetControllerAccount"})
		return
	}
	token, err := types.NewSessionJWTClaims(alias, contAcc)
	if err != nil {
		c.JSON(445, gin.H{"error": err.Error(), "where": "NewSessionJWTClaims"})
		return
	}
	c.JSON(200, gin.H{
		"token":   token,
		"origin":  origin,
		"address": contAcc.Address,
		"success": true,
	})
}

// SignInWithEmail registers a DIDDocument for a given email authorized jwt.
//
// @Summary Sign in with email
// @Description Registers a DIDDocument for a given email authorized jwt.
// @Accept  json
// @Produce  json
// @Param jwt query string true "JWT"
// @Success 200 {object} map[string]interface{} "Success response"
// @Failure 500 {object} map[string]string "Error message"
// @Router /signInWithEmail [post]
func SignInWithEmail(c *gin.Context) {
	// token := c.Query("jwt")
	// resp, err := sfs.Claims.ClaimWithEmail(token)
	// if err != nil {
	// 	c.JSON(500, gin.H{"error": err.Error(), "where": "ClaimWithEmail"})
	// 	return
	// }
	// c.JSON(200, mdw.StoreAuthCookies(c, resp, ""))
	c.JSON(500, gin.H{
		// "jwt":     token,
		// "ucw_id":  ucw,
		// "address": ucw,
	})
}
