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
// Basic example for the gdrive (Google Drive) service.
//
// Run from this directory:
//
//	GDRIVE_CREDENTIAL=file:/path/to/service-account-key.json \
//	GDRIVE_FOLDER="My Drive Folder" \
//	go run main.go
//
// Setup:
//  1. Create a service account in Google Cloud Console.
//  2. Download the JSON key and set GDRIVE_CREDENTIAL to "file:/abs/path/key.json".
//  3. Share the Drive folder with the service account email address.
//
// Required environment variables:
//
//	GDRIVE_CREDENTIAL  – path to Google service account JSON: "file:/path/to/key.json"
//	GDRIVE_FOLDER      – Drive folder name to use as the root
//
// Optional:
//
//	GDRIVE_WORK_DIR    – path prefix within the folder
package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/rgglez/go-storage/v5/pairs"
	"github.com/rgglez/go-storage/v5/types"

	"github.com/rgglez/go-storage/services/gdrive"
)

func main() {
	credential := requireEnv("GDRIVE_CREDENTIAL")
	folder := requireEnv("GDRIVE_FOLDER")

	ps := []types.Pair{
		pairs.WithCredential(credential),
		pairs.WithName(folder),
	}
	if wd := os.Getenv("GDRIVE_WORK_DIR"); wd != "" {
		ps = append(ps, pairs.WithWorkDir(wd))
	}

	store, err := gdrive.NewStorager(ps...)
	if err != nil {
		fatal("new", err)
	}
	fmt.Printf("storager: %s\n\n", store)

	const path = "go-storage-example/hello.txt"
	const content = "Hello from go-storage gdrive example!"

	n, err := store.Write(path, strings.NewReader(content), int64(len(content)))
	if err != nil {
		fatal("write", err)
	}
	fmt.Printf("write  %q  → %d bytes\n", path, n)

	o, err := store.Stat(path)
	if err != nil {
		fatal("stat", err)
	}
	size, _ := o.GetContentLength()
	fmt.Printf("stat   %q  → size=%d\n", path, size)

	var buf bytes.Buffer
	n, err = store.Read(path, &buf)
	if err != nil {
		fatal("read", err)
	}
	fmt.Printf("read   %q  → %d bytes: %q\n", path, n, buf.String())

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
