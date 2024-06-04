package vault

import (
	"context"
)

type Client interface {
	Assign(ctx context.Context, resp string) error
}
