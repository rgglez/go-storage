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
// Basic example for the gcs (Google Cloud Storage) service.
//
// Run from this directory:
//
//	GCS_CREDENTIAL=file:/path/to/service-account.json \
//	GCS_BUCKET=my-bucket \
//	GCS_PROJECT=my-gcp-project \
//	go run main.go
//
// Credential formats:
//
//	file:/path/to/service-account.json   – service account JSON file
//	base64:<base64-encoded-json>          – inline base64-encoded JSON
//	env:                                  – use GOOGLE_APPLICATION_CREDENTIALS / ADC
//
// Required environment variables:
//
//	GCS_CREDENTIAL  – credential string (see above)
//	GCS_BUCKET      – GCS bucket name
//	GCS_PROJECT     – GCP project ID
//
// Optional:
//
//	GCS_WORK_DIR    – key prefix used as working directory
package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/rgglez/go-storage/v5/pairs"
	"github.com/rgglez/go-storage/v5/types"

	"github.com/rgglez/go-storage/services/gcs/v3"
)

func main() {
	// ---- configuration ------------------------------------------------
	credential := requireEnv("GCS_CREDENTIAL")
	bucket := requireEnv("GCS_BUCKET")
	project := requireEnv("GCS_PROJECT")

	ps := []types.Pair{
		pairs.WithCredential(credential),
		pairs.WithName(bucket),
		gcs.WithProjectID(project),
	}
	if wd := os.Getenv("GCS_WORK_DIR"); wd != "" {
		ps = append(ps, pairs.WithWorkDir(wd))
	}

	// ---- create storager ----------------------------------------------
	_, store, err := gcs.New(ps...)
	if err != nil {
		fatal("new", err)
	}
	fmt.Printf("storager: %s\n\n", store)

	// ---- write --------------------------------------------------------
	const path = "go-storage-example/hello.txt"
	const content = "Hello from go-storage gcs example!"

	n, err := store.Write(path, strings.NewReader(content), int64(len(content)))
	if err != nil {
		fatal("write", err)
	}
	fmt.Printf("write  %q  → %d bytes\n", path, n)

	// ---- stat ---------------------------------------------------------
	o, err := store.Stat(path)
	if err != nil {
		fatal("stat", err)
	}
	size, _ := o.GetContentLength()
	fmt.Printf("stat   %q  → size=%d\n", path, size)

	// ---- read ---------------------------------------------------------
	var buf bytes.Buffer
	n, err = store.Read(path, &buf)
	if err != nil {
		fatal("read", err)
	}
	fmt.Printf("read   %q  → %d bytes: %q\n", path, n, buf.String())

	// ---- list ---------------------------------------------------------
	it, err := store.List("go-storage-example/")
	if err != nil {
		fatal("list", err)
	}
	fmt.Print("list:\n")
	for {
		obj, err := it.Next()
		if errors.Is(err, types.IterateDone) {
			break
		}
		if err != nil {
			fatal("list.Next", err)
		}
		fmt.Printf("  - %s\n", obj.Path)
	}

	// ---- delete -------------------------------------------------------
	if err := store.Delete(path); err != nil {
		fatal("delete", err)
	}
	fmt.Printf("delete %q  → ok\n", path)
}

func requireEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		fmt.Fprintf(os.Stderr, "ERROR: environment variable %s is required\n", key)
		os.Exit(1)
	}
	return v
}

func fatal(op string, err error) {
	fmt.Fprintf(os.Stderr, "ERROR [%s]: %v\n", op, err)
	os.Exit(1)
}

var _ = strings.NewReader // silence unused-import linter if content is changed
