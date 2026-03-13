module github.com/rgglez/go-storage/examples

go 1.25.0

require (
	github.com/rgglez/go-storage/services/fs/v4 v4.0.0-00010101000000-000000000000
	github.com/rgglez/go-storage/services/kodo/v3 v3.0.0-00010101000000-000000000000
	github.com/rgglez/go-storage/services/memory v0.0.0-00010101000000-000000000000
	github.com/rgglez/go-storage/services/minio v0.0.0-00010101000000-000000000000
	github.com/rgglez/go-storage/services/oss/v3 v3.0.0-00010101000000-000000000000
	github.com/rgglez/go-storage/services/s3/v3 v3.0.0-00010101000000-000000000000
	github.com/rgglez/go-storage/v5 v5.0.0
)

require (
	github.com/BurntSushi/toml v1.3.2 // indirect
	github.com/alex-ant/gomath v0.0.0-20160516115720-89013a210a82 // indirect
	github.com/aliyun/aliyun-oss-go-sdk v3.0.1+incompatible // indirect
	github.com/aws/aws-sdk-go-v2 v1.28.0 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.6.2 // indirect
	github.com/aws/aws-sdk-go-v2/config v1.27.19 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.17.19 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.16.6 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.10 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.10 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.0 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.3.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.11.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.3.12 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.11.12 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.17.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/s3 v1.55.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.20.12 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.24.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.28.13 // indirect
	github.com/aws/smithy-go v1.20.2 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/gofrs/flock v0.8.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/klauspost/compress v1.17.6 // indirect
	github.com/klauspost/cpuid/v2 v2.2.6 // indirect
	github.com/matishsiao/goInfo v0.0.0-20210923090445-da2e3fa8d45f // indirect
	github.com/minio/md5-simd v1.1.2 // indirect
	github.com/minio/minio-go/v7 v7.0.70 // indirect
	github.com/qingstor/go-mime v0.1.0 // indirect
	github.com/qiniu/go-sdk/v7 v7.21.1 // indirect
	github.com/rgglez/go-storage/credential v1.0.0 // indirect
	github.com/rgglez/go-storage/endpoint v1.2.1 // indirect
	github.com/rs/xid v1.5.0 // indirect
	golang.org/x/crypto v0.21.0 // indirect
	golang.org/x/net v0.23.0 // indirect
	golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/time v0.15.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
)

replace (
	github.com/rgglez/go-storage/services/fs/v4 => ../services/fs
	github.com/rgglez/go-storage/services/kodo/v3 => ../services/kodo
	github.com/rgglez/go-storage/services/memory => ../services/memory
	github.com/rgglez/go-storage/services/minio => ../services/minio
	github.com/rgglez/go-storage/services/oss/v3 => ../services/oss
	github.com/rgglez/go-storage/services/s3/v3 => ../services/s3
	github.com/rgglez/go-storage/v5 => ../
)
