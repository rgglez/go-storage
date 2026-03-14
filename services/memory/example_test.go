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
package memory_test

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/rgglez/go-storage/v5/pairs"
	"github.com/rgglez/go-storage/v5/types"

	"github.com/rgglez/go-storage/services/memory"
)

// ExampleNewStorager shows how to create an in-memory storager. No credentials
// or network access are required; the store lives entirely in process memory.
func ExampleNewStorager() {
	store, err := memory.NewStorager()
	if err != nil {
		panic(err)
	}
	fmt.Println(store)
	// Output:
	// memory
}

// ExampleStorage_Write shows how to write data into the store.
func ExampleStorage_Write() {
	store, _ := memory.NewStorager()

	content := "Hello, go-storage!"
	n, err := store.Write("greeting.txt", strings.NewReader(content), int64(len(content)))
	if err != nil {
		panic(err)
	}
	fmt.Printf("written %d bytes\n", n)
	// Output:
	// written 18 bytes
}

// ExampleStorage_Read shows how to read data back from the store.
func ExampleStorage_Read() {
	store, _ := memory.NewStorager()

	content := "Hello, go-storage!"
	store.Write("greeting.txt", strings.NewReader(content), int64(len(content))) //nolint

	var buf bytes.Buffer
	n, err := store.Read("greeting.txt", &buf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("read %d bytes: %s\n", n, buf.String())
	// Output:
	// read 18 bytes: Hello, go-storage!
}

// ExampleStorage_Read_withRange shows how to read a byte range using
// pairs.WithOffset and pairs.WithSize.
func ExampleStorage_Read_withRange() {
	store, _ := memory.NewStorager()

	content := "Hello, go-storage!"
	store.Write("greeting.txt", strings.NewReader(content), int64(len(content))) //nolint

	var buf bytes.Buffer
	_, err := store.Read("greeting.txt", &buf,
		pairs.WithOffset(7),
		pairs.WithSize(10),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(buf.String())
	// Output:
	// go-storage
}

// ExampleStorage_Stat shows how to retrieve object metadata.
func ExampleStorage_Stat() {
	store, _ := memory.NewStorager()

	content := "Hello, go-storage!"
	store.Write("greeting.txt", strings.NewReader(content), int64(len(content))) //nolint

	o, err := store.Stat("greeting.txt")
	if err != nil {
		panic(err)
	}
	size, _ := o.GetContentLength()
	fmt.Printf("name=%s size=%d\n", o.Path, size)
	// Output:
	// name=greeting.txt size=18
}

// ExampleStorage_Delete shows how to delete an object and confirm its removal.
func ExampleStorage_Delete() {
	store, _ := memory.NewStorager()

	store.Write("tmp.txt", strings.NewReader("data"), 4) //nolint

	err := store.Delete("tmp.txt")
	if err != nil {
		panic(err)
	}

	_, err = store.Stat("tmp.txt")
	fmt.Println(errors.Is(err, types.ErrNotImplemented) || err != nil)
	// Output:
	// true
}

// ExampleStorage_List shows how to iterate over objects in a directory prefix.
func ExampleStorage_List() {
	store, _ := memory.NewStorager()

	files := []string{"a.txt", "b.txt", "c.txt"}
	for _, f := range files {
		store.Write(f, strings.NewReader("x"), 1) //nolint
	}

	it, err := store.List("")
	if err != nil {
		panic(err)
	}

	var names []string
	for {
		o, err := it.Next()
		if err != nil && errors.Is(err, types.IterateDone) {
			break
		}
		if err != nil {
			panic(err)
		}
		names = append(names, o.Path)
	}
	fmt.Println(len(names))
	// Output:
	// 3
}

// ExampleStorage_Write_withContentType shows how to attach a MIME type to a
// written object using service-level pairs. Note: not all backends persist
// metadata set at write-time; check your backend's documentation.
func ExampleStorage_Write_withContentType() {
	store, _ := memory.NewStorager()

	content := `{"key":"value"}`
	n, err := store.Write("data.json",
		strings.NewReader(content),
		int64(len(content)),
		pairs.WithContentType("application/json"),
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("written %d bytes\n", n)
	// Output:
	// written 15 bytes
}
