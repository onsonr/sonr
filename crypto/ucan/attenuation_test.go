package ucan

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestAttenuationsContains(t *testing.T) {
	aContains := [][2]string{
		{
			`[
				{ "cap": "SUPER_USER", "dataset": "b5/world_bank_population"},
				{ "cap": "OVERWRITE", "api": "https://api.qri.cloud" }
			]`,
			`[
				{"cap": "SOFT_DELETE", "dataset": "b5/world_bank_population" }
			]`,
		},
		{
			`[
				{ "cap": "SUPER_USER", "dataset": "b5/world_bank_population"},
				{ "cap": "OVERWRITE", "api": "https://api.qri.cloud" }
			]`,
			`[
				{"cap": "SUPER_USER", "dataset": "b5/world_bank_population" }
			]`,
		},
	}

	for i, c := range aContains {
		t.Run(fmt.Sprintf("contains_%d", i), func(t *testing.T) {
			a := testAttenuations(c[0])
			b := testAttenuations(c[1])
			if !a.Contains(b) {
				t.Errorf("expected a attenuations to contain b attenuations")
			}
		})
	}

	aNotContains := [][2]string{
		{
			`[
				{ "cap": "SUPER_USER", "dataset": "b5/world_bank_population"},
				{ "cap": "OVERWRITE", "api": "https://api.qri.cloud" }
			]`,
			`[
				{ "cap": "CREATE", "dataset": "b5" }
			]`,
		},
	}

	for i, c := range aNotContains {
		t.Run(fmt.Sprintf("not_contains_%d", i), func(t *testing.T) {
			a := testAttenuations(c[0])
			b := testAttenuations(c[1])
			if a.Contains(b) {
				t.Errorf("expected a attenuations to NOT contain b attenuations")
			}
		})
	}
}

func mustJSON(data string, v interface{}) {
	if err := json.Unmarshal([]byte(data), v); err != nil {
		panic(err)
	}
}

func testAttenuations(data string) Attenuations {
	caps := NewNestedCapabilities("SUPER_USER", "OVERWRITE", "SOFT_DELETE", "REVISE", "CREATE")

	v := []map[string]string{}
	mustJSON(data, &v)

	var att Attenuations
	for _, x := range v {
		var cap Capability
		var rsc Resource
		for key, val := range x {
			switch key {
			case CapKey:
				cap = caps.Cap(val)
			default:
				rsc = NewStringLengthResource(key, val)
			}
		}
		att = append(att, Attenuation{cap, rsc})
	}

	return att
}

func TestNestedCapabilities(t *testing.T) {

}
