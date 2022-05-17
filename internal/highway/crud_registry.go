package highway

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	rtv1 "github.com/sonr-io/sonr/x/registry/types"
	rt "go.buf.build/grpc/go/sonr-io/blockchain/registry"
)

// QueryWhoIsHTTP
// @Summary
// @Schemes
// @Description QueryWhoIsHTTP
// @Tags Registry
// @Produce json
// @Success      200  {string}  message
// @Failure      500  {string}  message
// @Router /registry/who_is[GET]
func (s *HighwayServer) QueryWhoIsHTTP(c *gin.Context) {
	names, err := s.cosmos.QueryAllNames()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if names == nil {
		c.JSON(http.StatusOK, []rtv1.WhoIs{})
	}

	c.JSON(http.StatusOK, names)
}

// QueryWhoIsDIDHTTP
// @Summary
// @Schemes
// @Description QueryWhoIsDIDHTTP
// @Tags Registry
// @Produce json
// @Success      200  {string}  message
// @Failure      500  {string}  message
// @Router /registry/who_is/:did [GET]
func (s *HighwayServer) QueryWhoIsDIDHTTP(c *gin.Context) {
	if did := c.Param("did"); did != "" {
		res, err := s.cosmos.QueryName(did)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "did is required"})
}

// StartRegisterName starts the registration process for a Sonr Account.
// @Summary Start Register Name
// @Schemes
// @Description StartRegisterName starts the registration process and returns a PublicKeyCredentialCreationOptions. Initiating the registration process for a Sonr Account.
// @Tags Registry
// @Produce json
// @Success      200  {string}  message
// @Failure      500  {string}  message
// @Router /name/register/start/:username [GET]
func (s *HighwayServer) StartRegisterName(c *gin.Context) {
	if username := c.Param("username"); username != "" {
		if s.cosmos.NameExists(username) {
			c.JSON(http.StatusConflict, gin.H{"error": "username already exists"})
		}

		options, err := s.webauthn.SaveRegistrationSession(
			c.Request,
			c.Writer,
			username,
			s.cosmos.AccountName(),
		)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, options)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username is required"})
	}
}

// FinishRegisterName finishes the registration process for a Sonr Account.
// @Summary Finish Register Name
// @Schemes
// @Description FinishRegisterName finishes the registration process and returns a PublicKeyCredentialResponse. Succesfully registering a WebAuthn credential to a Sonr Account.
// @Tags Registry
// @Produce json
// @Success      200  {string}  message
// @Failure      500  {string}  message
// @Router /name/register/finish/:username [POST]
func (s *HighwayServer) FinishRegisterName(c *gin.Context) {
	var username string
	if username = c.Param("username"); username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username is required"})
	}

	cred, err := s.webauthn.FinishRegistrationSession(c.Request, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	// define a message to create a did
	msg := rtv1.NewMsgRegisterName(s.cosmos.Address(), username, *cred)

	// broadcast a transaction from account `alice` with the message
	// to create a did and return the response
	txResp, err := s.cosmos.BroadcastRegisterName(msg)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, txResp)
}

// StartAccessName accesses the user's existing credentials initiating the authentication process.
// @Summary Start Access Name
// @Schemes
// @Description StartAccessName accesses the user's existing credentials and returns a PublicKeyCredentialRequestOptions. Beggining the authentication process.
// @Tags Registry
// @Produce json
// @Success      200  {string}  message
// @Failure      500  {string}  message
// @Router /name/access/start/:username [GET]
func (s *HighwayServer) StartAccessName(c *gin.Context) {
	var username string
	if username = c.Param("username"); username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username is required"})
	}

	// Check if user exists and return error if it does not
	whoIs, err := s.cosmos.QueryName(username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	// Call Save to store the session data
	options, err := s.webauthn.SaveAuthenticationSession(c.Request, c.Writer, whoIs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, options)
}

// FinishAccessName finishes the authentication process successfully logging in the Sonr Account.
// @Summary Finish Access Name
// @Schemes
// @Description FinishAccessName finishes the authentication process and returns a PublicKeyCredentialResponse
//				successfully logging in a Sonr Account.
// @Tags Registry
// @Produce json
// @Success      200  {string}  message
// @Failure      500  {string}  message
// @Router /name/access/finish/:username [post]
func (s *HighwayServer) FinishAccessName(c *gin.Context) {
	var username string
	if username = c.Param("username"); username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username is required"})
	}

	cred, err := s.webauthn.FinishAuthenticationSession(c.Request, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, cred)
}

// UpdateName updates a name.
func (s *HighwayServer) UpdateName(ctx context.Context, req *rt.MsgUpdateName) (*rt.MsgUpdateNameResponse, error) {
	resp, err := s.cosmos.BroadcastUpdateName(rtv1.NewMsgUpdateNameFromBuf(req))
	if err != nil {
		return nil, err
	}

	return &rt.MsgUpdateNameResponse{
		Code:    resp.Code,
		Message: resp.Message,
		WhoIs:   rtv1.NewWhoIsToBuf(resp.WhoIs),
	}, nil
}

// UpdateNameHTTP updates a name on the Sonr blockchain registry.
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

	resp, err := s.UpdateName(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, resp)
}
