package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
	rt "go.buf.build/grpc/go/sonr-io/blockchain/registry"
)

// @Summary Start Register Name
// @Schemes
// @Description StartRegisterName starts the registration process and returns a PublicKeyCredentialCreationOptions. Initiating the registration process for a Sonr Account.
// @Tags Registry
// @Produce json
// @Success      200  {string}  message
// @Failure      500  {string}  message
// @Router /name/register/start/:username [get]
func (s *HighwayServer) StartRegisterName(c *gin.Context) {
	if username := c.Param("username"); username != "" {
		// Check if user exists and return error if it does
		if exists := s.Cosmos.NameExists(username); exists {
			c.JSON(http.StatusConflict, gin.H{"error": "username already exists"})
		}

		// Save Registration Session
		options, err := s.Webauthn.SaveRegistrationSession(c.Request, c.Writer, username, s.Cosmos.AccountName())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, options)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username is required"})
	}
}

// @Summary Finish Register Name
// @Schemes
// @Description FinishRegisterName finishes the registration process and returns a PublicKeyCredentialResponse. Succesfully registering a WebAuthn credential to a Sonr Account.
// @Tags Registry
// @Produce json
// @Success      200  {string}  message
// @Failure      500  {string}  message
// @Router /name/register/finish/:username [post]
func (s *HighwayServer) FinishRegisterName(c *gin.Context) {
	// get username
	username := c.Param("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username is required"})
	}

	// Finish Registration Session
	cred, err := s.Webauthn.FinishRegistrationSession(c.Request, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	// // define a message to create a did
	// msg := rtv1.NewMsgBuyNameAlias(s.Cosmos.Address(), username, *cred)

	// // broadcast a transaction from account `alice` with the message to create a did
	// // store response in txResp
	// txResp, err := s.Cosmos.BroadcastRegisterName(msg)
	// if err != nil {
	// 	c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	// }
	c.JSON(http.StatusOK, cred)
}

// @Summary Start Access Name
// @Schemes
// @Description StartAccessName accesses the user's existing credentials and returns a PublicKeyCredentialRequestOptions. Beggining the authentication process.
// @Tags Registry
// @Produce json
// @Success      200  {string}  message
// @Failure      500  {string}  message
// @Router /name/access/start/:username [get]
func (s *HighwayServer) StartAccessName(c *gin.Context) {
	// get username
	username := c.Param("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username is required"})
	}

	// Check if user exists and return error if it does not
	whoIs, err := s.Cosmos.QueryName(username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	// Call Save to store the session data
	options, err := s.Webauthn.SaveAuthenticationSession(c.Request, c.Writer, whoIs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, options)
}

// @Summary Finish Access Name
// @Schemes
// @Description FinishAccessName finishes the authentication process and returns a PublicKeyCredentialResponse. Succesfully logging in a Sonr Account.
// @Tags Registry
// @Produce json
// @Success      200  {string}  message
// @Failure      500  {string}  message
// @Router /name/access/finish/:username [post]
func (s *HighwayServer) FinishAccessName(c *gin.Context) {
	// get username
	username := c.Param("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username is required"})
	}

	// Finish the authentication session
	cred, err := s.Webauthn.FinishAuthenticationSession(c.Request, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	// handle successful login
	c.JSON(http.StatusOK, cred)
}

// @Summary Update Name
// @Schemes
// @Description UpdateName updates a name on the Sonr blockchain registry.
// @Tags Registry
// @Produce json
// @Success      200  {string}  message
// @Failure      500  {string}  message
// @Router /name/update [post]
func (s *HighwayServer) UpdateNameHTTP(c *gin.Context) {
	var req rt.MsgUpdateName
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	//c.JSON(http.StatusOK, resp)
}
