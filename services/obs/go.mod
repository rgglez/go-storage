module github.com/rgglez/go-storage/services/obs/v2

go 1.16

require (
	github.com/rgglez/go-storage/credential v1.0.0
	github.com/rgglez/go-storage/endpoint v1.2.1
	github.com/rgglez/go-storage/v5 v5.0.0
	github.com/google/uuid v1.4.0
	github.com/huaweicloud/huaweicloud-sdk-go-obs v3.23.9+incompatible
)

replace (
	github.com/rgglez/go-storage/credential => ../../credential
	github.com/rgglez/go-storage/endpoint => ../../endpoint
	github.com/rgglez/go-storage/v5 => ../../
)
