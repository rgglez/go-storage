# azfile — Azure File Storage

[![Go Reference](https://pkg.go.dev/badge/github.com/rgglez/go-storage/services/azfile.svg)](https://pkg.go.dev/github.com/rgglez/go-storage/services/azfile)
[![Status](https://img.shields.io/badge/status-stable-brightgreen)](https://github.com/rgglez/go-storage)

Backend for [Azure File Storage](https://azure.microsoft.com/en-us/products/storage/files/).

## Install

```bash
go get github.com/rgglez/go-storage/services/azfile
```

## Usage

```go
import (
    "github.com/rgglez/go-storage/services/azfile"
    ps "github.com/rgglez/go-storage/v5/pairs"
)

// Create both Servicer and Storager
srv, sto, err := azfile.New(
    ps.WithCredential("hmac:<account_name>:<account_key>"),
    ps.WithEndpoint("https://<account_name>.file.core.windows.net"),
    ps.WithName("<share_name>"),
    ps.WithWorkDir("/optional/dir/"),
)

// Create Storager only
sto, err := azfile.NewStorager(
    ps.WithCredential("hmac:<account_name>:<account_key>"),
    ps.WithEndpoint("https://<account_name>.file.core.windows.net"),
    ps.WithName("<share_name>"),
)
```

## Connection String

```go
import (
    "github.com/rgglez/go-storage/v5/services"
    _ "github.com/rgglez/go-storage/services/azfile" // register azfile factory
)

store, err := services.NewStoragerFromString(
    "azfile://my-share/docs/?credential=hmac:ACCOUNT_NAME:ACCOUNT_KEY&endpoint=https://ACCOUNT_NAME.file.core.windows.net",
)
```

| Component | Example | Notes |
|-----------|---------|-------|
| scheme | `azfile` | |
| name | `my-share` | File share name — placed right after `://` |
| work_dir | `/docs/` | Optional directory path |
| `credential` | `hmac:NAME:KEY` | Azure storage account name and key |
| `endpoint` | `https://ACCOUNT.file.core.windows.net` | Azure File service URL |

## Configuration

| Pair | Type | Required | Description |
|------|------|----------|-------------|
| `credential` | string | Yes | Authentication. Format: `hmac:<account_name>:<account_key>` |
| `endpoint` | string | Yes | Azure File service URL. Format: `https://<account>.file.core.windows.net` |
| `name` | string | Yes | File share name |
| `work_dir` | string | No | Working directory path. Defaults to `/` |

## Supported Interfaces

| Interface | Supported |
|-----------|-----------|
| `Storager` | Yes |
| `Servicer` | Yes |
| `Direr` | Yes |
| `Appender` | Yes |

## References

- [Azure File Storage documentation](https://azure.microsoft.com/en-us/products/storage/files/)
- [Go package reference](https://pkg.go.dev/github.com/rgglez/go-storage/services/azfile)


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
