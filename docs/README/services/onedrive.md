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
