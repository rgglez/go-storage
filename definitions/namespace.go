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
package definitions

type Namespace interface {
	Name() string
	Operations() []Operation
	HasFeature(name string) bool
	ListFeatures(ty ...string) []Feature
	ListPairs(name string) []Pair
}

func (s Service) Name() string {
	return NamespaceService
}

func (s Service) Operations() []Operation {
	return SortOperations(OperationsService)
}

func (s Service) HasFeature(name string) bool {
	return s.Features.Has(name)
}

func (s Service) ListFeatures(ty ...string) []Feature {
	fs := make([]Feature, 0)

	m := make(map[string]bool)
	for _, v := range ty {
		m[v] = true
	}

	for _, f := range FeaturesService {
		if s.Features.Has(f.Name) && m[f.Type] {
			fs = append(fs, f)
		}
	}
	return SortFeatures(fs)
}

func (s Storage) Name() string {
	return NamespaceStorage
}

func (s Storage) Operations() []Operation {
	return SortOperations(OperationsStorage)
}

func (s Storage) HasFeature(name string) bool {
	return s.Features.Has(name)
}

func (s Storage) ListFeatures(ty ...string) []Feature {
	fs := make([]Feature, 0)

	m := make(map[string]bool)
	for _, v := range ty {
		m[v] = true
	}

	for _, f := range FeaturesStorage {
		if s.Features.Has(f.Name) && m[f.Type] {
			fs = append(fs, f)
		}
	}
	return SortFeatures(fs)
}
