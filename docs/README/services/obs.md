# obs — Huawei Object Storage Service

[![Go Reference](https://pkg.go.dev/badge/github.com/rgglez/go-storage/services/obs/v2.svg)](https://pkg.go.dev/github.com/rgglez/go-storage/services/obs/v2)
[![Status](https://img.shields.io/badge/status-stable-brightgreen)](https://github.com/rgglez/go-storage)

Backend for [Huawei Cloud Object Storage Service (OBS)](https://www.huaweicloud.com/intl/en-us/product/obs.html).

## Install

```bash
go get github.com/rgglez/go-storage/services/obs/v2
```

## Usage

```go
import (
    "github.com/rgglez/go-storage/services/obs/v2"
    ps "github.com/rgglez/go-storage/v5/pairs"
)

// Create both Servicer and Storager
srv, sto, err := obs.New(
    ps.WithCredential("hmac:<access_key>:<secret_key>"),
    ps.WithEndpoint("https://obs.cn-north-4.myhuaweicloud.com"),
    ps.WithName("<bucket_name>"),
    ps.WithWorkDir("/optional/prefix/"),
)

// Create Storager only
sto, err := obs.NewStorager(
    ps.WithCredential("hmac:<access_key>:<secret_key>"),
    ps.WithEndpoint("https://obs.cn-north-4.myhuaweicloud.com"),
    ps.WithName("<bucket_name>"),
)
```

## Connection String

```go
import (
    "github.com/rgglez/go-storage/v5/services"
    _ "github.com/rgglez/go-storage/services/obs/v2" // register obs factory
)

store, err := services.NewStoragerFromString(
    "obs://my-bucket/data/?credential=hmac:ACCESS_KEY:SECRET_KEY&endpoint=https://obs.cn-north-4.myhuaweicloud.com",
)
```

| Component | Example | Notes |
|-----------|---------|-------|
| scheme | `obs` | |
| name | `my-bucket` | Bucket name — placed right after `://` |
| work_dir | `/data/` | Optional key prefix |
| `credential` | `hmac:AK:SK` | Huawei Cloud access key and secret key |
| `endpoint` | `https://obs.cn-north-4.myhuaweicloud.com` | OBS regional endpoint URL |

## Configuration

| Pair | Type | Required | Description |
|------|------|----------|-------------|
| `credential` | string | Yes | Authentication. Format: `hmac:<access_key>:<secret_key>` |
| `endpoint` | string | Yes | OBS endpoint URL (e.g. `https://obs.cn-north-4.myhuaweicloud.com`) |
| `name` | string | Yes | Bucket name |
| `work_dir` | string | No | Working directory (key prefix). Defaults to `/` |

## Supported Interfaces

| Interface | Supported |
|-----------|-----------|
| `Storager` | Yes |
| `Servicer` | Yes |
| `Multiparter` | Yes |

## References

- [Huawei OBS documentation](https://support.huaweicloud.com/intl/en-us/obs/index.html)
- [Go package reference](https://pkg.go.dev/github.com/rgglez/go-storage/services/obs/v2)


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
