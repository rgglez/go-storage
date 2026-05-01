# 🏷️ How to generate the packages for the sub-modules

```go-storage``` includes Go sub-modules within subdirectories of the project. The original author chose this organization because there were too many sub-modules and creating a separated Git project for each one was more difficult to maintain.

The Go toolchain resolves a sub-module tag by its path prefix (e.g. `services/fs/`) directly against any commit — no special branch or file isolation is needed.

Note: remember to replace the example values with your real case.

## ⚡ Quick path: Makefile targets

The root `Makefile` provides targets that automatically compute and create the next PATCH version tag, so you don't have to look up the current version manually.

### Create tags (local only)

| Command | What it does | Example result |
|---------|--------------|----------------|
| `make tag-service SERVICE=<name>` | Next PATCH tag for `services/<name>` | `services/oss/v3.0.7` |
| `make tag-credential` | Next PATCH tag for `credential` | `credential/v1.0.2` |
| `make tag-endpoint` | Next PATCH tag for `endpoint` | `endpoint/v1.2.2` |
| `make next-tag` | Next PATCH tag for the root module | `v5.0.1` |

### Push tags to origin

| Command | What it pushes |
|---------|----------------|
| `make push-tag-service SERVICE=<name>` | Latest `services/<name>/vX.Y.Z` tag |
| `make push-tag-credential` | Latest `credential/vX.Y.Z` tag |
| `make push-tag-endpoint` | Latest `endpoint/vX.Y.Z` tag |
| `make push-next-tag` | Latest root `vX.Y.Z` tag |

### Typical workflow

```bash
# 1. Verify what tags exist today
make latest-tags

# 2. Create the next patch tag locally
make tag-service SERVICE=oss
# Output: Creating tag: services/oss/v3.0.7

# 3. Inspect and push
git tag --sort=-v:refname | head -5
make push-tag-service SERVICE=oss
```

> **Note:** The `push-*` targets always push the *latest* tag of that module, which is the one just created by the matching `tag-*` target.

---

## 🧰 Manual preparation

1. Determine the directory (sub-module) which you want to release, for example `services/fs`.
1. Determine the version number you'll use (in [semantic](https://semver.org/) notation). For instance `v5.1.4`.
1. Commit or stash all your pending changes.

## 🔍 Verify `go.mod`

Ensure the sub-module's `go.mod` has the correct module path. For example:

```
module github.com/rgglez/go-storage/services/fs/v5
```

If it doesn't exist or the path is wrong, fix and commit it to `master` before tagging:

```bash
git add services/fs/go.mod
git commit -m "Configure go.mod for services/fs module"
```

## 🚀 Create and push the tag

Tags for Go sub-modules must include the sub-directory path as a prefix:

```bash
git tag services/fs/v5.1.4
git push origin services/fs/v5.1.4
```

That's all. The Go toolchain uses the tag prefix (`services/fs/`) to locate the module within the repository. No branch isolation or file removal is required.

## 📥 Using the module in another project

With the tag in place, you can pull in this module in another project:

```bash
go get github.com/rgglez/go-storage/services/fs/v5@v5.1.4
```


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
