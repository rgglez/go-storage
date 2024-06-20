module github.com/rgglez/go-storage/services/storj

go 1.16

require (
	github.com/rgglez/go-storage/credential v1.0.0
	github.com/rgglez/go-storage/v5 v5.0.0
	github.com/google/uuid v1.6.0
	storj.io/uplink v1.13.0
)

replace (
	github.com/rgglez/go-storage/credential => ../../credential
	github.com/rgglez/go-storage/v5 => ../../
)
