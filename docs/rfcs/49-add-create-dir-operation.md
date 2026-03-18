- Author: Xuanwo <github@xuanwo.io>
- Start Date: 2021-05-08
- RFC PR: [beyondstorage/specs#49](https://github.com/rgglez/specs/issues/49)
- Tracking Issue: N/A

# AOS-49: Add CreateDir Operation

- Updated By:
  - [GSP-837: Support Feature Flag](./837-support-feature-flag.md): Move `CreateDir` operation to `Storager`

## Background

Applications need the ability to create a directory. For now, our support is a bit wired.

In [fs](https://github.com/rgglez/go-service-fs), we support `CreateDir` by dirty hack:

```go
if s.isDirPath(rp) {
    // FIXME: Do we need to check r == nil && size == 0 ?
    return 0, s.createDir(rp)
}
```

In other storage service, user needs to create dir by special `content-type`:

```go
store.Write("abc/", nil, 0, ps.WithContentType("application/x-directory"))
```

We need to allow user create a directory in the same way.

## Proposal

So I propose to add a new operation `CreateDir` like we do on `append` / `multipart` Object.

```go
type Direr interface {
	CreateDir(path string, pairs ...Pair) (o *Object, err error)
}
```

`CreateDir` will return an Object with `dir` mode, and different service could have different implementations.

## Rationale

### Directory in Object Storage Services

Object Storage is a K-V Storage, and don't have the concept of directory natively. But most object storages support ListObjects via delimiter `/` to demonstrate a file system tree. With delimiter `/`, object storage services will organize objects end with `/` as common prefix.

## Compatibility

This proposal COULD break users who use `store.Write("abc/")` to create directory.

## Implementation

- Update specs to add `CreateDir` operations.
- Update go-storage to implement the changes.
- Update all services that support create dir.


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
