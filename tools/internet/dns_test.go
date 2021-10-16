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
	recs, err := hdnsResolver.LookupTXT(context.Background(), testVal)
	if err != nil {
		t.Errorf("LookupTXT(%q) failed: %v", testVal, err)
	}

	// Verify result
	println("[SUCCESS] - Test DNS Lookup")
	println(fmt.Sprintf("\t host: %s", testVal))
	println("\t value: %s \n")
	for _, v := range recs {
		v.Print()
	}
	t.Log(recs)
}
