package handler

import (
	"crypto/rand"
	"fmt"

	"github.com/go-webauthn/webauthn/protocol"

	"github.com/gofiber/fiber/v2"
	"github.com/sonrhq/core/internal/local"
	"github.com/sonrhq/core/internal/protocol/middleware"
	"github.com/sonrhq/core/x/identity"
	"github.com/sonrhq/core/x/service/types"
)

// ChallengeLength - Length of bytes to generate for a challenge.¡¡
const ChallengeLength = 32


func GetService(c *fiber.Ctx) error {
	q := middleware.ParseQuery(c)
	service, err := q.GetService()
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"success": true,
		"service": service,
	})
}

func ListServices(c *fiber.Ctx) error {
	serviceList, err := local.Context().GetAllServices(c.Context())
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"success":  true,
		"services": serviceList,
	})
}

func GetServiceAttestion(c *fiber.Ctx) error {
	q := middleware.ParseQuery(c)
	service, err := q.GetService()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
			"origin": q.Origin(),
			"alias": q.Alias(),
		})
	}

	ucw, err := local.Context().OldestUnclaimedWallet(c.Context())
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
			"origin": q.Origin(),
			"alias": q.Alias(),
		})
	}

	wc := identity.LoadClaimableWallet(ucw)
	chal, err := wc.IssueChallenge()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
			"origin": q.Origin(),
			"alias": q.Alias(),
		})
	}
	opts, err := service.GetCredentialCreationOptions(q.Alias(), chal, wc.Address(), q.IsMobile())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
			"origin": q.Origin(),
			"alias": q.Alias(),
		})
	}
	return c.JSON(fiber.Map{
		"alias":             q.Alias(),
		"attestion_options": opts,
		"origin":            q.Origin(),
		"challenge":         string(chal),
		"ucw_id":            int(ucw.Id),
	})

}

func VerifyServiceAttestion(c *fiber.Ctx) error {
	q := middleware.ParseQuery(c)
	if !q.HasAttestion() {
		return c.Status(400).SendString("Missing attestion.")
	}

	// Get the origin from the request.
	service, err := q.GetService()
	if err != nil {
		return c.SendStatus(fiber.ErrNotFound.Code)
	}
	claims, err := q.GetWalletClaims()
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	// Checking if the credential response is valid.
	cred, err := service.VerifyCreationChallenge(q.Attestion(), q.Challenge())
	if err != nil {
		return c.Status(403).SendString(fmt.Sprintf("Failed to verify attestion: %s", err.Error()))
	}

	cont, err := claims.Assign(cred, q.Alias())
	if err != nil {
		return c.Status(413).SendString(fmt.Sprintf("Failed to assign credential: %s", err.Error()))
	}
	usr := middleware.NewUser(cont, q.Alias())
	jwt, err := usr.JWT()
	if err != nil {
		return c.Status(412).SendString(fmt.Sprintf("Failed to create JWT: %s", err.Error()))
	}
	return c.JSON(fiber.Map{
		"success": true,
		"did":     cont.Did(),
		"primary": cont.PrimaryIdentity(),
		"tx_hash": cont.PrimaryTxHash(),
		"jwt":     jwt,
		"address": cont.Address(),
	})
}

func GetServiceAssertion(c *fiber.Ctx) error {
	q := middleware.ParseQuery(c)
	service, err := q.GetService()
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	doc, err := q.GetDID()
	if err != nil {
		return c.Status(405).SendString(err.Error())
	}

	chal, err := CreateChallenge()
	if err != nil {
		return c.Status(406).SendString(err.Error())
	}

	vms := doc.ListCredentialVerificationMethods()
	var creds []protocol.CredentialDescriptor
	for _, vm := range vms {
		cred, err := types.LoadCredentialFromVerificationMethod(vm)
		if err != nil {
			return c.Status(406).SendString(err.Error())
		}
		creds = append(creds, cred.CredentialDescriptor())
	}
	challenge, err := service.GetCredentialAssertionOptions(creds, chal, q.IsMobile())
	if err != nil {
		return c.Status(407).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"did":               doc.Id,
		"assertion_options": challenge,
		"origin":            q.Origin(),
	})
}

func VerifyServiceAssertion(c *fiber.Ctx) error {
	q := middleware.ParseQuery(c)
	if !q.HasAssertion() {
		return c.Status(400).SendString("Missing assertion.")
	}
	_, err := q.GetService()
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	doc, err := q.GetDID()
	if err != nil {
		return c.Status(405).SendString(err.Error())
	}
	// if err := service.VerifyAssertionChallenge(q.Assertion(), doc.ListCredentialVerificationMethods()...); err != nil {
	// 	return c.Status(403).SendString(err.Error())
	// }

	cont, err := identity.LoadController(doc)
	if err != nil {
		return c.Status(412).SendString(err.Error())
	}
	usr := middleware.NewUser(cont, doc.FindUsername())
	// Create the Claims
	jwt, err := usr.JWT()
	if err != nil {
		return c.Status(401).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"success": true,
		"did":     cont.Did(),
		"jwt":     jwt,
		"address": cont.Address(),
		"did_document": doc,
	})
}

// CreateChallenge creates a new challenge that should be signed and returned by the authenticator. The spec recommends
// using at least 16 bytes with 100 bits of entropy. We use 32 bytes.
func CreateChallenge() (challenge protocol.URLEncodedBase64, err error) {
	challenge = make([]byte, ChallengeLength)

	if _, err = rand.Read(challenge); err != nil {
		return nil, err
	}

	return challenge, nil
}
