# gdrive — Google Drive

[![Go Reference](https://pkg.go.dev/badge/github.com/rgglez/go-storage/services/gdrive.svg)](https://pkg.go.dev/github.com/rgglez/go-storage/services/gdrive)
[![Status](https://img.shields.io/badge/status-stable-brightgreen)](https://github.com/rgglez/go-storage)

Backend for [Google Drive](https://www.google.com/drive/).

## Install

```bash
go get github.com/rgglez/go-storage/services/gdrive
```

## Usage

```go
import (
    "github.com/rgglez/go-storage/services/gdrive"
    ps "github.com/rgglez/go-storage/v5/pairs"
)

// Create Storager
sto, err := gdrive.NewStorager(
    ps.WithCredential("file:/path/to/credentials.json"),
    ps.WithWorkDir("/optional/folder/"),
)
```

## Configuration

| Pair | Type | Required | Description |
|------|------|----------|-------------|
| `credential` | string | Yes | Path to OAuth2 credentials JSON. Format: `file:<path>` |
| `work_dir` | string | No | Working directory path in Drive. Defaults to `/` |

## Supported Interfaces

| Interface | Supported |
|-----------|-----------|
| `Storager` | Yes |
| `Copier` | Yes |
| `Mover` | Yes |
| `Direr` | Yes |

## References

- [Google Drive API documentation](https://developers.google.com/drive/api)
- [Go package reference](https://pkg.go.dev/github.com/rgglez/go-storage/services/gdrive)


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
