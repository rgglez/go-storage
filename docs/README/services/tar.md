# tar — TAR Archive Files

[![Go Reference](https://pkg.go.dev/badge/github.com/rgglez/go-storage/services/tar.svg)](https://pkg.go.dev/github.com/rgglez/go-storage/services/tar)
[![Status](https://img.shields.io/badge/status-beta-yellow)](https://github.com/rgglez/go-storage)

Read-access backend for TAR archive files.

> **Beta**: implemented but not fully integration-tested.

## Install

```bash
go get github.com/rgglez/go-storage/services/tar
```

## Usage

```go
import (
    "github.com/rgglez/go-storage/services/tar"
    ps "github.com/rgglez/go-storage/v5/pairs"
)

sto, err := tar.NewStorager(
    ps.WithName("/path/to/archive.tar"),
    ps.WithWorkDir("/optional/prefix/"),
)
```

## Configuration

| Pair | Type | Required | Description |
|------|------|----------|-------------|
| `name` | string | Yes | Path to the `.tar` file |
| `work_dir` | string | No | Path prefix inside the archive. Defaults to `/` |

## Supported Interfaces

| Interface | Supported |
|-----------|-----------|
| `Storager` | Yes (read-only) |

## References

- [Go package reference](https://pkg.go.dev/github.com/rgglez/go-storage/services/tar)
