[![Services Test Kodo](https://github.com/rgglez/go-storage/actions/workflows/services-test-kodo.yml/badge.svg)](https://github.com/rgglez/go-storage/actions/workflows/services-test-kodo.yml)

# kodo

[qiniu kodo](https://www.qiniu.com/products/kodo) service support for [go-storage](https://github.com/rgglez/go-storage).

## Install

```go
go get github.com/rgglez/go-storage/services/kodo/v3
```

## Usage

```go
import (
	"log"

	_ "github.com/rgglez/go-storage/services/kodo/v3"
	"github.com/rgglez/go-storage/v5/services"
)

func main() {
	store, err := services.NewStoragerFromString("kodo://bucket_name/path/to/workdir?credential=hmac:<access_key>:<secret_key>&endpoint=http:<domain>")
	if err != nil {
		log.Fatal(err)
	}
	
	// Write data from io.Reader into hello.txt
	n, err := store.Write("hello.txt", r, length)
}
```

- See more examples in [go-storage-example](https://github.com/rgglez/go-storage-example).
- Read [more docs](https://beyondstorage.io/docs/go-storage/services/kodo) about go-service-kodo.
