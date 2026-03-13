module github.com/rgglez/go-storage/services/azfile

go 1.25

require (
	github.com/Azure/azure-storage-file-go v0.8.0
	github.com/google/uuid v1.6.0
	github.com/rgglez/go-storage/credential v1.0.0
	github.com/rgglez/go-storage/endpoint v1.2.1
	github.com/rgglez/go-storage/v5 v5.0.0
)

require (
	github.com/Azure/azure-pipeline-go v0.2.1 // indirect
	github.com/Xuanwo/gg v0.3.0 // indirect
	github.com/Xuanwo/go-bufferpool v0.2.0 // indirect
	github.com/Xuanwo/templateutils v0.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gopherjs/gopherjs v1.17.2 // indirect
	github.com/jtolds/gls v4.20.0+incompatible // indirect
	github.com/mattn/go-ieproxy v0.0.0-20190610004146-91bb50d98149 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/smarty/assertions v1.15.0 // indirect
	github.com/smartystreets/goconvey v1.8.1 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	golang.org/x/net v0.0.0-20191209160850-c0dbc17a3553 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.3.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/rgglez/go-storage/credential => ../../credential
	github.com/rgglez/go-storage/endpoint => ../../endpoint
	github.com/rgglez/go-storage/v5 => ../../
)
