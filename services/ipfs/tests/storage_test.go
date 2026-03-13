// Copyright The go-storage Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package tests

import (
	"os"
	"testing"

	"github.com/rgglez/go-storage/v5/tests"
)

func TestStorager(t *testing.T) {
	if os.Getenv("STORAGE_IPFS_INTEGRATION_TEST") != "on" {
		t.Skipf("STORAGE_IPFS_INTEGRATION_TEST is not 'on', skipped")
	}
	tests.TestStorager(t, setupTest(t))
}

func TestCopier(t *testing.T) {
	if os.Getenv("STORAGE_IPFS_INTEGRATION_TEST") != "on" {
		t.Skipf("STORAGE_IPFS_INTEGRATION_TEST is not 'on', skipped")
	}
	tests.TestCopier(t, setupTest(t))
}

func TestMover(t *testing.T) {
	if os.Getenv("STORAGE_IPFS_INTEGRATION_TEST") != "on" {
		t.Skipf("STORAGE_IPFS_INTEGRATION_TEST is not 'on', skipped")
	}
	tests.TestMover(t, setupTest(t))
}

func TestStorageHttpSignerRead(t *testing.T) {
	if os.Getenv("STORAGE_IPFS_INTEGRATION_TEST") != "on" {
		t.Skipf("STORAGE_IPFS_INTEGRATION_TEST is not 'on', skipped")
	}
	tests.TestStorageHTTPSignerRead(t, setupTest(t))
}
