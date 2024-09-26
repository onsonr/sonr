package middleware

import (
	"fmt"
	"time"
)

const (
	OriginMacroonCaveat  MacroonCaveat = "origin"
	ScopesMacroonCaveat  MacroonCaveat = "scopes"
	SubjectMacroonCaveat MacroonCaveat = "subject"
	ExpMacroonCaveat     MacroonCaveat = "exp"
	TokenMacroonCaveat   MacroonCaveat = "token"
)

type MacroonCaveat string

func (c MacroonCaveat) Equal(other string) bool {
	return string(c) == other
}

func (c MacroonCaveat) String() string {
	return string(c)
}

func (c MacroonCaveat) Verify(value string) error {
	switch c {
	case OriginMacroonCaveat:
		return nil
	case ScopesMacroonCaveat:
		return nil
	case SubjectMacroonCaveat:
		return nil
	case ExpMacroonCaveat:
		// Check if the expiration time is still valid
		exp, err := time.Parse(time.RFC3339, value)
		if err != nil {
			return err
		}
		if time.Now().After(exp) {
			return fmt.Errorf("expired")
		}
		return nil
	case TokenMacroonCaveat:
		return nil
	default:
		return fmt.Errorf("unknown caveat: %s", c)
	}
}

var MacroonCaveats = []MacroonCaveat{OriginMacroonCaveat, ScopesMacroonCaveat, SubjectMacroonCaveat, ExpMacroonCaveat, TokenMacroonCaveat}
