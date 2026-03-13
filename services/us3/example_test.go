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
// Example code for the us3 (UCloud Object Storage) service.
//
// NOTE: This service is NOT YET IMPLEMENTED. NewStorager will panic at runtime.
// This file is compiled but not executed and serves as documentation for the
// intended API once the implementation is complete.
package us3_test

import (
	"fmt"

	"github.com/rgglez/go-storage/v5/pairs"

	"github.com/rgglez/go-storage/services/us3"
)

// ExampleNewStorager shows the intended initialization API for UCloud US3.
//
// WARNING: This service is not yet implemented. Calling NewStorager will panic.
//
//	export US3_CREDENTIAL=hmac:PUBLIC_KEY:PRIVATE_KEY
//	export US3_ENDPOINT=https://cn-bj.ufileos.com
func ExampleNewStorager() {
	// TODO: remove this guard when the implementation is complete.
	fmt.Println("us3: not yet implemented")
	return

	//nolint:govet // unreachable code is intentional — documents the intended API
	_, _ = us3.NewStorager(
		pairs.WithCredential("hmac:PUBLIC_KEY:PRIVATE_KEY"),
		pairs.WithEndpoint("https://cn-bj.ufileos.com"),
	)
}
