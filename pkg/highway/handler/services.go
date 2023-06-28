package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/sonrhq/core/internal/crypto"
	"github.com/sonrhq/core/internal/sfs"
	mdw "github.com/sonrhq/core/pkg/highway/middleware"
	servicetypes "github.com/sonrhq/core/x/service/types"
)

// GetCredentialCreationOptions returns the credential creation options to start account registration.
func GetCredentialCreationOptions(c *gin.Context) {
	origin := c.Param("origin")
	alias := c.Param("alias")

	// Get the service record from the origin
	record, err := mdw.GetServiceRecord(origin)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error(), "where": "GetServiceRecord"})
		return
	}

	ok, err := mdw.CheckAlias(alias)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error(), "where": "CheckAlias"})
		return
	}
	if !ok {
		c.JSON(400, gin.H{"error": "Desired alias already taken"})
		return
	}
	chal, err := protocol.CreateChallenge()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error(), "where": "CreateChallenge"})
		return
	}
	ucw, err := sfs.RandomUnclaimedWallet()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error(), "where": "RandomUnclaimedWallet"})
		return
	}
	attestionOpts, err := record.GetCredentialCreationOptions(alias, chal, ucw)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error(), "where": "GetCredentialCreationOptions"})
		return
	}

	c.JSON(200, gin.H{
		"attestion_options": attestionOpts,
		"challenge":         chal.String(),
		"origin":            origin,
		"alias":             alias,
		"ucw_id":            ucw,
		"address":           ucw,
	})
}

// RegisterCredentialForClaims registers a credential for a given Unclaimed Wallet.
func RegisterCredentialForClaims(c *gin.Context) {
	origin := c.Param("origin")
	alias := c.Param("alias")
	attestionResp := c.Query("attestation")
	challenge := c.Query("challenge")
	ucwDid := c.Query("ucw_id")
	// Get the service record from the origin
	record, err := mdw.GetServiceRecord(origin)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error(), "where": "GetServiceRecord"})
		return
	}
	ok, err := mdw.CheckAlias(alias)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error(), "where": "CheckAlias"})
		return
	}
	if !ok {
		c.JSON(400, gin.H{"error": "Desired alias already taken"})
		return
	}
	credential, err := record.VerifyCreationChallenge(attestionResp, challenge)
	if err != nil && credential == nil {
		c.JSON(500, gin.H{"error": err.Error(), "where": "VerifyCreationChallenge"})
		return
	}
	resp, err := sfs.ClaimAccount(ucwDid, crypto.SONRCoinType, credential, alias)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error(), "where": "ClaimAccount"})
		return
	}
	c.SetCookie("sonr-jwt", resp.JWT, 3600, "/", origin, false, true)
	c.JSON(200, gin.H{
		"did_document": resp.DIDDocument,
		"address":      resp.Address,
		"coin_type":    resp.CoinType,
		"success":       true,
		"account":      resp.Account,
	})
}

// GetCredentialAssertionOptions returns the credential assertion options to start account login.
func GetCredentialAssertionOptions(c *gin.Context) {
	origin := c.Param("origin")
	alias := c.Param("alias")
	record, err := mdw.GetServiceRecord(origin)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error(), "where": "GetServiceRecord"})
		return
	}

	didDoc, err := mdw.GetDID(alias)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error(), "where": "GetDID"})
		return
	}
	vms, err := servicetypes.GetCredentialDescriptorsForDIDDocument(didDoc)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error(), "where": "GetCredentialDescriptorsForDIDDocument"})
		return
	}
	chal, err := protocol.CreateChallenge()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error(), "where": "CreateChallenge"})
		return
	}
	assertionOpts, err := record.GetCredentialAssertionOptions(vms, chal)
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

// VerifyCredentialForAccount verifies a credential for a given account and returns the JWT Keyshare.
func VerifyCredentialForAccount(c *gin.Context) {
	origin := c.Param("origin")
	alias := c.Param("alias")
	assertionResp := c.Query("assertion")
	// Get the service record from the origin
	record, err := mdw.GetServiceRecord(origin)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error(), "where": "GetServiceRecord"})
		return
	}
	didDoc, err := mdw.GetDID(alias)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	cred, err := record.VerifyAssertionChallenge(assertionResp, didDoc.ListAuthenticationVerificationMethods()...)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	res, err := sfs.UnlockAccount(didDoc.Id, cred)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("sonr-jwt", res.JWT, 3600, "/", origin, false, true)
	c.SetCookie("sonr-did", didDoc.Id, 3600, "/", origin, false, true)
	c.SetCookie("sonr-alias", alias, 3600, "/", origin, false, true)
	c.JSON(200, gin.H{
		"did":     didDoc.Id,
		"didDoc":  didDoc,
		"account": res.Account,
		"success": true,
	})
}
