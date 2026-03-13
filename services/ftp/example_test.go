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
// Example code for the ftp (FTP/FTPS) service.
//
// These functions are compiled by "go test" but not executed (no // Output: comment).
// They serve as living, compile-checked documentation.
package ftp_test

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/rgglez/go-storage/v5/pairs"
	"github.com/rgglez/go-storage/v5/types"

	"github.com/rgglez/go-storage/services/ftp"
)

// ExampleNewStorager demonstrates how to initialize an FTP backend.
//
//	export FTP_ENDPOINT=ftp+addr://ftp.example.com:21
//	export FTP_CREDENTIAL=basic:USER:PASSWORD
func ExampleNewStorager() {
	ps := []types.Pair{
		pairs.WithEndpoint(os.Getenv("FTP_ENDPOINT")),
		pairs.WithCredential(os.Getenv("FTP_CREDENTIAL")),
	}
	if wd := os.Getenv("FTP_WORK_DIR"); wd != "" {
		ps = append(ps, pairs.WithWorkDir(wd))
	}

	store, err := ftp.NewStorager(ps...)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}
	fmt.Printf("storager: %s\n", store)
}

// ExampleStorage_Write demonstrates the Write → Stat → Read → List → Delete cycle.
// FTP credential format is "basic:USER:PASSWORD".
func ExampleStorage_Write() {
	store, err := ftp.NewStorager(
		pairs.WithEndpoint("ftp+addr://ftp.example.com:21"),
		pairs.WithCredential("basic:USER:PASSWORD"),
	)
	if err != nil {
		return
	}

	const path = "go-storage-example/hello.txt"
	const content = "Hello from go-storage ftp example!"

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
