package actor

import (
	"context"

	"github.com/di-dao/sonr/crypto/kss"
	"github.com/di-dao/sonr/crypto/mpc"
)

func GetValidatorKeyshare(ctx context.Context) (kss.Val, error) {
	ks, err := mpc.GenerateKss()
	if err != nil {
		return nil, err
	}
	return ks.Val(), nil
}
