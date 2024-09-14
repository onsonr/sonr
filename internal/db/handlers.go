package db

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/db/orm"
)

func (db *DB) ServeEcho(e *echo.Group) {
	e.GET("/accounts", db.HandleAccount)
	e.GET("/assets", db.HandleAsset)
	e.GET("/credentials", db.HandleCredential)
	e.GET("/keyshares", db.HandleKeyshare)
	e.GET("/permissions", db.HandlePermission)
	e.GET("/profiles", db.HandleProfile)
	e.GET("/properties", db.HandleProperty)
}

func (db *DB) HandleAccount(c echo.Context) error {
	data := new(orm.Account)
	if err := c.Bind(data); err != nil {
		return err
	}
	// Check the method for GET, POST, PUT, DELETE
	switch c.Request().Method {
	case echo.POST:

		if err := db.AddAccount(data); err != nil {
			return err
		}
		return c.JSON(200, "OK")
	case echo.PUT:

		if err := db.UpdateAccount(data); err != nil {
			return err
		}
		return c.JSON(200, "OK")
	case echo.DELETE:
		if err := db.DeleteAccount(data); err != nil {
			return err
		}
		return c.JSON(200, nil)
	}
	return c.JSON(200, data)
}

func (db *DB) HandleAsset(c echo.Context) error {
	data := new(orm.Asset)
	if err := c.Bind(data); err != nil {
		return err
	}

	switch c.Request().Method {
	case echo.POST:
		if err := db.AddAsset(data); err != nil {
			return err
		}
		return c.JSON(200, "OK")
	case echo.PUT:
		if err := db.UpdateAsset(data); err != nil {
			return err
		}
		return c.JSON(200, "OK")
	case echo.DELETE:
		if err := db.DeleteAsset(data); err != nil {
			return err
		}
		return c.JSON(200, "OK")
	}
	return c.JSON(200, data)
}

func (db *DB) HandleCredential(c echo.Context) error {
	data := new(orm.Credential)
	if err := c.Bind(data); err != nil {
		return err
	}

	switch c.Request().Method {
	case echo.POST:
		if err := db.AddCredential(data); err != nil {
			return err
		}
		return c.JSON(200, "OK")
	case echo.PUT:
		if err := db.UpdateCredential(data); err != nil {
			return err
		}
		return c.JSON(200, "OK")
	case echo.DELETE:
		if err := db.DeleteCredential(data); err != nil {
			return err
		}
	}
	return c.JSON(200, data)
}

func (db *DB) HandleKeyshare(c echo.Context) error {
	data := new(orm.Keyshare)
	if err := c.Bind(data); err != nil {
		return err
	}

	switch c.Request().Method {
	case echo.POST:
		if err := db.AddKeyshare(data); err != nil {
			return err
		}
		return c.JSON(200, "OK")
	case echo.PUT:
		if err := db.UpdateKeyshare(data); err != nil {
			return err
		}
		return c.JSON(200, "OK")
	case echo.DELETE:
		if err := db.DeleteKeyshare(data); err != nil {
			return err
		}
	}
	return c.JSON(200, data)
}

func (db *DB) HandlePermission(c echo.Context) error {
	data := new(orm.Permission)
	if err := c.Bind(data); err != nil {
		return err
	}

	switch c.Request().Method {
	case echo.POST:
		if err := db.AddPermission(data); err != nil {
			return err
		}
		return c.JSON(200, "OK")
	case echo.PUT:
		if err := db.UpdatePermission(data); err != nil {
			return err
		}
		return c.JSON(200, "OK")
	case echo.DELETE:
		if err := db.DeletePermission(data); err != nil {
			return err
		}
	}
	return c.JSON(200, data)
}

func (db *DB) HandleProfile(c echo.Context) error {
	data := new(orm.Profile)
	if err := c.Bind(data); err != nil {
		return err
	}

	switch c.Request().Method {
	case echo.POST:
		if err := db.AddProfile(data); err != nil {
			return err
		}
		return c.JSON(200, "OK")
	case echo.PUT:
		if err := db.UpdateProfile(data); err != nil {
			return err
		}
		return c.JSON(200, "OK")
	case echo.DELETE:
		if err := db.DeleteProfile(data); err != nil {
			return err
		}
	}
	return c.JSON(200, data)
}

func (db *DB) HandleProperty(c echo.Context) error {
	data := new(orm.Property)
	if err := c.Bind(data); err != nil {
		return err
	}

	switch c.Request().Method {
	case echo.POST:
		if err := db.AddProperty(data); err != nil {
			return err
		}
		return c.JSON(200, "OK")
	case echo.PUT:
		if err := db.UpdateProperty(data); err != nil {
			return err
		}
		return c.JSON(200, "OK")
	case echo.DELETE:
		if err := db.DeleteProperty(data); err != nil {
			return err
		}
	}
	return c.JSON(200, data)
}
