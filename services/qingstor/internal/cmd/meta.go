package main

import (
	def "github.com/rgglez/go-storage/v5/definitions"
	"github.com/rgglez/go-storage/v5/types"
)

var Metadata = def.Metadata{
	Name: "qingstor",
	Imports: []string{
		"github.com/rgglez/go-storage/v5/pkg/httpclient",
	},
	Pairs: []def.Pair{
		pairEncryptionCustomerAlgorithm,
		pairEncryptionCustomerKey,
		pairCopySourceEncryptionCustomerAlgorithm,
		pairCopySourceEncryptionCustomerKey,
		pairDisableUriCleaning,
		pairStorageClass,
		pairHTTPClientOptions,
	},
	Infos: []def.Info{
		infoObjectMetaStorageClass,
		infoObjectMetaEncryptionCustomerAlgorithm,
	},
	Factory: []def.Pair{
		def.PairCredential,
		def.PairEndpoint,
		pairHTTPClientOptions,
		def.PairName,
		pairDisableUriCleaning,
		def.PairLocation,
		def.PairWorkDir,
	},
	Service: def.Service{
		Features: types.ServiceFeatures{
			Create: true,
			Delete: true,
			Get:    true,
			List:   true,
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
			VirtualDir:  true,
			VirtualLink: true,

			CommitAppend:       true,
			CompleteMultipart:  true,
			Copy:               true,
			Create:             true,
			CreateAppend:       true,
			CreateDir:          true,
			CreateLink:         true,
			CreateMultipart:    true,
			Delete:             true,
			Fetch:              true,
			List:               true,
			ListMultipart:      true,
			Metadata:           true,
			Move:               true,
			QuerySignHTTPRead:  true,
			QuerySignHTTPWrite: true,
			Read:               true,
			Stat:               true,
			Write:              true,
			WriteAppend:        true,
			WriteMultipart:     true,
		},

		Create: []def.Pair{
			def.PairMultipartID,
			def.PairObjectMode,
		},
		CreateDir: []def.Pair{
			pairStorageClass,
		},
		Delete: []def.Pair{
			def.PairMultipartID,
			def.PairObjectMode,
		},
		Stat: []def.Pair{
			def.PairMultipartID,
			def.PairObjectMode,
		},
		List: []def.Pair{
			def.PairListMode,
		},
		Read: []def.Pair{
			def.PairOffset,
			def.PairIoCallback,
			def.PairSize,
			pairEncryptionCustomerAlgorithm,
			pairEncryptionCustomerKey,
		},
		Write: []def.Pair{
			def.PairContentMD5,
			def.PairContentType,
			def.PairIoCallback,
			pairStorageClass,
			pairEncryptionCustomerAlgorithm,
			pairEncryptionCustomerKey,
		},
		CreateAppend: []def.Pair{
			def.PairContentType,
			pairStorageClass,
		},
		WriteAppend: []def.Pair{
			def.PairContentMD5,
		},
		Copy: []def.Pair{
			pairEncryptionCustomerAlgorithm,
			pairEncryptionCustomerKey,
			pairCopySourceEncryptionCustomerAlgorithm,
			pairCopySourceEncryptionCustomerKey,
		},
		CreateMultipart: []def.Pair{
			pairEncryptionCustomerAlgorithm,
			pairEncryptionCustomerKey,
		},
		WriteMultipart: []def.Pair{
			pairEncryptionCustomerAlgorithm,
			pairEncryptionCustomerKey,
			def.PairIoCallback,
		},
		QuerySignHTTPRead: []def.Pair{
			def.PairOffset,
			pairEncryptionCustomerAlgorithm,
			pairEncryptionCustomerKey,
			def.PairSize,
		},
		QuerySignHTTPWrite: []def.Pair{
			def.PairContentMD5,
			def.PairContentType,
			pairEncryptionCustomerAlgorithm,
			pairEncryptionCustomerKey,
			pairStorageClass,
		},
	},
}

var pairEncryptionCustomerAlgorithm = def.Pair{
	Name:        "encryption_customer_algorithm",
	Type:        def.Type{Name: "string"},
	Description: "specifies the encryption algorithm. Only AES256 is supported now.",
}
var pairEncryptionCustomerKey = def.Pair{
	Name:        "encryption_customer_key",
	Type:        def.Type{Expr: "[]", Name: "byte"},
	Description: "is the customer-provided encryption key. For AES256 keys, the plaintext must be 32 bytes long.",
}
var pairCopySourceEncryptionCustomerAlgorithm = def.Pair{
	Name:        "copy_source_encryption_customer_algorithm",
	Type:        def.Type{Name: "string"},
	Description: "is the encryption algorithm for the source object. Only AES256 is supported now.",
}
var pairCopySourceEncryptionCustomerKey = def.Pair{
	Name:        "copy_source_encryption_customer_key",
	Type:        def.Type{Expr: "[]", Name: "byte"},
	Description: "is the customer-provided encryption key for the source object. For AES256 keys, the plaintext must be 32 bytes long.",
}
var pairDisableUriCleaning = def.Pair{
	Name: "disable_uri_cleaning",
	Type: def.Type{Name: "bool"},
}
var pairStorageClass = def.Pair{
	Name: "storage_class",
	Type: def.Type{Name: "string"},
}
var pairHTTPClientOptions = def.Pair{
	Name: "http_client_options",
	Type: def.Type{Expr: "*", Package: "httpclient", Name: "Options"},
}

var infoObjectMetaStorageClass = def.Info{
	Namespace: def.NamespaceObject,
	Category:  def.CategoryMeta,
	Name:      "storage_class",
	Type:      def.Type{Name: "string"},
}
var infoObjectMetaEncryptionCustomerAlgorithm = def.Info{
	Namespace: def.NamespaceObject,
	Category:  def.CategoryMeta,
	Name:      "encryption_customer_algorithm",
	Type:      def.Type{Name: "string"},
}
