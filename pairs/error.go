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
package pairs

import (
	"errors"
	"fmt"
)

var (
	// ErrPairTypeMismatch means the pair's type is not match
	ErrPairTypeMismatch = errors.New("pair type mismatch")
)

// Error represents error related to a pair.
type Error struct {
	Op  string
	Err error

	Key   string
	Type  string
	Value interface{}
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: key %s, type %s, value %s: %s", e.Op, e.Key, e.Type, e.Value, e.Err.Error())
}

// Unwrap implements xerrors.Wrapper
func (e *Error) Unwrap() error {
	return e.Err
}
