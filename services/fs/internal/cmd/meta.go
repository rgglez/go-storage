package main

import (
	def "github.com/rgglez/go-storage/v5/definitions"
	"github.com/rgglez/go-storage/v5/types"
)

var Metadata = def.Metadata{
	Name:    "fs",
	Service: def.Service{},
	Factory: []def.Pair{
		def.PairWorkDir,
	},
	Storage: def.Storage{
		Features: types.StorageFeatures{
			CommitAppend: true,
			Copy:         true,
			Create:       true,
			CreateAppend: true,
			CreateDir:    true,
			CreateLink:   true,
			Delete:       true,
			Fetch:        true,
			List:         true,
			Metadata:     true,
			Move:         true,
			Read:         true,
			Stat:         true,
			Write:        true,
			WriteAppend:  true,
		},

		Create: []def.Pair{
			def.PairObjectMode,
		},
		Delete: []def.Pair{
			def.PairObjectMode,
		},
		List: []def.Pair{
			def.PairContinuationToken,
			def.PairListMode,
		},
		Read: []def.Pair{
			def.PairOffset,
			def.PairIoCallback,
			def.PairSize,
		},
		Stat: []def.Pair{
			def.PairObjectMode,
		},
		Write: []def.Pair{
			def.PairContentMD5,
			def.PairContentType,
			def.PairOffset,
			def.PairIoCallback,
		},
	},
}
