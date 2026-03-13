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
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactory_FromString(t *testing.T) {
	f := &Factory{}

	err := f.FromString("hmac:ak:sk@http:xxx:xxx/bucket/dir/?disable_uri_cleaning")
	assert.NoError(t, err)
	assert.Equal(t, &Factory{
		Credential:         "hmac:ak:sk",
		DisableURICleaning: true,
		Endpoint:           "http:xxx:xxx",
		Name:               "bucket",
		WorkDir:            "/dir/",
	}, f)
}
