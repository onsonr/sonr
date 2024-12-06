package ucan

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/golang-jwt/jwt"
	"github.com/onsonr/sonr/crypto/mpc"
	"github.com/onsonr/sonr/crypto/ucan/attns/capability"
	"github.com/onsonr/sonr/crypto/ucan/attns/policytype"
	"github.com/onsonr/sonr/crypto/ucan/attns/resourcetype"
)

type ucanKeyshare struct {
	userShare *mpc.UserKeyshare
	valShare  *mpc.ValKeyshare

	addr      string
	issuerDID string
}

func (k ucanKeyshare) NewOriginToken(audienceDID string, att Attenuations, fct []Fact, notBefore, expires time.Time) (*Token, error) {
	return k.newToken(audienceDID, nil, att, fct, notBefore, expires)
}

func (k ucanKeyshare) NewAttenuatedToken(parent *Token, audienceDID string, att Attenuations, fct []Fact, nbf, exp time.Time) (*Token, error) {
	if !parent.Attenuations.Contains(att) {
		return nil, fmt.Errorf("scope of ucan attenuations must be less than it's parent")
	}
	return k.newToken(audienceDID, append(parent.Proofs, Proof(parent.Raw)), att, fct, nbf, exp)
}

func (k ucanKeyshare) newToken(audienceDID string, prf []Proof, att Attenuations, fct []Fact, nbf, exp time.Time) (*Token, error) {
	t := jwt.New(NewJWTSigningMethod("MPC256", k))

	// if _, err := did.Parse(audienceDID); err != nil {
	// 	return nil, fmt.Errorf("invalid audience DID: %w", err)
	// }

	t.Header[UCANVersionKey] = UCANVersion

	var (
		nbfUnix int64
		expUnix int64
	)

	if !nbf.IsZero() {
		nbfUnix = nbf.Unix()
	}
	if !exp.IsZero() {
		expUnix = exp.Unix()
	}

	// set our claims
	t.Claims = &Claims{
		StandardClaims: &jwt.StandardClaims{
			Issuer:    k.issuerDID,
			Audience:  audienceDID,
			NotBefore: nbfUnix,
			// set the expire time
			// see http://tools.ietf.org/html/draft-ietf-oauth-json-web-token-20#section-4.1.4
			ExpiresAt: expUnix,
		},
		Attenuations: att,
		Facts:        fct,
		Proofs:       prf,
	}

	raw, err := t.SignedString(nil)
	if err != nil {
		return nil, err
	}

	return &Token{
		Raw:          raw,
		Attenuations: att,
		Facts:        fct,
		Proofs:       prf,
	}, nil
}

func ComputeIssuerDID(pk []byte) (string, string, error) {
	addr, err := ComputeSonrAddr(pk)
	if err != nil {
		return "", "", err
	}
	return fmt.Sprintf("did:sonr:%s", addr), addr, nil
}

func ComputeSonrAddr(pk []byte) (string, error) {
	sonrAddr, err := bech32.ConvertAndEncode("idx", pk)
	if err != nil {
		return "", err
	}
	return sonrAddr, nil
}

// NewSmartAccount creates default attenuations for a smart account
func NewSmartAccount(
	accountAddr string,
) Attenuations {
	caps := AccountPermissions.GetCapabilities()
	return Attenuations{
		// Owner capabilities
		{Cap: caps.Cap(CapOwner.String()), Rsc: NewResource(ResAccount, accountAddr)},

		// Operation capabilities
		{Cap: caps.Cap(capability.CAPEXECUTE.String()), Rsc: NewResource(ResTransaction, fmt.Sprintf("%s:*", accountAddr))},
		{Cap: caps.Cap(capability.CAPPROPOSE.String()), Rsc: NewResource(ResTransaction, fmt.Sprintf("%s:*", accountAddr))},
		{Cap: caps.Cap(capability.CAPSIGN.String()), Rsc: NewResource(ResTransaction, fmt.Sprintf("%s:*", accountAddr))},

		// Policy capabilities
		{Cap: caps.Cap(capability.CAPSETPOLICY.String()), Rsc: NewResource(ResPolicy, fmt.Sprintf("%s:*", accountAddr))},
		{Cap: caps.Cap(capability.CAPSETTHRESHOLD.String()), Rsc: NewResource(ResPolicy, fmt.Sprintf("%s:threshold", accountAddr))},
	}
}

// NewSmartAccountPolicy creates attenuations for policy management
func NewSmartAccountPolicy(
	accountAddr string,
	policyType policytype.PolicyType,
) Attenuations {
	caps := AccountPermissions.GetCapabilities()
	return Attenuations{
		{
			Cap: caps.Cap(capability.CAPSETPOLICY.String()),
			Rsc: NewResource(
				ResPolicy,
				fmt.Sprintf("%s:%s", accountAddr, policyType),
			),
		},
	}
}

