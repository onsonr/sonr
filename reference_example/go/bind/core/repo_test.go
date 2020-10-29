package core

import (
	"reflect"
	"testing"
)

func TestRepo(t *testing.T) {
	path, clean := testingTempDir(t, "repo")
	defer clean()

	// check if repo is initialized
	if RepoIsInitialized(path) {
		t.Fatal("repo shouldn't be initialized")
	}

	cfg := testingConfig(t)

	// init repo
	err := InitRepo(path, cfg)
	if err != nil {
		t.Fatal(err)
	}

	if !RepoIsInitialized(path) {
		t.Fatal(err)
	}

	// open repo
	repo, err := OpenRepo(path)
	if err != nil {
		t.Fatal(err)
	}

	defer repo.Close()

	rootpath := repo.GetRootPath()
	if rootpath != path {
		t.Errorf("expected `%s` but got `%s`", path, rootpath)
	}

	repocfg, err := repo.GetConfig()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(repocfg, cfg) {
		t.Error("GetConfig value and original config should be equal")
	}
}
