package middleware

import (
	"context"

	"fmt"

	v1 "github.com/sonrhq/core/types/common"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sonrhq/core/internal/local"
	"github.com/sonrhq/core/x/identity/controller"
	"github.com/sonrhq/core/x/identity/types"
)

var (
	// ErrInvalidToken is returned when the token is invalid
	ErrInvalidToken = fmt.Errorf("invalid token")
)

type User struct {
	// DID of the user
	Did string `json:"_id"`

	// DID document of the primary identity
	Username string `json:"username"`

	Accounts []*v1.AccountInfo `json:"accounts"`

	// Controller
	controller.Controller
}

func NewUser(c controller.Controller, username string) *User {
	return &User{
		Did:        c.Did(),
		Username:   username,
		Controller: c,
	}
}

func (u *User) ListAccounts() ([]*v1.AccountInfo, error) {
	accs := make([]*v1.AccountInfo, 0)
	lclAccs, err := u.Controller.ListAccounts()
	if err != nil {
		return nil, err
	}
	for _, lclAcc := range lclAccs {
		accs = append(accs, lclAcc.ToProto())
	}
	return accs, nil
}

func (u *User) JWTClaims() jwt.MapClaims {
	return jwt.MapClaims{
		"did":      u.Did,
		"username": u.Username,
	}
}
func (u *User) PrimaryIdentity() (*types.DidDocument, error) {
	return local.Context().GetDID(context.Background(), u.Did)
}

func (u *User) JWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, u.JWTClaims())
	return token.SignedString(local.Context().SigningKey())
}

func FetchUser(c *fiber.Ctx) (*User, error) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	did, ok := claims["did"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid did")
	}
	username, ok := claims["username"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid username")
	}
	primary, err := local.Context().GetDID(context.Background(), did)
	if err != nil {
		return nil, err
	}
	cont, err := controller.LoadController(primary)
	if err != nil {
		return nil, err
	}
	return &User{
		Did:      did,
		Username: username,
		Controller: cont,
	}, nil
}
