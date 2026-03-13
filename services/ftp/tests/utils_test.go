package tests

import (
	"os"
	"testing"

	_ "github.com/rgglez/go-storage/services/ftp"
	ps "github.com/rgglez/go-storage/v5/pairs"
	"github.com/rgglez/go-storage/v5/services"
	"github.com/rgglez/go-storage/v5/types"
)

func initTest(t *testing.T) (store types.Storager) {
	t.Log("Setup test for ftp")

	store, err := services.NewStorager("ftp",
		ps.WithCredential(os.Getenv("STORAGE_FTP_CREDENTIAL")),
		ps.WithEndpoint(os.Getenv("STORAGE_FTP_ENDPOINT")),
	)
	if err != nil {
		t.Errorf("create storager: %v", err)
	}

	return
}
