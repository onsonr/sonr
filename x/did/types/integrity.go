package types

import "lukechampine.com/blake3"

func (g *GlobalIntegrity) Update(address string) bool {
	return true
}

func (g *GlobalIntegrity) getProof() (*Proof, error) {
	if g.Accumulator == nil {
		return NewProof(g.Controller, g.proofProperty(), g.seedKdf())
	}
	return &Proof{
		Issuer:      g.Controller,
		Property:    g.proofProperty(),
		Accumulator: g.Accumulator,
	}, nil
}

func (g *GlobalIntegrity) proofProperty() string {
	return "did:sonr:integrity"
}

func (g *GlobalIntegrity) seedKdf() []byte {
	res := blake3.Sum256([]byte(g.Seed))
	return res[:]
}
