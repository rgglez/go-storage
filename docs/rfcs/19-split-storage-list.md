- Author: Xuanwo <github@xuanwo.io>
- Start Date: 2020-04-09
- RFC PR: N/A
- Tracking Issue: N/A

# Proposal: Split storage list

- Updates:
  - [GSP-2](./2-use-callback-in-list-operations.md)
  - [GSP-12](./12-support-both-directory-and-prefix-based-list.md): Deprecates this RFC

## Background

proposal [support both directory and prefix based list] has been proved to be a failure by practice. In this proposal, we introduce `ObjectFunc` for prefix based list support, and add many restriction for the usage of `FileFunc`, `DirFunc` and `ObjectFunc`. The problem is user don't know whether this storage service is prefix based or directory based. So they always fallback to the directory based list method which is not suffcient for object storage service.

## Proposal

So I propose following changes:

- Split `List` into `ListDir` and `ListPrefix`
- Remove `List` from `Storager`
- Add interface `DirLister` for `ListDir`
- Add interface `PrefixLister` for `ListPrefix`

So user need to assert to interface `DirLister` to use `ListDir`.

At the same time, we should:

- Rename `ListSegments` to `ListPrefixSegments` to match prefix changes
- Remove `ListSegments` from `Segmenter`
- Add interface `PrefixSegmentsLister` for `ListSegments`

## Rationale

None.

## Compatibility

All API call to `List` will be broken.

## Implementation

Most of the work would be done by the author of this proposal.

[support both directory and prefix based list]: ./12-support-both-directory-and-prefix-based-list.md


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
