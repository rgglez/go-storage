// Copyright The go-storage Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
//go:build !windows
// +build !windows

package fs

// evalSymlinks returns the path name after the evaluation of any symbolic
// links.
// The original implementation can be referenced to filepath.EvalSymlinks,
// but it will return the current path other then ENOENT error while walkSymlinks.
// If path is relative the result will be relative to the current directory,
// unless one of the components is an absolute symbolic link.
// EvalSymlinks calls Clean on the result.
func evalSymlinks(path string) (string, error) {
	return walkSymlinks(path)
}
