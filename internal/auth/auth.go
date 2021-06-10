package auth

import (
	md "github.com/sonr-io/core/pkg/models"
)

type AuthService interface {
	BuildPrefix(val string) string
	GetUser(prefix string) *md.User
	PutUser(user *md.User)
	CreateSName(name string)
	ValidateUser(val string) bool
}
