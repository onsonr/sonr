package protocol_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/ipfs/go-datastore"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/sonr-io/sonr/pkg/protocol"
)

const IPFSShellUrl = ""

var cacheStore datastore.Datastore

func TestIPFSShell_GetData(t *testing.T) {
	type fields struct {
		Shell *shell.Shell
		cache datastore.Datastore
	}
	type args struct {
		ctx context.Context
		cid string
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := protocol.NewIPFSShell(IPFSShellUrl, cacheStore)
			got, err := i.GetData(tt.args.ctx, tt.args.cid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIPFSShell_PinFile(t *testing.T) {
	type fields struct {
		Shell *shell.Shell
		cache datastore.Datastore
	}
	type args struct {
		ctx    context.Context
		cidstr string
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := protocol.NewIPFSShell(IPFSShellUrl, cacheStore)
			if err := i.PinFile(tt.args.ctx, tt.args.cidstr); (err != nil) != tt.wantErr {
				t.Errorf("PinFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIPFSShell_PutData(t *testing.T) {
	type fields struct {
		Shell *shell.Shell
		cache datastore.Datastore
	}
	type args struct {
		ctx  context.Context
		data []byte
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := protocol.NewIPFSShell(IPFSShellUrl, cacheStore)

			got, err := i.PutData(tt.args.ctx, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("PutData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PutData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIPFSShell_RemoveFile(t *testing.T) {
	type fields struct {
		Shell *shell.Shell
		cache datastore.Datastore
	}
	type args struct {
		ctx    context.Context
		cidstr string
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := protocol.NewIPFSShell(IPFSShellUrl, cacheStore)
			if err := i.RemoveFile(tt.args.ctx, tt.args.cidstr); (err != nil) != tt.wantErr {
				t.Errorf("RemoveFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewIPFSShell(t *testing.T) {
	type args struct {
		url        string
		cacheStore datastore.Datastore
	}
	var tests []struct {
		name string
		args args
		want protocol.IPFS
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := protocol.NewIPFSShell(tt.args.url, tt.args.cacheStore); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIPFSShell() = %v, want %v", got, tt.want)
			}
		})
	}
}
