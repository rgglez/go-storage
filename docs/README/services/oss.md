# oss — Aliyun Object Storage Service

[![Go Reference](https://pkg.go.dev/badge/github.com/rgglez/go-storage/services/oss/v3.svg)](https://pkg.go.dev/github.com/rgglez/go-storage/services/oss/v3)
[![Status](https://img.shields.io/badge/status-stable-brightgreen)](https://github.com/rgglez/go-storage)

Backend for [Alibaba Cloud Object Storage Service (OSS)](https://www.aliyun.com/product/oss).

## Install

```bash
go get github.com/rgglez/go-storage/services/oss/v3
```

## Usage

```go
import (
    "github.com/rgglez/go-storage/services/oss/v3"
    ps "github.com/rgglez/go-storage/v5/pairs"
)

// Create both Servicer and Storager
srv, sto, err := oss.New(
    ps.WithCredential("hmac:<access_key_id>:<access_key_secret>"),
    ps.WithEndpoint("https://oss-cn-hangzhou.aliyuncs.com"),
    ps.WithName("<bucket_name>"),
    ps.WithWorkDir("/optional/prefix/"),
)

// Create Storager only
sto, err := oss.NewStorager(
    ps.WithCredential("hmac:<access_key_id>:<access_key_secret>"),
    ps.WithEndpoint("https://oss-cn-hangzhou.aliyuncs.com"),
    ps.WithName("<bucket_name>"),
)
```

## Configuration

| Pair | Type | Required | Description |
|------|------|----------|-------------|
| `credential` | string | Yes | Authentication method. See [Credential formats](#credential-formats) below |
| `endpoint` | string | Yes | OSS endpoint URL (e.g. `https://oss-cn-hangzhou.aliyuncs.com`) |
| `name` | string | Yes | Bucket name |
| `work_dir` | string | No | Working directory (key prefix). Defaults to `/` |

### Credential formats

| Format | Description |
|--------|-------------|
| `hmac:<access_key_id>:<access_key_secret>` | Static long-term credentials (RAM user or root account) |
| `env` | STS temporary credentials read from environment variables (see below) |

#### Static credentials (`hmac`)

```go
sto, err := oss.NewStorager(
    ps.WithCredential("hmac:<AccessKeyId>:<AccessKeySecret>"),
    ps.WithEndpoint("https://oss-cn-hangzhou.aliyuncs.com"),
    ps.WithName("<bucket_name>"),
)
```

#### STS temporary credentials (`env`)

When using Alibaba Cloud STS (Security Token Service), pass `"env"` as the credential value. The backend reads three environment variables:

| Variable | Description |
|----------|-------------|
| `ALIBABA_CLOUD_ACCESS_KEY_ID` | Temporary Access Key ID issued by STS |
| `ALIBABA_CLOUD_ACCESS_KEY_SECRET` | Temporary Access Key Secret issued by STS |
| `ALIBABA_CLOUD_SECURITY_TOKEN` | Security Token (required for STS) |

All three variables must be set; the backend returns an error if any is missing.

```bash
export ALIBABA_CLOUD_ACCESS_KEY_ID="STS.xxxxxxxx"
export ALIBABA_CLOUD_ACCESS_KEY_SECRET="yyyyyyyy"
export ALIBABA_CLOUD_SECURITY_TOKEN="CAISxxxxxxxxxxxxxxxx..."
```

```go
sto, err := oss.NewStorager(
    ps.WithCredential("env"),
    ps.WithEndpoint("https://oss-cn-hangzhou.aliyuncs.com"),
    ps.WithName("<bucket_name>"),
)
```

The security token is forwarded to the OSS SDK via `oss.SecurityToken(token)`. Refresh the environment variables and recreate the storager before the STS token expires.

## Supported Interfaces

| Interface | Supported |
|-----------|-----------|
| `Storager` | Yes |
| `Servicer` | Yes |
| `Appender` | Yes |
| `Multiparter` | Yes |
| `Direr` | Yes |
| `Linker` | Yes |

## References

- [Aliyun OSS documentation](https://www.alibabacloud.com/help/en/oss)
- [OSS endpoints list](https://www.alibabacloud.com/help/en/oss/user-guide/regions-and-endpoints)
- [Go package reference](https://pkg.go.dev/github.com/rgglez/go-storage/services/oss/v3)


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
