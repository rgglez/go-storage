# Connection Strings

go-storage supports initializing any storage backend from a single connection string, following a URL-like syntax. This is the recommended approach when the backend type and its parameters are determined at runtime (e.g., from a configuration file or environment variable).

## `NewStoragerFromString`

```go
func NewStoragerFromString(conn string, ps ...types.Pair) (types.Storager, error)
```

**Package**: `github.com/rgglez/go-storage/v5/services`

Creates a [`types.Storager`](../../types/) from a connection string. An optional list of [`types.Pair`](../../pairs/) values can be provided to override or supplement parameters encoded in the string.

**How it works**:

1. Parses the scheme (e.g., `s3`, `fs`, `oss`) and the rest of the connection string.
2. Looks up the registered `Factory` for that scheme in the service registry.
3. Calls `Factory.FromString()` with the remaining connection string, then `Factory.WithPairs()` with any extra pairs.
4. Calls `Factory.NewStorager()` to instantiate the backend.
5. If no factory is registered for the scheme, falls back to the legacy pair-based initializer.

The service package must be imported (even if unused) so its `init()` function runs and registers the factory:

```go
import _ "github.com/rgglez/go-storage/services/s3/v3"
```

### Companion functions

| Function | Description |
|----------|-------------|
| `NewStoragerFromString(conn, ps...)` | Create a `Storager` from a connection string |
| `NewServicerFromString(conn, ps...)` | Create a `Servicer` from a connection string (bucket/container management) |
| `NewStorager(ty, ps...)` | Create a `Storager` by service type name and pairs |
| `NewServicer(ty, ps...)` | Create a `Servicer` by service type name and pairs |
| `NewFactoryFromString(conn, ps...)` | Low-level: parse a connection string and return the populated `Factory` |

---

## Connection string format

```
<type>://<name><work_dir>[?key=value&...]
```

| Component | Required | Description |
|-----------|----------|-------------|
| `type` | Yes | Service identifier: `s3`, `fs`, `oss`, `gcs`, `cos`, … |
| `name` | Depends | Storage container (bucket name, etc.). Must **not** contain `/`. |
| `work_dir` | No | Path prefix inside the container. Must start with `/`. |
| `?key=value` | No | Additional pairs encoded as URL query parameters. |

### Supported parameter types in query strings

| Go type | Example |
|---------|---------|
| `string` | `location=us-east-1` |
| `bool` | `force_path_style=true` |
| `int` | `multipart_threshold=10` |
| `int64` | `offset=1048576` |
| `uint64` | `size=2097152` |
| `[]byte` | base64-encoded value |
| `time.Duration` | `timeout=30s` |

### Special query parameter prefixes

| Prefix | Meaning | Example |
|--------|---------|---------|
| `enable_` | Activate a feature flag | `enable_virtual_dir` |
| `default_` | Set a default pair for an operation | `default_storage_class=STANDARD_IA` |

---

## Examples

### Amazon S3

```go
package main

import (
    "fmt"
    "log"

    "github.com/rgglez/go-storage/v5/services"
    _ "github.com/rgglez/go-storage/services/s3/v3" // register s3 factory
)

func main() {
    // Minimal — credentials and region embedded in the connection string.
    store, err := services.NewStoragerFromString(
        "s3://my-bucket/data/?credential=hmac:ACCESS:SECRET&location=us-east-1",
    )
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(store.Metadata().GetName())
}
```

```go
// S3-compatible endpoint (e.g. MinIO, DigitalOcean Spaces).
store, err := services.NewStoragerFromString(
    "s3://my-bucket/uploads/?credential=hmac:ACCESS:SECRET&endpoint=https://nyc3.digitaloceanspaces.com&location=nyc3&force_path_style=true",
)
```

```go
// With extra pairs supplied at call time (pairs override connection string values).
import (
    ps "github.com/rgglez/go-storage/v5/pairs"
)

store, err := services.NewStoragerFromString(
    "s3://my-bucket/",
    ps.WithCredential("hmac:ACCESS:SECRET"),
    ps.WithLocation("eu-west-1"),
)
```

**Required parameters for S3**

| Parameter | Connection string key | Description |
|-----------|-----------------------|-------------|
| Credential | `credential` | `hmac:<access_key>:<secret_key>` |
| Bucket | `name` (path component) | Bucket name — place it immediately after `://` |
| Region | `location` | AWS region, e.g. `us-east-1` |

**Optional parameters for S3**

| Parameter | Connection string key | Description |
|-----------|-----------------------|-------------|
| Endpoint | `endpoint` | Custom S3-compatible endpoint URL |
| Work dir | `work_dir` (path component) | Key prefix within the bucket |
| Path style | `force_path_style` | `true` for MinIO / localhost deployments |
| Acceleration | `use_accelerate` | `true` to enable S3 Transfer Acceleration |

