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
// cloud-storage demonstrates that the same application code works unchanged
// against any S3-compatible cloud backend. Select the backend at runtime via
// the BACKEND environment variable.
//
// Supported backends: s3, oss, kodo, minio
//
// Run with Amazon S3:
//
//	BACKEND=s3 \
//	CREDENTIAL=hmac:ACCESS_KEY_ID:SECRET_ACCESS_KEY \
//	BUCKET=my-bucket \
//	LOCATION=us-east-1 \
//	go run .
//
// Run with Alibaba OSS:
//
//	BACKEND=oss \
//	CREDENTIAL=hmac:ACCESS_KEY_ID:ACCESS_KEY_SECRET \
//	ENDPOINT=https://oss-cn-hangzhou.aliyuncs.com \
//	BUCKET=my-bucket \
//	go run .
//
// Run with MinIO:
//
//	BACKEND=minio \
//	ENDPOINT=https://play.min.io \
//	CREDENTIAL=hmac:ACCESS_KEY:SECRET_KEY \
//	BUCKET=my-bucket \
//	go run .
//
// Run with Qiniu Kodo:
//
//	BACKEND=kodo \
//	CREDENTIAL=hmac:ACCESS_KEY:SECRET_KEY \
//	ENDPOINT=https://up.qiniup.com \
//	BUCKET=my-bucket \
//	go run .
//
// Environment variables:
//
//	BACKEND     – one of: s3, oss, kodo, minio (required)
//	CREDENTIAL  – HMAC credential "hmac:KEY:SECRET" (required)
//	BUCKET      – bucket/container name (required)
//	ENDPOINT    – service endpoint (required for oss, kodo, minio; optional for s3)
//	LOCATION    – region (required for s3; not used by other backends)
//	WORK_DIR    – key prefix, must start and end with "/" (optional)
package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/rgglez/go-storage/v5/pairs"
	"github.com/rgglez/go-storage/v5/types"

	"github.com/rgglez/go-storage/services/kodo/v3"
	"github.com/rgglez/go-storage/services/minio"
	"github.com/rgglez/go-storage/services/oss/v3"
	s3 "github.com/rgglez/go-storage/services/s3/v3"
)

func main() {
	backend := requireEnv("BACKEND")
	credential := requireEnv("CREDENTIAL")
	bucket := requireEnv("BUCKET")
	endpoint := os.Getenv("ENDPOINT")
	location := os.Getenv("LOCATION")
	workDir := os.Getenv("WORK_DIR")

	store, err := initBackend(backend, credential, bucket, endpoint, location, workDir)
	if err != nil {
		fatal("init", err)
	}
	fmt.Printf("backend: %s\nstorager: %s\n\n", backend, store)

	if err := runWorkload(store); err != nil {
		fmt.Fprintf(os.Stderr, "workload failed: %v\n", err)
		os.Exit(1)
	}
}

// initBackend constructs the appropriate Storager based on the backend name.
// This is the only function that knows about individual backends.
// Everything else uses only the generic types.Storager interface.
func initBackend(backend, credential, bucket, endpoint, location, workDir string) (types.Storager, error) {
	basePairs := []types.Pair{
		pairs.WithCredential(credential),
		pairs.WithName(bucket),
	}
	if workDir != "" {
		basePairs = append(basePairs, pairs.WithWorkDir(workDir))
	}

	switch backend {
	case "s3":
		ps := append(basePairs, pairs.WithLocation(location))
		if endpoint != "" {
			ps = append(ps, pairs.WithEndpoint(endpoint))
		}
		_, store, err := s3.New(ps...)
		return store, err

	case "oss":
		if endpoint == "" {
			return nil, fmt.Errorf("ENDPOINT is required for oss backend")
		}
		ps := append(basePairs, pairs.WithEndpoint(endpoint))
		_, store, err := oss.New(ps...)
		return store, err

	case "kodo":
		if endpoint == "" {
			return nil, fmt.Errorf("ENDPOINT is required for kodo backend")
		}
		ps := append(basePairs, pairs.WithEndpoint(endpoint))
		_, store, err := kodo.New(ps...)
		return store, err

	case "minio":
		if endpoint == "" {
			return nil, fmt.Errorf("ENDPOINT is required for minio backend")
		}
		ps := append(basePairs, pairs.WithEndpoint(endpoint))
		_, store, err := minio.New(ps...)
		return store, err

	default:
		return nil, fmt.Errorf("unknown backend %q; supported: s3, oss, kodo, minio", backend)
	}
}

// runWorkload executes a standard set of storage operations against any
// Storager implementation. The caller has no knowledge of which backend it
// is talking to — this is the core value of the go-storage abstraction.
func runWorkload(store types.Storager) error {
	const (
		path    = "go-storage-cloud-example/hello.txt"
		content = "Hello from go-storage cloud-storage example!"
	)

	// write
	n, err := store.Write(path, strings.NewReader(content), int64(len(content)))
	if err != nil {
		return fmt.Errorf("write: %w", err)
	}
	fmt.Printf("write  %q  → %d bytes\n", path, n)

	// stat
	o, err := store.Stat(path)
	if err != nil {
		return fmt.Errorf("stat: %w", err)
	}
	size, _ := o.GetContentLength()
	fmt.Printf("stat   %q  → size=%d\n", path, size)

	// read
	var buf bytes.Buffer
	n, err = store.Read(path, &buf)
	if err != nil {
		return fmt.Errorf("read: %w", err)
	}
	fmt.Printf("read   %q  → %d bytes: %q\n", path, n, buf.String())

	// list
	it, err := store.List("go-storage-cloud-example/")
	if err != nil {
		return fmt.Errorf("list: %w", err)
	}
	fmt.Print("list:\n")
	for {
		obj, err := it.Next()
		if errors.Is(err, types.IterateDone) {
			break
		}
		if err != nil {
			return fmt.Errorf("list.Next: %w", err)
		}
		fmt.Printf("  - %s\n", obj.Path)
	}

	// delete
	if err := store.Delete(path); err != nil {
		return fmt.Errorf("delete: %w", err)
	}
	fmt.Printf("delete %q  → ok\n", path)
	return nil
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
