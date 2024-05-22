package enclave

import (
	"github.com/di-dao/sonr/crypto/kss"
	"github.com/di-dao/sonr/internal/local"
)

func GetValKSS(recoveryCode []byte, valKssBytes []byte) (kss.Val, error) {
	_, err := local.Decrypt(valKssBytes, recoveryCode)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
