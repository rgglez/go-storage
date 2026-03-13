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
// Example code for the ocios (Oracle Cloud Infrastructure Object Storage) service.
//
// NOTE: This service is NOT YET IMPLEMENTED. NewStorager will panic at runtime.
package ocios_test

import (
	"fmt"

	"github.com/rgglez/go-storage/services/ocios"
)

// ExampleNewStorager documents the intended initialization API for OCI Object Storage.
//
// WARNING: This service is not yet implemented. Calling NewStorager will panic.
//
// Intended usage once implemented:
//
//	_, _ = ocios.NewStorager(
//	    pairs.WithCredential("hmac:ACCESS_KEY:SECRET_KEY"),
//	    pairs.WithEndpoint("https://objectstorage.us-phoenix-1.oraclecloud.com"),
//	)
func ExampleNewStorager() {
	_ = ocios.NewStorager // not yet implemented; see examples/basic/main.go
	fmt.Println("ocios: not yet implemented")
}
