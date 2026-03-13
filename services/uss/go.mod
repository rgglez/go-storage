module github.com/rgglez/go-storage/services/uss/v3

go 1.25

require (
	github.com/rgglez/go-storage/credential v1.0.0
	github.com/rgglez/go-storage/v5 v5.0.0
	github.com/upyun/go-sdk/v3 v3.0.4
)

require (
	github.com/Xuanwo/gg v0.3.0 // indirect
	github.com/Xuanwo/go-bufferpool v0.2.0 // indirect
	github.com/Xuanwo/templateutils v0.2.0 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	golang.org/x/sys v0.6.0 // indirect
)

replace (
	github.com/rgglez/go-storage/credential => ../../credential
	github.com/rgglez/go-storage/v5 => ../../
)
