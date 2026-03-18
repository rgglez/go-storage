# memory — In-Memory Storage

[![Go Reference](https://pkg.go.dev/badge/github.com/rgglez/go-storage/services/memory.svg)](https://pkg.go.dev/github.com/rgglez/go-storage/services/memory)
[![Status](https://img.shields.io/badge/status-stable-brightgreen)](https://github.com/rgglez/go-storage)

In-memory storage backend. Useful for testing and ephemeral data.

## Install

```bash
go get github.com/rgglez/go-storage/services/memory
```

## Usage

```go
import (
    "github.com/rgglez/go-storage/services/memory"
    ps "github.com/rgglez/go-storage/v5/pairs"
)

// Create an in-memory Storager
sto, err := memory.NewStorager(
    ps.WithWorkDir("/optional/prefix/"),
)
```

## Configuration

| Pair | Type | Required | Description |
|------|------|----------|-------------|
| `work_dir` | string | No | Virtual working directory. Defaults to `/` |

No credential or endpoint is required. All data lives in process memory.

## Supported Interfaces

| Interface | Supported |
|-----------|-----------|
| `Storager` | Yes |

## Notes

- Data is stored as a tree of in-memory objects; parent directories are created implicitly.
- All data is **lost when the process exits**.
- The backend is **goroutine-safe** (mutex-protected tree).

## References

- [Go package reference](https://pkg.go.dev/github.com/rgglez/go-storage/services/memory)
