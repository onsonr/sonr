package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-webauthn/webauthn/protocol"

	mdw "github.com/sonrhq/core/internal/highway/middleware"
)

// GetCredentialAttestationParams returns the credential creation options to start account registration.
//
// @Summary Get credential attestation parameters
// @Description Returns the credential creation options to start account registration.
// @Accept  json
// @Produce  json
// @Param origin path string true "Origin"
// @Param alias path string true "Alias"
// @Success 200 {object} map[string]interface{} "Success response"
// @Failure 500 {object} map[string]string "Error message"
// @Router /getCredentialAttestationParams/{origin}/{alias} [get]
func GetCredentialAttestationParams(c *gin.Context) {
	origin := c.Param("origin")
	alias := c.Param("alias")
	ok, err := mdw.CheckAliasAvailable(alias)
	if err != nil && !ok {
		c.JSON(500, gin.H{"error": err.Error(), "where": "CheckAliasAvailable"})
		return
	}
	// Get the service record from the origin
	rec, err := mdw.GetServiceRecord(origin)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error(), "where": "GetServiceRecord"})
		return
	}
	chal, err := protocol.CreateChallenge()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error(), "where": "CreateChallenge"})
		return
	}
	creOpts, err := rec.GetCredentialCreationOptions(alias, chal)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error(), "where": "GetCredentialCreationOptions"})
		return
	}
	c.JSON(200, gin.H{
		"origin":            origin,
		"alias":             alias,
		"attestion_options": creOpts,
		"challenge":         chal.String(),
	})
}

// GetCredentialAssertionParams returns the credential assertion options to start account login.
//
// @Summary Get credential assertion parameters
// @Description Returns the credential assertion options to start account login.
// @Accept  json
// @Produce  json
// @Param origin path string true "Origin"
// @Param alias path string true "Alias"
// @Success 200 {object} map[string]interface{} "Success response"
// @Failure 500 {object} map[string]string "Error message"
// @Router /getCredentialAssertionParams/{origin}/{alias} [get]
func GetCredentialAssertionParams(c *gin.Context) {
	origin := c.Param("origin")
	alias := c.Param("alias")
	record, err := mdw.GetServiceRecord(origin)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error(), "where": "GetServiceRecord"})
		return
	}
	notok, err := mdw.CheckAliasUnavailable(alias)
	if err != nil && notok {
		c.JSON(500, gin.H{"error": err.Error(), "where": "CheckAliasAvailable"})
		return
	}
	assertionOpts, chal, addr, err := mdw.IssueCredentialAssertionOptions(alias, record)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error(), "where": "GetCredentialAssertionOptions"})
		return
	}
	c.JSON(200, gin.H{
		"assertion_options": assertionOpts,
		"challenge":         chal.String(),
		"origin":            origin,
		"alias":             alias,
		"address":           addr,
	})
}

// GetEmailAssertionParams returns a JWT for the email controller. After it is confirmed, the user will claim one of their unclaimed Keyshares.
//
// @Summary Get email assertion parameters
// @Description Returns a JWT for the email controller. After it is confirmed, the user will claim one of their unclaimed Keyshares.
// @Accept  json
// @Produce  json
// @Param email query string true "Email"
// @Success 200 {object} map[string]interface{} "Success response"
// @Failure 500 {object} map[string]string "Error message"
// @Router /getEmailAssertionParams [get]
func GetEmailAssertionParams(c *gin.Context) {
	// email := c.Query("email")
	// ucw, err := sfs.Accounts.RandomUnclaimedWallet()
	// if err != nil {
	// 	c.JSON(500, gin.H{"error": err.Error(), "where": "RandomUnclaimedWallet"})
	// 	return
	// }
	// token, err := sfs.Claims.IssueEmailClaims(email, ucw)
	// if err != nil {
	// 	c.JSON(500, gin.H{"error": err.Error(), "where": "IssueEmailClaims"})
	// 	return
	// }
	c.JSON(500, gin.H{
		// "jwt":     token,
		// "ucw_id":  ucw,
		// "address": ucw,
	})
}
