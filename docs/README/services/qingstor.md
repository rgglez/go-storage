# qingstor — QingStor Object Storage

[![Go Reference](https://pkg.go.dev/badge/github.com/rgglez/go-storage/services/qingstor/v4.svg)](https://pkg.go.dev/github.com/rgglez/go-storage/services/qingstor/v4)
[![Status](https://img.shields.io/badge/status-stable-brightgreen)](https://github.com/rgglez/go-storage)

Backend for [QingStor Object Storage](https://www.qingcloud.com/products/qingstor/).

## Install

```bash
go get github.com/rgglez/go-storage/services/qingstor/v4
```

## Usage

```go
import (
    "github.com/rgglez/go-storage/services/qingstor/v4"
    ps "github.com/rgglez/go-storage/v5/pairs"
)

// Create both Servicer and Storager
srv, sto, err := qingstor.New(
    ps.WithCredential("hmac:<access_key_id>:<secret_access_key>"),
    ps.WithEndpoint("https://qingstor.com"),
    ps.WithName("<bucket_name>"),
    ps.WithLocation("<zone>"),     // e.g. "pek3b", "sh1b"
    ps.WithWorkDir("/optional/prefix/"),
)

// Create Storager only
sto, err := qingstor.NewStorager(
    ps.WithCredential("hmac:<access_key_id>:<secret_access_key>"),
    ps.WithEndpoint("https://qingstor.com"),
    ps.WithName("<bucket_name>"),
    ps.WithLocation("pek3b"),
)
```

## Configuration

| Pair | Type | Required | Description |
|------|------|----------|-------------|
| `credential` | string | Yes | Authentication. Format: `hmac:<access_key_id>:<secret_access_key>` |
| `endpoint` | string | Yes | QingStor API endpoint (e.g. `https://qingstor.com`) |
| `name` | string | Yes | Bucket name |
| `location` | string | Yes | Zone (e.g. `pek3b`, `sh1b`) |
| `work_dir` | string | No | Working directory (key prefix). Defaults to `/` |

## Supported Interfaces

| Interface | Supported |
|-----------|-----------|
| `Storager` | Yes |
| `Servicer` | Yes |
| `Multiparter` | Yes |
| `Appender` | Yes |
| `Copier` | Yes |
| `Reacher` | Yes |

## References

- [QingStor documentation](https://docsv4.qingcloud.com/user_guide/storage/object_storage/)
- [Go package reference](https://pkg.go.dev/github.com/rgglez/go-storage/services/qingstor/v4)
