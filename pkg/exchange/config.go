package exchange

import (
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/internal/wallet"
	"github.com/sonr-io/core/pkg/common"
)

// DNSMode is the DNS mode for ExchangeProtocol
type DNSMode int

const (
	DNSMode_NONE DNSMode = iota
	DNSMode_TEMP
	DNSMode_PERM
)

// IsNone returns true if DNSMode is None
func (m DNSMode) IsNone() bool {
	return m == DNSMode_NONE
}

// IsTemp returns true if DNSMode is Temp
func (m DNSMode) IsTemp() bool {
	return m == DNSMode_TEMP
}

// IsPerm returns true if DNSMode is Perm
func (m DNSMode) IsPerm() bool {
	return m == DNSMode_PERM
}

// ShouldCreate returns true if ExchangeProtocol should create DNS records
func (m DNSMode) ShouldCreate() bool {
	return m.IsTemp() || m.IsPerm()
}

// Option is a function that can be applied to ExchangeProtocol config
type Option func(*options)

// TempName sets temporary sname on DNS record
func TempName(sname string) Option {
	return func(o *options) {
		o.Mode = DNSMode_TEMP
		o.sNameVal = sname
	}
}

// PermName sets permanent sname on DNS record
func PermName(sname string) Option {
	return func(o *options) {
		o.Mode = DNSMode_PERM
		o.sNameVal = sname
	}
}

// options for ExchangeProtocol config
type options struct {
	Mode     DNSMode
	sNameVal string
}

// defaultOptions for ExchangeProtocol config
func defaultOptions() *options {
	return &options{
		Mode:     DNSMode_NONE,
		sNameVal: "",
	}
}

// Apply applies options to ExchangeProtocol config
func (o *options) Apply(p *ExchangeProtocol) error {
	if o.Mode.ShouldCreate() {
		logger.Infof("Registering SName on DNS Table: %s (%v)", o.sNameVal, o.Mode)
		// Get Record Prefix
		prefix, err := common.NewRecordPrefix(o.sNameVal)
		if err != nil {
			return err
		}

		// Get Public Key
		pub, err := wallet.Primary.GetSnrPubKey(wallet.Account)
		if err != nil {
			return err
		}

		// Convert to Base58
		pubStr, err := pub.String()
		if err != nil {
			return err
		}

		// Create DNS records
		authRec := host.NewNBAuthRecord(prefix, o.sNameVal, "test, test, test")
		nameRec := host.NewNBNameRecord(pubStr, o.sNameVal)
		p.authRecord = authRec
		p.nameRecord = nameRec

		// Register temporary name records
		_, err = p.Register(o.sNameVal, authRec, nameRec)
		if err != nil {
			return err
		}
	}
	return nil
}
