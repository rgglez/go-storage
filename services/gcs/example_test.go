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
// Example code for the gcs (Google Cloud Storage) service.
//
// These functions are compiled by "go test" but not executed (no // Output: comment).
// They serve as living, compile-checked documentation.
package gcs_test

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/rgglez/go-storage/v5/pairs"
	"github.com/rgglez/go-storage/v5/types"

	"github.com/rgglez/go-storage/services/gcs/v3"
)

// ExampleNewStorager demonstrates how to initialize a Google Cloud Storage backend.
//
// Create a service account with Storage permissions and download the JSON key.
//
//	export GCS_CREDENTIAL=file:/path/to/service-account-key.json
//	export GCS_BUCKET=my-bucket
//	export GCS_LOCATION=us-central1
func ExampleNewStorager() {
	ps := []types.Pair{
		pairs.WithCredential(os.Getenv("GCS_CREDENTIAL")),
		pairs.WithName(os.Getenv("GCS_BUCKET")),
		pairs.WithLocation(os.Getenv("GCS_LOCATION")),
	}
	if wd := os.Getenv("GCS_WORK_DIR"); wd != "" {
		ps = append(ps, pairs.WithWorkDir(wd))
	}

	_, store, err := gcs.New(ps...)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}
	fmt.Printf("storager: %s\n", store)
}

// ExampleStorage_Write demonstrates the Write → Stat → Read → List → Delete cycle.
func ExampleStorage_Write() {
	_, store, err := gcs.New(
		pairs.WithCredential("file:/path/to/service-account-key.json"),
		pairs.WithName("my-bucket"),
		pairs.WithLocation("us-central1"),
	)
	if err != nil {
		return
	}

	const path = "go-storage-example/hello.txt"
	const content = "Hello from go-storage gcs example!"

	n, err := store.Write(path, strings.NewReader(content), int64(len(content)))
	if err != nil {
		return
	}
	fmt.Printf("write  %q → %d bytes\n", path, n)

	o, err := store.Stat(path)
	if err != nil {
		return
	}
	size, _ := o.GetContentLength()
	fmt.Printf("stat   %q → size=%d\n", path, size)

	var buf bytes.Buffer
	n, err = store.Read(path, &buf)
	if err != nil {
		return
	}
	fmt.Printf("read   %q → %d bytes\n", path, n)

	it, err := store.List("go-storage-example/")
	if err != nil {
		return
	}
	for {
		obj, err := it.Next()
		if errors.Is(err, types.IterateDone) {
			break
		}
		if err != nil {
			return
		}
		fmt.Printf("  - %s\n", obj.Path)
	}

	if err := store.Delete(path); err != nil {
		return
	}
	fmt.Printf("delete %q → ok\n", path)
}

// ExampleServicer_Create demonstrates creating a new GCS bucket via the Servicer.
func ExampleServicer_Create() {
	srv, _, err := gcs.New(
		pairs.WithCredential("file:/path/to/service-account-key.json"),
		pairs.WithLocation("us-central1"),
	)
	if err != nil {
		return
	}

	store, err := srv.Create("new-bucket", pairs.WithLocation("us-central1"))
	if err != nil {
		return
	}
	fmt.Printf("created storager: %s\n", store)
}
