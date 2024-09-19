package orm

import "github.com/onsonr/sonr/internal/orm/keyalgorithm"

type COSEAlgorithmIdentifier int

func GetCoseIdentifier(alg keyalgorithm.KeyAlgorithm) COSEAlgorithmIdentifier {
	switch alg {
	case keyalgorithm.Es256:
		return COSEAlgorithmIdentifier(-7)
	case keyalgorithm.Es384:
		return COSEAlgorithmIdentifier(-35)
	case keyalgorithm.Es512:
		return COSEAlgorithmIdentifier(-36)
	case keyalgorithm.Eddsa:
		return COSEAlgorithmIdentifier(-8)
	case keyalgorithm.Es256k:
		return COSEAlgorithmIdentifier(-10)
	default:
		return COSEAlgorithmIdentifier(0)
	}
}
