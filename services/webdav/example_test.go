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
// Example code for the webdav (WebDAV) service.
//
// NOTE: This service is NOT YET IMPLEMENTED. NewStorager will panic at runtime.
package webdav_test

import (
	"fmt"

	"github.com/rgglez/go-storage/services/webdav"
)

// ExampleNewStorager documents the intended initialization API for WebDAV.
//
// WARNING: This service is not yet implemented. Calling NewStorager will panic.
//
// Intended usage once implemented:
//
//	_, _ = webdav.NewStorager(
//	    pairs.WithEndpoint("https://dav.example.com"),
//	    pairs.WithCredential("basic:USERNAME:PASSWORD"),
//	)
func ExampleNewStorager() {
	_ = webdav.NewStorager // not yet implemented; see examples/basic/main.go
	fmt.Println("webdav: not yet implemented")
}
