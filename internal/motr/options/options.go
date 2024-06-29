package options

import "github.com/di-dao/sonr/internal/vfd/models"

type LockOptions struct {
	Credential *models.Credential
	Initial    bool
}

type UnlockOptions struct {
	Credential *models.Credential
}
