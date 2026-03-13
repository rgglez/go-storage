// Basic example for the s3 (Amazon S3 / S3-compatible) service.
//
// Run from this directory:
//
//	S3_CREDENTIAL=hmac:ACCESS_KEY_ID:SECRET_ACCESS_KEY \
//	S3_BUCKET=my-bucket \
//	S3_LOCATION=us-east-1 \
//	go run main.go
//
// For S3-compatible services (MinIO, Ceph RGW, …) also set:
//
//	S3_ENDPOINT=https://play.min.io
//
// Required environment variables:
//
//	S3_CREDENTIAL  – HMAC credential: "hmac:ACCESS_KEY_ID:SECRET_KEY"
//	S3_BUCKET      – bucket name
//	S3_LOCATION    – AWS region or zone (e.g. "us-east-1")
//
// Optional:
//
//	S3_ENDPOINT    – custom endpoint for S3-compatible stores
//	S3_WORK_DIR    – key prefix to use as working directory (must start and end with "/")
package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/rgglez/go-storage/v5/pairs"
	"github.com/rgglez/go-storage/v5/types"

	"github.com/rgglez/go-storage/services/s3/v3"
)

func main() {
	// ---- configuration ------------------------------------------------
	credential := requireEnv("S3_CREDENTIAL")
	bucket := requireEnv("S3_BUCKET")
	location := requireEnv("S3_LOCATION")

	ps := []types.Pair{
		pairs.WithCredential(credential),
		pairs.WithName(bucket),
		pairs.WithLocation(location),
	}
	if ep := os.Getenv("S3_ENDPOINT"); ep != "" {
		ps = append(ps, pairs.WithEndpoint(ep))
	}
	if wd := os.Getenv("S3_WORK_DIR"); wd != "" {
		ps = append(ps, pairs.WithWorkDir(wd))
	}

	// ---- create storager ----------------------------------------------
	_, store, err := s3.New(ps...)
	if err != nil {
		fatal("new", err)
	}
	fmt.Printf("storager: %s\n\n", store)

	// ---- write --------------------------------------------------------
	const path = "go-storage-example/hello.txt"
	const content = "Hello from go-storage s3 example!"

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
