package tests

import (
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/rgglez/go-storage/v5/pairs"
	"github.com/rgglez/go-storage/v5/types"

	hdfs "github.com/rgglez/go-storage/services/hdfs"
)

func setupTest(t *testing.T) types.Storager {
	t.Log("Setup test for HDFS")

	store, err := hdfs.NewStorager(
		pairs.WithEndpoint(os.Getenv("STORAGE_HDFS_ENDPOINT")),
		pairs.WithWorkDir("/"+uuid.New().String()+"/"),
	)
	if err != nil {
		t.Errorf("new storager: %v", err)
	}

	return store
}
