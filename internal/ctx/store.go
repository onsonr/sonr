package ctx

import (
	"fmt"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

type WebBytes = protocol.URLEncodedBase64

type Session interface {
	ID() string
	Origin() string

	Address() string
	ChainID() string

	GetChallenge(subject string) (WebBytes, error)
	ValidateChallenge(challenge WebBytes, subject string) error

	IsState(State) bool
	SaveHTTP(c echo.Context) error
}

func defaultSession(id string, s *sessions.Session) *session {
	return &session{
		session: s,
		id:      id,
		origin:  "",
		address: "",
		chainID: "",
		state:   StateUnauthenticated,
	}
}

func NewSessionFromValues(vals map[interface{}]interface{}) *session {
	s := &session{
		id:        vals["id"].(string),
		origin:    vals["origin"].(string),
		address:   vals["address"].(string),
		chainID:   vals["chainID"].(string),
		state:     StateFromString(vals["state"].(string)),
		challenge: vals["challenge"].(WebBytes),
		subject:   vals["subject"].(string),
	}
	return s
}

type session struct {
	// Defaults
	session *sessions.Session
	id      string // Generated ksuid http cookie; Initialized on first request
	origin  string // Webauthn mapping to Relaying Party ID; Initialized on first request

	// Initialization
	address string // Webauthn mapping to User ID; Supplied by DWN frontend
	chainID string // Macaroon mapping to location; Supplied by DWN frontend

	// Authentication
	challenge WebBytes // Webauthn mapping to Challenge; Per session based on origin
	subject   string   // Webauthn mapping to User Displayable Name; Supplied by DWN frontend

	// State
	state State
}

func (s *session) ID() string {
	return s.id
}

func (s *session) Origin() string {
	return s.origin
}

func (s *session) Address() string {
	return s.address
}

func (s *session) ChainID() string {
	return s.chainID
}

func (s *session) GetChallenge(subject string) (WebBytes, error) {
	if s.challenge == nil {
		return nil, nil
	}
	return s.challenge, nil
}

func (s *session) ValidateChallenge(challenge WebBytes, subject string) error {
	if s.challenge == nil {
		return nil
	}
	if s.challenge.String() != challenge.String() {
		return fmt.Errorf("invalid challenge")
	}
	s.subject = subject
	s.state = StateAuthenticated
	return nil
}

func (s *session) IsState(state State) bool {
	return s.state == state
}

func (s *session) SaveHTTP(c echo.Context) error {
	sess, err := store.Get(c.Request(), s.id)
	if err != nil {
		return err
	}
	sess.Values = s.Values()
	err = sess.Save(c.Request(), c.Response().Writer)
	if err != nil {
		return err
	}
	return nil
}

func (s *session) Values() map[interface{}]interface{} {
	vals := make(map[interface{}]interface{})
	vals["id"] = s.id
	vals["address"] = s.address
	vals["chainID"] = s.chainID
	vals["state"] = s.state
	vals["challenge"] = s.challenge
	vals["subject"] = s.subject
	return vals
}

func GetSession(c echo.Context) Session {
	id, _ := getSessionID(c.Request().Context())
	sess, _ := store.Get(c.Request(), id)
	if sess.IsNew {
		s := defaultSession(id, sess)
		s.SaveHTTP(c)
		return s
	}
	s, _ := readSessionFromStore(c, id)
	return s
}
