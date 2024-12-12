package context

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/database/sessions"
)

// ╭───────────────────────────────────────────────────────╮
// │                  DB Setter Functions                  │
// ╰───────────────────────────────────────────────────────╯

// SetUserHandle sets the user handle in the session
func SetUserHandle(c echo.Context, handle string) error {
	sess, err := Get(c)
	if err != nil {
		return err
	}
	sess.Session().UserHandle = handle
	return sess.db.Save(sess.Session()).Error
}

// SetVaultAddress sets the vault address in the session
func SetVaultAddress(c echo.Context, address string) error {
	sess, err := Get(c)
	if err != nil {
		return err
	}
	sess.Session().VaultAddress = address
	return sess.db.Save(sess.Session()).Error
}

// ╭───────────────────────────────────────────────────────╮
// │                  DB Getter Functions                  │
// ╰───────────────────────────────────────────────────────╯

// GetID returns the session ID
func GetID(c echo.Context) (string, error) {
	sess, err := Get(c)
	if err != nil {
		return "", err
	}
	return sess.Session().ID, nil
}

// GetBrowserName returns the browser name
func GetBrowserName(c echo.Context) (string, error) {
	sess, err := Get(c)
	if err != nil {
		return "", err
	}
	return sess.Session().BrowserName, nil
}

// GetBrowserVersion returns the browser version
func GetBrowserVersion(c echo.Context) (string, error) {
	sess, err := Get(c)
	if err != nil {
		return "", err
	}
	return sess.Session().BrowserVersion, nil
}

// GetPlatform returns the platform
func GetPlatform(c echo.Context) (string, error) {
	sess, err := Get(c)
	if err != nil {
		return "", err
	}
	return sess.Session().Platform, nil
}

// GetUserHandle returns the user handle
func GetUserHandle(c echo.Context) (string, error) {
	sess, err := Get(c)
	if err != nil {
		return "", err
	}
	return sess.Session().UserHandle, nil
}

// GetVaultAddress returns the vault address
func GetVaultAddress(c echo.Context) (string, error) {
	sess, err := Get(c)
	if err != nil {
		return "", err
	}
	return sess.Session().VaultAddress, nil
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
