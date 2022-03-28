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
	"math/big"

	"github.com/pkg/errors"
	"github.com/sonr-io/core/crypto/common"
	"github.com/sonr-io/core/crypto/df"
	"github.com/sonr-io/core/crypto/pedersen"
	"github.com/sonr-io/core/crypto/qr"
)

// TODO probably doesn't make much sense if sec is unexported, remove
type KeyPair struct {
	Sec *SecKey
	Pub *PubKey
}

// SecKey is a secret key for the CL scheme.
type SecKey struct {
	RsaPrimes                  *qr.RSASpecialPrimes
	AttributesSpecialRSAPrimes *qr.RSASpecialPrimes
}

// NewSecKey accepts group g and commitment receiver cr, and returns
// new secret key for the CL scheme.
func NewSecKey(g *qr.RSASpecial, cr *df.Receiver) *SecKey {
	return &SecKey{
		RsaPrimes:                  g.GetPrimes(),
		AttributesSpecialRSAPrimes: cr.QRSpecialRSA.GetPrimes(),
	}
}

// PubKey is a public key for the CL scheme.
type PubKey struct {
	N              *big.Int
	S              *big.Int
	Z              *big.Int
	RsKnown        []*big.Int // one R corresponds to one attribute - these attributes are known to both - receiver and issuer
	RsCommitted    []*big.Int // issuer knows only commitments of these attributes
	RsHidden       []*big.Int // only receiver knows these attributes
	PedersenParams *pedersen.Params
	// the fields below are for commitments of the (committed) attributes
	N1 *big.Int
	G  *big.Int
	H  *big.Int
}

// NewPubKey accepts group g, parameters p and commitment receiver recv,
// and returns a public key for the CL scheme.
func NewPubKey(g *qr.RSASpecial, p *Params,
	attrs *AttrCount, recv *df.Receiver) (*PubKey,
	error) {
	S, Z, RsKnown, RsCommitted, RsHidden, err := generateQuadraticResidues(
		g, attrs.Known, attrs.Committed, attrs.Hidden)
	if err != nil {
		return nil, errors.Wrap(err, "error creating quadratic residues")
	}

	pp, err := pedersen.GenerateParams(int(p.RhoBitLen))
	if err != nil {
		return nil, errors.Wrap(err, "error creating Pedersen receiver")
	}

	return &PubKey{
		N:              g.N,
		S:              S,
		Z:              Z,
		RsKnown:        RsKnown,
		RsCommitted:    RsCommitted,
		RsHidden:       RsHidden,
		PedersenParams: pp,
		N1:             recv.QRSpecialRSA.N,
		G:              recv.G,
		H:              recv.H,
	}, nil
}

// GenerateUserMasterSecret generates a secret key that needs to be encoded into every user's credential as a
// sharing prevention mechanism.
func (k *PubKey) GenerateUserMasterSecret() *big.Int {
	return common.GetRandomInt(k.PedersenParams.Group.Q)
}

// GetContext concatenates public parameters and returns a corresponding number.
func (k *PubKey) GetContext() *big.Int {
	numbers := []*big.Int{k.N, k.S, k.Z}
	numbers = append(numbers, k.RsKnown...)
	numbers = append(numbers, k.RsCommitted...)
	numbers = append(numbers, k.RsHidden...)
	concatenated := common.ConcatenateNumbers(numbers...)
	return new(big.Int).SetBytes(concatenated)
}

// GenerateKeyPair takes and constructs a keypair containing public and
// secret key for the CL scheme.
func GenerateKeyPair(p *Params, attrs *AttrCount) (*KeyPair, error) {
	g, err := qr.NewRSASpecial(int(p.NLength) / 2)
	if err != nil {
		return nil, errors.Wrap(err, "error creating RSASpecial group")
	}

	// receiver for commitments of (committed) attributes:
	commRecv, err := df.NewReceiver(int(p.NLength/2), int(p.SecParam))
	if err != nil {
		return nil, errors.Wrap(err, "error creating DF commitment receiver")
	}

	sk := NewSecKey(g, commRecv)

	pk, err := NewPubKey(g, p, attrs, commRecv)
	if err != nil {
		return nil, err
	}

	return &KeyPair{
		Sec: sk,
		Pub: pk,
	}, nil
}
