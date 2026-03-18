# cos — Tencent Cloud Object Storage

[![Go Reference](https://pkg.go.dev/badge/github.com/rgglez/go-storage/services/cos/v3.svg)](https://pkg.go.dev/github.com/rgglez/go-storage/services/cos/v3)
[![Status](https://img.shields.io/badge/status-stable-brightgreen)](https://github.com/rgglez/go-storage)

Backend for [Tencent Cloud Object Storage (COS)](https://intl.cloud.tencent.com/product/cos).

## Install

```bash
go get github.com/rgglez/go-storage/services/cos/v3
```

## Usage

```go
import (
    "github.com/rgglez/go-storage/services/cos/v3"
    ps "github.com/rgglez/go-storage/v5/pairs"
)

// Create both Servicer and Storager
srv, sto, err := cos.New(
    ps.WithCredential("hmac:<secret_id>:<secret_key>"),
    ps.WithName("<bucket_name>"),
    ps.WithLocation("<region>"),   // e.g. "ap-beijing"
    ps.WithWorkDir("/optional/prefix/"),
)

// Create Storager only
sto, err := cos.NewStorager(
    ps.WithCredential("hmac:<secret_id>:<secret_key>"),
    ps.WithName("<bucket_name>"),
    ps.WithLocation("ap-beijing"),
)
```

## Configuration

| Pair | Type | Required | Description |
|------|------|----------|-------------|
| `credential` | string | Yes | Authentication. Format: `hmac:<secret_id>:<secret_key>` |
| `name` | string | Yes | Bucket name (including appid suffix, e.g. `my-bucket-1250000000`) |
| `location` | string | Yes | Region (e.g. `ap-beijing`, `ap-shanghai`) |
| `work_dir` | string | No | Working directory (key prefix). Defaults to `/` |

## Supported Interfaces

| Interface | Supported |
|-----------|-----------|
| `Storager` | Yes |
| `Servicer` | Yes |
| `Multiparter` | Yes |
| `Appender` | Yes |
| `Copier` | Yes |

## References

- [Tencent COS documentation](https://intl.cloud.tencent.com/document/product/436)
- [Go package reference](https://pkg.go.dev/github.com/rgglez/go-storage/services/cos/v3)
