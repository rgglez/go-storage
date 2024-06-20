module github.com/rgglez/go-storage/services/uss/v3

go 1.16

require (
	github.com/rgglez/go-storage/credential v1.0.0
	github.com/rgglez/go-storage/v5 v5.0.0
	github.com/upyun/go-sdk/v3 v3.0.4
)

replace (
	github.com/rgglez/go-storage/credential => ../../credential
	github.com/rgglez/go-storage/v5 => ../../
)
