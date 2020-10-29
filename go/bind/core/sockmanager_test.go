package core

import (
	"io/ioutil"
	"net"
	"path/filepath"
	"strings"
	"testing"
)

func TestSockmanager(t *testing.T) {
	tmpdir, err := ioutil.TempDir("", "sock")
	if err != nil {
		t.Fatal(err)
	}

	cases := []struct {
		Path        string
		ErrExpected bool
	}{
		{tmpdir, false}, // ok
		{filepath.Join(tmpdir, strings.Repeat("a", 110)), true},
		{"/" + strings.Repeat("a", 110), true},
		{"a", true},
		{"/foo/bar/buzz", true},
	}

	for _, tc := range cases {
		t.Run("new", func(t *testing.T) {
			_, err := NewSockManager(tc.Path)
			if !tc.ErrExpected && err != nil {
				t.Fatalf("expected err to be nil but got: %s", err)
			} else if tc.ErrExpected && err == nil {
				t.Fatalf("expected error, but it was nil")
			}
		})
	}
}

func TestSockmanagerNewPath(t *testing.T) {
	tmpdir, err := ioutil.TempDir("", "sock")
	if err != nil {
		t.Fatal(err)
	}

	sm, err := NewSockManager(tmpdir)
	if err != nil {
		t.Fatal(err)
	}

	// try to generate some path
	for i := 0; i < 100; i++ {
		_, err = sm.NewSockPath()
		if err != nil {
			t.Fatal(err)
		}
	}

	// reset
	sm.counter = 0

	// get new sock path
	sock, err := sm.NewSockPath()
	if err != nil {
		t.Fatal(err)
	}

	// try to listen on it
	l, err := net.Listen("unix", sock)
	if err != nil {
		t.Fatal(err)
	}

	defer l.Close()

	// try to create a sock on an already created one
	sm.counter = 0 // reset
	_, err = sm.NewSockPath()
	if err == nil {
		t.Fatal("new sock path should fail on an already created file")
	}
}
