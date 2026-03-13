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
// multi-backend demonstrates the vendor-neutral nature of go-storage: the same
// application code works unchanged regardless of which backend is selected at
// runtime.
//
// Two backends are shown side-by-side:
//   - memory  – in-process, zero configuration, useful for tests and caches
//   - fs      – local filesystem, rooted at a temporary directory
//
// Run:
//
//	go run .
package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/rgglez/go-storage/v5/pairs"
	"github.com/rgglez/go-storage/v5/types"

	fs "github.com/rgglez/go-storage/services/fs/v4"
	"github.com/rgglez/go-storage/services/memory"
)

func main() {
	// ---- build two completely different backends ----------------------
	memStore, err := memory.NewStorager()
	if err != nil {
		fatal("memory.NewStorager", err)
	}

	tmpDir, err := os.MkdirTemp("", "go-storage-multi-*")
	if err != nil {
		fatal("MkdirTemp", err)
	}
	defer os.RemoveAll(tmpDir)

	fsStore, err := fs.NewStorager(pairs.WithWorkDir(tmpDir))
	if err != nil {
		fatal("fs.NewStorager", err)
	}

	// ---- run the same workload against each backend ------------------
	for _, backend := range []types.Storager{memStore, fsStore} {
		fmt.Printf("=== backend: %s ===\n", backend)
		if err := runWorkload(backend); err != nil {
			fmt.Fprintf(os.Stderr, "workload failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Println()
	}
}

// runWorkload executes a standard set of storage operations against any
// Storager implementation. This is the key pattern: the caller has no
// knowledge of which backend it is talking to.
func runWorkload(store types.Storager) error {
	const (
		path    = "greeting.txt"
		content = "Hello, go-storage!"
	)

	// write
	n, err := store.Write(path, strings.NewReader(content), int64(len(content)))
	if err != nil {
		return fmt.Errorf("write: %w", err)
	}
	fmt.Printf("  write  %q  → %d bytes\n", path, n)

	// stat
	o, err := store.Stat(path)
	if err != nil {
		return fmt.Errorf("stat: %w", err)
	}
	size, _ := o.GetContentLength()
	fmt.Printf("  stat   %q  → size=%d\n", path, size)

	// read back
	var buf bytes.Buffer
	n, err = store.Read(path, &buf)
	if err != nil {
		return fmt.Errorf("read: %w", err)
	}
	fmt.Printf("  read   %q  → %d bytes: %q\n", path, n, buf.String())

	// list
	it, err := store.List("")
	if err != nil {
		return fmt.Errorf("list: %w", err)
	}
	fmt.Print("  list:\n")
	for {
		obj, err := it.Next()
		if errors.Is(err, types.IterateDone) {
			break
		}
		if err != nil {
			return fmt.Errorf("list.Next: %w", err)
		}
		fmt.Printf("    - %s\n", obj.Path)
	}

	// copy to a second path using only the Storager interface
	if err := copyObject(store, path, "copy-of-"+path); err != nil {
		return fmt.Errorf("copy: %w", err)
	}
	fmt.Printf("  copy   %q  → %q\n", path, "copy-of-"+path)

	// delete both objects
	for _, p := range []string{path, "copy-of-" + path} {
		if err := store.Delete(p); err != nil {
			return fmt.Errorf("delete %q: %w", p, err)
		}
	}
	fmt.Printf("  delete ok\n")

	return nil
}

// copyObject reads src and writes it to dst using only the generic Storager
// interface — no backend-specific APIs needed.
func copyObject(store types.Storager, src, dst string) error {
	// stat to get the content length required by Write
	o, err := store.Stat(src)
	if err != nil {
		return fmt.Errorf("stat src: %w", err)
	}
	size, ok := o.GetContentLength()
	if !ok {
		size = -1 // unknown size; not all backends support -1
	}

	var buf bytes.Buffer
	if _, err := store.Read(src, &buf); err != nil {
		return fmt.Errorf("read src: %w", err)
	}
	if _, err := store.Write(dst, &buf, size); err != nil {
		return fmt.Errorf("write dst: %w", err)
	}
	return nil
}

func fatal(op string, err error) {
	fmt.Fprintf(os.Stderr, "ERROR [%s]: %v\n", op, err)
	os.Exit(1)
}
