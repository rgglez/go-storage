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
package fs_test

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/rgglez/go-storage/v5/pairs"
	"github.com/rgglez/go-storage/v5/types"

	fs "github.com/rgglez/go-storage/services/fs/v4"
)

// ExampleNewStorager shows how to create a local-filesystem storager rooted at
// a given work directory.
func ExampleNewStorager() {
	store, err := fs.NewStorager(pairs.WithWorkDir("/tmp"))
	if err != nil {
		panic(err)
	}
	fmt.Println(store != nil)
	// Output:
	// true
}

// exampleStore is a package-level helper for the examples below. It creates a
// throwaway temp-dir so that examples are hermetic and leave no residue.
func exampleStore(t *testing.T) (types.Storager, func()) {
	t.Helper()
	dir := t.TempDir()
	store, err := fs.NewStorager(pairs.WithWorkDir(dir))
	if err != nil {
		t.Fatal(err)
	}
	return store, func() { os.RemoveAll(dir) }
}

// TestExample_Write_and_Read demonstrates write followed by read using the
// generic Storager interface. Running it as a test keeps the example
// compile-checked and verifiable without requiring a live filesystem path.
func TestExample_Write_and_Read(t *testing.T) {
	store, cleanup := exampleStore(t)
	defer cleanup()

	content := "Hello, go-storage!"
	n, err := store.Write("hello.txt", strings.NewReader(content), int64(len(content)))
	if err != nil {
		t.Fatal(err)
	}
	if n != int64(len(content)) {
		t.Fatalf("expected %d bytes written, got %d", len(content), n)
	}

	var buf bytes.Buffer
	n, err = store.Read("hello.txt", &buf)
	if err != nil {
		t.Fatal(err)
	}
	if buf.String() != content {
		t.Fatalf("expected %q, got %q", content, buf.String())
	}
	fmt.Printf("round-trip OK: %d bytes\n", n)
}

// TestExample_Read_withRange demonstrates reading a byte sub-range of an object.
func TestExample_Read_withRange(t *testing.T) {
	store, cleanup := exampleStore(t)
	defer cleanup()

	store.Write("data.txt", strings.NewReader("Hello, go-storage!"), 18) //nolint

	var buf bytes.Buffer
	_, err := store.Read("data.txt", &buf,
		pairs.WithOffset(7),
		pairs.WithSize(10),
	)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(buf.String())
}

// TestExample_Stat demonstrates retrieving object metadata.
func TestExample_Stat(t *testing.T) {
	store, cleanup := exampleStore(t)
	defer cleanup()

	store.Write("sample.txt", strings.NewReader("content"), 7) //nolint

	o, err := store.Stat("sample.txt")
	if err != nil {
		t.Fatal(err)
	}
	size, ok := o.GetContentLength()
	fmt.Printf("name=%s hasSize=%v size=%d\n", o.Path, ok, size)
}

// TestExample_Delete demonstrates deleting an object.
func TestExample_Delete(t *testing.T) {
	store, cleanup := exampleStore(t)
	defer cleanup()

	store.Write("to-delete.txt", strings.NewReader("bye"), 3) //nolint

	err := store.Delete("to-delete.txt")
	if err != nil {
		t.Fatal(err)
	}

	_, err = store.Stat("to-delete.txt")
	fmt.Println(err != nil)
}

// TestExample_List demonstrates listing objects under a path prefix.
func TestExample_List(t *testing.T) {
	store, cleanup := exampleStore(t)
	defer cleanup()

	for _, name := range []string{"a.txt", "b.txt", "c.txt"} {
		store.Write(name, strings.NewReader("x"), 1) //nolint
	}

	it, err := store.List("")
	if err != nil {
		t.Fatal(err)
	}

	count := 0
	for {
		_, err := it.Next()
		if errors.Is(err, types.IterateDone) {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		count++
	}
	fmt.Printf("listed %d objects\n", count)
}

// TestExample_CreateDir demonstrates creating a directory object.
func TestExample_CreateDir(t *testing.T) {
	store, cleanup := exampleStore(t)
	defer cleanup()

	direr, ok := store.(types.Direr)
	if !ok {
		t.Skip("storager does not implement Direr")
	}

	_, err := direr.CreateDir("subdir")
	if err != nil {
		t.Fatal(err)
	}

	o, err := store.Stat("subdir")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(o.Mode.IsDir())
}
