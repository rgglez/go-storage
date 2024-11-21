module fs

go 1.16

require (
	github.com/google/uuid v1.6.0
	github.com/qingstor/go-mime v0.1.0
	github.com/stretchr/testify v1.9.0
	github.com/rgglez/go-storage/v5 v5.0.0
	golang.org/x/sys v0.20.0
)

replace (
        github.com/rgglez/go-storage/v5 => ../../
)
