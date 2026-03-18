# us3 — UCloud Object Storage

[![Go Reference](https://pkg.go.dev/badge/github.com/rgglez/go-storage/services/us3.svg)](https://pkg.go.dev/github.com/rgglez/go-storage/services/us3)
[![Status](https://img.shields.io/badge/status-beta-yellow)](https://github.com/rgglez/go-storage)

Backend for [UCloud US3 (formerly UFile) Object Storage](https://www.ucloud.cn/site/product/ufile.html).

> **Beta**: implemented but not fully integration-tested.

## Install

```bash
go get github.com/rgglez/go-storage/services/us3
```

## Usage

```go
import (
    "github.com/rgglez/go-storage/services/us3"
    ps "github.com/rgglez/go-storage/v5/pairs"
)

srv, sto, err := us3.New(
    ps.WithCredential("hmac:<public_key>:<private_key>"),
    ps.WithEndpoint("https://cn-bj.ufileos.com"),
    ps.WithName("<bucket_name>"),
    ps.WithWorkDir("/optional/prefix/"),
)
```

## Configuration

| Pair | Type | Required | Description |
|------|------|----------|-------------|
| `credential` | string | Yes | Authentication. Format: `hmac:<public_key>:<private_key>` |
| `endpoint` | string | Yes | US3 endpoint URL |
| `name` | string | Yes | Bucket name |
| `work_dir` | string | No | Working directory (key prefix). Defaults to `/` |

## Supported Interfaces

| Interface | Supported |
|-----------|-----------|
| `Storager` | Yes |
| `Servicer` | Yes |
| `Multiparter` | Yes |

## References

- [UCloud US3 documentation](https://docs.ucloud.cn/ufile/README)
- [Go package reference](https://pkg.go.dev/github.com/rgglez/go-storage/services/us3)


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
