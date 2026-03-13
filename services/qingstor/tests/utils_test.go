package tests

import (
	"os"
	"testing"

	"github.com/google/uuid"

	"github.com/rgglez/go-storage/services/qingstor/v4"
	ps "github.com/rgglez/go-storage/v5/pairs"
	"github.com/rgglez/go-storage/v5/types"
)

func setupTest(t *testing.T) types.Storager {
	t.Log("Setup test for qingstor")

	store, err := qingstor.NewStorager(
		ps.WithCredential(os.Getenv("STORAGE_QINGSTOR_CREDENTIAL")),
		ps.WithEndpoint(os.Getenv("STORAGE_QINGSTOR_ENDPOINT")),
		ps.WithName(os.Getenv("STORAGE_QINGSTOR_NAME")),
		ps.WithWorkDir("/"+uuid.New().String()+"/"),
		ps.WithEnableVirtualDir(),
		ps.WithEnableVirtualLink(),
	)
	if err != nil {
		t.Errorf("new storager: %v", err)
	}
	return store
}
