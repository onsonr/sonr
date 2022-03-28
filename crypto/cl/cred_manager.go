/*
 * Copyright 2017 XLAB d.o.o.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package cl

import (
	"fmt"
	"log"
	"math/big"

	"github.com/pkg/errors"
	"github.com/sonr-io/core/crypto/common"
	"github.com/sonr-io/core/crypto/df"
	"github.com/sonr-io/core/crypto/pedersen"
	"github.com/sonr-io/core/crypto/qr"
	"github.com/sonr-io/core/crypto/schnorr"
)

// CredManager manages a single instance of anonymous credential.
//
// An instance of this struct should be created by a user before
// she wants a new credential to be issued, or an existing one to
// be updated or proved.
//
// When a user needs a new credential under a new nym, she also needs
// a new instance of CredManager.
type CredManager struct {
	Params             *Params
	PubKey             *PubKey
	RawCred            *RawCred
	nymCommitter       *pedersen.Committer // nym is actually a commitment to masterSecret
	Nym                *big.Int
	masterSecret       *big.Int
	Attrs              *Attrs
	CommitmentsOfAttrs []*big.Int // commitments of committedAttrs
	// V1 is a random element in credential - it is generated in GetCredRequest and needed when
	// proving the possesion of a credential - this is why it is stored in User and not in UserCredentialReceiver
	V1                        *big.Int            // v1 is random element in U; U = S^v1 * R_i^m_i where m_i are hidden attributes
	attrsCommitters           []*df.Committer     // committers for committedAttrs
	commitmentsOfAttrsProvers []*df.OpeningProver // for proving that you know how to open CommitmentsOfAttrs
	CredReqNonce              *big.Int
}

type Attrs struct {
	// attributes that are known to the credential receiver and issuer
	Known []*big.Int
	// attributes which are known only to the credential receiver
	Hidden []*big.Int
	// attributes for which the issuer knows only commitment
	Committed []*big.Int
}

func NewAttrs(known, committed, hidden []*big.Int) *Attrs {
	return &Attrs{
		Known:     known,
		Hidden:    hidden,
		Committed: committed,
	}
}

func (a *Attrs) join() []*big.Int {
	all := make([]*big.Int, 0)
	all = append(all, a.Known...)
	all = append(all, a.Hidden...)
	all = append(all, a.Committed...)
	return all
}

func checkBitLen(s []*big.Int, len int) bool {
	for _, ss := range s {
		if ss.BitLen() > len {
			return false
		}
	}

	return true
}

func NewCredManager(params *Params, pubKey *PubKey,
	masterSecret *big.Int, rawCred *RawCred) (*CredManager, error) {
	if err := rawCred.missingAttrs(); err != nil {
		return nil, errors.Wrap(err, "not all expected attributes"+
			" are present in the raw credential")
	}

	known := rawCred.GetKnownVals()
	committed := rawCred.GetCommittedVals()
	hidden := []*big.Int{} // currently not used

	attrs := NewAttrs(known, committed, hidden)
	if !checkBitLen(attrs.join(), int(params.AttrBitLen)) {
		return nil, fmt.Errorf("attributes length not ok")
	}

	attrsCommitters := make([]*df.Committer, len(attrs.Committed))
	commitmentsOfAttrs := make([]*big.Int, len(attrs.Committed))
	for i, attr := range attrs.Committed {
		committer := df.NewCommitter(pubKey.N1, pubKey.G, pubKey.H,
			pubKey.N1, int(params.SecParam))
		com, err := committer.GetCommitMsg(attr)
		if err != nil {
			return nil, fmt.Errorf("error when creating Pedersen commitment: %s", err)
		}
		commitmentsOfAttrs[i] = com
		attrsCommitters[i] = committer
	}
	commitmentsOfAttrsProvers := make([]*df.OpeningProver, len(commitmentsOfAttrs))
	for i := range commitmentsOfAttrs {
		prover := df.NewOpeningProver(attrsCommitters[i],
			int(params.ChallengeSpace))
		commitmentsOfAttrsProvers[i] = prover
	}

	credManager := CredManager{
		Params:                    params,
		PubKey:                    pubKey,
		RawCred:                   rawCred,
		Attrs:                     attrs,
		CommitmentsOfAttrs:        commitmentsOfAttrs,
		attrsCommitters:           attrsCommitters,
		commitmentsOfAttrsProvers: commitmentsOfAttrsProvers,
		masterSecret:              masterSecret,
	}
	credManager.generateNym()

	return &credManager, nil
}

// generateNym creates a pseudonym to be used with a given organization. Authentication can be done
// with respect to the pseudonym or not.
func (m *CredManager) generateNym() error {
	committer := pedersen.NewCommitter(m.PubKey.PedersenParams)
	nym, err := committer.GetCommitMsg(m.masterSecret)
	if err != nil {
		return fmt.Errorf("error when creating Pedersen commitment: %s", err)
	}
	m.Nym = nym
	m.nymCommitter = committer

	return nil
}

// GetCredRequest computes U and returns CredRequest which contains:
// - proof data for proving that nym was properly generated,
// - U and proof data that U was properly generated,
// - proof data for proving the knowledge of opening for commitments of attributes (for those attributes
// for which the committed value is known).
func (m *CredManager) GetCredRequest(nonceOrg *big.Int) (*CredRequest, error) {
	U, v1 := m.computeU()
	m.V1 = v1
	nymProver, uProver, err := m.getCredReqProvers(U)
	if err != nil {
		return nil, err
	}

	challenge := m.getCredReqChallenge(U, m.Nym, nonceOrg)
	commitmentsOfAttrsProofs := m.getCommitmentsOfAttrsProof(challenge)

	b := new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(m.Params.SecParam)), nil)
	nonce := common.GetRandomInt(b)
	m.CredReqNonce = nonce

	uProofRandomData, err := m.getUProofRandomData(uProver)
	if err != nil {
		return nil, err
	}

	return NewCredRequest(m.Nym, m.Attrs.Known, m.CommitmentsOfAttrs,
		schnorr.NewProof(nymProver.GetProofRandomData(), challenge,
			nymProver.GetProofData(challenge)), U,
		qr.NewRepresentationProof(uProofRandomData, challenge,
			uProver.GetProofData(challenge)),
		commitmentsOfAttrsProofs, nonce), nil
}

// Verify verifies anonymous credential cred, returning a boolean indicating
// success or failure of credential verification.
// When verification process fails due to misconfiguration, error is returned.
func (m *CredManager) Verify(cred *Cred, AProof *qr.RepresentationProof) (bool, error) {
	// check bit length of e:
	b1 := new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(m.Params.EBitLen-1)), nil)
	b22 := new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(m.Params.E1BitLen-1)), nil)
	b2 := new(big.Int).Add(b1, b22)

	if (cred.E.Cmp(b1) != 1) || (b2.Cmp(cred.E) != 1) {
		return false, fmt.Errorf("e is not of the proper bit length")
	}
	// check that e is prime
	if !cred.E.ProbablyPrime(20) {
		return false, fmt.Errorf("e is not prime")
	}

	v := new(big.Int).Add(m.V1, cred.V11)
	group := qr.NewRSApecialPublic(m.PubKey.N)
	// denom = S^v * R_1^attr_1 * ... * R_j^attr_j
	denom := group.Exp(m.PubKey.S, v) // s^v
	for i := 0; i < len(m.Attrs.Known); i++ {
		t1 := group.Exp(m.PubKey.RsKnown[i], m.Attrs.Known[i])
		denom = group.Mul(denom, t1)
	}

	for i := 0; i < len(m.Attrs.Committed); i++ {
		t1 := group.Exp(m.PubKey.RsCommitted[i], m.CommitmentsOfAttrs[i])
		denom = group.Mul(denom, t1)
	}

	for i := 0; i < len(m.Attrs.Hidden); i++ {
		t1 := group.Exp(m.PubKey.RsHidden[i], m.Attrs.Hidden[i])
		denom = group.Mul(denom, t1)
	}

	denomInv := group.Inv(denom)
	Q := group.Mul(m.PubKey.Z, denomInv)
	Q1 := group.Exp(cred.A, cred.E)
	if Q1.Cmp(Q) != 0 {
		return false, fmt.Errorf("Q should be A^e (mod n)")
	}

	// verify signature proof:
	ver := qr.NewRepresentationVerifier(group, int(m.Params.SecParam))
	ver.SetProofRandomData(AProof.ProofRandomData, []*big.Int{Q}, cred.A)
	// check challenge
	context := m.PubKey.GetContext()
	c := common.Hash(context, Q, cred.A, AProof.ProofRandomData, m.CredReqNonce)
	if AProof.Challenge.Cmp(c) != 0 {
		return false, fmt.Errorf("challenge is not correct")
	}

	ver.SetChallenge(AProof.Challenge)

	return ver.Verify(AProof.ProofData), nil
}

// Update updates credential.
func (m *CredManager) Update(c *RawCred) {
	m.RawCred = c
	m.Attrs.Known = m.RawCred.GetKnownVals()
}

// FilterAttributes returns only attributes to be revealed to the verifier.
func (m *CredManager) FilterAttributes(revealedKnownAttrsIndices,
	revealedCommitmentsOfAttrsIndices []int) ([]*big.Int, []*big.Int) {
	revealedKnownAttrs := []*big.Int{}
	revealedCommitmentsOfAttrs := []*big.Int{}
	for i := 0; i < len(m.Attrs.Known); i++ {
		if common.Contains(revealedKnownAttrsIndices, i) {
			revealedKnownAttrs = append(revealedKnownAttrs, m.Attrs.Known[i])
		}
	}
	for i := 0; i < len(m.CommitmentsOfAttrs); i++ {
		if common.Contains(revealedCommitmentsOfAttrsIndices, i) {
			revealedCommitmentsOfAttrs = append(revealedCommitmentsOfAttrs, m.CommitmentsOfAttrs[i])
		}
	}

	return revealedKnownAttrs, revealedCommitmentsOfAttrs
}

// randomize randomizes credential cred, and returns the
// randomized credential as a new credential.
func (m *CredManager) randomize(cred *Cred) *Cred {
	b := new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(m.Params.NLength+m.Params.SecParam)), nil)
	r := common.GetRandomInt(b)
	group := qr.NewRSApecialPublic(m.PubKey.N)
	t := group.Exp(m.PubKey.S, r)
	A := group.Mul(cred.A, t) // cred.A * S^r
	t = new(big.Int).Mul(cred.E, r)
	v11 := new(big.Int).Sub(cred.V11, t) // cred.v11 - e*r (in Z)

	t = new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(m.Params.EBitLen-1)), nil)
	log.Println(t)
	return NewCred(A, cred.E, v11)
}

func (m *CredManager) GetProofChallenge(credProofRandomData, nonceOrg *big.Int) *big.Int {
	context := m.PubKey.GetContext()
	l := []*big.Int{context, credProofRandomData, nonceOrg}
	//l = append(l, ...) // TODO: add other values

	return common.Hash(l...)
}

// BuildProof builds a proof of knowledge for the given credential.
func (m *CredManager) BuildProof(cred *Cred, revealedKnownAttrsIndices,
	revealedCommitmentsOfAttrsIndices []int, nonceOrg *big.Int) (*Cred,
	*qr.RepresentationProof, error) {
	if m.V1 == nil {
		return nil, nil, fmt.Errorf("v1 is not set (generated in GetCredRequest)")
	}
	rCred := m.randomize(cred)
	// Z = cred.A^cred.e * S^cred.v11 * R_1^m_1 * ... * R_l^m_l
	// Z = rCred.A^rCred.e * S^rCred.v11 * R_1^m_1 * ... * R_l^m_l
	group := qr.NewRSApecialPublic(m.PubKey.N)

	bases := []*big.Int{}
	unrevealedKnownAttrs := []*big.Int{}
	unrevealedCommitmentsOfAttrs := []*big.Int{}
	for i := 0; i < len(m.Attrs.Known); i++ {
		if !common.Contains(revealedKnownAttrsIndices, i) {
			bases = append(bases, m.PubKey.RsKnown[i])
			unrevealedKnownAttrs = append(unrevealedKnownAttrs, m.Attrs.Known[i])
		}
	}
	for i := 0; i < len(m.CommitmentsOfAttrs); i++ {
		if !common.Contains(revealedCommitmentsOfAttrsIndices, i) {
			bases = append(bases, m.PubKey.RsCommitted[i])
			unrevealedCommitmentsOfAttrs = append(unrevealedCommitmentsOfAttrs, m.CommitmentsOfAttrs[i])
		}
	}

	bases = append(bases, m.PubKey.RsHidden...)
	bases = append(bases, rCred.A)
	bases = append(bases, m.PubKey.S)

	secrets := append(unrevealedKnownAttrs, unrevealedCommitmentsOfAttrs...)
	secrets = append(secrets, m.Attrs.Hidden...)
	secrets = append(secrets, rCred.E)
	v := new(big.Int).Add(rCred.V11, m.V1)
	secrets = append(secrets, v)

	denom := big.NewInt(1)
	for i := 0; i < len(m.Attrs.Known); i++ {
		if common.Contains(revealedKnownAttrsIndices, i) {
			t1 := group.Exp(m.PubKey.RsKnown[i], m.Attrs.Known[i])
			denom = group.Mul(denom, t1)
		}
	}

	for i := 0; i < len(m.Attrs.Committed); i++ {
		if common.Contains(revealedCommitmentsOfAttrsIndices, i) {
			t1 := group.Exp(m.PubKey.RsCommitted[i], m.CommitmentsOfAttrs[i])
			denom = group.Mul(denom, t1)
		}
	}
	denomInv := group.Inv(denom)
	y := group.Mul(m.PubKey.Z, denomInv)

	prover := qr.NewRepresentationProver(group, int(m.Params.SecParam),
		secrets, bases, y)

	// boundary for m_tilde
	b_m := int(m.Params.AttrBitLen + m.Params.SecParam + m.Params.HashBitLen)
	// boundary for e
	b_e := int(m.Params.EBitLen + m.Params.SecParam + m.Params.HashBitLen)
	// boundary for v1
	b_v1 := int(m.Params.VBitLen + m.Params.SecParam + m.Params.HashBitLen)

	boundaries := []int{}
	for i := 0; i < len(unrevealedKnownAttrs); i++ {
		boundaries = append(boundaries, b_m)
	}
	for i := 0; i < len(unrevealedCommitmentsOfAttrs); i++ {
		boundaries = append(boundaries, b_m)
	}
	for range m.PubKey.RsHidden {
		boundaries = append(boundaries, b_m)
	}
	boundaries = append(boundaries, b_e)
	boundaries = append(boundaries, b_v1)

	proofRandomData, err := prover.GetProofRandomDataGivenBoundaries(boundaries, true)
	if err != nil {
		return nil, nil, fmt.Errorf("error when generating representation proof random data: %s", err)
	}

	challenge := m.GetProofChallenge(proofRandomData, nonceOrg)
	proofData := prover.GetProofData(challenge)

	return rCred, qr.NewRepresentationProof(proofRandomData, challenge, proofData), nil
}
