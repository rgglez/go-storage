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
// Example code for the azblob (Azure Blob Storage) service.
//
// These functions are compiled by "go test" but not executed (no // Output: comment).
// They serve as living, compile-checked documentation.
package azblob_test

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/rgglez/go-storage/v5/pairs"
	"github.com/rgglez/go-storage/v5/types"

	azblob "github.com/rgglez/go-storage/services/azblob/v3"
)

// ExampleNewStorager demonstrates how to initialize an Azure Blob Storage backend.
//
//	export AZBLOB_CREDENTIAL=hmac:ACCOUNT_NAME:ACCOUNT_KEY
//	export AZBLOB_ENDPOINT=https://ACCOUNT_NAME.blob.core.windows.net
//	export AZBLOB_CONTAINER=my-container
func ExampleNewStorager() {
	ps := []types.Pair{
		pairs.WithCredential(os.Getenv("AZBLOB_CREDENTIAL")),
		pairs.WithEndpoint(os.Getenv("AZBLOB_ENDPOINT")),
		pairs.WithName(os.Getenv("AZBLOB_CONTAINER")),
	}
	if wd := os.Getenv("AZBLOB_WORK_DIR"); wd != "" {
		ps = append(ps, pairs.WithWorkDir(wd))
	}

	_, store, err := azblob.New(ps...)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}
	fmt.Printf("storager: %s\n", store)
}

// ExampleStorage_Write demonstrates the Write → Stat → Read → List → Delete cycle.
func ExampleStorage_Write() {
	_, store, err := azblob.New(
		pairs.WithCredential("hmac:ACCOUNT_NAME:ACCOUNT_KEY"),
		pairs.WithEndpoint("https://ACCOUNT_NAME.blob.core.windows.net"),
		pairs.WithName("my-container"),
	)
	if err != nil {
		return
	}

	const path = "go-storage-example/hello.txt"
	const content = "Hello from go-storage azblob example!"

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

// ExampleServicer_Create demonstrates creating a new Azure Blob container
// using the Servicer returned by azblob.New.
func ExampleServicer_Create() {
	srv, _, err := azblob.New(
		pairs.WithCredential("hmac:ACCOUNT_NAME:ACCOUNT_KEY"),
		pairs.WithEndpoint("https://ACCOUNT_NAME.blob.core.windows.net"),
	)
	if err != nil {
		return
	}

	store, err := srv.Create("new-container")
	if err != nil {
		return
	}
	fmt.Printf("created storager: %s\n", store)
}
