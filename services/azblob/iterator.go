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
package azblob

import "github.com/Azure/azure-storage-blob-go/azblob"

type objectPageStatus struct {
	delimiter  string
	maxResults int32
	prefix     string
	marker     azblob.Marker
}

func (i *objectPageStatus) ContinuationToken() string {
	if i.marker.NotDone() {
		return *i.marker.Val
	}
	return ""
}

type storagePageStatus struct {
	marker     azblob.Marker
	maxResults int32
}

func (i *storagePageStatus) ContinuationToken() string {
	if i.marker.NotDone() {
		return *i.marker.Val
	}
	return ""
}
