# ftp — FTP

[![Go Reference](https://pkg.go.dev/badge/github.com/rgglez/go-storage/services/ftp.svg)](https://pkg.go.dev/github.com/rgglez/go-storage/services/ftp)
[![Status](https://img.shields.io/badge/status-stable-brightgreen)](https://github.com/rgglez/go-storage)

Backend for FTP servers (File Transfer Protocol).

## Install

```bash
go get github.com/rgglez/go-storage/services/ftp
```

## Usage

```go
import (
    "github.com/rgglez/go-storage/services/ftp"
    ps "github.com/rgglez/go-storage/v5/pairs"
)

// Anonymous connection (defaults)
sto, err := ftp.NewStorager(
    ps.WithEndpoint("tcp:ftp.example.com:21"),
)

// Authenticated connection
sto, err := ftp.NewStorager(
    ps.WithEndpoint("tcp:ftp.example.com:21"),
    ps.WithCredential("basic:<username>:<password>"),
    ps.WithWorkDir("/upload/"),
)
```

## Configuration

| Pair | Type | Required | Description |
|------|------|----------|-------------|
| `endpoint` | string | No | FTP server address. Format: `tcp:<host>:<port>`. Defaults to `localhost:21` |
| `credential` | string | No | Authentication. Format: `basic:<user>:<password>`. Defaults to anonymous login |
| `work_dir` | string | No | Working directory on the FTP server. Defaults to `/` |

## Supported Interfaces

| Interface | Supported |
|-----------|-----------|
| `Storager` | Yes |
| `Direr` | Yes |

## Notes

- Connects immediately on `NewStorager`; returns an error if the server is unreachable.
- Connection timeout is 5 seconds.
- Supports file, folder, and symlink object modes.

## References

- [FTP RFC 959](https://tools.ietf.org/html/rfc959)
- [Go package reference](https://pkg.go.dev/github.com/rgglez/go-storage/services/ftp)


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
