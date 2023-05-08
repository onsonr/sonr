package middleware

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sonrhq/core/internal/local"
	"github.com/sonrhq/core/x/identity"
	idtypes "github.com/sonrhq/core/x/identity/types"
	srvtypes "github.com/sonrhq/core/x/service/types"
)

type QueryOptions struct {
	origin    string
	alias     string
	did       string
	address   string
	assertion string
	attestion string
	isMobile  bool
	ucw_id    uint64
	challenge string
}

func ParseQuery(c *fiber.Ctx) *QueryOptions {
	origin := c.Params("origin", "localhost")
	alias := c.Query("alias", "admin")
	did := c.Query("did", "")
	address := c.Query("address", "")
	assertion := c.Query("assertion", "")
	attestion := c.Query("attestion", "")
	isMobile := c.QueryBool("mobile", false)
	ucw_id := c.QueryInt("ucw_id", 0)
	challenge := c.Query("challenge", "")

	return &QueryOptions{
		origin:    origin,
		alias:     alias,
		did:       did,
		address:   address,
		assertion: assertion,
		attestion: attestion,
		isMobile:  isMobile,
		ucw_id:    uint64(ucw_id),
		challenge: challenge,
	}
}

func (q *QueryOptions) Alias() string {
	return q.alias
}

func (q *QueryOptions) DID() string {
	return q.did
}

func (q *QueryOptions) Address() string {
	return q.address
}

func (q *QueryOptions) Assertion() string {
	return q.assertion
}

func (q *QueryOptions) Attestion() string {
	return q.attestion
}

func (q *QueryOptions) Challenge() string {
	return q.challenge
}

func (q *QueryOptions) IsMobile() bool {
	return q.isMobile
}

func (q *QueryOptions) Origin() string {
	return q.origin
}

func (q *QueryOptions) UCWID() uint64 {
	return q.ucw_id
}

func (q *QueryOptions) HasDID() bool {
	return q.did != ""
}

func (q *QueryOptions) HasAddress() bool {
	return q.address != ""
}

func (q *QueryOptions) HasAssertion() bool {
	return q.assertion != ""
}

func (q *QueryOptions) HasAttestion() bool {
	return q.attestion != ""
}

func (q *QueryOptions) HasAlias() bool {
	return q.alias != ""
}

func (q *QueryOptions) HasChallenge() bool {
	return q.challenge != ""
}

func (q *QueryOptions) HasOrigin() bool {
	return q.origin != ""
}

func (q *QueryOptions) HasUCWID() bool {
	return q.ucw_id != 0
}

func (q *QueryOptions) HasQuery() bool {
	return q.HasAlias() || q.HasDID() || q.HasAddress() || q.HasAssertion() || q.HasAttestion()
}

func (q *QueryOptions) GetService() (*srvtypes.ServiceRecord, error) {
	if q.HasOrigin() {
		return local.Context().GetService(context.Background(), q.Origin())
	}
	return nil, fmt.Errorf("no origin provided as parameter")
}

func (q *QueryOptions) GetDID() (*idtypes.DidDocument, error) {
	if q.HasDID() {
		return local.Context().GetDID(context.Background(), q.DID())
	}
	if q.HasAddress() {
		return local.Context().GetDIDByOwner(context.Background(), q.Address())
	}
	if q.HasAlias() {
		return local.Context().GetDIDByAlias(context.Background(), q.Alias())
	}
	return nil, fmt.Errorf("no did, alias, or address provided as query option")
}

func (q *QueryOptions) GetWalletClaims() (identity.WalletClaims, error) {
	ucw, err := local.Context().GetUnclaimedWallet(context.Background(), q.UCWID())
	if err != nil {
		return nil, fmt.Errorf("failed to find unclaimed wallet: %w", err)
	}
	claims := identity.LoadClaimableWallet(ucw)
	return claims, nil
}
