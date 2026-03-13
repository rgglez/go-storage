module github.com/rgglez/go-storage/services/kodo/v3

go 1.25

require (
	github.com/google/uuid v1.6.0
	github.com/qiniu/go-sdk/v7 v7.21.1
	github.com/rgglez/go-storage/credential v1.0.0
	github.com/rgglez/go-storage/endpoint v1.2.1
	github.com/rgglez/go-storage/v5 v5.0.0
)

require (
	github.com/BurntSushi/toml v1.3.2 // indirect
	github.com/Xuanwo/gg v0.3.0 // indirect
	github.com/Xuanwo/go-bufferpool v0.2.0 // indirect
	github.com/Xuanwo/templateutils v0.2.0 // indirect
	github.com/alex-ant/gomath v0.0.0-20160516115720-89013a210a82 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gofrs/flock v0.8.1 // indirect
	github.com/gopherjs/gopherjs v1.17.2 // indirect
	github.com/jtolds/gls v4.20.0+incompatible // indirect
	github.com/matishsiao/goInfo v0.0.0-20210923090445-da2e3fa8d45f // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/smarty/assertions v1.15.0 // indirect
	github.com/smartystreets/goconvey v1.8.1 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4 // indirect
	golang.org/x/sys v0.12.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/rgglez/go-storage/credential => ../../credential
	github.com/rgglez/go-storage/endpoint => ../../endpoint
	github.com/rgglez/go-storage/v5 => ../../
)
