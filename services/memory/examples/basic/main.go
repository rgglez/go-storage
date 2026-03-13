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
// Basic example for the memory (in-process) service.
//
// Run from this directory:
//
//	go run main.go
//
// No credentials or network access required. The store is created fresh on
// every run and discarded when the process exits.
package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/rgglez/go-storage/v5/types"

	"github.com/rgglez/go-storage/services/memory"
)

func main() {
	// ---- create storager ----------------------------------------------
	store, err := memory.NewStorager()
	if err != nil {
		fatal("new storager", err)
	}
	fmt.Printf("storager: %s\n\n", store)

	// ---- write --------------------------------------------------------
	const path = "hello.txt"
	const content = "Hello, go-storage!"

	n, err := store.Write(path, strings.NewReader(content), int64(len(content)))
	if err != nil {
		fatal("write", err)
	}
	fmt.Printf("write  %q  → %d bytes\n", path, n)

	// ---- stat ---------------------------------------------------------
	o, err := store.Stat(path)
	if err != nil {
		fatal("stat", err)
	}
	size, _ := o.GetContentLength()
	fmt.Printf("stat   %q  → size=%d\n", path, size)

	// ---- read ---------------------------------------------------------
	var buf bytes.Buffer
	n, err = store.Read(path, &buf)
	if err != nil {
		fatal("read", err)
	}
	fmt.Printf("read   %q  → %d bytes: %q\n", path, n, buf.String())

	// ---- list ---------------------------------------------------------
	it, err := store.List("")
	if err != nil {
		fatal("list", err)
	}
	fmt.Print("list:\n")
	for {
		obj, err := it.Next()
		if errors.Is(err, types.IterateDone) {
			break
		}
		if err != nil {
			fatal("list.Next", err)
		}
		fmt.Printf("  - %s\n", obj.Path)
	}

	// ---- delete -------------------------------------------------------
	if err := store.Delete(path); err != nil {
		fatal("delete", err)
	}
	fmt.Printf("delete %q  → ok\n", path)
}

func fatal(op string, err error) {
	fmt.Fprintf(os.Stderr, "ERROR [%s]: %v\n", op, err)
	os.Exit(1)
}
