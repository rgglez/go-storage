package tests

import (
	"os"
	"testing"

	"github.com/rgglez/go-storage/v5/tests"
)

func TestStorage(t *testing.T) {
	if os.Getenv("STORAGE_ONEDRIVE_INTEGRATION_TEST") != "on" {
		t.Skipf("STORAGE_ONEDRIVE_INTEGRATION_TEST is not 'on', skipped")
	}
	tests.TestStorager(t, setupTest(t))
}
