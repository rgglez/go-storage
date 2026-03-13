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
	def "github.com/rgglez/go-storage/v5/definitions"
)

//go:generate go run .
func main() {
	def.GenerateIterator("../../../types/iterator.generated.go")
	def.GenerateInfo("../../../types/info.generated.go")
	def.GeneratePair("../../../pairs/generated.go")
	def.GenerateOperation("../../../types/operation.generated.go")
	def.GenerateObject("../../../types/object.generated.go")
	def.GenerateNamespace("../../../definitions/namespace.generated.go")
}
