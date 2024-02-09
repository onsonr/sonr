package config_test

import (
	"testing"

	"github.com/sonrhq/sonr/config"
)

func TestOpen(t *testing.T) {
	c, err := config.LoadConfig()
	if err != nil {
		t.Errorf("Error loading config: %s", err)
	}
	if c.Name != "sonr" {
		t.Errorf("Expected name to be sonr, got %s", c.Name)
	}
}
