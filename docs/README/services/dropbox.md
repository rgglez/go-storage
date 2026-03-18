# dropbox — Dropbox

[![Go Reference](https://pkg.go.dev/badge/github.com/rgglez/go-storage/services/dropbox/v3.svg)](https://pkg.go.dev/github.com/rgglez/go-storage/services/dropbox/v3)
[![Status](https://img.shields.io/badge/status-stable-brightgreen)](https://github.com/rgglez/go-storage)

Backend for [Dropbox](https://www.dropbox.com/).

## Install

```bash
go get github.com/rgglez/go-storage/services/dropbox/v3
```

## Usage

```go
import (
    "github.com/rgglez/go-storage/services/dropbox/v3"
    ps "github.com/rgglez/go-storage/v5/pairs"
)

// Create Storager
sto, err := dropbox.NewStorager(
    ps.WithCredential("token:<access_token>"),
    ps.WithWorkDir("/optional/folder/"),
)
```

## Configuration

| Pair | Type | Required | Description |
|------|------|----------|-------------|
| `credential` | string | Yes | OAuth2 access token. Format: `token:<access_token>` |
| `work_dir` | string | No | Working directory path inside Dropbox. Defaults to `/` |

## Supported Interfaces

| Interface | Supported |
|-----------|-----------|
| `Storager` | Yes |
| `Copier` | Yes |
| `Mover` | Yes |

## Getting a Dropbox Access Token

1. Go to [Dropbox App Console](https://www.dropbox.com/developers/apps).
2. Create a new app with **Full Dropbox** access.
3. Under **Settings → OAuth 2**, generate an access token.

## References

- [Dropbox API v2 documentation](https://www.dropbox.com/developers/documentation/http/documentation)
- [Go package reference](https://pkg.go.dev/github.com/rgglez/go-storage/services/dropbox/v3)
