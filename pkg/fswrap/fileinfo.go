// Copyright The go-storage Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package fswrap

import (
	"os"
	"time"

	"github.com/rgglez/go-storage/v5/types"
)

type fileInfoWrapper struct {
	object *types.Object
}

func (o fileInfoWrapper) Name() string {
	return o.object.Path
}

func (o fileInfoWrapper) Size() int64 {
	return o.object.MustGetContentLength()
}

func (o fileInfoWrapper) Mode() os.FileMode {
	return formatFileMode(o.object.Mode)
}

func (o fileInfoWrapper) ModTime() time.Time {
	return o.object.MustGetLastModified()
}

func (o fileInfoWrapper) IsDir() bool {
	return o.object.Mode.IsDir()
}

// Sys will return internal Object.
func (o fileInfoWrapper) Sys() interface{} {
	return o.object
}

func formatFileMode(om types.ObjectMode) os.FileMode {
	var m os.FileMode

	if om.IsDir() {
		m |= os.ModeDir
	}
	if om.IsAppend() {
		m |= os.ModeAppend
	}
	if om.IsLink() {
		m |= os.ModeSymlink
	}
	return m
}
