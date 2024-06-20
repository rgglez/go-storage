module github.com/rgglez/go-storage/services/cos/v3

go 1.16

require (
	github.com/rgglez/go-storage/credential v1.0.0
	github.com/rgglez/go-storage/v5 v5.0.0
	github.com/google/uuid v1.4.0
	github.com/tencentyun/cos-go-sdk-v5 v0.7.50
)

replace (
	github.com/rgglez/go-storage/credential => ../../credential
	github.com/rgglez/go-storage/v5 => ../../
)
