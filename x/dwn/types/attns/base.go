package attns

import (
	"github.com/onsonr/sonr/x/dwn/types/attns/capability"
	"github.com/onsonr/sonr/x/dwn/types/attns/resourcetype"
	"github.com/ucan-wg/go-ucan"
)

const (
	CapOwner        = capability.CAPOWNER
	CapOperator     = capability.CAPOPERATOR
	CapServer       = capability.CAPOBSERVER
	CapExecute      = capability.CAPEXECUTE
	CapPropose      = capability.CAPPROPOSE
	CapSign         = capability.CAPSIGN
	CapSetPolicy    = capability.CAPSETPOLICY
	CapSetThreshold = capability.CAPSETTHRESHOLD
	CapRecover      = capability.CAPRECOVER
	CapSocial       = capability.CAPSOCIAL

	ResAccount     = resourcetype.RESACCOUNT
	ResTransaction = resourcetype.RESTRANSACTION
	ResPolicy      = resourcetype.RESPOLICY
	ResRecovery    = resourcetype.RESRECOVERY
	ResVault       = resourcetype.RESVAULT

	PolicyThreshold = resourcetype.RESPOLICY
	PolicySet       = resourcetype.RESPOLICY
	PolicyGet       = resourcetype.RESPOLICY
	PolicyDelete    = resourcetype.RESPOLICY
)

// NewVaultResource creates a new resource identifier
func NewResource(resType resourcetype.ResourceType, path string) ucan.Resource {
	return ucan.NewStringLengthResource(string(resType), path)
}
