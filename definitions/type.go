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
package definitions

type Type struct {
	Expr    string // Raw expr that before type name, e.g. `[]`, `...`, `[]*`
	Package string
	Name    string
}

func (t Type) FullName(pkg ...string) string {
	// The type is a builtin type, we can use directly.
	if t.Package == "" {
		return t.Expr + t.Name
	}
	// The types package name is the same with input one.
	if len(pkg) > 0 && t.Package == pkg[0] {
		return t.Expr + t.Name
	}
	return t.Expr + t.Package + "." + t.Name
}
