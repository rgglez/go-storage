# ocios — Oracle Cloud Infrastructure Object Storage

[![Go Reference](https://pkg.go.dev/badge/github.com/rgglez/go-storage/services/ocios.svg)](https://pkg.go.dev/github.com/rgglez/go-storage/services/ocios)
[![Status](https://img.shields.io/badge/status-stable-brightgreen)](https://github.com/rgglez/go-storage)

Backend for [Oracle Cloud Infrastructure (OCI) Object Storage](https://www.oracle.com/cloud/storage/object-storage/).

## Install

```bash
go get github.com/rgglez/go-storage/services/ocios
```

## Usage

```go
import (
    "github.com/rgglez/go-storage/services/ocios"
    ps "github.com/rgglez/go-storage/v5/pairs"
)

// Create both Servicer and Storager
srv, sto, err := ocios.New(
    ps.WithCredential("file:/home/user/.oci/config"),
    ps.WithName("<bucket_name>"),
    ps.WithLocation("<region>"),   // e.g. "us-ashburn-1"
    ps.WithWorkDir("/optional/prefix/"),
)
```

## Configuration

| Pair | Type | Required | Description |
|------|------|----------|-------------|
| `credential` | string | Yes | OCI config file path. Format: `file:<path>` |
| `name` | string | Yes | Bucket name |
| `location` | string | Yes | OCI region identifier (e.g. `us-ashburn-1`, `eu-frankfurt-1`) |
| `work_dir` | string | No | Working directory (key prefix). Defaults to `/` |

## Supported Interfaces

| Interface | Supported |
|-----------|-----------|
| `Storager` | Yes |
| `Servicer` | Yes |
| `Multiparter` | Yes |

## Notes

> **Alpha/stub status**: some operations may not be fully implemented.

## References

- [OCI Object Storage documentation](https://docs.oracle.com/en-us/iaas/Content/Object/home.htm)
- [Go package reference](https://pkg.go.dev/github.com/rgglez/go-storage/services/ocios)
