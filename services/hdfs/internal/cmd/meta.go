package main

import (
	def "github.com/rgglez/go-storage/v5/definitions"
	"github.com/rgglez/go-storage/v5/types"
)

var Metadata = def.Metadata{
	Name:    "hdfs",
	Service: def.Service{},
	Factory: []def.Pair{
		def.PairEndpoint,
		def.PairWorkDir,
	},
	Storage: def.Storage{
		Features: types.StorageFeatures{
			Create:    true,
			CreateDir: true,
			Delete:    true,
			List:      true,
			Metadata:  true,
			Move:      true,
			Read:      true,
			Stat:      true,
			Write:     true,
		},

		Create: []def.Pair{
			def.PairObjectMode,
		},
		Delete: []def.Pair{
			def.PairObjectMode,
		},
		List: []def.Pair{
			def.PairListMode,
			def.PairContinuationToken,
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
			def.PairIoCallback,
		},
	},
}
