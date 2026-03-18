# ipfs — InterPlanetary File System

[![Go Reference](https://pkg.go.dev/badge/github.com/rgglez/go-storage/services/ipfs.svg)](https://pkg.go.dev/github.com/rgglez/go-storage/services/ipfs)
[![Status](https://img.shields.io/badge/status-stable-brightgreen)](https://github.com/rgglez/go-storage)

Backend for [IPFS (InterPlanetary File System)](https://ipfs.io).

## Install

```bash
go get github.com/rgglez/go-storage/services/ipfs
```

## Usage

```go
import (
    "github.com/rgglez/go-storage/services/ipfs"
    ps "github.com/rgglez/go-storage/v5/pairs"
)

// Connect to a local IPFS node
sto, err := ipfs.NewStorager(
    ps.WithEndpoint("http:localhost:5001"),
    ps.WithWorkDir("/ipfs/optional/path/"),
)
```

## Configuration

| Pair | Type | Required | Description |
|------|------|----------|-------------|
| `endpoint` | string | No | IPFS HTTP API address. Defaults to `http://localhost:5001` |
| `work_dir` | string | No | Working directory (IPFS path). Defaults to `/` |

## Supported Interfaces

| Interface | Supported |
|-----------|-----------|
| `Storager` | Yes |

## Notes

- Requires a running IPFS daemon with the HTTP API accessible.
- Objects in IPFS are content-addressed; once written, their CID is immutable.

## References

- [IPFS documentation](https://docs.ipfs.tech)
- [IPFS HTTP API reference](https://docs.ipfs.tech/reference/kubo/rpc/)
- [Go package reference](https://pkg.go.dev/github.com/rgglez/go-storage/services/ipfs)
