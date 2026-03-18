# zip — ZIP Archive Files

[![Go Reference](https://pkg.go.dev/badge/github.com/rgglez/go-storage/services/zip.svg)](https://pkg.go.dev/github.com/rgglez/go-storage/services/zip)
[![Status](https://img.shields.io/badge/status-alpha-red)](https://github.com/rgglez/go-storage)

Read-access backend for ZIP archive files.

> **Alpha**: still under active development; API may change and some operations may not be implemented.

## Install

```bash
go get github.com/rgglez/go-storage/services/zip
```

## Usage

```go
import (
    "github.com/rgglez/go-storage/services/zip"
    ps "github.com/rgglez/go-storage/v5/pairs"
)

sto, err := zip.NewStorager(
    ps.WithName("/path/to/archive.zip"),
    ps.WithWorkDir("/optional/prefix/"),
)
```

## Configuration

| Pair | Type | Required | Description |
|------|------|----------|-------------|
| `name` | string | Yes | Path to the `.zip` file |
| `work_dir` | string | No | Path prefix inside the archive. Defaults to `/` |

## Supported Interfaces

| Interface | Supported |
|-----------|-----------|
| `Storager` | Partial (read-only) |

## References

- [ZIP file format specification](https://pkware.cachefly.net/webdocs/casestudies/APPNOTE.TXT)
- [Go package reference](https://pkg.go.dev/github.com/rgglez/go-storage/services/zip)
