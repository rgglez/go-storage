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
package hdfs

import (
	"context"
	"errors"
	"io"

	"github.com/colinmarc/hdfs/v2"
	"github.com/rgglez/go-storage/v5/types"
)

const defaultListObjectLimit = 100

type listDirInput struct {
	rp  string
	dir *hdfs.FileReader

	continuationToken string
}

func (i *listDirInput) ContinuationToken() string {
	return i.continuationToken
}

func (s *Storage) listDirNext(ctx context.Context, page *types.ObjectPage) (err error) {
	input := page.Status.(*listDirInput)

	if input.dir == nil {
		input.dir, err = s.hdfs.Open(input.rp)
		if err != nil {
			return
		}
	}

	fileList, err := input.dir.Readdir(defaultListObjectLimit)

	if err != nil && errors.Is(err, io.EOF) {
		_ = input.dir.Close()
		input.dir = nil
		return types.IterateDone
	}

	for _, f := range fileList {
		o := s.newObject(true)
		o.ID = input.rp
		o.Path = f.Name()

		if f.Mode().IsDir() {
			o.Mode |= types.ModeDir
		}

		if f.Mode().IsRegular() {
			o.Mode |= types.ModeRead
		}

		o.SetContentLength(f.Size())

		page.Data = append(page.Data, o)
		input.continuationToken = o.Path
	}

	return
}
