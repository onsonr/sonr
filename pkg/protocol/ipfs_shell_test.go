package protocol_test

import (
	"context"
	"encoding/json"
	"github.com/sonr-io/sonr/pkg/protocol"
	"github.com/sonr-io/sonr/x/schema/types"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/ipfs/go-datastore"
	shell "github.com/ipfs/go-ipfs-api"
)

const IPFSShellUrl = "localhost:5001"

var cacheStore datastore.Datastore

func TestIPFSShell_PutData(t *testing.T) {
	type fields struct {
		Shell *shell.Shell
		cache datastore.Datastore
	}
	type args struct {
		ctx    context.Context
		schema *types.SchemaDefinition
	}
	var tests = []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr error
	}{
		{name: "test#1", fields: fields{
			Shell: nil,
			cache: nil,
		}, args: args{
			ctx: nil,
			schema: &types.SchemaDefinition{
				Creator: "did:snr:7Prd74ry1Uct87nZqL3ny7aR7Cg46JamVbJgk8azVgUm",
				Label:   "test-label",
			},
		}, want: "QmW4Ghk82fyq4LsoBKwH5o66Zb1sEpZ735Tmn1yA7o1uGu", wantErr: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := protocol.NewIPFSShell(IPFSShellUrl, cacheStore)

			data, err := json.Marshal(tt.args.schema)
			assert.NoError(t, err)

			got, err := i.PutData(tt.args.ctx, data)

			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr, err)
			} else {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
