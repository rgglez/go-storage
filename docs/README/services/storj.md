# storj — Storj Decentralized Cloud Storage

[![Go Reference](https://pkg.go.dev/badge/github.com/rgglez/go-storage/services/storj.svg)](https://pkg.go.dev/github.com/rgglez/go-storage/services/storj)
[![Status](https://img.shields.io/badge/status-alpha-red)](https://github.com/rgglez/go-storage)

Backend for [Storj Decentralized Cloud Storage](https://www.storj.io/).

> **Alpha**: still under active development; API may change and some operations may not be implemented.

## Install

```bash
go get github.com/rgglez/go-storage/services/storj
```

## Usage

```go
import (
    "github.com/rgglez/go-storage/services/storj"
    ps "github.com/rgglez/go-storage/v5/pairs"
)

srv, sto, err := storj.New(
    ps.WithCredential("token:<access_grant>"),
    ps.WithName("<bucket_name>"),
    ps.WithWorkDir("/optional/prefix/"),
)
```

## Configuration

| Pair | Type | Required | Description |
|------|------|----------|-------------|
| `credential` | string | Yes | Storj access grant (serialized macaroon). Format: `token:<access_grant>` |
| `name` | string | Yes | Bucket name |
| `work_dir` | string | No | Working directory (key prefix). Defaults to `/` |

## Supported Interfaces

| Interface | Supported |
|-----------|-----------|
| `Storager` | Partial |
| `Servicer` | Partial |

## References

- [Storj documentation](https://docs.storj.io/)
- [Storj access grants](https://docs.storj.io/dcs/getting-started/quickstart-uplink-cli/uploading-your-first-object/create-first-access-grant)
- [Go package reference](https://pkg.go.dev/github.com/rgglez/go-storage/services/storj)
