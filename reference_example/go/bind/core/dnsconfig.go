// from: https://gist.github.com/cs8425/107e01a0652f1f1f6e033b5b68364b5e

//nolint
package core

import (
	_ "net"
	"sync"
	"time"
	_ "unsafe"
)

// DNS resolve workaround for Android & iOS in pure go

//go:linkname defaultNS net.defaultNS
var defaultNS []string

// need to keep sync with go version
//go:linkname resolvConf net.resolvConf
var resolvConf resolverConfig

// copy from /src/net/dnsconfig_unix.go
type dnsConfig struct {
	servers    []string      // server addresses (in host:port form) to use
	search     []string      // rooted suffixes to append to local name
	ndots      int           // number of dots in name to trigger absolute lookup
	timeout    time.Duration // wait before giving up on a query, including retries
	attempts   int           // lost packets before giving up on server
	rotate     bool          // round robin among servers
	unknownOpt bool          // anything unknown was encountered
	lookup     []string      // OpenBSD top-level database "lookup" order
	err        error         // any error that occurs during open of resolv.conf
	mtime      time.Time     // time of resolv.conf modification
	soffset    uint32        // used by serverOffset
}

// copy from /src/net/dnsclient_unix.go
type resolverConfig struct {
	initOnce sync.Once // guards init of resolverConfig

	// ch is used as a semaphore that only allows one lookup at a
	// time to recheck resolv.conf.
	ch          chan struct{} // guards lastChecked and modTime
	lastChecked time.Time     // last time resolv.conf was checked

	mu        sync.RWMutex // protects dnsConfig
	dnsConfig *dnsConfig   // parsed resolv.conf structure used in lookups
}

//go:linkname (*resolverConfig).tryUpdate net.(*resolverConfig).tryUpdate
func (conf *resolverConfig) tryUpdate(name string)

// Need an empty .s file (dnsconfig_empty.s)

func setDefaultNS(addrs []string, loadFromSystem bool) {
	if resolvConf.dnsConfig == nil {
		resolvConf.tryUpdate("")
	}

	if loadFromSystem {
		now := time.Now()
		resolvConf.lastChecked = now.Add(-6 * time.Second)
		resolvConf.dnsConfig.mtime = now
	}

	resolvConf.dnsConfig.servers = addrs
}

func SetDNSPair(primary, secondary string, loadFromSystem bool) {
	setDefaultNS([]string{primary, secondary}, loadFromSystem)
}

func init() {
	// Set DNSWatch as DNS provider by default (privacy-focused provider)
	SetDNSPair("84.200.69.80:53", "84.200.70.40:53", false)
}
