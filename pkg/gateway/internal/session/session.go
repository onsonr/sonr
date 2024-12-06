package session

import "github.com/labstack/echo/v4"

// SetUserHandle sets the user handle in the session
func SetUserHandle(c echo.Context, handle string) error {
	sess, err := Get(c)
	if err != nil {
		return err
	}
	sess.Session().UserHandle = handle
}
