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
package randbytes

import (
	"io"
	"math/rand"
	"time"
)

// Rand creates a stream of non-crypto quality random bytes
type Rand struct {
	rand.Source
}

// NewRand creates a new random reader with a time source.
func NewRand() io.Reader {
	return &Rand{rand.NewSource(time.Now().UnixNano())}
}

// Read satisfies io.Reader
func (r *Rand) Read(p []byte) (n int, err error) {
	todo := len(p)
	offset := 0
	for {
		val := r.Int63()
		for i := 0; i < 7; i++ {
			p[offset] = byte(val)
			todo--
			if todo == 0 {
				return len(p), nil
			}
			offset++
			val >>= 8
		}
	}
}
