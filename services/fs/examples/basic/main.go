// Basic example for the fs (local filesystem) service.
//
// Run from this directory:
//
//	go run main.go
//
// Optional environment variables:
//
//	WORK_DIR  – absolute path to use as work directory (default: a temp dir)
package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/rgglez/go-storage/v5/pairs"
	"github.com/rgglez/go-storage/v5/types"

	fs "github.com/rgglez/go-storage/services/fs/v4"
)

func main() {
	// ---- configuration ------------------------------------------------
	workDir := os.Getenv("WORK_DIR")
	if workDir == "" {
		dir, err := os.MkdirTemp("", "go-storage-fs-example-*")
		if err != nil {
			fatal("create temp dir", err)
		}
		defer os.RemoveAll(dir)
		workDir = dir
	}

	// ---- create storager ----------------------------------------------
	store, err := fs.NewStorager(pairs.WithWorkDir(workDir))
	if err != nil {
		fatal("new storager", err)
	}
	fmt.Printf("storager: %s\n\n", store)

	// ---- write --------------------------------------------------------
	const path = "hello.txt"
	const content = "Hello, go-storage!"

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
	fmt.Printf("stat   %q  → size=%d mode=%s\n", path, size, o.Mode)

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

func fatal(op string, err error) {
	fmt.Fprintf(os.Stderr, "ERROR [%s]: %v\n", op, err)
	os.Exit(1)
}
