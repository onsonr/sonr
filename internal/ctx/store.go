package ctx

import (
	"github.com/labstack/echo/v4"
)

type Session interface {
	ID() string
	Address() string
	Authenticated() bool
	ChainID() string
}

type session struct {
	id            string
	address       string
	authenticated bool
	chainID       string
}

func (s *session) ID() string {
	return s.id
}

func (s *session) Address() string {
	return s.address
}

func (s *session) Authenticated() bool {
	return s.authenticated
}

func (s *session) ChainID() string {
	return s.chainID
}

func GetSession(c echo.Context) Session {
	id, _ := getSessionID(c.Request().Context())
	s, _ := readSessionFromStore(c, id)
	return s
}

func readSessionFromStore(c echo.Context, id string) (*session, error) {
	d := &session{
		id:            id,
		authenticated: false,
	}
	sess, err := store.Get(c.Request(), id)
	if err != nil {
		return nil, err
	}

	for k, v := range sess.Values {
		switch k {
		case "address":
			d.address = v.(string)
		case "authenticated":
			d.authenticated = v.(bool)
		case "chainID":
			d.chainID = v.(string)
		}
	}
	return d, nil
}

func writeSessionToStore(
	c echo.Context,
	id string,
	authenticated bool,
	address string,
	chainID string,
) error {
	sess, err := store.Get(c.Request(), id)
	if err != nil {
		return err
	}
	sess.Values["address"] = address
	sess.Values["authenticated"] = authenticated
	sess.Values["chainID"] = chainID
	err = sess.Save(c.Request(), c.Response().Writer)
	if err != nil {
		return err
	}
	return nil
}
