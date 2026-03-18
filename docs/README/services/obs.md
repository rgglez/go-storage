# obs — Huawei Object Storage Service

[![Go Reference](https://pkg.go.dev/badge/github.com/rgglez/go-storage/services/obs/v2.svg)](https://pkg.go.dev/github.com/rgglez/go-storage/services/obs/v2)
[![Status](https://img.shields.io/badge/status-stable-brightgreen)](https://github.com/rgglez/go-storage)

Backend for [Huawei Cloud Object Storage Service (OBS)](https://www.huaweicloud.com/intl/en-us/product/obs.html).

## Install

```bash
go get github.com/rgglez/go-storage/services/obs/v2
```

## Usage

```go
import (
    "github.com/rgglez/go-storage/services/obs/v2"
    ps "github.com/rgglez/go-storage/v5/pairs"
)

// Create both Servicer and Storager
srv, sto, err := obs.New(
    ps.WithCredential("hmac:<access_key>:<secret_key>"),
    ps.WithEndpoint("https://obs.cn-north-4.myhuaweicloud.com"),
    ps.WithName("<bucket_name>"),
    ps.WithWorkDir("/optional/prefix/"),
)

// Create Storager only
sto, err := obs.NewStorager(
    ps.WithCredential("hmac:<access_key>:<secret_key>"),
    ps.WithEndpoint("https://obs.cn-north-4.myhuaweicloud.com"),
    ps.WithName("<bucket_name>"),
)
```

## Configuration

| Pair | Type | Required | Description |
|------|------|----------|-------------|
| `credential` | string | Yes | Authentication. Format: `hmac:<access_key>:<secret_key>` |
| `endpoint` | string | Yes | OBS endpoint URL (e.g. `https://obs.cn-north-4.myhuaweicloud.com`) |
| `name` | string | Yes | Bucket name |
| `work_dir` | string | No | Working directory (key prefix). Defaults to `/` |

## Supported Interfaces

| Interface | Supported |
|-----------|-----------|
| `Storager` | Yes |
| `Servicer` | Yes |
| `Multiparter` | Yes |

## References

- [Huawei OBS documentation](https://support.huaweicloud.com/intl/en-us/obs/index.html)
- [Go package reference](https://pkg.go.dev/github.com/rgglez/go-storage/services/obs/v2)
