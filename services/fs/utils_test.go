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
package fs

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rgglez/go-storage/v5/services"
)

func TestNewClient(t *testing.T) {
	c, err := NewStorager()
	assert.NotNil(t, c)
	assert.NoError(t, err)
}

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

func BenchmarkStorage_getAbsPath(b *testing.B) {
	store := &Storage{
		workDir: "/abc/def",
	}

	b.StartTimer()
	for i := 0; i < b.N; i += 1 {
		_ = store.getAbsPath("xyz/nml")
	}
	b.StopTimer()
}
