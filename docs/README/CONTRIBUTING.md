# Contributing to go-storage

First off, thanks for taking the time to contribute!

> **Important**: The original author ([Xuanwo](https://github.com/Xuanwo)) abandoned this project. All new contributions should be directed to the active fork at **[rgglez/go-storage](https://github.com/rgglez/go-storage)**. Please do not open issues or pull requests against the original repository.

Before contributing, please read the [Architecture Overview](docs/README/README_ARCHITECTURE.md) to understand how the system is structured, particularly the code generation pipeline and the manual vs. generated split in service implementations.

## Did you find a bug?

- Ensure the bug was not already reported by searching for existing issues in both:
  - The active fork: [rgglez/go-storage Issues](https://github.com/rgglez/go-storage/issues)
  - The original (archived) repository: [beyondstorage/go-storage Issues](https://github.com/beyondstorage/go-storage/issues)
- If the bug is not reported yet, open a new issue at [rgglez/go-storage](https://github.com/rgglez/go-storage/issues) with the following information:
  - Bug description
  - Library commit ID
  - Minimal reproduction code

## Did you write a patch that fixes a bug?

- Open a new GitHub pull request with the patch.
- Ensure the PR description clearly describes the problem and solution. Include the relevant issue number if applicable.
- Add unittest for this bug.

## Do you intend to implement a new service?

- `Storager` must be implemented, others can be optional.
- Add support in `coreutils.Open`.
- Add unittests as best effort.

## Do you intend to change public API?

- Open a new Github Issue for discuss.
- After achieve consensus, add a proposal in `docs/design` and submit a PR.
- Implement a proposal and change status to `candidate`

> In next release, relevant proposal statue will be updated to `finished`


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
