# go-storage: Architecture Overview

> **Module**: `github.com/rgglez/go-storage/v5`
> **Go version**: 1.16+
> **Description**: Unified, vendor-neutral storage library for Go — write once, run against any storage backend.

---

## Table of Contents

1. [Project Overview](#1-project-overview)
2. [Directory Structure](#2-directory-structure)
3. [Core Interfaces (`types/`)](#3-core-interfaces-types)
4. [Pair System (`pairs/`)](#4-pair-system-pairs)
5. [Object and Metadata Model](#5-object-and-metadata-model)
6. [Error Handling](#6-error-handling)
7. [Iterator Pattern](#7-iterator-pattern)
8. [Service Registry and Factory](#8-service-registry-and-factory)
9. [Code Generation Pipeline (`definitions/`)](#9-code-generation-pipeline-definitions)
10. [Implementing a Service](#10-implementing-a-service)
11. [S3 Service — Worked Example](#11-s3-service--worked-example)
12. [Utility Packages (`pkg/`)](#12-utility-packages-pkg)
13. [Module and Dependency Structure](#13-module-and-dependency-structure)
14. [Design Principles](#14-design-principles)
15. [Full Architecture Diagram](#15-full-architecture-diagram)

---

## 1. Project Overview

`go-storage` is a **generative, vendor-agnostic storage abstraction layer**. A user writes code once against a unified interface and can swap storage backends (AWS S3, GCS, Azure Blob, local filesystem, FTP, IPFS, etc.) by changing only the initialization string.

It is a maintained fork of the original [BeyondStorage/go-storage](https://github.com/beyondstorage/go-storage) project.

All **27** service backends live in [`services/`](../../services/) as independent Go modules. **Supported service tiers:**

| Tier | Services |
|------|----------|
| Stable (18) | azblob, azfile, bos, cos, dropbox, fs, ftp, gcs, gdrive, ipfs, kodo, memory, minio, obs, ocios, oss, qingstor, s3 |
| Beta (5) | cephfs, hdfs, tar, us3, uss |
| Alpha (4) | onedrive, storj, webdav, zip |

---

## 2. Directory Structure

```
go-storage/
├── types/                      # Core interfaces, object model, generated types
│   ├── operation.generated.go  # Servicer & Storager interfaces (generated)
│   ├── object.go               # Object struct and ObjectMode bitflags
│   ├── object.generated.go     # Object metadata getters/setters (generated)
│   ├── iterator.generated.go   # Iterator types (generated)
│   ├── info.generated.go       # StorageMeta and StorageFeatures (generated)
│   └── error.go                # Core error types
│
├── pairs/                      # Type-safe option/configuration pairs
│   └── generated.go            # WithXxx() pair constructors (generated)
│
├── definitions/                # Code generation specifications and generators
│   ├── operations.go           # Operation definitions (hand-maintained)
│   ├── pairs.go                # Global pair definitions (hand-maintained)
│   ├── infos.go                # Metadata field definitions (hand-maintained)
│   ├── features.go             # Feature flag definitions
│   ├── gen_service.go          # Generator: service-specific code
│   ├── gen_op.go               # Generator: Servicer/Storager interfaces
│   ├── gen_pair.go             # Generator: pair constructors
│   ├── gen_iterator.go         # Generator: iterator types
│   ├── gen_info.go             # Generator: StorageMeta types
│   └── namespace.generated.go  # Namespace utilities (generated)
│
├── services/                   # Storage backend implementations (27 services)
│   ├── s3/                     # AWS S3
│   ├── gcs/                    # Google Cloud Storage
│   ├── azblob/                 # Azure Blob Storage
│   ├── fs/                     # Local filesystem
│   ├── memory/                 # In-memory (testing/ephemeral)
│   ├── ftp/                    # FTP
│   ├── hdfs/                   # Hadoop Distributed File System
│   ├── qingstor/               # QingStor Object Storage
│   ├── ...                     # 19 other services
│   ├── factory.go              # Factory interface & service registry
│   ├── new.go                  # NewStorager / NewServicer / connection string parsing
│   └── error.go                # Common service errors
│
├── pkg/                        # Shared utility packages
│   ├── iowrap/                 # I/O decorators and callbacks
│   ├── httpclient/             # HTTP client helpers
│   ├── fswrap/                 # Filesystem wrappers
│   ├── headers/                # HTTP header utilities
│   └── randbytes/              # Random byte generation
│
├── credential/                 # Credential parsing and representation
├── endpoint/                   # Endpoint/URL parsing
├── internal/cmd/definitions/   # go generate entry point for the whole module
├── docs/                       # RFCs and specification documents
└── go.mod                      # Single module for the core library
```

> **Multi-module note**: Each service under `services/<name>/` has its **own `go.mod`**, allowing independent versioning and avoiding transitive dependencies on cloud SDKs for users who only need one backend.

---

## 3. Core Interfaces (`types/`)

All interfaces are defined (and mostly generated) in `types/`.

### 3.1 `Storager`

The primary interface for object-level operations. All storage backends implement this.

```go
type Storager interface {
    String() string
    Features() StorageFeatures

    // Basic operations
    Read(path string, w io.Writer, pairs ...Pair) (n int64, err error)
    ReadWithContext(ctx context.Context, path string, w io.Writer, pairs ...Pair) (n int64, err error)

    Write(path string, r io.Reader, size int64, pairs ...Pair) (n int64, err error)
    WriteWithContext(ctx context.Context, path string, r io.Reader, size int64, pairs ...Pair) (n int64, err error)

    Stat(path string, pairs ...Pair) (o *Object, err error)
    StatWithContext(ctx context.Context, path string, pairs ...Pair) (o *Object, err error)

    Delete(path string, pairs ...Pair) (err error)
    DeleteWithContext(ctx context.Context, path string, pairs ...Pair) (err error)

    List(path string, pairs ...Pair) (oi *ObjectIterator, err error)
    ListWithContext(ctx context.Context, path string, pairs ...Pair) (oi *ObjectIterator, err error)

    // Extended (capability-gated)
    Copy(src, dst string, pairs ...Pair) error
    Move(src, dst string, pairs ...Pair) error
    CreateDir(path string, pairs ...Pair) (*Object, error)
    CreateLink(path, target string, pairs ...Pair) (*Object, error)

    // Multipart upload
    CreateMultipart(path string, pairs ...Pair) (*Object, error)
    WriteMultipart(o *Object, r io.Reader, size int64, index int, pairs ...Pair) (n int64, part *Part, err error)
    CompleteMultipart(o *Object, parts []*Part, pairs ...Pair) error
    ListMultipart(o *Object, pairs ...Pair) (*PartIterator, error)

    // Append mode
    CreateAppend(path string, pairs ...Pair) (*Object, error)
    WriteAppend(o *Object, r io.Reader, size int64, pairs ...Pair) (n int64, err error)
    CommitAppend(o *Object, pairs ...Pair) error

    // Pre-signed URLs
    QuerySignHTTPRead(path string, expire time.Duration, pairs ...Pair) (req *http.Request, err error)
    QuerySignHTTPWrite(path string, size int64, expire time.Duration, pairs ...Pair) (req *http.Request, err error)
    // ...
}
```

Each operation has both a synchronous and a `WithContext` variant. Services only need to implement the methods they advertise in their `Features()` — unimplemented methods are covered by the embedded `UnimplementedStorager` struct that returns `ErrNotImplemented`.

### 3.2 `Servicer`

Manages containers (buckets, file shares, etc.) within an account.

```go
type Servicer interface {
    String() string
    Features() ServiceFeatures

    Create(name string, pairs ...Pair) (store Storager, err error)
    Delete(name string, pairs ...Pair) error
    Get(name string, pairs ...Pair) (store Storager, err error)
    List(pairs ...Pair) (sti *StoragerIterator, err error)
    // ...WithContext variants for each
}
```

### 3.3 `Factory`

Decouples service initialization from usage. Registered per backend type.

```go
type Factory interface {
    FromString(conn string) error
    FromMap(m map[string]interface{}) error
    WithPairs(ps ...types.Pair) error
    NewServicer() (types.Servicer, error)
    NewStorager() (types.Storager, error)
}
```

### 3.4 `StorageFeatures` and `ServiceFeatures`

Bitfield structs that describe which operations a backend supports. Callers can query capabilities before attempting an operation:

```go
features := store.Features()
if features.Copy {
    err = store.Copy(src, dst)
} else {
    // fallback: read + write
}
```

---

## 4. Pair System (`pairs/`)

Pairs are the mechanism for passing **optional, typed arguments** to any operation. They replace configuration structs while preserving type safety.

### 4.1 Core Type

```go
// types/pair.go
type Pair struct {
    Key   string
    Value interface{}
}
```

### 4.2 Generated Constructors (`pairs/generated.go`)

The code generator produces a typed constructor for every defined pair:

```go
func WithContentType(v string) types.Pair {
    return types.Pair{Key: "content_type", Value: v}
}

func WithOffset(v int64) types.Pair {
    return types.Pair{Key: "offset", Value: v}
}

func WithSize(v int64) types.Pair {
    return types.Pair{Key: "size", Value: v}
}

func WithIoCallback(v func([]byte)) types.Pair {
    return types.Pair{Key: "io_callback", Value: v}
}
```

### 4.3 Usage

```go
n, err := store.Read(
    "data/report.pdf",
    writer,
    pairs.WithOffset(1024),
    pairs.WithSize(4096),
    pairs.WithIoCallback(func(bs []byte) { log.Printf("read %d bytes", len(bs)) }),
)
```

### 4.4 Default Pairs

Services accept `DefaultServicePairs` and `DefaultStoragePairs` at initialization time to set default values for specific operations:

```go
// Apply SSE-KMS encryption to all Write calls by default
store, _ := services.NewStoragerFromString(
    "s3://my-bucket/data",
    pairs.WithDefaultStoragePairs(types.DefaultStoragePairs{
        Write: []types.Pair{
            s3pairs.WithServerSideEncryption("aws:kms"),
            s3pairs.WithServerSideEncryptionAwsKmsKeyId("arn:aws:kms:..."),
        },
    }),
)
```

---

## 5. Object and Metadata Model

### 5.1 `Object` Struct

```go
type Object struct {
    client Storager      // Back-reference for lazy stat() calls
    ID     string        // Service-specific unique identifier
    Path   string        // Logical path
    Mode   ObjectMode    // Type bitflags

    // Metadata fields (private, accessed via getters)
    contentLength int64
    contentMd5    string
    contentType   string
    etag          string
    lastModified  time.Time
    linkTarget    string
    multipartID   string
    // ...

    systemMetadata interface{}  // Service-specific extended metadata

    // Lazy-loading state (atomic)
    done uint32
    mu   sync.Mutex
    bit  uint64          // Bitfield: which fields have been set
}
```

### 5.2 `ObjectMode` Bitflags

```go
const (
    ModeDir    ObjectMode = 1 << iota  // Is a directory / prefix
    ModeRead                           // Is a readable object
    ModeLink                           // Is a symbolic link
    ModePart                           // Is an in-progress multipart upload
    ModeBlock                          // Is a block-structured object
    ModePage                           // Is a page-structured object
    ModeAppend                         // Is an appendable object
)
```

### 5.3 Lazy Metadata Loading

Metadata is loaded **at most once** per `Object`, on first access, using `sync/atomic`:

```go
func (o *Object) GetContentLength() (int64, bool) {
    o.stat()
    if o.bit & bitContentLength == 0 {
        return 0, false
    }
    return o.contentLength, true
}

func (o *Object) stat() {
    if atomic.LoadUint32(&o.done) == 0 {
        o.statSlow()
    }
}

func (o *Object) statSlow() {
    o.mu.Lock()
    defer o.mu.Unlock()
    if o.done == 0 {
        ob, err := o.client.Stat(o.Path)
        if err == nil {
            o.clone(ob)
        }
        atomic.StoreUint32(&o.done, 1)
    }
}
```

This pattern means listing operations can return lightweight `Object` instances; full metadata is only fetched if the caller actually reads it.

### 5.4 Metadata Accessors (Generated)

```go
// Optional getter — returns zero value + false if not set
func (o *Object) GetContentType() (string, bool)
func (o *Object) GetLastModified() (time.Time, bool)

// Must getter — panics if not set (use when you know the field exists)
func (o *Object) MustGetContentLength() int64

// Setter — marks the bitfield
func (o *Object) SetContentType(v string) *Object
```

### 5.5 System Metadata

Service-specific metadata (e.g., S3 SSE details, storage class) is stored in the `systemMetadata interface{}` field. Each service generates typed accessors:

```go
// services/s3/generated.go
type ObjectSystemMetadata struct {
    StorageClass                         string
    ServerSideEncryption                 string
    ServerSideEncryptionAwsKmsKeyID      string
    ServerSideEncryptionBucketKeyEnabled bool
    // ...
}

func GetObjectSystemMetadata(o *types.Object) ObjectSystemMetadata
```

---

## 6. Error Handling

### 6.1 Core Types

```go
// types/error.go
type OperationError struct {
    Op  string
    Err error
}

func (e *OperationError) Error() string
func (e *OperationError) Unwrap() error
func (e *OperationError) IsInternalError() bool
```

### 6.2 Sentinel Errors

```go
var (
    ErrNotImplemented       = errors.New("not implemented")
    ErrServiceNotRegistered = errors.New("service not registered")
    IterateDone             = errors.New("iterate done")  // not a real error
)
```

### 6.3 Service-Specific Errors

Each service maps SDK-native errors to standard `go-storage` error codes in its `errors.go`:

```go
// services/s3/errors.go
func (s *Storage) formatError(op string, err error, args ...interface{}) error {
    if err == nil {
        return nil
    }
    var ae smithy.APIError
    if errors.As(err, &ae) {
        switch ae.ErrorCode() {
        case "NoSuchKey":
            return services.ErrObjectNotExist.WithStack(err)
        case "AccessDenied":
            return services.ErrPermissionDenied.WithStack(err)
        }
    }
    return services.ErrUnexpected.WithStack(err)
}
```

---

## 7. Iterator Pattern

Iterators provide **uniform, lazy, paginated** access to lists of objects, storages, parts, or blocks, regardless of the backend's native pagination API.

### 7.1 Interface

```go
type NextObjectFunc func(ctx context.Context, page *ObjectPage) error

type ObjectIterator struct {
    ctx   context.Context
    next  NextObjectFunc
    index int
    done  bool
    o     ObjectPage
}

type ObjectPage struct {
    Status Continuable
    Data   []*Object
}

func (it *ObjectIterator) Next() (object *Object, err error)
```

`Next()` returns the next object, transparently fetching the next page when the current cache is exhausted. It returns `IterateDone` (a sentinel, not an error) when there are no more items.

### 7.2 `Continuable` Interface

Services that support resumable iteration implement `Continuable`:

```go
type Continuable interface {
    ContinuationToken() string
}
```

This enables clients to save progress and resume listing from a checkpoint (e.g., after a crash).

### 7.3 Iterator Types (all generated)

| Iterator | Page type | Used for |
|----------|-----------|----------|
| `ObjectIterator` | `ObjectPage` | Listing objects/files |
| `StoragerIterator` | `StoragerPage` | Listing buckets/containers |
| `PartIterator` | `PartPage` | Listing multipart upload parts |
| `BlockIterator` | `BlockPage` | Listing committed blocks |

### 7.4 Usage Example

```go
iter, err := store.List("prefix/")
for {
    obj, err := iter.Next()
    if err != nil && errors.Is(err, types.IterateDone) {
        break
    }
    if err != nil {
        return err
    }
    fmt.Println(obj.Path)
}
```

---

## 8. Service Registry and Factory

### 8.1 Registration

At program startup, each service's `init()` function registers itself:

```go
// services/s3/generated.go (generated)
func init() {
    services.RegisterFactory("s3", &Factory{})
}
```

Three maps are maintained in `services/`:

```go
var factoryRegistry    map[string]Factory
var servicerFnMap      map[string]NewServicerFunc
var storagerFnMap      map[string]NewStoragerFunc
```

### 8.2 Initialization Paths

```go
// From a connection string (most common)
store, err := services.NewStoragerFromString(
    "s3://my-bucket/workdir?force_path_style=true",
    pairs.WithCredential(cred),
)

// From a type + map (config file use case)
store, err := services.NewStorager("s3", map[string]interface{}{
    "name":     "my-bucket",
    "work_dir": "/workdir",
})

// Two-step via factory (most control)
f, err := services.NewFactory("s3", pairs.WithCredential(cred))
svc, err := f.NewServicer()
store, err := svc.Get("my-bucket")
```

### 8.3 Connection String Format

```
<type>://<name>/<work_dir>?<key>=<value>&<key>=<value>
```

Examples:
```
s3://my-bucket/data/2024?force_path_style=true&storage_class=STANDARD_IA
gcs://my-bucket/backups
fs:///tmp/storage
azblob://container/path?endpoint=https://account.blob.core.windows.net
```

Query string values are automatically coerced to the declared type of the pair (string, bool, int64, []byte, etc.).

---

## 9. Code Generation Pipeline (`definitions/`)

This is the most distinctive aspect of go-storage's architecture. Almost all boilerplate is **generated from a single source of truth**.

### 9.1 Two-Level Generation

There are two distinct generation stages:

#### Stage 1 — Global generation (for `types/` and `pairs/`)

Triggered by `go generate` in `internal/cmd/definitions/`:

```go
// internal/cmd/definitions/main.go
func main() {
    def.GenerateIterator("../../../types/iterator.generated.go")
    def.GenerateInfo("../../../types/info.generated.go")
    def.GeneratePair("../../../pairs/generated.go")
    def.GenerateOperation("../../../types/operation.generated.go")
    def.GenerateObject("../../../types/object.generated.go")
    def.GenerateNamespace("../../../definitions/namespace.generated.go")
}
```

Input: hand-maintained spec files in `definitions/` (`operations.go`, `pairs.go`, `infos.go`).
Output: `types/*.generated.go`, `pairs/generated.go`.

#### Stage 2 — Per-service generation

Each service has its own mini-generator at `services/<name>/internal/cmd/`:

```go
// services/s3/internal/cmd/main.go
func main() {
    def.GenerateService(Metadata, "generated.go")
}
```

Input: `meta.go` (hand-maintained) — declares the service's capabilities.
Output: `services/s3/generated.go`.

### 9.2 `meta.go` Structure

The heart of per-service generation. The developer declares exactly what the service supports. All 27 services in this repository use this V2 generator — there are no legacy `service.toml` files.

```go
var Metadata = def.Metadata{
    Name: "s3",

    // Optional: extra package import paths needed by custom pair types.
    // Use when a pair's Type.Package refers to a package outside the
    // standard generator imports (context, io, net/http, strings, time,
    // errors, services, types).
    Imports: []string{
        "github.com/rgglez/go-storage/v5/pkg/httpclient",
    },

    Pairs: []def.Pair{ /* service-specific pairs */ },
    Infos: []def.Info{ /* system metadata fields */ },
    Factory: []def.Pair{ /* constructor pairs */ },

    // Service: must be set (even as def.Service{}) for storage-only services
    // to avoid a nil-interface panic in buildFeaturePairs().
    Service: def.Service{
        Features: types.ServiceFeatures{Create: true, Delete: true, Get: true, List: true},
        Create: []def.Pair{def.PairLocation},
        // ...
    },

    Storage: def.Storage{
        Features: types.StorageFeatures{
            VirtualDir: true, Read: true, Write: true, Delete: true,
            List: true, Stat: true, CreateMultipart: true,
            // ...
        },
        Read: []def.Pair{
            def.PairOffset, def.PairSize, def.PairIoCallback,
            pairServerSideEncryptionCustomerAlgorithm,
        },
        Write: []def.Pair{
            def.PairContentType, def.PairContentMD5, pairStorageClass,
            // ...
        },
        // one entry per declared feature
    },
}
```

### 9.3 What `GenerateService` produces

From `meta.go`, the generator outputs `generated.go` containing:

| Generated artifact | Description |
|---|---|
| `const Type` | Service type string (e.g., `"s3"`) |
| `ObjectSystemMetadata` | Typed struct for service-specific object metadata |
| `StorageSystemMetadata` | Typed struct for service-specific storage metadata |
| `pairStorage<Op>` structs | Parsed pair sets for each operation |
| `parse<Op>Pairs()` functions | Validates & extracts pairs for each operation |
| Public `Read(...)` wrappers | Validates features, calls private `read(...)` |
| `Factory` struct | Implements `services.Factory` interface |
| `init()` | Registers the factory in the global registry |

### 9.4 Manual vs. Generated Split

```
services/s3/
│
├── internal/cmd/meta.go      HAND-WRITTEN: what the service can do
│
├── generated.go              GENERATED: boilerplate (DO NOT EDIT)
│   ├─ Public method wrappers (Read, Write, Stat, ...)
│   ├─ Pair parsing structs
│   ├─ Feature validation
│   ├─ SystemMetadata types
│   └─ Factory + init()
│
├── storage.go                HAND-WRITTEN: real implementation
│   └─ Private methods: read(), write(), stat(), delete(), list(), ...
│
├── service.go                HAND-WRITTEN: real implementation
│   └─ Private methods: create(), delete(), get(), list()
│
├── utils.go                  HAND-WRITTEN: SDK helpers, formatters
├── iterator.go               HAND-WRITTEN: pagination state machines
└── errors.go                 HAND-WRITTEN: error mapping
```

The generated public wrapper calls the private method. For example:

```go
// generated.go  (DO NOT EDIT)
func (s *Storage) Read(path string, w io.Writer, pairs ...Pair) (int64, error) {
    ctx := context.Background()
    return s.ReadWithContext(ctx, path, w, pairs...)
}

func (s *Storage) ReadWithContext(ctx context.Context, path string, w io.Writer, pairs ...Pair) (n int64, err error) {
    if !s.features.Read {
        return 0, services.NewOperationNotImplementedError("read")
    }
    opt, err := s.parseStorageReadPairs(pairs...)
    if err != nil {
        return 0, err
    }
    return s.read(ctx, path, w, opt)
}

// storage.go  (hand-written)
func (s *Storage) read(ctx context.Context, path string, w io.Writer, opt pairStorageRead) (n int64, err error) {
    // Actual AWS SDK call here
    input := &s3sdk.GetObjectInput{ Bucket: &s.name, Key: aws.String(rp) }
    if opt.HasOffset { input.Range = ... }
    // ...
}
```

---

## 10. Implementing a Service

To add a new backend, a developer follows these steps:

1. **Create the module**:
   ```
   services/mybackend/
   ├── go.mod
   ├── Makefile
   └── internal/cmd/
       ├── main.go    # calls def.GenerateService(Metadata, "generated.go")
       └── meta.go    # declares capabilities
   ```

2. **Write `meta.go`**: Declare features, pairs per operation, and any custom metadata fields.

3. **Run `go generate`**: This produces `generated.go`.

4. **Implement manual code**:
   - `service.go` — service-level operations (`create`, `delete`, `get`, `list`)
   - `storage.go` — object operations (`read`, `write`, `stat`, `delete`, `list`, ...)
   - `utils.go` — SDK helpers, credential/endpoint handling, request formatters
   - `iterator.go` — pagination state machine (`nextObjectPage`, `nextStoragePage`, etc.)
   - `errors.go` — map backend errors to `go-storage` error codes

5. **Write tests** in `services/mybackend/tests/`.

Only the private lowercase methods need to be written by hand. All public API surface, pair parsing, and feature gating is generated.

---

## 11. S3 Service — Worked Example

### Initialization flow

```
services.NewStoragerFromString("s3://bucket/data")
  │
  ├─ Parse connection string → type="s3", name="bucket", work_dir="/data"
  │
  ├─ Lookup factoryRegistry["s3"] → *s3.Factory
  │
  ├─ factory.FromString(conn)   → sets pairs
  ├─ factory.WithPairs(extra...) → merges pairs
  │
  └─ factory.NewStorager()
       ├─ Build aws.Config (credentials, region, endpoint options)
       ├─ Create s3.Client from SDK
       └─ Return &Storage{service: client, name: "bucket", workDir: "/data"}
```

### Read operation flow

```
store.Read("report.pdf", writer, pairs.WithOffset(1024))
  │
  ├─ generated.go: ReadWithContext()
  │    ├─ Check features.Read == true
  │    ├─ parseStorageReadPairs() → pairStorageRead{Offset: 1024, HasOffset: true}
  │    └─ call s.read(ctx, "report.pdf", writer, opt)
  │
  └─ storage.go: read()
       ├─ rp = s.getAbsPath("report.pdf")  → "/data/report.pdf"
       ├─ Build GetObjectInput{Bucket, Key, Range: "bytes=1024-"}
       ├─ s.service.GetObject(ctx, input)
       ├─ if opt.HasIoCallback: wrap reader with iowrap.CallbackReadCloser
       └─ io.Copy(writer, body)
```

### Iterator flow (object listing)

```
store.List("prefix/")
  │
  └─ generated.go → storage.go: list()
       ├─ Create objectPageStatus{} with prefix, delimiter, continuation token
       └─ Return types.NewObjectIterator(ctx, s.nextObjectPage, status)
            │
            iter.Next() calls:
            └─ iterator.go: nextObjectPage(ctx, page)
                 ├─ s.service.ListObjectsV2(ctx, input)
                 ├─ Append *types.Object for each result to page.Data
                 ├─ Update status.ContinuationToken
                 └─ Return IterateDone when IsTruncated == false
```

---

## 12. Utility Packages (`pkg/`)

| Package | Purpose |
|---------|---------|
| `pkg/iowrap` | Wraps `io.Reader`/`io.Writer` with progress callbacks, size limits, and transformation |
| `pkg/httpclient` | Shared HTTP client with timeout and retry configuration |
| `pkg/fswrap` | Wraps `os.File` and filesystem operations for the `fs` service |
| `pkg/headers` | Parse and build HTTP headers (Content-Type, Content-MD5, Range, etc.) |
| `pkg/randbytes` | Generate random byte slices for testing |

### `iowrap` in detail

A key utility used by almost every service:

```go
// Wrap a reader with a progress callback
r = iowrap.CallbackReadCloser(r, func(bs []byte) {
    atomic.AddInt64(&bytesRead, int64(len(bs)))
})

// Wrap a writer with a size limit
w = iowrap.LimitWriter(w, maxBytes)
```

---

## 13. Module and Dependency Structure

### Core module (`go.mod`)

```
module github.com/rgglez/go-storage/v5

require:
  github.com/Xuanwo/gg             v0.3.0   // Code generation DSL
  github.com/Xuanwo/templateutils  v0.2.0   // Template parsing for gen_service.go
  github.com/golang/mock           v1.6.0   // Interface mock generation
  github.com/google/uuid           v1.6.0
  github.com/sirupsen/logrus       v1.9.3
  github.com/stretchr/testify      v1.9.0
```

### Service modules (e.g., `services/s3/go.mod`)

```
module github.com/rgglez/go-storage/services/s3/v3

require:
  github.com/rgglez/go-storage/v5  v5.x.x
  github.com/aws/aws-sdk-go-v2     v1.x.x
  // ...

replace:
  github.com/rgglez/go-storage/v5 => ../../
```

Each service module uses a `replace` directive in development so it builds against the local core library. The service module paths follow the pattern `github.com/rgglez/go-storage/services/<name>` (with a major version suffix when applicable, e.g. `/v3`, `/v4`).

This means: importing `go-storage/v5` in your project does **not** pull in any cloud SDK. You only pay the dependency cost for the services you actually use.

---

## 14. Design Principles

| Principle | How it's applied |
|-----------|-----------------|
| **Write once, run anywhere** | Single `Storager` interface works against all backends |
| **Zero-cost abstraction** | Unused backends add no transitive dependencies (separate go.mod per service) |
| **Explicit capability model** | `Features()` returns a struct; callers can branch on available operations |
| **Type-safe configuration** | Generated `WithXxx()` pair constructors, no `interface{}` in user code |
| **Minimal manual boilerplate** | Only business logic is hand-written; all API surface is generated |
| **Lazy metadata** | `Object` fields are fetched on-demand and cached atomically |
| **Uniform pagination** | All listing APIs return an `*ObjectIterator` with a consistent `Next()` pattern |
| **Context-first for I/O** | Every remote operation has a `WithContext` variant |
| **Error categorization** | Backends map native errors to standard codes; callers inspect with `errors.Is` |
| **Resumable iteration** | Iterators expose a `ContinuationToken()` for checkpoint/resume |

---

## 15. Full Architecture Diagram

```
┌─────────────────────────────────────────────────────────────────────┐
│                           User Application                          │
│                                                                     │
│  store, _ := services.NewStoragerFromString("s3://bucket/data")     │
│  store.Write("file.txt", reader, size)                              │
│  iter, _ := store.List("prefix/")                                   │
└───────────────────────────────┬─────────────────────────────────────┘
                                │
                    ┌───────────▼────────────┐
                    │   services/new.go      │
                    │  Connection string     │
                    │  parsing & dispatch    │
                    └───────────┬────────────┘
                                │  lookup by type
                    ┌───────────▼────────────┐
                    │  services/factory.go   │
                    │  Global Factory        │
                    │  Registry (map)        │
                    └───────────┬────────────┘
                                │  factory.NewStorager()
          ┌─────────────────────┼──────────────────────────┐
          │                     │                          │
┌─────────▼──────┐   ┌──────────▼────────┐    ┌───────────▼────────┐
│  services/s3   │   │  services/gcs     │    │  services/fs       │
│                │   │                   │    │                    │
│ generated.go   │   │  generated.go     │    │  generated.go      │
│ (DO NOT EDIT)  │   │  (DO NOT EDIT)    │    │  (DO NOT EDIT)     │
│  ┌──────────┐  │   │   ┌──────────┐   │    │   ┌──────────┐    │
│  │ Wrappers │  │   │   │ Wrappers │   │    │   │ Wrappers │    │
│  │ Pair     │  │   │   │ Pair     │   │    │   │ Pair     │    │
│  │ parsing  │  │   │   │ parsing  │   │    │   │ parsing  │    │
│  │ Feature  │  │   │   │ Feature  │   │    │   │ Feature  │    │
│  │ gating   │  │   │   │ gating   │   │    │   │ gating   │    │
│  └────┬─────┘  │   │   └────┬─────┘   │    │   └────┬─────┘    │
│       │ calls  │   │        │ calls   │    │        │ calls    │
│  ┌────▼─────┐  │   │   ┌────▼─────┐  │    │   ┌────▼─────┐   │
│  │storage.go│  │   │   │storage.go│  │    │   │storage.go│   │
│  │service.go│  │   │   │service.go│  │    │   │service.go│   │
│  │utils.go  │  │   │   │utils.go  │  │    │   │utils.go  │   │
│  │iter.go   │  │   │   │iter.go   │  │    │   │iter.go   │   │
│  └────┬─────┘  │   │   └────┬─────┘  │    │   └────┬─────┘   │
└───────┼────────┘   └────────┼─────── ┘    └────────┼─────────┘
        │                     │                       │
   AWS SDK v2            GCS SDK                  os package
        │
        ▼
   Amazon S3 API


┌─────────────────── Code Generation Pipeline ────────────────────────┐
│                                                                     │
│  definitions/           internal/cmd/definitions/    types/         │
│  ┌───────────────┐      ┌────────────────────┐      ┌───────────┐  │
│  │ operations.go │─────▶│                    │─────▶│operation  │  │
│  │ pairs.go      │      │  GenerateOperation │      │.generated │  │
│  │ infos.go      │─────▶│  GeneratePair      │─────▶│.go        │  │
│  │ features.go   │      │  GenerateIterator  │      │           │  │
│  └───────────────┘      │  GenerateInfo      │─────▶│iterator   │  │
│                         │  GenerateObject    │      │.generated │  │
│                         │  GenerateNamespace │      │.go etc.   │  │
│                         └────────────────────┘      └───────────┘  │
│                                                                     │
│  services/s3/internal/cmd/                                          │
│  ┌───────────────┐      ┌────────────────────┐                     │
│  │  meta.go      │─────▶│  GenerateService   │─────▶ generated.go  │
│  │  (manual)     │      │  (reads meta +     │      (per service)  │
│  └───────────────┘      │   existing .go)    │                     │
│                         └────────────────────┘                     │
└─────────────────────────────────────────────────────────────────────┘
```

---

*This document was generated from analysis of the source code at `github.com/rgglez/go-storage/v5`.*


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
