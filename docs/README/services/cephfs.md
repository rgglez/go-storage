# cephfs — Ceph Filesystem

[![Go Reference](https://pkg.go.dev/badge/github.com/rgglez/go-storage/services/cephfs.svg)](https://pkg.go.dev/github.com/rgglez/go-storage/services/cephfs)
[![Status](https://img.shields.io/badge/status-beta-yellow)](https://github.com/rgglez/go-storage)

Backend for [Ceph Filesystem (CephFS)](https://docs.ceph.com/en/latest/cephfs/).

> **Beta**: implemented but not fully integration-tested.

## Install

```bash
go get github.com/rgglez/go-storage/services/cephfs
```

## Usage

```go
import (
    "github.com/rgglez/go-storage/services/cephfs"
    ps "github.com/rgglez/go-storage/v5/pairs"
)

sto, err := cephfs.NewStorager(
    ps.WithEndpoint("tcp:<monitor_host>:<port>"),
    ps.WithCredential("hmac:<client_id>:<client_key>"),
    ps.WithWorkDir("/optional/path/"),
)
```

## Configuration

| Pair | Type | Required | Description |
|------|------|----------|-------------|
| `endpoint` | string | No | Ceph monitor address |
| `credential` | string | No | Client ID and key |
| `work_dir` | string | No | Working directory in CephFS. Defaults to `/` |

## Supported Interfaces

| Interface | Supported |
|-----------|-----------|
| `Storager` | Yes |

## References

- [Ceph documentation](https://docs.ceph.com/en/latest/cephfs/)
- [Go package reference](https://pkg.go.dev/github.com/rgglez/go-storage/services/cephfs)


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
