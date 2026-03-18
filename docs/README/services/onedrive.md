# onedrive — Microsoft OneDrive

[![Go Reference](https://pkg.go.dev/badge/github.com/rgglez/go-storage/services/onedrive.svg)](https://pkg.go.dev/github.com/rgglez/go-storage/services/onedrive)
[![Status](https://img.shields.io/badge/status-alpha-red)](https://github.com/rgglez/go-storage)

Backend for [Microsoft OneDrive](https://www.microsoft.com/en-ww/microsoft-365/onedrive/online-cloud-storage).

> **Alpha**: still under active development; API may change and some operations may not be implemented.

## Install

```bash
go get github.com/rgglez/go-storage/services/onedrive
```

## Usage

```go
import (
    "github.com/rgglez/go-storage/services/onedrive"
    ps "github.com/rgglez/go-storage/v5/pairs"
)

sto, err := onedrive.NewStorager(
    ps.WithCredential("token:<access_token>"),
    ps.WithWorkDir("/optional/folder/"),
)
```

## Configuration

| Pair | Type | Required | Description |
|------|------|----------|-------------|
| `credential` | string | Yes | OAuth2 access token. Format: `token:<access_token>` |
| `work_dir` | string | No | Working directory path in OneDrive. Defaults to `/` |

## Supported Interfaces

| Interface | Supported |
|-----------|-----------|
| `Storager` | Partial |

## References

- [Microsoft OneDrive API documentation](https://learn.microsoft.com/en-us/onedrive/developer/rest-api/)
- [Go package reference](https://pkg.go.dev/github.com/rgglez/go-storage/services/onedrive)
