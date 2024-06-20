module github.com/rgglez/go-storage/services/gdrive

go 1.16

require (
	github.com/rgglez/go-storage/credential v1.0.0
	github.com/rgglez/go-storage/v5 v5.0.0
	github.com/dgraph-io/ristretto v0.1.1
	github.com/google/uuid v1.6.0
	golang.org/x/oauth2 v0.21.0
	google.golang.org/api v0.183.0
)

replace (
	github.com/rgglez/go-storage/credential => ../../credential
	github.com/rgglez/go-storage/v5 => ../../
)
