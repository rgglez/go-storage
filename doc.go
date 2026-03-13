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
/*
Package storage intends to provide a unified storage layer for Golang.

Goals

- Production ready: high test coverage, enterprise storage software adaptation, semantic versioning, well documented.

- High performance: more code generation, less runtime reflect.

- Vendor agnostic: more generic abstraction, less internal details.

Examples

The most common case to use a Storager service could be following:

1. Init a storager.

    store, err := fs.NewStorager(pairs.WithWorkDir("/tmp"))
	if err != nil {
		log.Fatalf("service init failed: %v", err)
	}

2. Use Storager API to maintain data.

	var buf bytes.Buffer

	n, err := store.Read("path/to/file", &buf)
	if err != nil {
		log.Printf("storager read: %v", err)
	}

*/
package storage
