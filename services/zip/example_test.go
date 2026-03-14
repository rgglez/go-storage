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
//
// Example code for the zip (ZIP archive storage) service.
//
// NOTE: This service is NOT YET IMPLEMENTED. NewStorager will panic at runtime.
package zip_test

import (
	"fmt"

	"github.com/rgglez/go-storage/v5/zip"
)

// ExampleNewStorager documents the intended initialization API for ZIP archive storage.
//
// WARNING: This service is not yet implemented. Calling NewStorager will panic.
//
// Intended usage once implemented:
//
//	_, _ = zip.NewStorager(
//	    pairs.WithEndpoint("file:/path/to/archive.zip"),
//	)
func ExampleNewStorager() {
	_ = zip.NewStorager // not yet implemented; see examples/basic/main.go
	fmt.Println("zip: not yet implemented")
}
