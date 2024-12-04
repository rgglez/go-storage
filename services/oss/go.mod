module github.com/rgglez/go-storage/services/oss/v3

go 1.16

require (
	github.com/aliyun/aliyun-oss-go-sdk v3.0.1+incompatible
	github.com/google/uuid v1.6.0
	github.com/rgglez/go-storage/credential v1.0.0
	github.com/rgglez/go-storage/endpoint v1.2.1
	github.com/rgglez/go-storage/v5 v5.0.0
)

replace (
	github.com/rgglez/go-storage/credential => ../../credential
	github.com/rgglez/go-storage/endpoint => ../../endpoint
	github.com/rgglez/go-storage/v5 => ../../
)
