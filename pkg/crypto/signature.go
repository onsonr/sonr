package crypto

import (
	"errors"

	"github.com/taurusgroup/multi-party-sig/pkg/ecdsa"
	"github.com/taurusgroup/multi-party-sig/pkg/math/curve"
)

const (
	// These fields provide convenient access to each of the words of the
	// secp256k1 curve group order N to improve code readability.
	//
	// The group order of the curve per [SECG] is:
	// 0xffffffff ffffffff ffffffff fffffffe baaedce6 af48a03b bfd25e8c d0364141
	orderWordZero  uint32 = 0xd0364141
	orderWordOne   uint32 = 0xbfd25e8c
	orderWordTwo   uint32 = 0xaf48a03b
	orderWordThree uint32 = 0xbaaedce6
	orderWordFour  uint32 = 0xfffffffe
	orderWordFive  uint32 = 0xffffffff
	orderWordSix   uint32 = 0xffffffff
	orderWordSeven uint32 = 0xffffffff

	// These fields provide convenient access to each of the words of the two's
	// complement of the secp256k1 curve group order N to improve code
	// readability.
	//
	// The two's complement of the group order is:
	// 0x00000000 00000000 00000000 00000001 45512319 50b75fc4 402da173 2fc9bebf
	orderComplementWordZero  uint32 = (^orderWordZero) + 1
	orderComplementWordOne   uint32 = ^orderWordOne
	orderComplementWordTwo   uint32 = ^orderWordTwo
	orderComplementWordThree uint32 = ^orderWordThree
	//orderComplementWordFour  uint32 = ^orderWordFour  // unused
	//orderComplementWordFive  uint32 = ^orderWordFive  // unused
	//orderComplementWordSix   uint32 = ^orderWordSix   // unused
	//orderComplementWordSeven uint32 = ^orderWordSeven // unused

	// These fields provide convenient access to each of the words of the
	// secp256k1 curve group order N / 2 to improve code readability and avoid
	// the need to recalculate them.
	//
	// The half order of the secp256k1 curve group is:
	// 0x7fffffff ffffffff ffffffff ffffffff 5d576e73 57a4501d dfe92f46 681b20a0
	halfOrderWordZero  uint32 = 0x681b20a0
	halfOrderWordOne   uint32 = 0xdfe92f46
	halfOrderWordTwo   uint32 = 0x57a4501d
	halfOrderWordThree uint32 = 0x5d576e73
	halfOrderWordFour  uint32 = 0xffffffff
	halfOrderWordFive  uint32 = 0xffffffff
	halfOrderWordSix   uint32 = 0xffffffff
	halfOrderWordSeven uint32 = 0x7fffffff

	// uint32Mask is simply a mask with all bits set for a uint32 and is used to
	// improve the readability of the code.
	uint32Mask = 0xffffffff
)

// constantTimeEq returns 1 if a == b or 0 otherwise in constant time.
func constantTimeEq(a, b uint32) uint32 {
	return uint32((uint64(a^b) - 1) >> 63)
}

// constantTimeNotEq returns 1 if a != b or 0 otherwise in constant time.
func constantTimeNotEq(a, b uint32) uint32 {
	return ^uint32((uint64(a^b)-1)>>63) & 1
}

// constantTimeLess returns 1 if a < b or 0 otherwise in constant time.
func constantTimeLess(a, b uint32) uint32 {
	return uint32((uint64(a) - uint64(b)) >> 63)
}

// constantTimeLessOrEq returns 1 if a <= b or 0 otherwise in constant time.
func constantTimeLessOrEq(a, b uint32) uint32 {
	return uint32((uint64(a) - uint64(b) - 1) >> 63)
}

// constantTimeGreater returns 1 if a > b or 0 otherwise in constant time.
func constantTimeGreater(a, b uint32) uint32 {
	return constantTimeLess(b, a)
}

// constantTimeGreaterOrEq returns 1 if a >= b or 0 otherwise in constant time.
func constantTimeGreaterOrEq(a, b uint32) uint32 {
	return constantTimeLessOrEq(b, a)
}

// constantTimeMin returns min(a,b) in constant time.
func constantTimeMin(a, b uint32) uint32 {
	return b ^ ((a ^ b) & -constantTimeLess(a, b))
}

