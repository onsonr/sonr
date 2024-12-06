package ucan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttenuationPresetConstructor(t *testing.T) {
	tests := []struct {
		name    string
		data    map[string]interface{}
		wantErr bool
	}{
		{
			name: "valid smart account attenuation",
			data: map[string]interface{}{
				"preset": "account",
				"cap":    string(CapOwner),
				"type":   string(ResAccount),
				"path":   "/accounts/123",
			},
			wantErr: false,
		},
		{
			name: "valid vault attenuation",
			data: map[string]interface{}{
				"preset": "vault",
				"cap":    string(CapOperator),
				"type":   string(ResVault),
				"path":   "/vaults/456",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			preset, data, err := ParseAttenuationData(tt.data)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)

			constructor, err := GetPresetConstructor(preset.String())
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)

			attenuation, err := constructor(data)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.NotNil(t, attenuation)
		})
	}
}
