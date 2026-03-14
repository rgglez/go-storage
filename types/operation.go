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
package types

import "strings"

// ListMode is the type for list, underlying type is int.
type ListMode uint8

const (
	// ListModeDir means this list will use dir type.
	ListModeDir ListMode = 1 << iota
	// ListModePrefix means this list will use prefix type.
	ListModePrefix
	ListModePart
	ListModeBlock
)

// String implement Stringer for ListMode.
//
// An object with Dir,Part will print like "dir|part"
func (l ListMode) String() string {
	s := make([]string, 0)
	if l.IsDir() {
		s = append(s, "dir")
	}
	if l.IsPrefix() {
		s = append(s, "prefix")
	}
	if l.IsPart() {
		s = append(s, "part")
	}
	if l.IsBlock() {
		s = append(s, "block")
	}
	return strings.Join(s, "|")
}

func (l ListMode) IsDir() bool {
	return l&ListModeDir != 0
}

func (l ListMode) IsPrefix() bool {
	return l&ListModePrefix != 0
}

func (l ListMode) IsPart() bool {
	return l&ListModePart != 0
}

func (l ListMode) IsBlock() bool {
	return l&ListModeBlock != 0
}