// IsOverHalfOrder returns whether or not the scalar exceeds the group order
// divided by 2 in constant time.
func IsOverHalfOrder(S curve.Scalar) bool {
	// The intuition here is that the scalar is greater than half of the group
	// order if one of the higher individual words is greater than the
	// corresponding word of the half group order and all higher words in the
	// scalar are equal to their corresponding word of the half group order.
	//
	// Note that the words 4, 5, and 6 are all the max uint32 value, so there is
	// no need to test if those individual words of the scalar exceeds them,
	// hence, only equality is checked for them.
	// result := constantTimeGreater(S.Act(curve.Point), halfOrderWordSeven)
	// highWordsEqual := constantTimeEq(s.n[7], halfOrderWordSeven)
	// highWordsEqual &= constantTimeEq(s.n[6], halfOrderWordSix)
	// highWordsEqual &= constantTimeEq(s.n[5], halfOrderWordFive)
	// highWordsEqual &= constantTimeEq(s.n[4], halfOrderWordFour)
	// result |= highWordsEqual & constantTimeGreater(s.n[3], halfOrderWordThree)
	// highWordsEqual &= constantTimeEq(s.n[3], halfOrderWordThree)
	// result |= highWordsEqual & constantTimeGreater(s.n[2], halfOrderWordTwo)
	// highWordsEqual &= constantTimeEq(s.n[2], halfOrderWordTwo)
	// result |= highWordsEqual & constantTimeGreater(s.n[1], halfOrderWordOne)
	// highWordsEqual &= constantTimeEq(s.n[1], halfOrderWordOne)
	// result |= highWordsEqual & constantTimeGreater(s.n[0], halfOrderWordZero)

	// return result != 0
	return true
}

// SerializeECDSA marshals an ECDSA signature to DER format for use with the CMP protocol
func SerializeECDSA(sig *ecdsa.Signature) []byte {
	rBuf, _ := sig.R.MarshalBinary()
	sBuf, _ := sig.S.MarshalBinary()
	// The format of a DER encoded signature is as follows:
	//
	// 0x30 <total length> 0x02 <length of R> <R> 0x02 <length of S> <S>
	//   - 0x30 is the ASN.1 identifier for a sequence.
	//   - Total length is 1 byte and specifies length of all remaining data.
	//   - 0x02 is the ASN.1 identifier that specifies an integer follows.
	//   - Length of R is 1 byte and specifies how many bytes R occupies.
	//   - R is the arbitrary length big-endian encoded number which
	//     represents the R value of the signature.  DER encoding dictates
	//     that the value must be encoded using the minimum possible number
	//     of bytes.  This implies the first byte can only be null if the
	//     highest bit of the next byte is set in order to prevent it from
	//     being interpreted as a negative number.
	//   - 0x02 is once again the ASN.1 integer identifier.
	//   - Length of S is 1 byte and specifies how many bytes S occupies.
	//   - S is the arbitrary length big-endian encoded number which
	//     represents the S value of the signature.  The encoding rules are
	//     identical as those for R.

	// Ensure the S component of the signature is less than or equal to the half
	// order of the group because both S and its negation are valid signatures
	// modulo the order, so this forces a consistent choice to reduce signature
	// malleability.
	// TODO

	// Ensure the encoded bytes for the R and S components are canonical per DER
	// by trimming all leading zero bytes so long as the next byte does not have
	// the high bit set and it's not the final byte.
	canonR, canonS := rBuf[:], sBuf[:]
	for len(canonR) > 1 && canonR[0] == 0x00 && canonR[1]&0x80 == 0 {
		canonR = canonR[1:]
	}
	for len(canonS) > 1 && canonS[0] == 0x00 && canonS[1]&0x80 == 0 {
		canonS = canonS[1:]
	}

	// Total length of returned signature is 1 byte for each magic and length
	// (6 total), plus lengths of R and S.
	totalLen := 6 + len(canonR) + len(canonS)
	b := make([]byte, 0, totalLen)
	b = append(b, asn1SequenceID)
	b = append(b, byte(totalLen-2))
	b = append(b, asn1IntegerID)
	b = append(b, byte(len(canonR)))
	b = append(b, canonR...)
	b = append(b, asn1IntegerID)
	b = append(b, byte(len(canonS)))
	b = append(b, canonS...)
	return b
}

