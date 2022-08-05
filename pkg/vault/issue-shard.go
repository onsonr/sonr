package vault

import (
	"errors"

	"github.com/sonr-io/sonr/pkg/crypto/did"
)

func (v *vaultImpl) IssueShard(shardPrefix, dscPub, dscShard string) (did.Service, error) {
	return did.Service{}, errors.New("unimplemented")
}
