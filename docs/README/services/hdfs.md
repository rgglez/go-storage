# hdfs — Hadoop Distributed File System

[![Go Reference](https://pkg.go.dev/badge/github.com/rgglez/go-storage/services/hdfs.svg)](https://pkg.go.dev/github.com/rgglez/go-storage/services/hdfs)
[![Status](https://img.shields.io/badge/status-beta-yellow)](https://github.com/rgglez/go-storage)

Backend for [Hadoop Distributed File System (HDFS)](https://hadoop.apache.org/docs/stable/hadoop-project-dist/hadoop-hdfs/HdfsDesign.html).

> **Beta**: implemented but not fully integration-tested.

## Install

```bash
go get github.com/rgglez/go-storage/services/hdfs
```

## Usage

```go
import (
    "github.com/rgglez/go-storage/services/hdfs"
    ps "github.com/rgglez/go-storage/v5/pairs"
)

sto, err := hdfs.NewStorager(
    ps.WithEndpoint("tcp:<namenode_host>:9000"),
    ps.WithWorkDir("/user/myuser/"),
)
```

## Connection String

```go
import (
    "github.com/rgglez/go-storage/v5/services"
    _ "github.com/rgglez/go-storage/services/hdfs" // register hdfs factory
)

store, err := services.NewStoragerFromString(
    "hdfs:///user/myuser/?endpoint=tcp:namenode.example.com:9000",
)
```

| Component | Example | Notes |
|-----------|---------|-------|
| scheme | `hdfs` | |
| work_dir | `/user/myuser/` | Optional working directory in HDFS |
| `endpoint` | `tcp:host:9000` | HDFS NameNode address |

## Configuration

| Pair | Type | Required | Description |
|------|------|----------|-------------|
| `endpoint` | string | Yes | HDFS NameNode address. Format: `tcp:<host>:<port>` |
| `work_dir` | string | No | Working directory in HDFS. Defaults to `/` |

## Supported Interfaces

| Interface | Supported |
|-----------|-----------|
| `Storager` | Yes |
| `Direr` | Yes |

## References

- [HDFS documentation](https://hadoop.apache.org/docs/stable/hadoop-project-dist/hadoop-hdfs/HdfsDesign.html)
- [Go package reference](https://pkg.go.dev/github.com/rgglez/go-storage/services/hdfs)


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
