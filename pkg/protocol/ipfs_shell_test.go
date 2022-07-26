package protocol_test

import (
	"reflect"
	"testing"

	"github.com/ipfs/go-cid"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/sonr-io/sonr/pkg/protocol"
)

func TestIPFSShell_GetData(t *testing.T) {
	type fields struct {
		Shell *shell.Shell
	}
	type args struct {
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
			I := protocol.IPFSShell{
				Shell: tt.fields.Shell,
			}
			got, err := I.GetData(tt.args.cid)
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
	}
	type args struct {
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
			I := protocol.IPFSShell{
				Shell: tt.fields.Shell,
			}
			if err := I.PinFile(tt.args.cidstr); (err != nil) != tt.wantErr {
				t.Errorf("PinFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIPFSShell_PutData(t *testing.T) {
	type fields struct {
		Shell *shell.Shell
	}
	type args struct {
		data []byte
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    *cid.Cid
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			I := protocol.IPFSShell{
				Shell: tt.fields.Shell,
			}
			got, err := I.PutData(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("PutData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PutData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIPFSShell_RemoveFile(t *testing.T) {
	type fields struct {
		Shell *shell.Shell
	}
	type args struct {
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
			I := protocol.IPFSShell{
				Shell: tt.fields.Shell,
			}
			if err := I.RemoveFile(tt.args.cidstr); (err != nil) != tt.wantErr {
				t.Errorf("RemoveFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewIPFSShell(t *testing.T) {
	type args struct {
		url string
	}
	var tests []struct {
		name string
		args args
		want protocol.IPFS
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := protocol.NewIPFSShell(tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIPFSShell() = %v, want %v", got, tt.want)
			}
		})
	}
}
