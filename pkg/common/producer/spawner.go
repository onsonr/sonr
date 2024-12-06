package producer

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/crypto/mpc"
)

func NewKeyset(c echo.Context) (mpc.Keyset, error) {
	ks, err := mpc.NewKeyset()
	if err != nil {
		return nil, err
	}
	return ks, nil
}

//
// func GetKeyset(c echo.Context) (mpc.Keyset, error) {
// 	cc, ok := c.(*SignerContext)
// 	if !ok {
// 		return nil, errors.New("not an SignerContext")
// 	}
// 	if !cc.hasKeyset {
// 		return nil, fmt.Errorf("keyset not found")
// 	}
// 	if cc.keyset == nil {
// 		return nil, fmt.Errorf("keyset is nil")
// 	}
// 	return cc.keyset, nil
// }
//
// func NewSource(c echo.Context) (mpc.KeyshareSource, error) {
// 	cc, ok := c.(*SignerContext)
// 	if !ok {
// 		return nil, errors.New("not an SignerContext")
// 	}
// 	if !cc.hasKeyset {
// 		return nil, fmt.Errorf("keyset not found")
// 	}
// 	if cc.keyset == nil {
// 		return nil, fmt.Errorf("keyset is nil")
// 	}
// 	src, err := mpc.NewSource(cc.keyset)
// 	if err != nil {
// 		return nil, err
// 	}
// 	cc.signer = src
// 	cc.hasSigner = true
// 	return src, nil
// }
//
// func GetSource(c echo.Context) (mpc.KeyshareSource, error) {
// 	cc, ok := c.(*SignerContext)
// 	if !ok {
// 		return nil, errors.New("not an SignerContext")
// 	}
// 	if !cc.hasSigner {
// 		return nil, fmt.Errorf("signer not found")
// 	}
// 	if cc.signer == nil {
// 		return nil, fmt.Errorf("signer is nil")
// 	}
// 	return cc.signer, nil
// }
