package context

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/database/sessions"
)

func InsertCredential(c echo.Context, handle string) error {
	cred, err := ExtractCredential(c.FormValue("credential"))
	if err != nil {
		return err
	}
	sess, err := Get(c)
	if err != nil {
		return err
	}
	return sess.db.Save(cred.ToDBModel(handle, c.Request().Host)).Error
}

func InsertProfile(c echo.Context) error {
	sess, err := Get(c)
	if err != nil {
		return err
	}
	handle := c.FormValue("handle")
	firstName := c.FormValue("first_name")
	lastName := c.FormValue("last_name")
	return sess.db.Save(&sessions.User{
		Handle: handle,
		Name:   fmt.Sprintf("%s %s", firstName, lastName),
	}).Error
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
	if err := sess.db.Model(&sessions.Session{}).Where("user_handle = ?", handle).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}
