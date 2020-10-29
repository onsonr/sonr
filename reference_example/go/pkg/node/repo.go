package node

import (
	ipfs_repo "github.com/ipfs/go-ipfs/repo"
)

var _ ipfs_repo.Repo = (*MobileRepo)(nil)

type MobileRepo struct {
	ipfs_repo.Repo
	Path string

	// extra config
	EnablePubsubExperiment bool
	EnableNamesysPubsub    bool
}
