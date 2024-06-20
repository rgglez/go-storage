module github.com/rgglez/go-storage/services/bos/v2

go 1.16

require (
	github.com/baidubce/bce-sdk-go v0.9.183
	github.com/rgglez/go-storage/credential v1.0.0
	github.com/rgglez/go-storage/endpoint v1.2.0
	github.com/rgglez/go-storage/v5 v5.0.0
	github.com/google/uuid v1.6.0
)

replace (
	github.com/rgglez/go-storage/credential => ../../credential
	github.com/rgglez/go-storage/endpoint => ../../endpoint
	github.com/rgglez/go-storage/v5 => ../../
)
