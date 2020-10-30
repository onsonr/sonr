package core

import (
	"io/ioutil"
	"os"
	"testing"
)

type cleanFunc func()

func testingConfig(t *testing.T) *Config {
	t.Helper()

	cfg, err := NewDefaultConfig()
	if err != nil {
		t.Fatal(err)
	}

	// do not bootstrap
	err = cfg.SetKey("Bootstrap", []byte("[]"))
	if err != nil {
		t.Fatal(err)
	}

	return cfg
}

func testingRepo(t *testing.T, path string) (*Repo, cleanFunc) {
	t.Helper()

	cfg := testingConfig(t)

	// init repo
	if err := InitRepo(path, cfg); err != nil {
		t.Fatal(err)
	}

	// repo should be initialized
	if !RepoIsInitialized(path) {
		t.Fatal("repo hasn't been init successfully")
	}

	// open repo
	repo, err := OpenRepo(path)
	if err != nil {
		t.Fatal(err)
	}

	return repo, func() {
		repo.Close()
	}
}

func testingNode(t *testing.T, path string) (*Node, cleanFunc) {
	repo, cleanRepo := testingRepo(t, path)

	// create new node
	node, err := NewNode(repo)
	if err != nil {
		t.Fatal(err)
	}

	return node, func() {
		node.Close()
		cleanRepo()
	}

}

func testingSockmanager(t *testing.T) (*SockManager, cleanFunc) {
	t.Helper()

	path, clean := testingTempDir(t, "sm")
	sm, err := NewSockManager(path)
	if err != nil {
		clean()
		t.Fatal(err)
	}

	return sm, clean
}

func testingTempDir(t *testing.T, name string) (string, cleanFunc) {
	t.Helper()

	path, err := ioutil.TempDir("", name)
	if err != nil {
		t.Fatal(err)
	}

	return path, func() {
		os.RemoveAll(path)
	}
}
