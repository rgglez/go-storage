# azblob — Azure Blob Storage

[![Go Reference](https://pkg.go.dev/badge/github.com/rgglez/go-storage/services/azblob/v3.svg)](https://pkg.go.dev/github.com/rgglez/go-storage/services/azblob/v3)
[![Status](https://img.shields.io/badge/status-stable-brightgreen)](https://github.com/rgglez/go-storage)

Backend for [Azure Blob Storage](https://docs.microsoft.com/en-us/azure/storage/blobs/storage-blobs-introduction).

## Install

```bash
go get github.com/rgglez/go-storage/services/azblob/v3
```

## Usage

```go
import (
    "github.com/rgglez/go-storage/services/azblob/v3"
    ps "github.com/rgglez/go-storage/v5/pairs"
)

// Create both Servicer and Storager
srv, sto, err := azblob.New(
    ps.WithCredential("hmac:<account_name>:<account_key>"),
    ps.WithEndpoint("https://<account_name>.blob.core.windows.net"),
    ps.WithName("<container_name>"),
    ps.WithWorkDir("/optional/prefix/"),
)

// Create Storager only
sto, err := azblob.NewStorager(
    ps.WithCredential("hmac:<account_name>:<account_key>"),
    ps.WithEndpoint("https://<account_name>.blob.core.windows.net"),
    ps.WithName("<container_name>"),
)
```

## Configuration

| Pair | Type | Required | Description |
|------|------|----------|-------------|
| `credential` | string | Yes | Authentication. Format: `hmac:<account_name>:<account_key>` |
| `endpoint` | string | Yes | Azure Blob service URL. Format: `https://<account>.blob.core.windows.net` |
| `name` | string | Yes | Container name |
| `work_dir` | string | No | Working directory (key prefix). Defaults to `/` |

## Supported Interfaces

| Interface | Supported |
|-----------|-----------|
| `Storager` | Yes |
| `Servicer` | Yes |
| `Appender` | Yes |
| `Multiparter` | Yes |
| `Copier` | Yes |

## Storage Classes

| Constant | Azure Tier | Description |
|----------|-----------|-------------|
| `azblob.StorageClassArchive` | Archive | Lowest cost, highest latency |
| `azblob.StorageClassCool` | Cool | Infrequent access |
| `azblob.StorageClassHot` | Hot | Frequent access |
| `azblob.StorageClassNone` | (default) | No tier specified |

```go
import "github.com/rgglez/go-storage/services/azblob/v3"

n, err := sto.Write("path/to/blob", r, size,
    azblob.WithStorageClass(azblob.StorageClassCool),
)
```

## Server-Side Encryption

Customer-provided key (AES-256):

```go
n, err := sto.Write("path", r, size,
    azblob.WithServerSideEncryptionCustomerAlgorithm("AES256"),
    azblob.WithServerSideEncryptionCustomerKey([]byte{...32 bytes...}),
    azblob.WithServerSideEncryptionCustomerScope("my-scope"),
)
```

## Limits

| Operation | Limit |
|-----------|-------|
| Single write (`Write`) | 5 000 MB |
| Single append chunk (`WriteAppend`) | 4 MB |
| Max append operations | 50 000 |

## References

- [Azure Blob Storage documentation](https://docs.microsoft.com/en-us/azure/storage/blobs/)
- [Go package reference](https://pkg.go.dev/github.com/rgglez/go-storage/services/azblob/v3)
