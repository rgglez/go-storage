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
// Basic example for the ftp service.
//
// Run from this directory:
//
//	FTP_ENDPOINT=tcp+addr://ftp.example.com:21 \
//	FTP_CREDENTIAL=basic:user:password \
//	FTP_WORK_DIR=/upload \
//	go run main.go
//
// Required environment variables:
//
//	FTP_ENDPOINT    – TCP endpoint, e.g. "tcp+addr://host:21"
//	FTP_CREDENTIAL  – basic credential, e.g. "basic:user:password"
//
// Optional:
//
//	FTP_WORK_DIR    – remote work directory (default: "/")
package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/rgglez/go-storage/v5/pairs"
	"github.com/rgglez/go-storage/v5/types"

	"github.com/rgglez/go-storage/services/ftp"
)

func main() {
	// ---- configuration ------------------------------------------------
	endpoint := requireEnv("FTP_ENDPOINT")
	credential := requireEnv("FTP_CREDENTIAL")
	workDir := os.Getenv("FTP_WORK_DIR")

	ps := []types.Pair{
		pairs.WithEndpoint(endpoint),
		pairs.WithCredential(credential),
	}
	if workDir != "" {
		ps = append(ps, pairs.WithWorkDir(workDir))
	}

	// ---- create storager ----------------------------------------------
	store, err := ftp.NewStorager(ps...)
	if err != nil {
		fatal("new storager", err)
	}
	fmt.Printf("storager: %s\n\n", store)

	// ---- write --------------------------------------------------------
	const path = "go-storage-example.txt"
	const content = "Hello from go-storage ftp example!"

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
	it, err := store.List("")
	if err != nil {
		fatal("list", err)
	}
	fmt.Print("list (first 5):\n")
	count := 0
	for {
		obj, err := it.Next()
		if errors.Is(err, types.IterateDone) {
			break
		}
		if err != nil {
			fatal("list.Next", err)
		}
		fmt.Printf("  - %s\n", obj.Path)
		count++
		if count >= 5 {
			fmt.Println("  ...")
			break
		}
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

// Ensure the pairs package is used (suppresses "imported and not used" on
// re-imports when the caller adjusts the example).
var _ = strings.NewReader
