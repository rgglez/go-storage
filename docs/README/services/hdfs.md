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
