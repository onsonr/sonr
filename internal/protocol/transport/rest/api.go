package rest

import (
	"context"
	"encoding/base64"
	"errors"
	"regexp"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/sonrhq/core/internal/protocol/packages/controller"
	"github.com/sonrhq/core/internal/protocol/packages/resolver"
	v1 "github.com/sonrhq/core/types/highway/v1"
)

// ! ||--------------------------------------------------------------------------------||
// ! ||                          Auth API Rest Implementation                          ||
// ! ||--------------------------------------------------------------------------------||

func (htt *HttpTransport) Keygen(c *fiber.Ctx) error {
	req := new(v1.KeygenRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	// Get the origin from the request.
	// uuid := req.Uuid
	origin := regexp.MustCompile(`[^a-zA-Z]+`).ReplaceAllString(req.Origin, "")
	service, _ := resolver.GetService(context.Background(), origin)
	if service == nil {
		sentry.CaptureEvent(&sentry.Event{
			Message: "Provided Service is nil",
			Extra: map[string]interface{}{
				"origin": origin,
			},
		})
		// Try to get the service from the localhost
		service, _ = resolver.GetService(context.Background(), "localhost")
	}
	// Check if service is still nil - return internal server error
	if service == nil {
		sentry.CaptureException(errors.New("Localhost not found on blockchain"))
		return c.Status(500).SendString("Internal Server Error.")
	}

	// Generate the keypair.
	cred, _ := service.VerifyCreationChallenge(req.CredentialResponse)
	// if err != nil {
	// 	return c.Status(403).SendString(err.Error())
	// }

	cont, acc, err := controller.NewController(context.Background(), controller.WithInitialAccounts(req.InitialAccounts...), controller.WithWebauthnCredential(cred))
	if err != nil {
		sentry.CaptureException(err)
		return c.Status(500).SendString(err.Error())
	}
	accs := make([]*v1.Account, 0)
	lclAccs := cont.ListLocalAccounts()
	for _, lclAcc := range lclAccs {
		accs = append(accs, lclAcc.ToProto())
	}

	res := &v1.KeygenResponse{
		Success:  true,
		Did:      acc.Did(),
		Primary:  cont.PrimaryIdentity(),
		Accounts: accs,
	}
	return c.JSON(res)
}

func (htt *HttpTransport) Login(c *fiber.Ctx) error {
	req := new(v1.LoginRequest)
	if err := c.BodyParser(req); err != nil {
		sentry.CaptureException(err)
		return c.Status(400).SendString(err.Error())
	}

	// Get the origin from the request.
	origin := regexp.MustCompile(`[^a-zA-Z]+`).ReplaceAllString(req.Origin, "")
	_, err := resolver.GetService(context.Background(), origin)
	if err != nil {
		sentry.CaptureException(err)
		return c.Status(404).SendString(err.Error())
	}
	return nil
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                        Query Method for the Highway API                        ||
// ! ||--------------------------------------------------------------------------------||

func QueryAccount(c *fiber.Ctx) error {
	did := c.Params("did")

	// Get the origin from the request.
	doc, err := resolver.GetDID(context.Background(), did)
	if err != nil {
		sentry.CaptureException(err)
		return c.Status(404).SendString(err.Error())
	}
	resp := &v1.QueryDocumentResponse{
		Success:        (doc != nil),
		AccountAddress: doc.DIDIdentifier(),
		DidDocument:    doc,
	}
	return c.JSON(resp)
}

func (htt *HttpTransport) QueryDocument(c *fiber.Ctx) error {
	did := c.Params("did")
	// Get the origin from the request.
	doc, err := resolver.GetDID(context.Background(), did)
	if err != nil {
		sentry.CaptureException(err)
		return c.Status(404).SendString(err.Error())
	}
	resp := &v1.QueryDocumentResponse{
		Success:        (doc != nil),
		AccountAddress: doc.DIDIdentifier(),
		DidDocument:    doc,
	}
	return c.JSON(resp)
}

func (htt *HttpTransport) QueryService(c *fiber.Ctx) error {
	origin := c.Params("origin", "localhost")

	// Regex remove non-alphabet characters.
	origin = regexp.MustCompile(`[^a-zA-Z]+`).ReplaceAllString(origin, "")

	// Get the origin from the request.
	service, err := resolver.GetService(context.Background(), origin)
	if err != nil {
		sentry.CaptureException(err)
		return c.Status(404).SendString(err.Error())
	}

	challenge, err := service.IssueChallenge()
	if err != nil {
		sentry.CaptureException(err)
		return c.Status(500).SendString(err.Error())
	}
	resp := &v1.QueryServiceResponse{
		Challenge: string(challenge),
		RpName:    "Sonr",
		RpId:      service.Origin,
	}
	return c.JSON(resp)
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                        Accounts API Rest Implementation                        ||
// ! ||--------------------------------------------------------------------------------||
func (htt *HttpTransport) CreateAccount(c *fiber.Ctx) error {
	return nil
}

func (htt *HttpTransport) ListAccounts(c *fiber.Ctx) error {
	return nil
}

func (htt *HttpTransport) GetAccount(c *fiber.Ctx) error {
	return nil
}

func (htt *HttpTransport) DeleteAccount(c *fiber.Ctx) error {
	return nil
}

func (htt *HttpTransport) SignMessage(c *fiber.Ctx) error {
	return nil
}

func (htt *HttpTransport) VerifyMessage(c *fiber.Ctx) error {
	return nil
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                          Vault API Rest Implementation                         ||
// ! ||--------------------------------------------------------------------------------||

func (htt *HttpTransport) AddShare(c *fiber.Ctx) error {
	req := new(v1.AddShareRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err := resolver.InsertKeyShare(resolver.NewBasicStoreItem(req.Key, []byte(req.Value)))
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(&v1.AddShareResponse{
		Success: true,
		Key:     req.Key,
	})
}

func (htt *HttpTransport) SyncShare(c *fiber.Ctx) error {
	req := new(v1.SyncShareRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	record, err := resolver.GetKeyShare(req.Key)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.JSON(&v1.SyncShareResponse{
		Success: true,
		Value:   base64.StdEncoding.EncodeToString(record),
		Key:     req.Key,
	})
}

func (htt *HttpTransport) RefreshShare(c *fiber.Ctx) error {
	return c.Status(500).JSON(fiber.Map{"error": "not implemented"})
}
