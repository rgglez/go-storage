# bos — Baidu Object Storage

[![Go Reference](https://pkg.go.dev/badge/github.com/rgglez/go-storage/services/bos/v2.svg)](https://pkg.go.dev/github.com/rgglez/go-storage/services/bos/v2)
[![Status](https://img.shields.io/badge/status-stable-brightgreen)](https://github.com/rgglez/go-storage)

Backend for [Baidu Object Storage (BOS)](https://cloud.baidu.com/product/bos.html).

## Install

```bash
go get github.com/rgglez/go-storage/services/bos/v2
```

## Usage

```go
import (
    "github.com/rgglez/go-storage/services/bos/v2"
    ps "github.com/rgglez/go-storage/v5/pairs"
)

// Create both Servicer and Storager
srv, sto, err := bos.New(
    ps.WithCredential("hmac:<access_key>:<secret_key>"),
    ps.WithEndpoint("https://bj.bcebos.com"),
    ps.WithName("<bucket_name>"),
    ps.WithWorkDir("/optional/prefix/"),
)

// Create Storager only
sto, err := bos.NewStorager(
    ps.WithCredential("hmac:<access_key>:<secret_key>"),
    ps.WithEndpoint("https://bj.bcebos.com"),
    ps.WithName("<bucket_name>"),
)
```

## Configuration

| Pair | Type | Required | Description |
|------|------|----------|-------------|
| `credential` | string | Yes | Authentication. Format: `hmac:<access_key>:<secret_key>` |
| `endpoint` | string | Yes | BOS endpoint URL (e.g. `https://bj.bcebos.com`) |
| `name` | string | Yes | Bucket name |
| `work_dir` | string | No | Working directory (key prefix). Defaults to `/` |
| `location` | string | No | Region (e.g. `bj`, `gz`, `su`) |

## Supported Interfaces

| Interface | Supported |
|-----------|-----------|
| `Storager` | Yes |
| `Servicer` | Yes |
| `Multiparter` | Yes |
| `Appender` | Yes |

## References

- [Baidu Object Storage documentation](https://cloud.baidu.com/doc/BOS/index.html)
- [Go package reference](https://pkg.go.dev/github.com/rgglez/go-storage/services/bos/v2)
