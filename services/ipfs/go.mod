module github.com/rgglez/go-storage/services/ipfs

go 1.16

require (
	github.com/rgglez/go-storage/endpoint v1.2.0
	github.com/rgglez/go-storage/v5 v5.0.0
	github.com/google/uuid v1.3.0
	github.com/ipfs/go-ipfs-api v0.6.0
	github.com/ipfs/go-ipfs-cmds v0.6.0
)

replace (
	github.com/rgglez/go-storage/endpoint => ../../endpoint
	github.com/rgglez/go-storage/v5 => ../../
)
