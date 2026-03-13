# go-storage

[![Go dev](https://pkg.go.dev/badge/github.com/rgglez/go-storage/v5)](https://pkg.go.dev/gitub.com/v5)
[![License](https://img.shields.io/badge/license-apache%20v2-blue.svg)](https://github.com/rgglez/go-storage/blob/master/LICENSE)
[![Build Test](https://github.com/rgglez/go-storage/actions/workflows/build-test.yml/badge.svg)](https://github.com/rgglez/go-storage/actions/workflows/build-test.yml)
[![Cross Build](https://github.com/rgglez/go-storage/actions/workflows/cross-build.yml/badge.svg)](https://github.com/rgglez/go-storage/actions/workflows/cross-build.yml)
[![Unit Test](https://github.com/rgglez/go-storage/actions/workflows/unit-test.yml/badge.svg)](https://github.com/rgglez/go-storage/actions/workflows/unit-test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/rgglez/go-storage)](https://goreportcard.com/report/github.com/rgglez/go-storage)
![GitHub stars](https://img.shields.io/github/stars/rgglez/go-storage?style=social)
![GitHub forks](https://img.shields.io/github/forks/rgglez/go-storage?style=social)

A **vendor-neutral** storage library for Golang.

## Table of Contents

- [About this fork](#about-this-fork)
- [Vision](#vision)
- [Goals](#goals)
- [Documentation](#documentation)
  - [`docs/README/` - Project documents](#docsreadme--project-documents)
  - [`docs/rfcs/` - Request for Comments](#docsrfcs--request-for-comments)
  - [`docs/spec/` - Specifications](#docsspec--specifications)
- [Makefile](#makefile)
- [Features](#features)
  - [Widely native services support](#widely-native-services-support)
  - [Complete and easily extensible interface](#complete-and-easily-extensible-interface)
  - [Comprehensive metadata](#comprehensive-metadata)
  - [Strong Typing Everywhere](#strong-typing-everywhere)
  - [Server-Side Encrypt](#server-side-encrypt)
- [License](#license)


## About this fork

The [original project](https://github.com/beyondstorage/go-storage) seems to be [dead](https://github.com/beyondstorage/go-storage/issues/1382) and [broken](https://github.com/beyondstorage/go-storage/issues/1263). My changes are listed in the [CHANGELOG.md](docs/CHANGELOG.md) file

## Vision

Write once, run on every storage service.

## Goals

- Vendor agnostic
- Production ready
- High performance

## Documentation

### `docs/README/` — Project documents

| File | Description |
|------|-------------|
| [ARCHITECTURE.md](docs/README/ARCHITECTURE.md) | Detailed architecture overview: package structure, interfaces, code generation pipeline, pair system, iterator pattern, and how to implement a new service. **Recommended reading before contributing.** |
| [CONTRIBUTING.md](docs/README/CONTRIBUTING.md) | Contribution guidelines: how to report bugs, submit patches, implement new services, or propose API changes. |
| [TAGS_AND_SUBMODULES.md](docs/README/TAGS_AND_SUBMODULES.md) | How to create and push Go module tags for sub-modules in this monorepo. |
| [CHANGELOG.md](docs/README/CHANGELOG.md) | History of changes and releases. |
| [CODE_OF_CONDUCT.md](docs/README/CODE_OF_CONDUCT.md) | Community code of conduct. |

### `docs/rfcs/` — Request for Comments

Design decision records covering the evolution of the library. Each RFC is numbered and describes the motivation, proposal, and rationale behind a change. Notable examples:

| RFC | Topic |
|-----|-------|
| [RFC-1](docs/rfcs/1-unify-storager-behavior.md) | Unify Storager behavior |
| [RFC-11](docs/rfcs/11-error-handling.md) | Error handling design |
| [RFC-25](docs/rfcs/25-object-mode.md) | Object mode bitflags |
| [RFC-48](docs/rfcs/48-service-registry.md) | Service registry pattern |
| [RFC-700](docs/rfcs/700-config-features-and-defaultpairs-via-connection-string.md) | Connection string configuration |
| [RFC-840](docs/rfcs/840-convert-to-monorepo.md) | Conversion to monorepo |
| [RFC-970](docs/rfcs/970-service-factory.md) | Service factory design |

The full list of RFCs is in [`docs/rfcs/`](docs/rfcs/).

### `docs/spec/` — Specifications

Behavioral specifications that services must conform to:

| File | Description |
|------|-------------|
| [spec/1-error-handling.md](docs/spec/1-error-handling.md) | Error handling specification |
| [spec/2-proposal.md](docs/spec/2-proposal.md) | Proposal process specification |

## Makefile

The root [Makefile](Makefile) provides the following targets:

| Target | Description |
|--------|-------------|
| `make help` | Lists available targets with a short description. |
| `make check` | Runs static analysis (alias for `vet`). |
| `make vet` | Runs `go vet ./...` to report suspicious constructs. |
| `make format` | Formats all Go source files in-place with `gofmt`. |
| `make generate` | Runs `go generate ./...` to regenerate all code-generated files, then formats them. |
| `make build` | Full build for the current module: runs `tidy`, `generate`, `format`, `check`, and `go build ./...`. |
| `make build-all` | Iterates over every `go.mod` in the monorepo and runs `make build` in each sub-module directory. |
| `make test` | Runs the unit test suite with race detection and produces `coverage.txt` and `coverage.html` reports. |
| `make test-all` | Iterates over every `go.mod` in the monorepo and runs `make test` in each sub-module directory. |
| `make tidy` | Runs `go mod tidy` and `go mod verify` for the current module. |
| `make tidy-all` | Iterates over every `go.mod` in the monorepo and runs `make tidy` in each sub-module directory. |
| `make clean` | Deletes all `generated.go` files across the repository. |

## Features

### Widely native services support

All services live in this monorepo under [`services/`](services/). Each service is an independent Go module so you only pay the dependency cost for the backends you actually use.

**18** stable services:

- [azblob](services/azblob/): [Azure Blob storage](https://docs.microsoft.com/en-us/azure/storage/blobs/)
- [azfile](services/azfile/): [Azure File Storage](https://azure.microsoft.com/en-us/products/storage/files/)
- [bos](services/bos/): [Baidu Object Storage](https://cloud.baidu.com/product/bos.html)
- [cos](services/cos/): [Tencent Cloud Object Storage](https://cloud.tencent.com/product/cos)
- [dropbox](services/dropbox/): [Dropbox](https://www.dropbox.com)
- [fs](services/fs/): Local file system
- [ftp](services/ftp/): FTP
- [gcs](services/gcs/): [Google Cloud Storage](https://cloud.google.com/storage/)
- [gdrive](services/gdrive/): [Google Drive](https://www.google.com/drive/)
- [ipfs](services/ipfs/): [InterPlanetary File System](https://ipfs.io)
- [kodo](services/kodo/): [Qiniu Kodo](https://www.qiniu.com/products/kodo)
- [memory](services/memory/): In-memory storage (testing / ephemeral data)
- [minio](services/minio/): [MinIO](https://min.io)
- [obs](services/obs/): [Huawei Object Storage Service](https://www.huaweicloud.com/product/obs.html)
- [ocios](services/ocios/): [Oracle Cloud Infrastructure Object Storage](https://www.oracle.com/cloud/storage/object-storage/)
- [oss](services/oss/): [Aliyun Object Storage](https://www.aliyun.com/product/oss)
- [qingstor](services/qingstor/): [QingStor Object Storage](https://www.qingcloud.com/products/qingstor/)
- [s3](services/s3/): [Amazon S3](https://aws.amazon.com/s3/)

**5** beta services (implemented but not fully integration-tested):

- [cephfs](services/cephfs/): [Ceph Filesystem](https://docs.ceph.com/en/latest/cephfs/)
- [hdfs](services/hdfs/): [Hadoop Distributed File System](https://hadoop.apache.org/docs/stable/hadoop-project-dist/hadoop-hdfs/HdfsDesign.html)
- [tar](services/tar/): tar archive files
- [us3](services/us3/): [UCloud Object Storage](https://www.ucloud.cn/site/product/ufile.html)
- [uss](services/uss/): [UPYUN Storage Service](https://www.upyun.com/products/file-storage)

**4** alpha services (still under development):

- [onedrive](services/onedrive/): [Microsoft OneDrive](https://www.microsoft.com/en-ww/microsoft-365/onedrive/online-cloud-storage)
- [storj](services/storj/): [Storj](https://www.storj.io/)
- [webdav](services/webdav/): [WebDAV](http://www.webdav.org/)
- [zip](services/zip/): zip archive files

More service ideas could be found at [Service Integration Tracking](https://github.com/rgglez/go-storage/issues/536).

### Complete and easily extensible interface

Basic operations

- Metadata: get `Storager` metadata
```go
meta := store.Metadata()
_ := meta.GetWorkDir() // Get object WorkDir
_, ok := meta.GetWriteSizeMaximum() // Get the maximum size for write operation
```
- Read: read `Object` content
```go
// Read 2048 byte at the offset 1024 into the io.Writer.
n, err := store.Read("path", w, pairs.WithOffset(1024), pairs.WithSize(2048))
```
- Write: write content into `Object`
```go
// Write 2048 byte from io.Reader
n, err := store.Write("path", r, 2048)
```
- Stat: get `Object` metadata or check existences
```go
o, err := store.Stat("path")
if errors.Is(err, services.ErrObjectNotExist) {
	// object is not exist
}
length, ok := o.GetContentLength() // get the object content length.
```
- Delete: delete an `Object`
```go
err := store.Delete("path") // Delete the object "path"
```
- List: list `Object` in given prefix or dir
```go
it, err := store.List("path")
for {
	o, err := it.Next()
	if err != nil && errors.Is(err, types.IterateDone) {
        // the list is over
    }
    length, ok := o.GetContentLength() // get the object content length.
}
```

Extended operations

- Copy: copy a `Object` inside storager
```go
err := store.(Copier).Copy(src, dst) // Copy an object from src to dst.
```
- Move: move a `Object` inside storager
```go
err := store.(Mover).Move(src, dst) // Move an object from src to dst.
```
- Reach: generate a public accessible url to an `Object`
```go
url, err := store.(Reacher).Reach("path") // Generate an url to the object.
```
- Dir: Dir `Object` support
```go
o, err := store.(Direr).CreateDir("path") // Create a dir object.
```

Large file manipulation

- Multipart: allow doing multipart uploads
```go
ms := store.(Multiparter)

// Create a multipart object.
o, err := ms.CreateMultipart("path")
// Write 1024 bytes from io.Reader into a multipart at index 1
n, part, err := ms.WriteMultipart(o, r, 1024, 1)
// Complete a multipart object.
err := ms.CompleteMultipart(o, []*Part{part})
```
- Append: allow appending to an object
```go
as := store.(Appender)

// Create an appendable object.
o, err := as.CreateAppend("path")
// Write 1024 bytes from io.Reader.
n, err := as.WriteAppend(o, r, 1024)
// Commit an append object.
err = as.CommitAppend(o)
```
- Block: allow combining an object with block ids
```go
bs := store.(Blocker)

// Create a block object.
o, err := bs.CreateBlock("path")
// Write 1024 bytes from io.Reader with block id "id-abc"
n, err := bs.WriteBlock(o, r, 1024, "id-abc")
// Combine block via block ids.
err := bs.CombineBlock(o, []string{"id-abc"})
```
- Page: allow doing random writes
```go
ps := store.(Pager)

// Create a page object.
o, err := ps.CreatePage("path")
// Write 1024 bytes from io.Reader at offset 2048
n, err := ps.WritePage(o, r, 1024, 2048)
```

### Comprehensive metadata

Global object metadata

- `id`: unique key in service
- `name`: relative path towards service's work dir
- `mode`: object mode can be a combination of `read`, `dir`, `part` and [more](https://github.com/rgglez/go-storage/blob/master/types/object.go#L11)
- `etag`: entity tag as defined in [rfc2616](https://tools.ietf.org/html/rfc2616#section-14.19)
- `content-length`: object's content size.
- `content-md5`: md5 digest as defined in [rfc2616](https://tools.ietf.org/html/rfc2616#section-14.15)
- `content-type`: media type as defined in [rfc2616](https://tools.ietf.org/html/rfc2616#section-14.17)
- `last-modified`: object's last updated time.

System object metadata

Service system object metadata like `storage-class` and so on.

```go
o, err := store.Stat("path")

// Get service system metadata via API provides by go-service-s3.
om := s3.GetObjectSystemMetadata(o)
_ = om.StorageClass // this object's storage class
_ = om.ServerSideEncryptionCustomerAlgorithm // this object's sse algorithm
```

### Strong Typing Everywhere

Self maintained codegen ([`definitions/`](definitions/)) helps to generate all our APIs, pairs and metadata.

Generated pairs which can be used as API optional arguments.

```go
func WithContentMd5(v string) Pair {
    return Pair{
        Key:   "content_md5",
        Value: v,
    }
}
```

Generated object metadata which can be used to get content md5 from object.

```go
func (o *Object) GetContentMd5() (string, bool) {
    o.stat()

    if o.bit&objectIndexContentMd5 != 0 {
        return o.contentMd5, true
    }

    return "", false
}
```

### Server-Side Encrypt

Server-Side Encrypt supports via system pair and system metadata, and we can use [Default Pairs](https://beyondstorage.io/docs/go-storage/pairs#default-pairs) to simplify the job.

```go

func NewS3SseC(key []byte) (types.Storager, error) {
    defaultPairs := s3.DefaultStoragePairs{
        Write: []types.Pair{
            // Required, must be AES256
            s3.WithServerSideEncryptionCustomerAlgorithm(s3.ServerSideEncryptionAes256),
            // Required, your AES-256 key, a 32-byte binary value
            s3.WithServerSideEncryptionCustomerKey(key),
        },
        // Now you have to provide customer key to read encrypted data
        Read: []types.Pair{
            // Required, must be AES256
            s3.WithServerSideEncryptionCustomerAlgorithm(s3.ServerSideEncryptionAes256),
            // Required, your AES-256 key, a 32-byte binary value
            s3.WithServerSideEncryptionCustomerKey(key),
        }}

    return s3.NewStorager(..., s3.WithDefaultStoragePairs(defaultPairs))
}
```

## License

This project is licensed under the [Apache License 2.0](LICENSE).

```
Copyright The go-storage Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
