package tests

import (
	"os"
	"testing"

	"github.com/google/uuid"

	minio "github.com/rgglez/go-storage/services/minio"
	ps "github.com/rgglez/go-storage/v5/pairs"
	"github.com/rgglez/go-storage/v5/types"
)

func setupTest(t *testing.T) types.Storager {
	t.Log("Setup test for minio")

	store, err := minio.NewStorager(
		ps.WithCredential(os.Getenv("STORAGE_MINIO_CREDENTIAL")),
		ps.WithEndpoint(os.Getenv("STORAGE_MINIO_ENDPOINT")),
		ps.WithName(os.Getenv("STORAGE_MINIO_NAME")),
		ps.WithWorkDir("/"+uuid.New().String()),
	)
	if err != nil {
		t.Errorf("new storager: %v", err)
	}
	return store
}
