package rest

import (
	"context"
	"encoding/base64"
	"regexp"

	"github.com/gofiber/fiber/v2"
	"github.com/sonrhq/core/internal/local"
	"github.com/sonrhq/core/x/identity/controller"
	"github.com/sonrhq/core/types/crypto"
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
	sess, err := htt.SessionStore.Get(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	sess.Set("username", req.Username)

	// Get the origin from the request.
	// uuid := req.Uuid
	service, _ := local.Context().GetService(context.Background(), req.Origin)
	if service == nil {
		// Try to get the service from the localhost
		service, _ = local.Context().GetService(context.Background(), "localhost")
	}
	sess.Set("service", service.Origin)

	// Check if service is still nil - return internal server error
	if service == nil {
		return c.Status(500).SendString("Internal Server Error.")
	}

	// Checking if the credential response is valid.
	cred, err := service.VerifyCreationChallenge(req.CredentialResponse)
	if err != nil {
		c.Status(400).SendString(err.Error())
	}

	// Create a new controller with the credential.
	cont, err := controller.NewController(controller.WithWebauthnCredential(cred))
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	accs := make([]*v1.Account, 0)
	lclAccs, err := cont.ListAccounts()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	for _, lclAcc := range lclAccs {
		accs = append(accs, lclAcc.ToProto())
	}
	res := &v1.KeygenResponse{
		Success:  true,
		Did:      cont.Did(),
		Primary:  cont.PrimaryIdentity(),
		Accounts: accs,
	}
	usr := NewUser(cont)
	usrBytes, err := usr.Marshal()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	sess.Set("user", usrBytes)
	err = sess.Save()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(res)
}

func (htt *HttpTransport) Login(c *fiber.Ctx) error {
	req := new(v1.LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	// Get the origin from the request.
	origin := regexp.MustCompile(`[^a-zA-Z]+`).ReplaceAllString(req.Origin, "")
	_, err := local.Context().GetService(context.Background(), origin)
	if err != nil {
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
	doc, err := local.Context().GetDID(context.Background(), did)
	if err != nil {

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
	doc, err := local.Context().GetDID(context.Background(), did)
	if err != nil {

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
	username := c.Params("username", "admin")
	// Get the origin from the request.
	service, err := local.Context().GetService(context.Background(), origin)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	challenge, err := service.GetCredentialCreationOptions(username)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	resp := &v1.QueryServiceResponse{
		CredentialOptions: challenge,
		RpName:            "Sonr",
		RpId:              service.Origin,
	}
	return c.JSON(resp)
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                        Accounts API Rest Implementation                        ||
// ! ||--------------------------------------------------------------------------------||
func (htt *HttpTransport) CreateAccount(c *fiber.Ctx) error {
	sess, err := htt.SessionStore.Get(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	usrVal := sess.Get("user")
	if usrVal == nil {
		return c.Status(401).SendString("Unauthorized")
	}
	usrBytes, ok := usrVal.([]byte)
	if !ok {
		return c.Status(500).SendString("Internal Server Error")
	}
	usr, err := LoadUser(usrBytes)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	req := new(v1.CreateAccountRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	cont, err := controller.LoadController(usr.PrimaryIdentity)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	ct := crypto.CoinTypeFromName(req.CoinType)
	acc, err := cont.CreateAccount(req.Name, ct)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	res := &v1.CreateAccountResponse{
		Success: true,
		Accounts: []*v1.Account{
			acc.ToProto(),
		},
	}
	return c.JSON(res)
}

func (htt *HttpTransport) ListAccounts(c *fiber.Ctx) error {
	sess, err := htt.SessionStore.Get(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	usrVal := sess.Get("user")
	if usrVal == nil {
		return c.Status(401).SendString("Unauthorized")
	}
	usrBytes, ok := usrVal.([]byte)
	if !ok {
		return c.Status(500).SendString("Internal Server Error")
	}
	usr, err := LoadUser(usrBytes)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	cont, err := controller.LoadController(usr.PrimaryIdentity)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	accs, err := cont.ListAccounts()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	res := &v1.ListAccountsResponse{
		Success: true,
	}
	for _, acc := range accs {
		res.Accounts = append(res.Accounts, acc.ToProto())
	}
	return c.JSON(res)
}

func (htt *HttpTransport) GetAccount(c *fiber.Ctx) error {
	sess, err := htt.SessionStore.Get(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	usrVal := sess.Get("user")
	if usrVal == nil {
		return c.Status(401).SendString("Unauthorized")
	}
	usrBytes, ok := usrVal.([]byte)
	if !ok {
		return c.Status(500).SendString("Internal Server Error")
	}
	usr, err := LoadUser(usrBytes)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	req := new(v1.GetAccountRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	cont, err := controller.LoadController(usr.PrimaryIdentity)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	acc, err := cont.GetAccount(req.Address)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	res := &v1.GetAccountResponse{
		Success:  true,
		CoinType: acc.CoinType().Name(),
		Accounts: []*v1.Account{
			acc.ToProto(),
		},
	}
	return c.JSON(res)
}

func (htt *HttpTransport) SignMessage(c *fiber.Ctx) error {
	sess, err := htt.SessionStore.Get(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	usrVal := sess.Get("user")
	if usrVal == nil {
		return c.Status(401).SendString("Unauthorized")
	}
	usrBytes, ok := usrVal.([]byte)
	if !ok {
		return c.Status(500).SendString("Internal Server Error")
	}
	usr, err := LoadUser(usrBytes)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	req := new(v1.SignMessageRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	cont, err := controller.LoadController(usr.PrimaryIdentity)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	bz, err := base64.RawStdEncoding.DecodeString(req.Message)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	sig, err := cont.Sign(req.Did, bz)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	res := &v1.SignMessageResponse{
		Success:   true,
		Signature: base64.RawStdEncoding.EncodeToString(sig),
		Message:   req.Message,
		Did:       req.Did,
	}
	return c.JSON(res)
}

func (htt *HttpTransport) VerifyMessage(c *fiber.Ctx) error {
	sess, err := htt.SessionStore.Get(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	usrVal := sess.Get("user")
	if usrVal == nil {
		return c.Status(401).SendString("Unauthorized")
	}
	usrBytes, ok := usrVal.([]byte)
	if !ok {
		return c.Status(500).SendString("Internal Server Error")
	}
	usr, err := LoadUser(usrBytes)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	req := new(v1.VerifyMessageRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	bz, err := base64.RawStdEncoding.DecodeString(req.Message)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	sig, err := base64.RawStdEncoding.DecodeString(req.Signature)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	cont, err := controller.LoadController(usr.PrimaryIdentity)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	ok, err = cont.Verify(req.Did, bz, sig)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	res := &v1.VerifyMessageResponse{
		Success: ok,
		Did:     req.Did,
	}
	return c.JSON(res)
}

func (htt *HttpTransport) SendMail(c *fiber.Ctx) error {
	sess, err := htt.SessionStore.Get(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	usrVal := sess.Get("user")
	if usrVal == nil {
		return c.Status(401).SendString("Unauthorized")
	}
	usrBytes, ok := usrVal.([]byte)
	if !ok {
		return c.Status(500).SendString("Internal Server Error")
	}
	usr, err := LoadUser(usrBytes)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	req := new(v1.SendMailRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	cont, err := controller.LoadController(usr.PrimaryIdentity)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	err = cont.SendMail(req.FromAddress, req.ToAddress, req.Message)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	res := &v1.SendMailResponse{
		Success: true,
	}
	return c.JSON(res)
}

func (htt *HttpTransport) ReadMail(c *fiber.Ctx) error {
	sess, err := htt.SessionStore.Get(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	usrVal := sess.Get("user")
	if usrVal == nil {
		return c.Status(401).SendString("Unauthorized")
	}
	usrBytes, ok := usrVal.([]byte)
	if !ok {
		return c.Status(500).SendString("Internal Server Error")
	}
	usr, err := LoadUser(usrBytes)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	req := new(v1.ReadMailRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	cont, err := controller.LoadController(usr.PrimaryIdentity)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	msgs, err := cont.ReadMail(req.AccountAddress)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	fromBodyMap := make(map[string]string)
	for _, msg := range msgs {
		fromBodyMap[msg.Sender] = msg.Content
	}
	res := &v1.ReadMailResponse{
		Success:  true,
		Messages: fromBodyMap,
	}
	return c.JSON(res)
}
