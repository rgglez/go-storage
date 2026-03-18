- Author: npofsi <npofsi@outlook.com>
- Start Date: 2021-08-02
- RFC PR: https://github.com/rgglez/go-credential/pull/3
- Tracking Issue: [beyondstorage/go-credential#1](https://github.com/rgglez/go-credential/issues/1)

# RFC-3: Add protocol basic


## Background

Some services don't support other credentials, but user or password.


## Proposal

Add protocol `basic`. Like `hmac`, `basic` have two parameters, corresponding to user and password of an account.

For example, go-service-ftp need a account to sign in, like

`ftp://xxx?credential=basic:user:password`

## Rationale

- Account is the only certification to some platform.
- Account is a basic method to identify quests.

## Compatibility

Will just add a choose to use protocol `basic`.

## Implementation

Just need to parse `basic` like `hmac`.

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
