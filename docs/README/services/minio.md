# minio — MinIO

[![Go Reference](https://pkg.go.dev/badge/github.com/rgglez/go-storage/services/minio.svg)](https://pkg.go.dev/github.com/rgglez/go-storage/services/minio)
[![Status](https://img.shields.io/badge/status-stable-brightgreen)](https://github.com/rgglez/go-storage)

Backend for [MinIO](https://min.io) object storage.

## Install

```bash
go get github.com/rgglez/go-storage/services/minio
```

## Usage

```go
import (
    "github.com/rgglez/go-storage/services/minio"
    ps "github.com/rgglez/go-storage/v5/pairs"
)

// Create both Servicer and Storager
srv, sto, err := minio.New(
    ps.WithCredential("hmac:<access_key>:<secret_key>"),
    ps.WithEndpoint("http:localhost:9000"),
    ps.WithName("<bucket_name>"),
    ps.WithWorkDir("/optional/prefix/"),
)

// Create Storager only
sto, err := minio.NewStorager(
    ps.WithCredential("hmac:<access_key>:<secret_key>"),
    ps.WithEndpoint("http:localhost:9000"),
    ps.WithName("<bucket_name>"),
)
```

## Configuration

| Pair | Type | Required | Description |
|------|------|----------|-------------|
| `credential` | string | Yes | Authentication. Format: `hmac:<access_key>:<secret_key>` |
| `endpoint` | string | Yes | MinIO server address. Format: `http:<host>:<port>` or `https:<host>:<port>` |
| `name` | string | Yes | Bucket name |
| `location` | string | No | Region/location (e.g. `us-east-1`). Defaults to `us-east-1` |
| `work_dir` | string | No | Working directory (key prefix). Defaults to `/` |

## Supported Interfaces

| Interface | Supported |
|-----------|-----------|
| `Storager` | Yes |
| `Servicer` | Yes |
| `Multiparter` | Yes |
| `Copier` | Yes |

## References

- [MinIO documentation](https://min.io/docs/minio/linux/index.html)
- [Go package reference](https://pkg.go.dev/github.com/rgglez/go-storage/services/minio)


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
