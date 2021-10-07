package internet_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/sonr-io/core/tools/internet"
)

func TestLookup(t *testing.T) {
	// Create resolver
	testVal := "_redirect.snr"
	hdnsResolver := internet.NewHDNSResolver()

	// Test with a valid domain
	rec, err := hdnsResolver.LookupTXT(context.Background(), testVal)
	if err != nil {
		t.Errorf("LookupTXT(%q) failed: %v", testVal, err)
	}

	// Verify result
	println("[SUCCESS]")
	println(fmt.Sprintf("\t host: %s", testVal))
	println(fmt.Sprintf("\t value: %s \n", rec.Record))
	t.Log(rec)
}
