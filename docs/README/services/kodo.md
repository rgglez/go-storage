# kodo — Qiniu Kodo

[![Go Reference](https://pkg.go.dev/badge/github.com/rgglez/go-storage/services/kodo/v3.svg)](https://pkg.go.dev/github.com/rgglez/go-storage/services/kodo/v3)
[![Status](https://img.shields.io/badge/status-stable-brightgreen)](https://github.com/rgglez/go-storage)

Backend for [Qiniu Kodo Object Storage](https://www.qiniu.com/en/products/kodo).

## Install

```bash
go get github.com/rgglez/go-storage/services/kodo/v3
```

## Usage

```go
import (
    "github.com/rgglez/go-storage/services/kodo/v3"
    ps "github.com/rgglez/go-storage/v5/pairs"
)

// Create both Servicer and Storager
srv, sto, err := kodo.New(
    ps.WithCredential("hmac:<access_key>:<secret_key>"),
    ps.WithName("<bucket_name>"),
    ps.WithLocation("<zone>"),     // e.g. "z0" (East China), "z1" (North China)
    ps.WithWorkDir("/optional/prefix/"),
)

// Create Storager only
sto, err := kodo.NewStorager(
    ps.WithCredential("hmac:<access_key>:<secret_key>"),
    ps.WithName("<bucket_name>"),
    ps.WithLocation("z0"),
)
```

## Configuration

| Pair | Type | Required | Description |
|------|------|----------|-------------|
| `credential` | string | Yes | Authentication. Format: `hmac:<access_key>:<secret_key>` |
| `name` | string | Yes | Bucket name |
| `location` | string | Yes | Zone/region (e.g. `z0`, `z1`, `z2`, `na0`, `as0`) |
| `work_dir` | string | No | Working directory (key prefix). Defaults to `/` |

## Supported Interfaces

| Interface | Supported |
|-----------|-----------|
| `Storager` | Yes |
| `Servicer` | Yes |
| `Multiparter` | Yes |

## References

- [Qiniu Kodo documentation](https://developer.qiniu.com/kodo)
- [Go package reference](https://pkg.go.dev/github.com/rgglez/go-storage/services/kodo/v3)


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
