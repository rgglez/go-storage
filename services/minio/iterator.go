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
package minio

import (
	"strconv"

	"github.com/minio/minio-go/v7"
)

type storagePageStatus struct {
	buckets []minio.BucketInfo
}

func (i *storagePageStatus) ContinuationToken() string {
	return ""
}

type objectPageStatus struct {
	bufferSize int
	counter    int
	options    minio.ListObjectsOptions

	objChan <-chan minio.ObjectInfo
}

func (i *objectPageStatus) ContinuationToken() string {
	return strconv.Itoa(i.counter)
}
