package context

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/gateway/models"
	"github.com/onsonr/sonr/internal/gateway/repository"
)

func InsertCredential(c echo.Context, handle string, cred *models.CredentialDescriptor) error {
	sess, err := Get(c)
	if err != nil {
		return err
	}

	params := repository.InsertCredentialParams{
		Handle:       handle,
		CredentialID: cred.ID,
		Origin:       c.Request().Host,
		Type:         cred.Type,
		Transports:   cred.GetTransportsString(),
	}

	_, err = sess.db.InsertCredential(context.Background(), params)
	return err
}

func InsertProfile(c echo.Context, addr string, handle string, name string) error {
	sess, err := Get(c)
	if err != nil {
		return err
	}

	// Set Address as cookie
	c.SetCookie(&http.Cookie{
		Name:   "sonr.address",
		Value:  addr,
		Path:   "/",
		Secure: true,
	})

	params := repository.InsertUserParams{
		Address: addr,
		Handle:  handle,
		Origin:  c.Request().Host,
		Name:    name,
	}

	_, err = sess.db.InsertUser(context.Background(), params)
	return err
}

func VerifyIsHumanSum(c echo.Context) bool {
	sum := c.FormValue("is_human")
	sess, err := Get(c)
	if err != nil {
		return false
	}
	sumInt, err := strconv.Atoi(sum)
	if err != nil {
		return false
	}
	// Get the current session
	sessionSum := sess.Session().IsHumanFirst + sess.Session().IsHumanLast
	return sessionSum == sumInt
}

// GetProfile returns the current user profile from the address cookie
func GetProfile(c echo.Context) (repository.User, error) {
	sess, err := Get(c)
	if err != nil {
		return repository.User{}, err
	}
	addr, err := c.Cookie("sonr.address")
	if err != nil {
		return repository.User{}, err
	}

	return sess.db.GetUserByAddress(context.Background(), addr.Value)
}

// ╭───────────────────────────────────────────────────────╮
// │                  DB Getter Functions                  │
// ╰───────────────────────────────────────────────────────╯

// SessionID returns the session ID
func SessionID(c echo.Context) (string, error) {
	sess, err := Get(c)
	if err != nil {
		return "", err
	}
	return sess.Session().ID, nil
}

// HandleExists checks if a handle already exists in any session
func HandleExists(c echo.Context, handle string) (bool, error) {
	sess, err := Get(c)
	if err != nil {
		return false, err
	}

	return sess.db.CheckHandleExists(context.Background(), handle)
}
