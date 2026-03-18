# ipfs — InterPlanetary File System

[![Go Reference](https://pkg.go.dev/badge/github.com/rgglez/go-storage/services/ipfs.svg)](https://pkg.go.dev/github.com/rgglez/go-storage/services/ipfs)
[![Status](https://img.shields.io/badge/status-stable-brightgreen)](https://github.com/rgglez/go-storage)

Backend for [IPFS (InterPlanetary File System)](https://ipfs.io).

## Install

```bash
go get github.com/rgglez/go-storage/services/ipfs
```

## Usage

```go
import (
    "github.com/rgglez/go-storage/services/ipfs"
    ps "github.com/rgglez/go-storage/v5/pairs"
)

// Connect to a local IPFS node
sto, err := ipfs.NewStorager(
    ps.WithEndpoint("http:localhost:5001"),
    ps.WithWorkDir("/ipfs/optional/path/"),
)
```

## Configuration

| Pair | Type | Required | Description |
|------|------|----------|-------------|
| `endpoint` | string | No | IPFS HTTP API address. Defaults to `http://localhost:5001` |
| `work_dir` | string | No | Working directory (IPFS path). Defaults to `/` |

## Supported Interfaces

| Interface | Supported |
|-----------|-----------|
| `Storager` | Yes |

## Notes

- Requires a running IPFS daemon with the HTTP API accessible.
- Objects in IPFS are content-addressed; once written, their CID is immutable.

## References

- [IPFS documentation](https://docs.ipfs.tech)
- [IPFS HTTP API reference](https://docs.ipfs.tech/reference/kubo/rpc/)
- [Go package reference](https://pkg.go.dev/github.com/rgglez/go-storage/services/ipfs)


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
