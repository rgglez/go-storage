# How to generate the packages for the sub-modules

```go-storage``` includes Go "sub-modules" (actually modules in themselves, in fact) within subdirectories of the project. The original author chose this organization because there were too many sub-modules and creating a separated Git project for each one was more difficult to mantain.

Note: remember to replace the example values for your real case.

## Preparation

1. Determine the directory (sub-module) which you want to release, for example ```services/fs```
1. Determine the version number you'll use (in [semantic](https://semver.org/) notation). For instance ```v4.1.4```.
1. Commit or stash all your pending changes.

## Git operations

1. Create a new branch for the submodule:

   * Create and switch to a branch that will contain only the services/fs subdirectory. This branch will allow you to create tags specific to the submodule. For example:

      ```bash
      git checkout -b services-fs-v4.1.4
      ```

1. Remove all other files (temporarily):

   * In this branch, remove everything except the services/fs directory. You can do this by running:

      ```bash
      git rm -rf --cached .
      git reset -- services/fs
      git add services/fs
      git commit -m "Isolated services/fs module for tagging"
      ```

   This way, the branch will only track files within services/fs.

   *Note that the following step is optional, mainly for new modules.*

1. Add and commit ```go.mod``` in the subdirectory:

   * Ensure the services/fs subdirectory has a go.mod file with the correct module path. For example:

      ```bash
      module github.com/rgglez/go-storage/services/fs/v4
      ```

   * Commit this change if it hasn’t been committed yet:

      ```bash
      git add services/fs/go.mod
      git commit -m "Configure go.mod for services/fs module"
      ```

1. Create and push the tag:

   * Now that only the services/fs files are in this branch, you can create and push a tag specific to this module.

      ```bash
      git tag services/fs/v4.1.4
      git push origin services-fs-v4.1.4 --tags
      ```

## Clean up

Once you have created the tag, you can do one of the following steps to clean the branch up and return to the master branch. Otherwise you will get this error if you execute ```git checkout master``` inmediately.

```bash
The following untracked working tree files would be overwritten by checkout
```

### Why this error happens

The error occurs because you have untracked files in your working directory that will conflict with files in the branch you're trying to switch to (in this case, master). Git is preventing you from losing those untracked files by overwriting them.

After you ran the ```git rm -rf --cached .``` command, you effectively removed the files from Git’s index, making them untracked. However, the files still exist in your working directory. When you try to switch back to another branch (like master), Git sees that the files in the working directory would conflict with tracked files in the master branch.

### Option 1: Commit or stash your current changes

If the untracked files are important and you want to keep them, you can either commit them to your current branch or stash them before switching branches.

1. Commit the changes:

   * If you want to keep these changes, add and commit them to the current branch:

      ```bash
      git add .
      git commit -m "Committing untracked changes before switching branches"
      ```

   * Then you can switch to master without issues:

      ```bash
      git checkout master
      ```

1. Stash the changes:

   * If you don’t want to commit these changes but still want to keep them temporarily, stash them:

      ```bash
      git stash -u  # -u includes untracked files
      ```

   * After stashing, you should be able to switch branches:

      ```bash
      git checkout master
      ```

   * Later, you can apply the stashed changes (if needed) by running:

      ```bash
      git stash apply
      ```

### Option 2: discard untracked changes

If the untracked files aren’t important and you don’t need them, you can forcefully discard them to switch branches.

1. Remove untracked files:

   * You can remove all untracked files using the following command:

      ```bash
      git clean -fd
      ```

   * ```-f``` Force deletion of untracked files.
   * ```-d``` Remove untracked directories.

1. Switch to master:

   * Once the untracked files are deleted, you can switch to the master branch without any conflicts:

      ```bash
      git checkout master
      ```

### Option 3: backup untracked files (manual backup)

If you’re unsure whether you need the untracked files, you can manually copy them to a backup location and then remove them from the working directory.

1. Copy files:

   * Manually copy the untracked files or directories (like services/fs) to another location outside the Git repository.

1. Remove untracked files:

   * Run the following to clean up the working directory:

      ```bash
      git clean -fd
      ```

1. Switch to master:

   * After removing the untracked files, you can switch to the master branch:

      ```bash
      git checkout master
      ```

## Using the module in another project

With the tag in place, you can now pull in this module in another project:

```bash
go get github.com/rgglez/go-storage/services/fs/v4@v4.1.4
```
