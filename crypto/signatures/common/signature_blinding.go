//
// Copyright Coinbase, Inc. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0
//

package common

import (
	"github.com/di-dao/core/crypto/core/curves"
)

// SignatureBlinding is a value used for computing blind signatures
type SignatureBlinding = curves.PairingScalar
