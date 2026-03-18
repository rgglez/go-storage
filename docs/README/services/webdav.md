# webdav — WebDAV

[![Go Reference](https://pkg.go.dev/badge/github.com/rgglez/go-storage/services/webdav.svg)](https://pkg.go.dev/github.com/rgglez/go-storage/services/webdav)
[![Status](https://img.shields.io/badge/status-alpha-red)](https://github.com/rgglez/go-storage)

Backend for [WebDAV](http://www.webdav.org/) protocol servers (e.g. Nextcloud, ownCloud, Apache with WebDAV module).

> **Alpha**: still under active development; API may change and some operations may not be implemented.

## Install

```bash
go get github.com/rgglez/go-storage/services/webdav
```

## Usage

```go
import (
    "github.com/rgglez/go-storage/services/webdav"
    ps "github.com/rgglez/go-storage/v5/pairs"
)

sto, err := webdav.NewStorager(
    ps.WithEndpoint("https://dav.example.com/remote.php/dav/files/user/"),
    ps.WithCredential("basic:<username>:<password>"),
    ps.WithWorkDir("/optional/folder/"),
)
```

## Configuration

| Pair | Type | Required | Description |
|------|------|----------|-------------|
| `endpoint` | string | Yes | WebDAV server base URL |
| `credential` | string | Yes | Authentication. Format: `basic:<user>:<password>` |
| `work_dir` | string | No | Working directory path. Defaults to `/` |

## Supported Interfaces

| Interface | Supported |
|-----------|-----------|
| `Storager` | Partial |

## References

- [WebDAV RFC 4918](https://tools.ietf.org/html/rfc4918)
- [Go package reference](https://pkg.go.dev/github.com/rgglez/go-storage/services/webdav)
