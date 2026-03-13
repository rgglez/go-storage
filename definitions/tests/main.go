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
	"github.com/rgglez/go-storage/v5/types"
)

//go:generate go run .
func main() {
	def.GenerateService(metadata, "generated_test.go")
}

var metadata = def.Metadata{
	Name: "main",
	Pairs: []def.Pair{
		pairDisableUriCleaning,
		pairStorageClass,
		pairStringPair,
	},
	Infos: []def.Info{
		infoObjectMetaStorageClass,
		infoStorageMetaQueriesPerSecond,
	},
	Factory: []def.Pair{
		def.PairCredential,
		def.PairEndpoint,
		def.PairName,
		def.PairLocation,
		def.PairWorkDir,
		pairDisableUriCleaning,
	},
	Service: def.Service{
		Features: types.ServiceFeatures{
			Delete: true,
		},
		Create: []def.Pair{
			def.PairLocation,
		},
		Delete: []def.Pair{
			def.PairLocation,
		},
		Get: []def.Pair{
			def.PairLocation,
		},
		List: []def.Pair{
			def.PairLocation,
		},
	},
	Storage: def.Storage{
		Features: types.StorageFeatures{
			Read: true,
		},

		Delete: []def.Pair{
			def.PairMultipartID,
			def.PairObjectMode,
		},
		List: []def.Pair{
			def.PairListMode,
		},
		Read: []def.Pair{
			def.PairIoCallback,
			def.PairSize,
			def.PairOffset,
		},
		Create: []def.Pair{
			def.PairObjectMode,
		},
		Stat: []def.Pair{
			def.PairObjectMode,
		},
		Write: []def.Pair{
			def.PairContentMD5,
			def.PairContentType,
			def.PairIoCallback,
			pairStorageClass,
		},
	},
}

var pairDisableUriCleaning = def.Pair{
	Name: "disable_uri_cleaning",
	Type: def.Type{Name: "bool"},
}

var pairStorageClass = def.Pair{
	Name:        "storage_class",
	Type:        def.Type{Name: "string"},
	Defaultable: true,
}

var pairStringPair = def.Pair{
	Name:        "string_pair",
	Type:        def.Type{Name: "string"},
	Description: "tests connection string",
}

var infoObjectMetaStorageClass = def.Info{
	Namespace:   def.NamespaceObject,
	Category:    def.CategoryMeta,
	Name:        "storage_class",
	Type:        def.Type{Name: "string"},
	Description: "is the storage class for this object",
}
var infoStorageMetaQueriesPerSecond = def.Info{
	Namespace:   def.NamespaceStorage,
	Category:    def.CategoryMeta,
	Name:        "queries_per_second",
	Type:        def.Type{Name: "int64"},
	Description: "tests storage system metadata",
}
