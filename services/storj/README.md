[![Services Test Storj](https://github.com/rgglez/go-storage/actions/workflows/services-test-storj.yml/badge.svg)](https://github.com/rgglez/go-storage/actions/workflows/services-test-storj.yml)

# storj

[Storj DCS] (Decentralized Cloud Storage) support for [go-storage].

[Storj DCS]: https://www.storj.io
[go-storage]: https://github.com/rgglez/go-storage

## Install

```go
go get github.com/rgglez/go-storage/services/storj
```

## Usage

```go
import (
	"log"

	_ "github.com/rgglez/go-storage/services/storj"
	"github.com/rgglez/go-storage/v5/services"
)
s
func main() {
	store, err := services.NewStoragerFromString("storj://bucket_name/path/to/workdir")
	if err != nil {
		log.Fatal(err)
	}
	
	// Write data from io.Reader into hello.txt
	n, err := store.Write("hello.txt", r, length)
}
```

- See more examples in [go-storage-example](https://github.com/rgglez/go-storage-example).
- Read [more docs](https://beyondstorage.io/docs/go-storage/services/storj) about go-service-storj.