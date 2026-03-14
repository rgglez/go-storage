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
package hdfs

import (
	"errors"
	"os"
	"testing"

	"github.com/rgglez/go-storage/v5/services"
	"github.com/stretchr/testify/assert"
)

func TestFormatOsError(t *testing.T) {
	testErr := errors.New("test error")
	tests := []struct {
		name     string
		input    error
		expected error
	}{
		{
			"not found",
			os.ErrNotExist,
			services.ErrObjectNotExist,
		},
		{
			"not supported error",
			testErr,
			services.ErrUnexpected,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := formatError(tt.input)
			assert.True(t, errors.Is(err, tt.expected))
		})
	}
}

func TestGetAbsPath(t *testing.T) {
	cases := []struct {
		name         string
		base         string
		path         string
		expectedPath string
	}{
		{"direct path", "", "abc", "abc"},
		{"under direct path", "", "root/abc", "root/abc"},
		{"under direct path", "", "root/abc/", "root/abc"},
		{"under root", "/", "abc", "/abc"},
		{"under exist dir", "/root", "abc", "/root/abc"},
		{"under new dir", "/root", "abc/", "/root/abc"},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			client := Storage{workDir: tt.base}

			getPath := client.getAbsPath(tt.path)
			assert.Equal(t, tt.expectedPath, getPath)
		})
	}
}
