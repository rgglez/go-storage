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

	"github.com/google/uuid"

	s3 "github.com/rgglez/go-storage/services/s3/v3"
	ps "github.com/rgglez/go-storage/v5/pairs"
	"github.com/rgglez/go-storage/v5/types"
)

func setupTest(t *testing.T) types.Storager {
	t.Log("Setup test for s3")

	store, err := s3.NewStorager(
		ps.WithCredential(os.Getenv("STORAGE_S3_CREDENTIAL")),
		ps.WithName(os.Getenv("STORAGE_S3_NAME")),
		ps.WithLocation(os.Getenv("STORAGE_S3_LOCATION")),
		ps.WithEndpoint(os.Getenv("STORAGE_S3_ENDPOINT")),
		ps.WithWorkDir("/"+uuid.New().String()+"/"),
		ps.WithEnableVirtualDir(),
		ps.WithEnableVirtualLink(),
		s3.WithForcePathStyle(),
	)
	if err != nil {
		t.Errorf("new storager: %v", err)
	}
	return store
}
