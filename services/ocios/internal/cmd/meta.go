package main

import (
	def "github.com/rgglez/go-storage/v5/definitions"
	"github.com/rgglez/go-storage/v5/types"
)

var Metadata = def.Metadata{
	Name:  "ocios",
	Pairs: []def.Pair{},
	Infos: []def.Info{},
	Factory: []def.Pair{
		def.PairWorkDir,
	},
	Service: def.Service{
		Features: types.ServiceFeatures{},
	},
	Storage: def.Storage{
		Features: types.StorageFeatures{
			Create:   true,
			Delete:   true,
			List:     true,
			Metadata: true,
			Read:     true,
			Stat:     true,
			Write:    true,
		},

		Create: []def.Pair{
			def.PairObjectMode,
		},
		Delete: []def.Pair{
			def.PairObjectMode,
		},
		List: []def.Pair{
			def.PairListMode,
		},
		Read: []def.Pair{
			def.PairOffset,
			def.PairIoCallback,
			def.PairSize,
		},
		Write: []def.Pair{
			def.PairContentMD5,
			def.PairContentType,
			def.PairIoCallback,
		},
		Stat: []def.Pair{
			def.PairObjectMode,
		},
	},
}
