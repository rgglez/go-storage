# ocios — Oracle Cloud Infrastructure Object Storage

[![Go Reference](https://pkg.go.dev/badge/github.com/rgglez/go-storage/services/ocios.svg)](https://pkg.go.dev/github.com/rgglez/go-storage/services/ocios)
[![Status](https://img.shields.io/badge/status-stable-brightgreen)](https://github.com/rgglez/go-storage)

Backend for [Oracle Cloud Infrastructure (OCI) Object Storage](https://www.oracle.com/cloud/storage/object-storage/).

## Install

```bash
go get github.com/rgglez/go-storage/services/ocios
```

## Usage

```go
import (
    "github.com/rgglez/go-storage/services/ocios"
    ps "github.com/rgglez/go-storage/v5/pairs"
)

// Create both Servicer and Storager
srv, sto, err := ocios.New(
    ps.WithCredential("file:/home/user/.oci/config"),
    ps.WithName("<bucket_name>"),
    ps.WithLocation("<region>"),   // e.g. "us-ashburn-1"
    ps.WithWorkDir("/optional/prefix/"),
)
```

## Connection String

```go
import (
    "github.com/rgglez/go-storage/v5/services"
    _ "github.com/rgglez/go-storage/services/ocios" // register ocios factory
)

store, err := services.NewStoragerFromString(
    "ocios://my-bucket/data/?credential=file:/home/user/.oci/config&location=us-ashburn-1",
)
```

| Component | Example | Notes |
|-----------|---------|-------|
| scheme | `ocios` | |
| name | `my-bucket` | Bucket name — placed right after `://` |
| work_dir | `/data/` | Optional key prefix |
| `credential` | `file:/home/user/.oci/config` | Path to OCI configuration file |
| `location` | `us-ashburn-1` | Required OCI region identifier |

## Configuration

| Pair | Type | Required | Description |
|------|------|----------|-------------|
| `credential` | string | Yes | OCI config file path. Format: `file:<path>` |
| `name` | string | Yes | Bucket name |
| `location` | string | Yes | OCI region identifier (e.g. `us-ashburn-1`, `eu-frankfurt-1`) |
| `work_dir` | string | No | Working directory (key prefix). Defaults to `/` |

## Supported Interfaces

| Interface | Supported |
|-----------|-----------|
| `Storager` | Yes |
| `Servicer` | Yes |
| `Multiparter` | Yes |

## Notes

> **Alpha/stub status**: some operations may not be fully implemented.

## References

- [OCI Object Storage documentation](https://docs.oracle.com/en-us/iaas/Content/Object/home.htm)
- [Go package reference](https://pkg.go.dev/github.com/rgglez/go-storage/services/ocios)


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
