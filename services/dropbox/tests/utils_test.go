package tests

import (
	"os"
	"testing"

	"github.com/google/uuid"

	dropbox "github.com/rgglez/go-storage/services/dropbox/v3"
	ps "github.com/rgglez/go-storage/v5/pairs"
	"github.com/rgglez/go-storage/v5/types"
)

func setupTest(t *testing.T) types.Storager {
	t.Log("Setup test for dropbox")

	store, err := dropbox.NewStorager(
		ps.WithCredential(os.Getenv("STORAGE_DROPBOX_CREDENTIAL")),
		ps.WithWorkDir("/"+uuid.New().String()),
	)
	if err != nil {
		t.Errorf("new storager: %v", err)
	}

	t.Cleanup(func() {
		err = store.Delete("")
		if err != nil {
			t.Errorf("cleanup: %v", err)
		}
	})
	return store
}
