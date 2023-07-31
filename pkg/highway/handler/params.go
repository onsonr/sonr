package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/highlight/highlight/sdk/highlight-go"
	mdw "github.com/sonrhq/core/pkg/highway/middleware"
)

// GetCredentialAttestationParams returns the credential creation options to start account registration.
func GetCredentialAttestationParams(c *gin.Context) {
	origin := c.Param("origin")
	alias := c.Param("alias")
	ok, err := mdw.CheckAliasAvailable(alias)
	if err != nil && !ok {
		highlight.RecordError(c.Request.Context(), err)
		c.JSON(500, gin.H{"error": err.Error(), "where": "CheckAliasAvailable"})
		return
	}
	// Get the service record from the origin
	rec, err := mdw.GetServiceRecord(origin)
	if err != nil {
		highlight.RecordError(c.Request.Context(), err)
		c.JSON(500, gin.H{"error": err.Error(), "where": "GetServiceRecord"})
		return
	}
	chal, err := protocol.CreateChallenge()
	if err != nil {
		highlight.RecordError(c.Request.Context(), err)
		c.JSON(500, gin.H{"error": err.Error(), "where": "CreateChallenge"})
		return
	}
	creOpts, err := rec.GetCredentialCreationOptions(alias, chal)
	if err != nil {
		highlight.RecordError(c.Request.Context(), err)
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
func GetCredentialAssertionParams(c *gin.Context) {
	origin := c.Param("origin")
	alias := c.Param("alias")
	record, err := mdw.GetServiceRecord(origin)
	if err != nil {
		highlight.RecordError(c.Request.Context(), err)
		c.JSON(500, gin.H{"error": err.Error(), "where": "GetServiceRecord"})
		return
	}
	notok, err := mdw.CheckAliasUnavailable(alias)
	if err != nil && notok {
		highlight.RecordError(c.Request.Context(), err)
		c.JSON(500, gin.H{"error": err.Error(), "where": "CheckAliasAvailable"})
		return
	}
	assertionOpts, chal, addr, err := mdw.IssueCredentialAssertionOptions(alias, record)
	if err != nil {
		highlight.RecordError(c.Request.Context(), err)
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
