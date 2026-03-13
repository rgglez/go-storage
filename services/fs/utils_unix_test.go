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
//go:build linux || darwin
// +build linux darwin

package fs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAbsPath(t *testing.T) {
	cases := []struct {
		name         string
		base         string
		path         string
		expectedPath string
	}{
		{"under root", "/", "abc", "/abc"},
		{"under sub dir", "/root", "abc", "/root/abc"},
		{"with dir path", "/root", "abc/", "/root/abc"},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			client := Storage{workDir: tt.base}

			gotPath := client.getAbsPath(tt.path)
			assert.Equal(t, tt.expectedPath, gotPath)
		})
	}
}
