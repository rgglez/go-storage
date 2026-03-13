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
package s3

import (
	"strconv"
)

type objectPageStatus struct {
	delimiter string
	maxKeys   int64
	prefix    string

	// Only used for object
	continuationToken string

	// Only used for part object
	keyMarker      string
	uploadIdMarker string

	expectedBucketOwner string
}

// getServiceContinuationToken equals aws.String, but return nil while empty.
//
// NOTES:
//   aws will return "InvalidArgument: The continuation token provided is incorrect" if
//   input's ContinuationToken is set to "".
func (i objectPageStatus) getServiceContinuationToken() *string {
	if i.continuationToken == "" {
		return nil
	}
	return &i.continuationToken
}

func (i *objectPageStatus) ContinuationToken() string {
	if i.uploadIdMarker != "" {
		return i.continuationToken + "/" + i.uploadIdMarker
	}
	return i.continuationToken
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
	key              string
	maxParts         int64
	partNumberMarker string
	uploadId         string

	expectedBucketOwner string
}

func (i *partPageStatus) ContinuationToken() string {
	return i.partNumberMarker
}
