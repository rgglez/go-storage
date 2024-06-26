module github.com/rgglez/go-storage/services/tar

go 1.16

require (
	github.com/rgglez/go-storage/endpoint v1.2.1
	github.com/rgglez/go-storage/v5 v5.0.0
	github.com/stretchr/testify v1.9.0
)

replace (
	github.com/rgglez/go-storage/endpoint => ../../endpoint
	github.com/rgglez/go-storage/v5 => ../../
)
