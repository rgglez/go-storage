# fs — Local File System

[![Go Reference](https://pkg.go.dev/badge/github.com/rgglez/go-storage/services/fs/v4.svg)](https://pkg.go.dev/github.com/rgglez/go-storage/services/fs/v4)
[![Status](https://img.shields.io/badge/status-stable-brightgreen)](https://github.com/rgglez/go-storage)

Backend for the local file system. Supports cross-platform operation (Linux, macOS, Windows).

## Install

```bash
go get github.com/rgglez/go-storage/services/fs/v4
```

## Usage

```go
import (
    "github.com/rgglez/go-storage/services/fs/v4"
    ps "github.com/rgglez/go-storage/v5/pairs"
)

// Create Storager (no Servicer for fs)
sto, err := fs.NewStorager(
    ps.WithWorkDir("/path/to/work/dir"),
)
```

## Connection String

```go
import (
    "github.com/rgglez/go-storage/v5/services"
    _ "github.com/rgglez/go-storage/services/fs/v4" // register fs factory
)

// Root of the filesystem
store, err := services.NewStoragerFromString("fs:///")

// Specific working directory
store, err := services.NewStoragerFromString("fs:///tmp/myapp")
```

| Component | Example | Notes |
|-----------|---------|-------|
| scheme | `fs` | |
| work_dir | `/tmp/myapp` | Absolute path used as the root for all operations; defaults to `/` |

No credential, endpoint, or name required.

## Configuration

| Pair | Type | Required | Description |
|------|------|----------|-------------|
| `work_dir` | string | No | Base directory for all operations. Created automatically if absent. Defaults to `/` |

No credential or endpoint required — the local filesystem is used directly.

## Special Paths

The `fs` backend recognizes these special paths regardless of `work_dir`:

| Path | Description |
|------|-------------|
| `/dev/stdin` | Standard input |
| `/dev/stdout` | Standard output |
| `/dev/stderr` | Standard error |

## Supported Interfaces

| Interface | Supported |
|-----------|-----------|
| `Storager` | Yes |
| `Copier` | Yes |
| `Mover` | Yes |
| `Appender` | Yes |
| `Direr` | Yes |
| `Linker` | Yes |
| `Fetcher` | Yes |

## Example: Copy a local file

```go
sto, _ := fs.NewStorager(ps.WithWorkDir("/tmp"))

// Read
r, _, _ := sto.Read("input.txt", nil)

// Write
n, _ := sto.Write("output.txt", r, size)

// Copy
_ = sto.(types.Copier).Copy("input.txt", "copy.txt")

// Create directory
_, _ = sto.(types.Direr).CreateDir("subdir")
```

## Notes

- The `work_dir` is evaluated through symlinks on creation.
- `Stat` uses `Lstat` internally (does **not** follow symlinks).
- Files are created with mode `0664`; directories with `0755`.

## References

- [Go package reference](https://pkg.go.dev/github.com/rgglez/go-storage/services/fs/v4)


## License

Copyright 2024 go-storage authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
