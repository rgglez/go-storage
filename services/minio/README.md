[![Services Test Minio](https://github.com/rgglez/go-storage/actions/workflows/services-test-minio.yml/badge.svg)](https://github.com/rgglez/go-storage/actions/workflows/services-test-minio.yml)

# minio

[MinIO](https://min.io/) is an open source cloud-native high-performance object storage service. 
This project will use minio's native SDK to implement [go-storage](https://github.com/rgglez/go-storage/), 
enabling users to manipulate data on minio servers through a unified interface.

## Install

```go
go get github.com/rgglez/go-storage/services/minio
```

## Usage

```go
import (
	"log"

	_ "github.com/rgglez/go-storage/services/minio"
	"github.com/rgglez/go-storage/v5/services"
)

func main() {
	store, err := services.NewStoragerFromString("minio://<bucket_name>/<work_dir>?credential=hmac:<access_key>:<secret_key>&endpoint=https:<host>:<port>")
	if err != nil {
		log.Fatal(err)
	}
	
	// Write data from io.Reader into hello.txt
	n, err := store.Write("hello.txt", r, length)
}
```

- See more examples in [go-storage-example](https://github.com/rgglez/go-storage-example).
- Read [more docs](https://beyondstorage.io/docs/go-storage/services/minio) about go-service-minio.


## License

Copyright 2024 go-storage authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
