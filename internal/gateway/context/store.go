package context

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/gateway/models"
)

func InsertCredential(c echo.Context, handle string, cred *models.CredentialDescriptor) error {
	sess, err := Get(c)
	if err != nil {
		return err
	}
	return sess.db.Save(cred.ToDBModel(handle, c.Request().Host)).Error
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

	return sess.db.Save(&models.User{
		Origin:  c.Request().Host,
		Address: addr,
		Handle:  handle,
		Name:    name,
	}).Error
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
	var session models.Session
	if err := sess.db.Where("id = ?", sess.Session().ID).First(&session).Error; err != nil {
		return false
	}
	sessionSum := sess.Session().IsHumanFirst + sess.Session().IsHumanLast
	return sessionSum == sumInt
}

// GetProfile returns the current user profile from the address cookie
func GetProfile(c echo.Context) (models.User, error) {
	sess, err := Get(c)
	if err != nil {
		return models.User{}, err
	}
	addr, err := c.Cookie("sonr.address")
	if err != nil {
		return models.User{}, err
	}
	var user models.User
	if err := sess.db.Where("address = ?", addr).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
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

	var count int64
	if err := sess.db.Model(&models.User{}).Where("handle = ?", handle).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
