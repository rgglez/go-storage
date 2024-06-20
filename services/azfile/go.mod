module github.com/rgglez/go-storage/services/azfile

go 1.16

require (
	github.com/Azure/azure-storage-file-go v0.8.0
	github.com/rgglez/go-storage/credential v1.0.0
	github.com/rgglez/go-storage/endpoint v1.2.0
	github.com/rgglez/go-storage/v5 v5.0.0
	github.com/google/uuid v1.4.0
	github.com/pkg/errors v0.9.1 // indirect
)

replace (
	github.com/rgglez/go-storage/credential => ../../credential
	github.com/rgglez/go-storage/endpoint => ../../endpoint
	github.com/rgglez/go-storage/v5 => ../../
)
