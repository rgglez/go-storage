# uss — UPYUN Storage Service

[![Go Reference](https://pkg.go.dev/badge/github.com/rgglez/go-storage/services/uss/v3.svg)](https://pkg.go.dev/github.com/rgglez/go-storage/services/uss/v3)
[![Status](https://img.shields.io/badge/status-beta-yellow)](https://github.com/rgglez/go-storage)

Backend for [UPYUN Storage Service (USS)](https://www.upyun.com/products/file-storage).

> **Beta**: implemented but not fully integration-tested.

## Install

```bash
go get github.com/rgglez/go-storage/services/uss/v3
```

## Usage

```go
import (
    "github.com/rgglez/go-storage/services/uss/v3"
    ps "github.com/rgglez/go-storage/v5/pairs"
)

sto, err := uss.NewStorager(
    ps.WithCredential("basic:<operator>:<password>"),
    ps.WithName("<service_name>"),
    ps.WithWorkDir("/optional/prefix/"),
)
```

## Configuration

| Pair | Type | Required | Description |
|------|------|----------|-------------|
| `credential` | string | Yes | Authentication. Format: `basic:<operator>:<password>` |
| `name` | string | Yes | UPYUN service (bucket) name |
| `work_dir` | string | No | Working directory (key prefix). Defaults to `/` |

## Supported Interfaces

| Interface | Supported |
|-----------|-----------|
| `Storager` | Yes |
| `Multiparter` | Yes |

## References

- [UPYUN USS documentation](https://help.upyun.com/knowledge-base/rest_api/)
- [Go package reference](https://pkg.go.dev/github.com/rgglez/go-storage/services/uss/v3)
