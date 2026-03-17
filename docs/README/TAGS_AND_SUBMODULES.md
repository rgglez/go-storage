# 🏷️ How to generate the packages for the sub-modules

```go-storage``` includes Go sub-modules within subdirectories of the project. The original author chose this organization because there were too many sub-modules and creating a separated Git project for each one was more difficult to maintain.

The Go toolchain resolves a sub-module tag by its path prefix (e.g. `services/fs/`) directly against any commit — no special branch or file isolation is needed.

Note: remember to replace the example values with your real case.

## 🧰 Preparation

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
