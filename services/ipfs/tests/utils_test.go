package tests

import (
	"os"
	"testing"

	"github.com/google/uuid"
	ipfs "github.com/rgglez/go-storage/services/ipfs"

	"github.com/rgglez/go-storage/v5/pairs"
	"github.com/rgglez/go-storage/v5/types"
)

func setupTest(t *testing.T) types.Storager {
	t.Log("Setup test for IPFS")

	store, err := ipfs.NewStorager(
		pairs.WithEndpoint(os.Getenv("STORAGE_IPFS_ENDPOINT")),
		ipfs.WithGateway(os.Getenv("STORAGE_IPFS_GATEWAY")),
		pairs.WithWorkDir("/"+uuid.New().String()+"/"),
	)
	if err != nil {
		t.Errorf("new storager: %v", err)
	}
	return store
}
