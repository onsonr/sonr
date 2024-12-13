package context

import (
	"net/http"
	"strconv"

	"github.com/go-webauthn/webauthn/protocol"
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

func SetIsHumanSum(c echo.Context, isHumanSum int) error {
	sess, err := Get(c)
	if err != nil {
		return err
	}
	challenge, err := protocol.CreateChallenge()
	if err != nil {
		return err
	}
	return sess.db.Save(&models.Session{
		Challenge:  challenge.String(),
		IsHumanSum: isHumanSum,
	}).Error
}

func VerifyIsHumanSum(c echo.Context) error {
	sum := c.FormValue("is_human")
	sess, err := Get(c)
	if err != nil {
		return err
	}
	sumInt, err := strconv.Atoi(sum)
	if err != nil {
		return err
	}
	// Get the current session
	var session models.Session
	if err := sess.db.Where("id = ?", sess.Session().ID).First(&session).Error; err != nil {
		return err
	}
	if session.IsHumanSum != sumInt {
		return echo.NewHTTPError(400, "invalid human sum")
	}
	return nil
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
