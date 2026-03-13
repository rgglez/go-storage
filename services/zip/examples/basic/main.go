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
// Basic example for the zip (ZIP archive storage) service.
//
// WARNING: This service is NOT YET IMPLEMENTED.
// This program exits immediately with an error message.
// See https://github.com/rgglez/go-storage/issues for implementation status.
//
// Intended usage (once implemented):
//
//	ZIP_FILE=file:/path/to/archive.zip \
//	go run main.go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stderr, "ERROR: zip service is not yet implemented.")
	fmt.Fprintln(os.Stderr, "The ZIP archive storage backend has been defined but the implementation")
	fmt.Fprintln(os.Stderr, "is pending.")
	fmt.Fprintln(os.Stderr, "See https://github.com/rgglez/go-storage for updates.")
	os.Exit(1)
}
