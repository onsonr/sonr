package session

import "github.com/labstack/echo/v4"

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

// SetUserArchitecture sets the user architecture in the session
func SetUserArchitecture(c echo.Context, arch string) error {
	sess, err := Get(c)
	if err != nil {
		return err
	}
	sess.Session().UserArchitecture = arch
	return sess.db.Save(sess.Session()).Error
}

// SetPlatform sets the platform in the session
func SetPlatform(c echo.Context, platform string) error {
	sess, err := Get(c)
	if err != nil {
		return err
	}
	sess.Session().Platform = platform
	return sess.db.Save(sess.Session()).Error
}

// SetPlatformVersion sets the platform version in the session
func SetPlatformVersion(c echo.Context, version string) error {
	sess, err := Get(c)
	if err != nil {
		return err
	}
	sess.Session().PlatformVersion = version
	return sess.db.Save(sess.Session()).Error
}

// SetDeviceModel sets the device model in the session
func SetDeviceModel(c echo.Context, model string) error {
	sess, err := Get(c)
	if err != nil {
		return err
	}
	sess.Session().DeviceModel = model
	return sess.db.Save(sess.Session()).Error
}