---

### Local file system (`fs`)

The `fs` backend needs no credentials or endpoint — only an optional `work_dir`.

```go
package main

import (
    "log"

    "github.com/rgglez/go-storage/v5/services"
    _ "github.com/rgglez/go-storage/services/fs/v4" // register fs factory
)

func main() {
    // Root of the filesystem.
    store, err := services.NewStoragerFromString("fs:///")
    if err != nil {
        log.Fatal(err)
    }
    _ = store

    // Specific working directory.
    store, err = services.NewStoragerFromString("fs:///tmp/myapp")
    if err != nil {
        log.Fatal(err)
    }
    _ = store
}
```

```go
// Read a file via the storager.
n, err := store.Read("report.pdf", w)
```

**Parameters for `fs`**

| Parameter | Connection string key | Description |
|-----------|-----------------------|-------------|
| Work dir | path component after `fs://` | Absolute path used as the root for all operations. Defaults to `/`. |

> The `fs` scheme has no `name` component — use the path directly: `fs:///path/to/dir`.

---

### Alibaba Cloud OSS (`oss`)

```go
package main

import (
    "log"

    "github.com/rgglez/go-storage/v5/services"
    _ "github.com/rgglez/go-storage/services/oss/v3" // register oss factory
)

func main() {
    // Static credentials.
    store, err := services.NewStoragerFromString(
        "oss://my-bucket/data/?credential=hmac:ACCESS_KEY_ID:ACCESS_KEY_SECRET&endpoint=https://oss-cn-hangzhou.aliyuncs.com",
    )
    if err != nil {
        log.Fatal(err)
    }
    _ = store
}
```

```go
// STS temporary credentials: pass "env" as credential value.
// Set ALIBABA_CLOUD_ACCESS_KEY_ID, ALIBABA_CLOUD_ACCESS_KEY_SECRET,
// and ALIBABA_CLOUD_SECURITY_TOKEN in your environment before running.
store, err := services.NewStoragerFromString(
    "oss://my-bucket/?credential=env&endpoint=https://oss-cn-shanghai.aliyuncs.com",
)
```

```go
// With extra pairs supplied at call time.
import (
    ps "github.com/rgglez/go-storage/v5/pairs"
)

store, err := services.NewStoragerFromString(
    "oss://my-bucket/backups/",
    ps.WithCredential("hmac:ACCESS_KEY_ID:ACCESS_KEY_SECRET"),
    ps.WithEndpoint("https://oss-cn-beijing.aliyuncs.com"),
)
```

**Required parameters for OSS**

| Parameter | Connection string key | Description |
|-----------|-----------------------|-------------|
| Credential | `credential` | `hmac:<access_key_id>:<access_key_secret>` or `env` |
| Bucket | `name` (path component) | Bucket name — place it immediately after `://` |
| Endpoint | `endpoint` | OSS regional endpoint URL |

**Optional parameters for OSS**

| Parameter | Connection string key | Description |
|-----------|-----------------------|-------------|
| Work dir | `work_dir` (path component) | Key prefix within the bucket |

---

## Error handling

```go
import (
    "errors"

    "github.com/rgglez/go-storage/v5/services"
)

store, err := services.NewStoragerFromString("s3://my-bucket/?...")
if err != nil {
    var initErr services.InitError
    if errors.As(err, &initErr) {
        // initErr.Type  — service type (e.g. "s3")
        // initErr.Op    — failed operation (e.g. "new_factory", "parse_conn")
        // initErr.Err   — underlying error
    }
    if errors.Is(err, services.ErrServiceNotRegistered) {
        // The blank-import for the service package is missing.
    }
    log.Fatal(err)
}
```

Common errors:

| Error | Cause |
|-------|-------|
| `ErrServiceNotRegistered` | The service package was not imported (missing blank import). |
| `ErrConnectionStringInvalid` | The connection string does not follow the `type://…` format. |
| Service-specific init error | A required pair (`credential`, `location`, `endpoint`, …) is missing. |

---

## Learn more

The specific instructions for each service is provided in the respective service's documentation.

---

## References

- [`services/factory.go`](../../services/factory.go) — `NewStoragerFromString` implementation
- [RFC-90](../rfcs/90-re-support-initialization-via-connection-string.md) — Connection string design
- [RFC-700](../rfcs/700-config-features-and-defaultpairs-via-connection-string.md) — Feature flags and default pairs via connection string
- [RFC-970](../rfcs/970-service-factory.md) — Service factory design
- [S3 backend docs](services/s3.md)
- [fs backend docs](services/fs.md)
- [OSS backend docs](services/oss.md)
