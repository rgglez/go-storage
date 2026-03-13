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
