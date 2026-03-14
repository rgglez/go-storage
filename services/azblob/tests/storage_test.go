// Copyright The go-storage Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package tests

import (
	"os"
	"testing"

	"github.com/rgglez/go-storage/v5/tests"
)

func TestStorage(t *testing.T) {
	if os.Getenv("STORAGE_AZBLOB_INTEGRATION_TEST") != "on" {
		t.Skipf("STORAGE_AZBLOB_INTEGRATION_TEST is not 'on', skipped")
	}
	tests.TestStorager(t, setupTest(t))
}

func TestAppend(t *testing.T) {
	if os.Getenv("STORAGE_AZBLOB_INTEGRATION_TEST") != "on" {
		t.Skipf("STORAGE_AZBLOB_INTEGRATION_TEST is not 'on', skipped")
	}
	tests.TestAppender(t, setupTest(t))
}

func TestDir(t *testing.T) {
	if os.Getenv("STORAGE_AZBLOB_INTEGRATION_TEST") != "on" {
		t.Skipf("STORAGE_AZBLOB_INTEGRATION_TEST is not 'on', skipped")
	}
	tests.TestDirer(t, setupTest(t))
}
