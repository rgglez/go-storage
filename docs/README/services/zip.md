# zip — ZIP Archive Files

[![Go Reference](https://pkg.go.dev/badge/github.com/rgglez/go-storage/services/zip.svg)](https://pkg.go.dev/github.com/rgglez/go-storage/services/zip)
[![Status](https://img.shields.io/badge/status-alpha-red)](https://github.com/rgglez/go-storage)

Read-access backend for ZIP archive files.

> **Alpha**: still under active development; API may change and some operations may not be implemented.

## Install

```bash
go get github.com/rgglez/go-storage/services/zip
```

## Usage

```go
import (
    "github.com/rgglez/go-storage/services/zip"
    ps "github.com/rgglez/go-storage/v5/pairs"
)

sto, err := zip.NewStorager(
    ps.WithName("/path/to/archive.zip"),
    ps.WithWorkDir("/optional/prefix/"),
)
```

## Configuration

| Pair | Type | Required | Description |
|------|------|----------|-------------|
| `name` | string | Yes | Path to the `.zip` file |
| `work_dir` | string | No | Path prefix inside the archive. Defaults to `/` |

## Supported Interfaces

| Interface | Supported |
|-----------|-----------|
| `Storager` | Partial (read-only) |

## References

- [ZIP file format specification](https://pkware.cachefly.net/webdocs/casestudies/APPNOTE.TXT)
- [Go package reference](https://pkg.go.dev/github.com/rgglez/go-storage/services/zip)


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
