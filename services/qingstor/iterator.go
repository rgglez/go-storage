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
package qingstor

import (
	"strconv"
)

type objectPageStatus struct {
	delimiter    string
	limit        int
	marker       string
	prefix       string
	partIdMarker string
}

func (i *objectPageStatus) ContinuationToken() string {
	if i.partIdMarker != "" {
		return i.marker + "/" + i.partIdMarker
	}
	return i.marker
}

type storagePageStatus struct {
	limit    int
	offset   int
	location string
}

func (i *storagePageStatus) ContinuationToken() string {
	return strconv.FormatInt(int64(i.offset), 10)
}

type partPageStatus struct {
	prefix           string
	limit            int
	partNumberMarker int
	uploadID         string
}

func (i *partPageStatus) ContinuationToken() string {
	return strconv.FormatInt(int64(i.partNumberMarker), 10)
}