// SmartAccountCapabilities defines the capability hierarchy
func SmartAccountCapabilities() []string {
	return []string{
		CapOwner.String(),
		CapOperator.String(),
		CapObserver.String(),
		CapExecute.String(),
		CapPropose.String(),
		CapSign.String(),
		CapSetPolicy.String(),
		CapSetThreshold.String(),
		CapRecover.String(),
		CapSocial.String(),
	}
}

// CreateVaultAttenuations creates default attenuations for a smart account
func NewService(
	origin string,
) Attenuations {
	caps := ServicePermissions.GetCapabilities()
	return Attenuations{
		// Owner capabilities
		{Cap: caps.Cap(capability.CAPOWNER.String()), Rsc: NewResource(resourcetype.RESACCOUNT, origin)},

		// Operation capabilities
		{Cap: caps.Cap(capability.CAPEXECUTE.String()), Rsc: NewResource(resourcetype.RESTRANSACTION, fmt.Sprintf("%s:*", origin))},
		{Cap: caps.Cap(capability.CAPPROPOSE.String()), Rsc: NewResource(resourcetype.RESTRANSACTION, fmt.Sprintf("%s:*", origin))},
		{Cap: caps.Cap(capability.CAPSIGN.String()), Rsc: NewResource(resourcetype.RESTRANSACTION, fmt.Sprintf("%s:*", origin))},

		// Policy capabilities
		{Cap: caps.Cap(capability.CAPSETPOLICY.String()), Rsc: NewResource(resourcetype.RESPOLICY, fmt.Sprintf("%s:*", origin))},
		{Cap: caps.Cap(capability.CAPSETTHRESHOLD.String()), Rsc: NewResource(resourcetype.RESPOLICY, fmt.Sprintf("%s:threshold", origin))},
	}
}

// ServiceCapabilities defines the capability hierarchy
func ServiceCapabilities() []string {
	return []string{
		CapOwner.String(),
		CapOperator.String(),
		CapObserver.String(),
		CapExecute.String(),
		CapPropose.String(),
		CapSign.String(),
		CapResolver.String(),
		CapProducer.String(),
	}
}

// NewVault creates default attenuations for a smart account
func NewVault(
	kss mpc.Keyset,
) Attenuations {
	accountAddr, err := mpc.ComputeSonrAddr(kss.User().GetPublicKey())
	if err != nil {
		return nil
	}
	caps := VaultPermissions.GetCapabilities()
	return Attenuations{
		// Owner capabilities
		{Cap: caps.Cap(capability.CAPOWNER.String()), Rsc: NewResource(resourcetype.RESACCOUNT, accountAddr)},

		// Operation capabilities
		{Cap: caps.Cap(capability.CAPEXECUTE.String()), Rsc: NewResource(resourcetype.RESTRANSACTION, fmt.Sprintf("%s:*", accountAddr))},
		{Cap: caps.Cap(capability.CAPPROPOSE.String()), Rsc: NewResource(resourcetype.RESTRANSACTION, fmt.Sprintf("%s:*", accountAddr))},
		{Cap: caps.Cap(capability.CAPSIGN.String()), Rsc: NewResource(resourcetype.RESTRANSACTION, fmt.Sprintf("%s:*", accountAddr))},

		// Policy capabilities
		{Cap: caps.Cap(capability.CAPSETPOLICY.String()), Rsc: NewResource(resourcetype.RESPOLICY, fmt.Sprintf("%s:*", accountAddr))},
		{Cap: caps.Cap(capability.CAPSETTHRESHOLD.String()), Rsc: NewResource(resourcetype.RESPOLICY, fmt.Sprintf("%s:threshold", accountAddr))},
	}
}

// NewVaultPolicy creates attenuations for policy management
func NewVaultPolicy(
	accountAddr string,
	policyType policytype.PolicyType,
) Attenuations {
	caps := VaultPermissions.GetCapabilities()
	return Attenuations{
		{
			Cap: caps.Cap(capability.CAPSETPOLICY.String()),
			Rsc: NewResource(
				resourcetype.RESPOLICY,
				fmt.Sprintf("%s:%s", accountAddr, policyType),
			),
		},
	}
}

// VaultCapabilities defines the capability hierarchy
func VaultCapabilities() []string {
	return []string{
		CapOwner.String(),
		CapOperator.String(),
		CapObserver.String(),
		CapAuthenticate.String(),
		CapAuthorize.String(),
		CapDelegate.String(),
		CapInvoke.String(),
		CapExecute.String(),
		CapRecover.String(),
	}
}
