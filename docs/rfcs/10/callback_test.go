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
	"bytes"
	"io"
	"testing"
)

// Use 4K read for benchmark
var n int64 = 4 * 1024
var content = bytes.Repeat([]byte{'x'}, int(n))

func BenchmarkPlainReader(b *testing.B) {
	b.SetBytes(n)
	for i := 0; i < b.N; i++ {
		r := bytes.NewReader(content)
		_, _ = io.ReadAll(r)
	}
}

type IntCallbackReader struct {
	r  io.Reader
	fn func(int)
}

func (r IntCallbackReader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	r.fn(n)
	return n, err
}

func BenchmarkIntCallbackReader(b *testing.B) {
	b.SetBytes(n)
	for i := 0; i < b.N; i++ {
		x := 0
		r := IntCallbackReader{
			r: bytes.NewReader(content),
			fn: func(i int) {
				x += i
			},
		}
		_, _ = io.ReadAll(r)
	}
}

type BytesCallbackReader struct {
	r  io.Reader
	fn func([]byte)
}

func (r BytesCallbackReader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	r.fn(p[:n])
	return n, err
}

func BenchmarkBytesCallbackReader(b *testing.B) {
	b.SetBytes(n)
	for i := 0; i < b.N; i++ {
		x := 0
		r := BytesCallbackReader{
			r: bytes.NewReader(content),
			fn: func(i []byte) {
				x += len(i)
			},
		}
		_, _ = io.ReadAll(r)
	}
}
