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

type Metadata struct {
	Name string

	// Imports lists additional package import paths needed by custom pair types.
	Imports []string

	Pairs   []Pair
	Infos   []Info
	Factory []Pair

	Service Namespace
	Storage Namespace
}

func (m Metadata) Normalize() Metadata {
	m.buildDefaultPairs()
	m.buildFeaturePairs()
	return m
}

func (m *Metadata) buildDefaultPairs() {
	dp := make([]Pair, 0)
	for _, v := range m.Pairs {
		if !v.Defaultable {
			continue
		}
		dp = append(dp, Pair{
			Name:        "default_" + v.Name,
			Type:        v.Type,
			Description: "default value for " + v.Name,
		})
	}
	m.Pairs = append(m.Pairs, dp...)
	m.Factory = append(m.Factory, dp...)
}

func (m *Metadata) buildFeaturePairs() {
	dp := make(map[string]bool)
	for _, v := range []Namespace{m.Service, m.Storage} {
		for _, f := range v.ListFeatures(FeatureTypeVirtual) {
			dp["enable_"+f.Name] = true
		}
	}

	for name := range dp {
		m.Factory = append(m.Factory, PairMap[name])
	}
}
