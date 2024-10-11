package ctx

import "github.com/labstack/echo/v4"

type AuthState string

const (
	Visitor       AuthState = "visitor"
	Authenticated AuthState = "authenticated"
	Expired       AuthState = "expired"

	PendingCredentials AuthState = "pending_credentials"
	PendingAssertion   AuthState = "pending_assertion"
)

func (s AuthState) String() string {
	return string(s)
}

func GetAuthState(c echo.Context) AuthState {
	vals := c.Request().Header.Values("Authorization")
	if len(vals) == 0 {
		return Visitor
	}
	s := AuthState(c.Request().Header.Get("Authorization"))
	return s
}

func readSessionFromStore(c echo.Context, id string) (*session, error) {
	sess, err := store.Get(c.Request(), id)
	if err != nil {
		return nil, err
	}
	return NewSessionFromValues(sess.Values), nil
}

func writeSessionToStore(
	c echo.Context,
	id string,
) error {
	sess, err := store.Get(c.Request(), id)
	if err != nil {
		return err
	}
	s := defaultSession(id, sess)
	err = s.SaveHTTP(c)
	if err != nil {
		return err
	}
	return nil
}
