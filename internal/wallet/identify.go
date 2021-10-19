package wallet

import (
	"context"
	"net"
	"time"

	"github.com/pkg/errors"
)

const (
	// 53 is the default DNS port
	DIAL_TIMEOUT        = time.Millisecond * time.Duration(10000)
	HDNS_RESOLVER_ONE   = "103.196.38.38:53" // Hardcoded Public DNS Resolver for HDNS #1
	HDNS_RESOLVER_TWO   = "103.196.38.39:53" // Hardcoded Public DNS Resolver for HDNS #2
	HDNS_RESOLVER_THREE = "103.196.38.40:53" // Hardcoded Public DNS Resolver for HDNS #3
	DOMAIN              = "snr"
	API_ADDRESS         = "https://www.namebase.io/api/v0/"
	DNS_ENDPOINT        = "dns/domains/snr/nameserver"
	FINGERPRINT_DIVIDER = "v=0;fingerprint="
	AUTH_DIVIDER        = "._auth."
	NB_DNS_API_URL      = API_ADDRESS + DNS_ENDPOINT
)

// hdnsResolver is a DNS Resolver that resolves SName records.
type hdnsResolver struct {
	resolver *net.Resolver
}

// SignedMetadata is a struct to be used for signing metadata.
type SignedMetadata struct {
	Timestamp int64
	PublicKey []byte
	NodeId    string
}

// SignedUUID is a struct to be converted into a UUID.
type SignedUUID struct {
	Timestamp int64
	Signature []byte
	Value     string
}

func newResolver() *hdnsResolver {
	return &hdnsResolver{
		resolver: &net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				// Create Dialer
				d := net.Dialer{
					Timeout: DIAL_TIMEOUT,
				}

				// Dial First Resolver
				c, err := d.DialContext(ctx, network, HDNS_RESOLVER_ONE)
				if err == nil {
					return c, nil
				}

				// Dial Second Resolver
				c, err = d.DialContext(ctx, network, HDNS_RESOLVER_TWO)
				if err == nil {
					return c, nil
				}

				// Dial Third Resolver
				c, err = d.DialContext(ctx, network, HDNS_RESOLVER_THREE)
				if err == nil {
					return c, nil
				}

				// Return Error if we failed to dial all three resolvers
				return nil, errors.WithMessage(err, "Failed to resolve on all 3 DNS HTTP Servers")
			},
		},
	}
}