const (
	// asn1SequenceID is the ASN.1 identifier for a sequence and is used when
	// parsing and serializing signatures encoded with the Distinguished
	// Encoding Rules (DER) format per section 10 of [ISO/IEC 8825-1].
	asn1SequenceID = 0x30

	// asn1IntegerID is the ASN.1 identifier for an integer and is used when
	// parsing and serializing signatures encoded with the Distinguished
	// Encoding Rules (DER) format per section 10 of [ISO/IEC 8825-1].
	asn1IntegerID = 0x02
)

// - The R and S values must be in the valid range for secp256k1 scalars:
//   - Negative values are rejected
//   - Zero is rejected
//   - Values greater than or equal to the secp256k1 group order are rejected
func ParseDERSignature(sig []byte) (*ecdsa.Signature, error) {
	// The format of a DER encoded signature for secp256k1 is as follows:
	//
	// 0x30 <total length> 0x02 <length of R> <R> 0x02 <length of S> <S>
	//   - 0x30 is the ASN.1 identifier for a sequence
	//   - Total length is 1 byte and specifies length of all remaining data
	//   - 0x02 is the ASN.1 identifier that specifies an integer follows
	//   - Length of R is 1 byte and specifies how many bytes R occupies
	//   - R is the arbitrary length big-endian encoded number which
	//     represents the R value of the signature.  DER encoding dictates
	//     that the value must be encoded using the minimum possible number
	//     of bytes.  This implies the first byte can only be null if the
	//     highest bit of the next byte is set in order to prevent it from
	//     being interpreted as a negative number.
	//   - 0x02 is once again the ASN.1 integer identifier
	//   - Length of S is 1 byte and specifies how many bytes S occupies
	//   - S is the arbitrary length big-endian encoded number which
	//     represents the S value of the signature.  The encoding rules are
	//     identical as those for R.
	//
	// NOTE: The DER specification supports specifying lengths that can occupy
	// more than 1 byte, however, since this is specific to secp256k1
	// signatures, all lengths will be a single byte.
	const (
		// minSigLen is the minimum length of a DER encoded signature and is
		// when both R and S are 1 byte each.
		//
		// 0x30 + <1-byte> + 0x02 + 0x01 + <byte> + 0x2 + 0x01 + <byte>
		minSigLen = 8

		// maxSigLen is the maximum length of a DER encoded signature and is
		// when both R and S are 33 bytes each.  It is 33 bytes because a
		// 256-bit integer requires 32 bytes and an additional leading null byte
		// might be required if the high bit is set in the value.
		//
		// 0x30 + <1-byte> + 0x02 + 0x21 + <33 bytes> + 0x2 + 0x21 + <33 bytes>
		maxSigLen = 72

		// sequenceOffset is the byte offset within the signature of the
		// expected ASN.1 sequence identifier.
		sequenceOffset = 0

		// dataLenOffset is the byte offset within the signature of the expected
		// total length of all remaining data in the signature.
		dataLenOffset = 1

		// rTypeOffset is the byte offset within the signature of the ASN.1
		// identifier for R and is expected to indicate an ASN.1 integer.
		rTypeOffset = 2

		// rLenOffset is the byte offset within the signature of the length of
		// R.
		rLenOffset = 3

		// rOffset is the byte offset within the signature of R.
		rOffset = 4
	)

	// The signature must adhere to the minimum and maximum allowed length.
	sigLen := len(sig)
	if sigLen < minSigLen {
		return nil, errors.New("malformed signature: too short")
	}
	if sigLen > maxSigLen {
		return nil, errors.New("malformed signature: too long")
	}

	// The signature must start with the ASN.1 sequence identifier.
	if sig[sequenceOffset] != asn1SequenceID {
		return nil, errors.New("malformed signature: format has wrong type")
	}

	// The signature must indicate the correct amount of data for all elements
	// related to R and S.
	if int(sig[dataLenOffset]) != sigLen-2 {
		return nil, errors.New("malformed signature: bad length")
	}

	// Calculate the offsets of the elements related to S and ensure S is inside
	// the signature.
	//
	// rLen specifies the length of the big-endian encoded number which
	// represents the R value of the signature.
	//
	// sTypeOffset is the offset of the ASN.1 identifier for S and, like its R
	// counterpart, is expected to indicate an ASN.1 integer.
	//
	// sLenOffset and sOffset are the byte offsets within the signature of the
	// length of S and S itself, respectively.
	rLen := int(sig[rLenOffset])
	sTypeOffset := rOffset + rLen
	sLenOffset := sTypeOffset + 1
	if sTypeOffset >= sigLen {
		return nil, errors.New("malformed signature: S type indicator missing")
	}
	if sLenOffset >= sigLen {
		return nil, errors.New("malformed signature: S length missing")
	}

	// The lengths of R and S must match the overall length of the signature.
	//
	// sLen specifies the length of the big-endian encoded number which
	// represents the S value of the signature.
	sOffset := sLenOffset + 1
	sLen := int(sig[sLenOffset])
	if sOffset+sLen != sigLen {
		return nil, errors.New("malformed signature: invalid S length")
	}

	// R elements must be ASN.1 integers.
	if sig[rTypeOffset] != asn1IntegerID {
		return nil, errors.New("malformed signature: R integer marker invalid")
	}

	// Zero-length integers are not allowed for R.
	if rLen == 0 {
		return nil, errors.New("malformed signature: R length is zero")
	}

	// R must not be negative.
	if sig[rOffset]&0x80 != 0 {

		return nil, errors.New("malformed signature: R is negative")
	}

	// Null bytes at the start of R are not allowed, unless R would otherwise be
	// interpreted as a negative number.
	if rLen > 1 && sig[rOffset] == 0x00 && sig[rOffset+1]&0x80 == 0 {
		return nil, errors.New("malformed signature: R value has too much padding")
	}

	// S elements must be ASN.1 integers.
	if sig[sTypeOffset] != asn1IntegerID {
		return nil, errors.New("malformed signature: S integer marker invalid")
	}

	// Zero-length integers are not allowed for S.
	if sLen == 0 {
		return nil, errors.New("malformed signature: S length is zero")
	}

	// S must not be negative.
	if sig[sOffset]&0x80 != 0 {
		return nil, errors.New("malformed signature: S is negative")
	}

	// Null bytes at the start of S are not allowed, unless S would otherwise be
	// interpreted as a negative number.
	if sLen > 1 && sig[sOffset] == 0x00 && sig[sOffset+1]&0x80 == 0 {
		return nil, errors.New("malformed signature: S value has too much padding")
	}

	// The signature is validly encoded per DER at this point, however, enforce
	// additional restrictions to ensure R and S are in the range [1, N-1] since
	// valid ECDSA signatures are required to be in that range per spec.
	//
	// Also note that while the overflow checks are required to make use of the
	// specialized mod N scalar type, rejecting zero here is not strictly
	// required because it is also checked when verifying the signature, but
	// there really isn't a good reason not to fail early here on signatures
	// that do not conform to the ECDSA spec.

	// Strip leading zeroes from R.
	rBytes := sig[rOffset : rOffset+rLen]
	for len(rBytes) > 0 && rBytes[0] == 0x00 {
		rBytes = rBytes[1:]
	}

	// R must be in the range [1, N-1].  Notice the check for the maximum number
	// of bytes is required because SetByteSlice truncates as noted in its
	// comment so it could otherwise fail to detect the overflow.
	crv := curve.Secp256k1{}
	r := crv.NewPoint()
	if r.UnmarshalBinary(rBytes) != nil {
		return nil, errors.New("malformed signature: R is not in the range [1, N-1]")
	}

	// Strip leading zeroes from S.
	sBytes := sig[sOffset : sOffset+sLen]
	for len(sBytes) > 0 && sBytes[0] == 0x00 {
		sBytes = sBytes[1:]
	}

	// S must be in the range [1, N-1].  Notice the check for the maximum number
	// of bytes is required because SetByteSlice truncates as noted in its
	// comment so it could otherwise fail to detect the overflow.
	s := crv.NewScalar()
	if s.UnmarshalBinary(sBytes) != nil {
		return nil, errors.New("malformed signature: S is not in the range [1, N-1]")
	}

	// Create and return the signature.
	return &ecdsa.Signature{R: r, S: s}, nil
}
