package ucan

import (
	"context"
	"fmt"

	"github.com/onsonr/sonr/crypto/mpc"
	"github.com/onsonr/sonr/crypto/ucan/attns/capability"
	"github.com/onsonr/sonr/crypto/ucan/attns/policytype"
	"github.com/onsonr/sonr/crypto/ucan/attns/resourcetype"
)

// CtxKey defines a distinct type for context keys used by the access
// package
type CtxKey string

// TokenCtxKey is the key for adding an access UCAN to a context.Context
const TokenCtxKey CtxKey = "UCAN"

// CtxWithToken adds a UCAN value to a context
func CtxWithToken(ctx context.Context, t Token) context.Context {
	return context.WithValue(ctx, TokenCtxKey, t)
}

// FromCtx extracts a token from a given context if one is set, returning nil
// otherwise
func FromCtx(ctx context.Context) *Token {
	iface := ctx.Value(TokenCtxKey)
	if ref, ok := iface.(*Token); ok {
		return ref
	}
	return nil
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
	kss mpc.KeyEnclave,
) Attenuations {
	accountAddr := kss.Address()
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
