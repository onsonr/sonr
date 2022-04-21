package highway

import (
	context "context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	rtv1 "github.com/sonr-io/blockchain/x/registry/types"
	rt "go.buf.build/sonr-io/grpc-gateway/sonr-io/blockchain/registry"
)

// RegisterNameStart starts the registration process for webauthn on http
func (s *HighwayServer) StartRegisterName(c *gin.Context) {
	if username := c.Param("username"); username != "" {
		// Check if user exists and return error if it does
		if exists := s.cosmos.ExistsName(username); exists {
			c.JSON(http.StatusConflict, gin.H{"error": "username already exists"})
		}

		creator := s.cosmos.AccountName()

		// Save Registration Session
		options, err := s.webauthn.SaveRegistrationSession(c.Request, c.Writer, username, creator)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, options)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username is required"})
	}
}

// FinishRegisterName handles the registration of a new credential
func (s *HighwayServer) FinishRegisterName(c *gin.Context) {
	// get username
	username := c.Param("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username is required"})
	}

	// Finish Registration Session
	cred, err := s.webauthn.FinishRegistrationSession(c.Request, username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// define a message to create a did
	msg := rtv1.NewMsgRegisterName(s.cosmos.Address(), username, *cred)

	// broadcast a transaction from account `alice` with the message to create a did
	// store response in txResp
	txResp, err := s.cosmos.BroadcastRegisterName(msg)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, txResp)
}

// StartAccessName accesses the user's existing credentials and returns a PublicKeyCredentialRequestOptions
func (s *HighwayServer) StartAccessName(c *gin.Context) {
	// get username
	username := c.Param("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username is required"})
	}

	// Check if user exists and return error if it does not
	whoIs, err := s.cosmos.QueryName(username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Call Save to store the session data
	options, err := s.webauthn.SaveAuthenticationSession(c.Request, c.Writer, whoIs)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, options)
}

// FinishAccessName handles the login of a credential and returns a PublicKeyCredentialResponse
func (s *HighwayServer) FinishAccessName(c *gin.Context) {
	// get username
	username := c.Param("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username is required"})
	}

	// Finish the authentication session
	cred, err := s.webauthn.FinishAuthenticationSession(c.Request, username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// handle successful login
	c.JSON(http.StatusOK, cred)
}

// UpdateName updates a name.
func (s *HighwayServer) UpdateName(ctx context.Context, req *rt.MsgUpdateName) (*rt.MsgUpdateNameResponse, error) {
	// Broadcast the Transaction
	resp, err := s.cosmos.BroadcastUpdateName(rtv1.NewMsgUpdateNameFromBuf(req))
	if err != nil {
		return nil, err
	}
	fmt.Println(resp.String())
	return &rt.MsgUpdateNameResponse{
		Code:    resp.Code,
		Message: resp.Message,
		WhoIs:   rtv1.NewWhoIsToBuf(resp.WhoIs),
	}, nil
}
