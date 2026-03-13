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
package ftp

import (
	"context"
	"fmt"
	"strconv"

	"github.com/rgglez/go-storage/v5/types"
)

func (s *Storage) listDirNext(ctx context.Context, page *types.ObjectPage) (err error) {
	input := page.Status.(*listDirInput)
	if input.objList == nil {
		input.objList, err = s.connection.List(input.rp)
		if err != nil {
			return err
		}
	}
	if !input.started {
		input.counter, err = strconv.Atoi(input.continuationToken)
		if err != nil {
			input.counter = 0
		}
		input.started = true
	}
	n := len(input.objList)
	input.continuationToken = fmt.Sprintf("%x", input.counter)
	if input.counter >= n {
		return types.IterateDone
	}

	v := input.objList[input.counter]

	obj, err := s.formatFileObject(v, input.rp)
	if err != nil {
		return err
	}

	page.Data = append(page.Data, obj)

	input.counter++

	return
}
