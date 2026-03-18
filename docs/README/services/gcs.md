# gcs — Google Cloud Storage

[![Go Reference](https://pkg.go.dev/badge/github.com/rgglez/go-storage/services/gcs/v3.svg)](https://pkg.go.dev/github.com/rgglez/go-storage/services/gcs/v3)
[![Status](https://img.shields.io/badge/status-stable-brightgreen)](https://github.com/rgglez/go-storage)

Backend for [Google Cloud Storage (GCS)](https://cloud.google.com/storage/).

## Install

```bash
go get github.com/rgglez/go-storage/services/gcs/v3
```

## Usage

```go
import (
    "github.com/rgglez/go-storage/services/gcs/v3"
    ps "github.com/rgglez/go-storage/v5/pairs"
)

// Using a service account key file
srv, sto, err := gcs.New(
    ps.WithCredential("file:/path/to/service-account.json"),
    ps.WithName("<bucket_name>"),
    ps.WithProjectID("my-gcp-project"),
    ps.WithWorkDir("/optional/prefix/"),
)

// Using a base64-encoded service account key
srv, sto, err := gcs.New(
    ps.WithCredential("base64:<base64-encoded-json>"),
    ps.WithName("<bucket_name>"),
    ps.WithProjectID("my-gcp-project"),
)

// Using Application Default Credentials (ADC)
srv, sto, err := gcs.New(
    ps.WithCredential("env"),
    ps.WithName("<bucket_name>"),
    ps.WithProjectID("my-gcp-project"),
)

// Create Storager only
sto, err := gcs.NewStorager(
    ps.WithCredential("env"),
    ps.WithName("<bucket_name>"),
)
```

## Configuration

| Pair | Type | Required | Description |
|------|------|----------|-------------|
| `credential` | string | Yes | Auth method: `file:<path>`, `base64:<json>`, or `env` |
| `name` | string | Yes | Bucket name |
| `project_id` | string | Yes (Servicer) | GCP project ID |
| `work_dir` | string | No | Working directory (key prefix). Defaults to `/` |

### Credential formats

| Format | Description |
|--------|-------------|
| `file:/path/to/key.json` | Path to a service account JSON key file |
| `base64:<encoded>` | Base64-encoded service account JSON |
| `env` | Application Default Credentials (reads `GOOGLE_APPLICATION_CREDENTIALS`, well-known files, or metadata server) |

## Supported Interfaces

| Interface | Supported |
|-----------|-----------|
| `Storager` | Yes |
| `Servicer` | Yes |
| `Multiparter` | Yes |
| `Copier` | Yes |

## Storage Classes

| Constant | GCS Class |
|----------|-----------|
| `gcs.StorageClassStandard` | STANDARD |
| `gcs.StorageClassNearLine` | NEARLINE |
| `gcs.StorageClassColdLine` | COLDLINE |
| `gcs.StorageClassArchive` | ARCHIVE |

```go
n, err := sto.Write("path/to/object", r, size,
    gcs.WithStorageClass(gcs.StorageClassNearLine),
)
```

## Server-Side Encryption

Customer-supplied encryption key (CSEK):

```go
n, err := sto.Write("path", r, size,
    gcs.WithServerSideEncryptionCustomerAlgorithm("AES256"),
    gcs.WithServerSideEncryptionCustomerKey([]byte{...32 bytes...}),
)
```

## References

- [Google Cloud Storage documentation](https://cloud.google.com/storage/docs)
- [Application Default Credentials](https://cloud.google.com/docs/authentication/application-default-credentials)
- [Go package reference](https://pkg.go.dev/github.com/rgglez/go-storage/services/gcs/v3)
