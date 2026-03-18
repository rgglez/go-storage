- Author: Xuanwo <github@xuanwo.io>
- Start Date: 2021-05-06
- RFC PR: [beyondstorage/specs#46](https://github.com/rgglez/specs/issues/46)
- Tracking Issue: N/A

# AOS-46: Idempotent Storager Delete Operation

## Background

After [AOS-25] has been introduced, we use `Delete` to handle all object delete operations. But their behavior is not well-defined:

- File system alike service may return `not exist` for object that not exist.
- Object storage alike service always return success no matter the object exist or not.

The problem is much more serious for `multipart` and `append` object:

- For `fs` and `qingstor`, `append` object is exist and readable once they have been created.
- For `dropbox`, `append` object is exist and readable after they have been `CommitAppend`.

If user tried to delete an `append` object, he will get errors under `dropbox`.

## Proposal

So I propose to make Storager's `Delete` operation idempotent.

`idempotent` means:

- Without outside changes, `Delete` operation could be executed on the same path multiple times and always get the same results.
- `Delete` operation will not return `ObjectNotExist` anymore.

For service provider:

- Don't NEED to check the file exist or not.
- SHOULD omit `ObjectNotExist` related error. (Especially for `multipart` and `append` object)

## Rationale

### Alternative Way: Make sure ObjectNotExist returned

The alternative way is make `Delete` more strict: no matter what the object mode is, we always check the object before delete them.

For object storage service, we need to `stat` object and return ObjectNotExist directly if stat returns ObjectNotExist.

This will need extra requests.

### What if user wants to know whether a file has been deleted or not?

User can stat the object by self before delete, or we can provide an operation called `CheckedDelete(path string) (deleted bool, err error)` in another interface.

### Corner Cases

`dropbox` returns a `upload-session-id` for `create_append` operation, the file is visitable after `commit_append`.

So:

- If there is an object exist, the `delete` operation could delete it by mistake.
- If the object returned by `create_append` is missing, user could not resume the uploads.

## Compatibility

For `fs` and `dropbox`: `Delete` will not return `ObjectNotExist` anymore. 

## Implementation

- Update [go-integration-test](https://github.com/rgglez/go-integration-test)
  - Add a case that delete an object twice, and should not meet error.
- Make sure all service implement delete correctly.

[AOS-25]: ./25-object-mode.md


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
