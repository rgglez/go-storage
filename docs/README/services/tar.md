# tar — TAR Archive Files

[![Go Reference](https://pkg.go.dev/badge/github.com/rgglez/go-storage/services/tar.svg)](https://pkg.go.dev/github.com/rgglez/go-storage/services/tar)
[![Status](https://img.shields.io/badge/status-beta-yellow)](https://github.com/rgglez/go-storage)

Read-access backend for TAR archive files.

> **Beta**: implemented but not fully integration-tested.

## Install

```bash
go get github.com/rgglez/go-storage/services/tar
```

## Usage

```go
import (
    "github.com/rgglez/go-storage/services/tar"
    ps "github.com/rgglez/go-storage/v5/pairs"
)

sto, err := tar.NewStorager(
    ps.WithName("/path/to/archive.tar"),
    ps.WithWorkDir("/optional/prefix/"),
)
```

## Connection String

```go
import (
    "github.com/rgglez/go-storage/v5/services"
    _ "github.com/rgglez/go-storage/services/tar" // register tar factory
)

store, err := services.NewStoragerFromString(
    "tar:///path/to/archive.tar/subdir/",
)
```

| Component | Example | Notes |
|-----------|---------|-------|
| scheme | `tar` | |
| name | `/path/to/archive.tar` | Absolute path to the `.tar` file — placed right after `://` |
| work_dir | `/subdir/` | Optional path prefix inside the archive |

No credential or endpoint required.

## Configuration

| Pair | Type | Required | Description |
|------|------|----------|-------------|
| `name` | string | Yes | Path to the `.tar` file |
| `work_dir` | string | No | Path prefix inside the archive. Defaults to `/` |

## Supported Interfaces

| Interface | Supported |
|-----------|-----------|
| `Storager` | Yes (read-only) |

## References

- [Go package reference](https://pkg.go.dev/github.com/rgglez/go-storage/services/tar)


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
