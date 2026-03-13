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
package randbytes

import (
	"crypto/rand"
	"io"
	"testing"
)

func TestRand(t *testing.T) {
	buf := make([]byte, 16)
	n, err := NewRand().Read(buf)
	if err != nil {
		t.Fatalf("Error reading: %v", err)
	}
	if n != len(buf) {
		t.Fatalf("Short read: %v", n)
	}
	t.Logf("Read %x", buf)
}

const toCopy = 1024 * 1024

func BenchmarkRand(b *testing.B) {
	b.SetBytes(toCopy)
	r := NewRand()
	for i := 0; i < b.N; i++ {
		_, _ = io.CopyN(io.Discard, r, toCopy)
	}
}

func BenchmarkCrypto(b *testing.B) {
	b.SetBytes(toCopy)
	for i := 0; i < b.N; i++ {
		_, _ = io.CopyN(io.Discard, rand.Reader, toCopy)
	}
}
