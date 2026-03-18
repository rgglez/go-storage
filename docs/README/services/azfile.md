# azfile — Azure File Storage

[![Go Reference](https://pkg.go.dev/badge/github.com/rgglez/go-storage/services/azfile.svg)](https://pkg.go.dev/github.com/rgglez/go-storage/services/azfile)
[![Status](https://img.shields.io/badge/status-stable-brightgreen)](https://github.com/rgglez/go-storage)

Backend for [Azure File Storage](https://azure.microsoft.com/en-us/products/storage/files/).

## Install

```bash
go get github.com/rgglez/go-storage/services/azfile
```

## Usage

```go
import (
    "github.com/rgglez/go-storage/services/azfile"
    ps "github.com/rgglez/go-storage/v5/pairs"
)

// Create both Servicer and Storager
srv, sto, err := azfile.New(
    ps.WithCredential("hmac:<account_name>:<account_key>"),
    ps.WithEndpoint("https://<account_name>.file.core.windows.net"),
    ps.WithName("<share_name>"),
    ps.WithWorkDir("/optional/dir/"),
)

// Create Storager only
sto, err := azfile.NewStorager(
    ps.WithCredential("hmac:<account_name>:<account_key>"),
    ps.WithEndpoint("https://<account_name>.file.core.windows.net"),
    ps.WithName("<share_name>"),
)
```

## Configuration

| Pair | Type | Required | Description |
|------|------|----------|-------------|
| `credential` | string | Yes | Authentication. Format: `hmac:<account_name>:<account_key>` |
| `endpoint` | string | Yes | Azure File service URL. Format: `https://<account>.file.core.windows.net` |
| `name` | string | Yes | File share name |
| `work_dir` | string | No | Working directory path. Defaults to `/` |

## Supported Interfaces

| Interface | Supported |
|-----------|-----------|
| `Storager` | Yes |
| `Servicer` | Yes |
| `Direr` | Yes |
| `Appender` | Yes |

## References

- [Azure File Storage documentation](https://azure.microsoft.com/en-us/products/storage/files/)
- [Go package reference](https://pkg.go.dev/github.com/rgglez/go-storage/services/azfile)
