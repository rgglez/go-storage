module github.com/rgglez/go-storage/services/gcs/v3

go 1.16

require (
	cloud.google.com/go/storage v1.29.0
	github.com/rgglez/go-storage/credential v1.0.0
	github.com/rgglez/go-storage/v5 v5.0.0
	github.com/google/uuid v1.4.0
	golang.org/x/oauth2 v0.0.0-20221014153046-6fdb5e3db783
	google.golang.org/api v0.109.0
)

replace (
	github.com/rgglez/go-storage/credential => ../../credential
	github.com/rgglez/go-storage/v5 => ../../
)
