package session

import "github.com/labstack/echo/v4"

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

// SetFirstName sets the first name in the session
func SetFirstName(c echo.Context, name string) error {
	sess, err := Get(c)
	if err != nil {
		return err
	}
	sess.Session().FirstName = name
	return sess.db.Save(sess.Session()).Error
}

// SetLastInitial sets the last initial in the session
func SetLastInitial(c echo.Context, initial string) error {
	sess, err := Get(c)
	if err != nil {
		return err
	}
	sess.Session().LastInitial = initial
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
