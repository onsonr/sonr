package ctx

import "github.com/labstack/echo/v4"

type State string

const (
	StateAuthenticated      State = "authenticated"
	StateUnauthenticated    State = "unauthenticated"
	StatePendingCredentials State = "pending_credentials"
	StatePendingAssertion   State = "pending_assertion"
	StateDisabled           State = "disabled"
	StateDisconnected       State = "disconnected"
)

func (s State) String() string {
	return string(s)
}

func StateFromString(s string) State {
	switch s {
	case StateAuthenticated.String():
		return StateAuthenticated
	case StateUnauthenticated.String():
		return StateUnauthenticated
	case StatePendingCredentials.String():
		return StatePendingCredentials
	case StatePendingAssertion.String():
		return StatePendingAssertion
	case StateDisabled.String():
		return StateDisabled
	case StateDisconnected.String():
		return StateDisconnected
	default:
		return State("")
	}
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
