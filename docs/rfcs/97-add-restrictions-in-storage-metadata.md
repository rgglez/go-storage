- Author: JinnyYi <github.com/JinnyYi>
- Start Date: 2021-06-09
- RFC PR: [beyondstorage/specs#97](https://github.com/rgglez/specs/issues/97)
- Tracking Issue: [beyondstorage/go-storage#599](https://github.com/rgglez/go-storage/issues/599)

# GSP-97: Add Restrictions In Storage Metadata

## Background

We use `Write` to handle single write operation. And the size is not unlimited.

- S3 alike storage services have 5GB limit for a single PUT operation.
- azblob limited to 5000 MiB for put blob (<https://docs.microsoft.com/en-us/rest/api/storageservices/put-blob>).

If we upload a file with size out of limit in a single operation, we will get an error like `Request Entity Too Large` in azblob.

Similar restriction also with `copy`,`append`, etc. We should figure out the error before sending request to reduce the consumption of network resources.

They are all storage level restrictions and shared across the whole service.

For `append` and `multipart`, related metadata are misplaced at the object level, the current design leads to following problems:

- Enforce users must get an `Object` to check the maximum, they can't check before creating a multipart object.
- Service should set those storage level restrictions every time creates an object, which makes it hard to maintain.

Actually, we have `StorageMeta` to carry storage metadata which is generated by `go generate cmd/definitions`:

```go
type StorageMeta struct {
	location string
	Name     string
	WorkDir  string
	
	// bit used as a bitmap for object value, 0 means not set, 1 means set 
	bit uint64
	m   map[string]interface{}
}
```

But the current `StorageMeta` cannot solve the above problem:

- The restrictions should be global storage metadata.
- For now, `m` is unused, and it's not convenient to hold the definite metadata.

## Proposal

So I propose to add restrictions in storage metadata.

### Add global storage metadata

Add storage level restrictions as storage metadata into [info_storage_meta.toml]

- `copy-size-maximum`: Maximum size for copy operation.
- `fetch-size-maximum`: Maximum size for fetch operation.
- `move-size-maximum`: Maximum size for move operation.
- `write-size-maximum`: Maximum size for write operation.
- `append-number-maximum`: Max append numbers in append operation.
- `append-size-maximum`: Max append size in per append operation.
- `append-total-size-maximum`: Max append total size in append operation.
- `multipart-number-maximum`: Maximum part number in multipart operation.
- `multipart-size-maximum`: Maximum part size defined by storage service.
- `multipart-size-minimum`: Minimum part size defined by storage service.

### Deprecate the Existing Storage Metadata in Object Metadata

The following append and multipart related metadata in the current object metadata, actually belongs to storage metadata, should be deprecated in object metadata.

```go
type Object struct {
	...
    // AppendNumberMaximum Max append numbers in append operation
    appendNumberMaximum int
    // AppendSizeMaximum Max append size in per append operation
    appendSizeMaximum int64
    // AppendTotalSizeMaximum Max append total size in append operation
    appendTotalSizeMaximum int64
	...
    // MultipartNumberMaximum Maximum part number in multipart operation
    multipartNumberMaximum int
    // MultipartSizeMaximum Maximum part size defined by storager
    multipartSizeMaximum int64
    // MultipartSizeMinimum Minimum part size defined by storager
    multipartSizeMinimum int64
	...
}
```


Based on above, the added storage metadata and corresponding `set/get` functions could be generated and added into `StorageMeta`.

```go
type StorageMeta struct {
	location               string
	Name                   string
	WorkDir                string
	copySizeMaximum        int64
	fetchSizeMaximum       int64
	moveSizeMaximum        int64
	writeSizeMaximum       int64
	appendNumberMaximum    int64
	appendSizeMaximum      int64
	appendTotalSizeMaximum int64
	multipartNumberMaximum int64
	multipartSizeMaximum   int64
	multipartSizeMinimum   int64
	
	// bit used as a bitmap for object value, 0 means not set, 1 means set 
	bit uint64
	m   map[string]interface{}
}
```

For services:

- References to `Object.appendNumberMaximum`, etc SHOULD to be updated.
- `metadata` SHOULD return the added storage metadata if the service imposes limits on the resources and features.

## Rationale

### Alternative Way: Ignore size limit

The alternative way is to send request directly without size check and get an error when the `size` larger than the limit.

This will need invalid request and get the ambiguous error.

## Compatibility

This change will not break services and users. We could migrate as follows:

- Add fields in `StorageMeta` and mark append and multipart related meta as deprecated in `Object`.
- Release a new version for [go-storage] and all services bump to this version with all references to `Object.appendNumberMaximum`, etc updated.
- Remove deprecated fields in `Object` in the next major version.

## Implementation

- `specs`
  - Add storage meta in [info_storage_meta.toml].
  - Mark append and multipart related meta as deprecated in [info_object_meta.toml].
- `go-storage`
  - Generate the added storage metadata and corresponding `get/set` functions into `StorageMeta`.
  - Add comments start with `Deprecated` at the deprecated object metadata in generate template.
- `go-service-*`
  - Update all references to `Object.appendNumberMaximum`, etc.
  - `metadata` should return the added meta if the service imposes the limits.
  - Check `size` for write related operations and return `ErrRestrictionDissatisfied` when `size` out of limit.


[go-storage]: https://github.com/rgglez/go-storage
[info_storage_meta.toml]: ../definitions/info_storage_meta.toml
[info_object_meta.toml]: ../definitions/info_object_meta.toml