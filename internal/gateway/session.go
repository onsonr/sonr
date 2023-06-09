package gateway

import (
	"github.com/kataras/go-sessions/v3"
)

// Session represents an authenticated session. In a real application, this might include additional fields such as the user ID, the session expiry time, etc.
type Session struct {
	UserAlias  string
	Credential string
	Address    string
	IsMobile   bool
	Did        string
	Name       string
	Days       string
	Secret     string
}

// LoadSession gets all the session values and verifies their types.
func LoadSession(sess *sessions.Session) *Session {
	return &Session{
		UserAlias:  sess.GetString("UserAlias"),
		Credential: sess.GetString("Credential"),
		Address:    sess.GetString("Address"),
		IsMobile:   sess.GetBooleanDefault("IsMobile", false),
		Did:        sess.GetString("Did"),
	}
}

func (s *Session) Save(sess *sessions.Session) {
	sess.Set("UserAlias", s.UserAlias)
	sess.Set("Credential", s.Credential)
	sess.Set("Address", s.Address)
	sess.Set("IsMobile", s.IsMobile)
	sess.Set("Did", s.Did)
}

func defaultSession() *Session {
	return &Session{
		UserAlias:  "",
		Credential: "",
		Address:    "",
		IsMobile:   false,
		Did:        "",
		Name:       "sonr",
		Days:       "1",
		Secret:     "dsads£2132215£%%Ssdsa",
	}
}

type SessionValue func(s *Session)

func WithUserAlias(alias string) SessionValue {
	return func(s *Session) {
		s.UserAlias = alias
	}
}

func WithCredential(credential []byte) SessionValue {
	return func(s *Session) {
		s.Credential = string(credential)
	}
}

func WithAddress(address string) SessionValue {
	return func(s *Session) {
		s.Address = address
	}
}

func WithIsMobile(isMobile bool) SessionValue {
	return func(s *Session) {
		s.IsMobile = isMobile
	}
}

func WithDid(did string) SessionValue {
	return func(s *Session) {
		s.Did = did
	}
}
