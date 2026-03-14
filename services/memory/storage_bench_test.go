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
package memory

import (
	"bytes"
	"io"
	"testing"

	"github.com/google/uuid"

	"github.com/rgglez/go-storage/v5/pkg/randbytes"
	"github.com/rgglez/go-storage/v5/types"
)

func setup(b *testing.B, size int64) (store *Storage, path string) {
	root := newObject("", nil, types.ModeDir)
	root.parent = root

	store = &Storage{
		root:    root,
		workDir: "/",
	}

	path = uuid.NewString()
	content, err := io.ReadAll(io.LimitReader(randbytes.NewRand(), size))
	if err != nil {
		b.Fatal(err)
	}

	_, err = store.Write(path, bytes.NewReader(content), size)
	if err != nil {
		b.Fatal(err)
	}
	return
}

func BenchmarkStorage_Read(b *testing.B) {
	cases := []struct {
		name string
		size int64
	}{
		{"64B", 64},
		{"4k", 4 * 1024},
		{"64M", 64 * 1024 * 1024},
	}
	for _, v := range cases {
		b.Run(v.name, func(b *testing.B) {
			store, path := setup(b, v.size)

			b.SetBytes(v.size)
			for i := 0; i < b.N; i++ {
				_, _ = store.Read(path, io.Discard)
			}
		})
	}
}

func BenchmarkStorage_Write(b *testing.B) {
	cases := []struct {
		name string
		size int64
	}{
		{"64B", 64},
		{"4k", 4 * 1024},
		{"64M", 64 * 1024 * 1024},
	}
	for _, v := range cases {
		b.Run(v.name, func(b *testing.B) {
			store, _ := setup(b, v.size)

			path := uuid.NewString()
			content, err := io.ReadAll(io.LimitReader(randbytes.NewRand(), v.size))
			if err != nil {
				b.Fatal(err)
			}

			b.SetBytes(v.size)
			for i := 0; i < b.N; i++ {
				_, _ = store.Write(path, bytes.NewReader(content), v.size)
			}
		})
	}
}
