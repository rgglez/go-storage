module github.com/rgglez/go-storage/examples

go 1.25

require (
	github.com/rgglez/go-storage/services/fs/v4 v4.0.0-00010101000000-000000000000
	github.com/rgglez/go-storage/services/memory v0.0.0-00010101000000-000000000000
	github.com/rgglez/go-storage/v5 v5.0.0
)

require (
	github.com/qingstor/go-mime v0.1.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
)

replace (
	github.com/rgglez/go-storage/services/fs/v4 => ../services/fs
	github.com/rgglez/go-storage/services/memory => ../services/memory
	github.com/rgglez/go-storage/v5 => ../
)
