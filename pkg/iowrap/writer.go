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
package iowrap

import (
	"io"
)

// CallbackWriter will create a new CallbackifyWriter.
func CallbackWriter(w io.Writer, fn func([]byte)) *CallbackifyWriter {
	return &CallbackifyWriter{
		w:  w,
		fn: fn,
	}
}

// CallbackifyWriter will execute callback func in Write.
type CallbackifyWriter struct {
	w  io.Writer
	fn func([]byte)
}

// Write will write into underlying Writer.
func (w *CallbackifyWriter) Write(p []byte) (int, error) {
	n, err := w.w.Write(p)
	w.fn(p[:n])
	return n, err
}
