package types

import (
	"time"

	dns "github.com/Focinfi/go-dns-resolver"
)

type DNSOption func(*DNSOptions)

func WithRecordTypes(types ...dns.QueryType) DNSOption {
	return func(opts *DNSOptions) {
		opts.RecordTypes = types
	}
}

func WithDomains(domains ...string) DNSOption {
	return func(opts *DNSOptions) {
		opts.Domains = domains
	}
}

type DNSOptions struct {
	RecordTypes []dns.QueryType
	Domains     []string
}

func DefaultDNSOptions() *DNSOptions {
	return &DNSOptions{
		RecordTypes: []dns.QueryType{dns.TypeTXT, dns.TypeA},
		Domains:     []string{"sonr", "welcome.nb"},
	}
}

func (o *DNSOptions) Apply(opts ...DNSOption) *dns.Result {
	for _, opt := range opts {
		opt(o)
	}
	params := DefaultParams()
	resolver := dns.NewResolver(params.DnsResolverIp)
	resolver.Targets(o.Domains...).Types(o.RecordTypes...)
	return resolver.Lookup()
}

type DNSRecord struct {
	Name     string
	Type     string
	Content  string
	Ttl      time.Duration
	Priority uint16
}

func NewDNSRecordFromResultItem(target string, item *dns.ResultItem) *DNSRecord {
	return &DNSRecord{
		Name:     target,
		Type:     item.Type,
		Content:  item.Content,
		Ttl:      item.Ttl,
		Priority: item.Priority,
	}
}
