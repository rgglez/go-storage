- Author: Xuanwo <github@xuanwo.io>
- Start Date: 2020-03-18
- RFC PR: N/A
- Tracking Issue: N/A

# Proposal: Proposal process

- Updated By:
  - [GSP-118](./118-update-rfc-format-and-process.md): Updates the proposal spec and process

## Background

[storage]'s development is proposal driven. We need to explain why and how we make changes to [storage] so that we can understand why we are at here.

## Proposal

So I propose following process procedure:

**Simple changes**

- Send PR directly.

**BUG Fix**

- Create an issue 
- Send a related PR to resolve it

**Big changes**

- Create an issue
- Send a PR with proposal
- Implement proposal

All steps do not need to be done by the same person. For example, issue could be created by user A, and proposal written by user B, and implemented by user C. 

Changes level could be increased while needed. For example, user A sends an one line simple change, but it found out that we need a whole refactor on this package. At this time, we will need to follow the **Big changes** procedure.

Proposal's spec will be presented in spec [2-proposal].

## Rationale

None

## Compatibility

None

## Implementation

No code related changes.

[storage]: https://github.com/Xuanwo/storage
[2-proposal]: ../spec/2-proposal.md


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
