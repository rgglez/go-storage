# cos — Tencent Cloud Object Storage

[![Go Reference](https://pkg.go.dev/badge/github.com/rgglez/go-storage/services/cos/v3.svg)](https://pkg.go.dev/github.com/rgglez/go-storage/services/cos/v3)
[![Status](https://img.shields.io/badge/status-stable-brightgreen)](https://github.com/rgglez/go-storage)

Backend for [Tencent Cloud Object Storage (COS)](https://intl.cloud.tencent.com/product/cos).

## Install

```bash
go get github.com/rgglez/go-storage/services/cos/v3
```

## Usage

```go
import (
    "github.com/rgglez/go-storage/services/cos/v3"
    ps "github.com/rgglez/go-storage/v5/pairs"
)

// Create both Servicer and Storager
srv, sto, err := cos.New(
    ps.WithCredential("hmac:<secret_id>:<secret_key>"),
    ps.WithName("<bucket_name>"),
    ps.WithLocation("<region>"),   // e.g. "ap-beijing"
    ps.WithWorkDir("/optional/prefix/"),
)

// Create Storager only
sto, err := cos.NewStorager(
    ps.WithCredential("hmac:<secret_id>:<secret_key>"),
    ps.WithName("<bucket_name>"),
    ps.WithLocation("ap-beijing"),
)
```

## Configuration

| Pair | Type | Required | Description |
|------|------|----------|-------------|
| `credential` | string | Yes | Authentication. Format: `hmac:<secret_id>:<secret_key>` |
| `name` | string | Yes | Bucket name (including appid suffix, e.g. `my-bucket-1250000000`) |
| `location` | string | Yes | Region (e.g. `ap-beijing`, `ap-shanghai`) |
| `work_dir` | string | No | Working directory (key prefix). Defaults to `/` |

## Supported Interfaces

| Interface | Supported |
|-----------|-----------|
| `Storager` | Yes |
| `Servicer` | Yes |
| `Multiparter` | Yes |
| `Appender` | Yes |
| `Copier` | Yes |

## References

- [Tencent COS documentation](https://intl.cloud.tencent.com/document/product/436)
- [Go package reference](https://pkg.go.dev/github.com/rgglez/go-storage/services/cos/v3)


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
