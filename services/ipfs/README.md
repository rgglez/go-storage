[![Services Test Ipfs](https://github.com/rgglez/go-storage/actions/workflows/services-test-ipfs.yml/badge.svg)](https://github.com/rgglez/go-storage/actions/workflows/services-test-ipfs.yml)

# ipfs

[InterPlanetary File System(IPFS)](https://ipfs.io/) support for [go-storage](https://github.com/rgglez/go-storage).

## Install

```go
go get github.com/rgglez/go-storage/services/ipfs
```

## Usage

```go
import (
	"log"

	_ "github.com/rgglez/go-storage/services/ipfs"
	"github.com/rgglez/go-storage/v5/services"
)

func main() {
	store, err := services.NewStoragerFromString("ipfs:///path/to/workdir?endpoint=<ipfs_http_api_endpoint>&gateway=<ipfs_http_gateway>")
	if err != nil {
		log.Fatal(err)
	}
	
	// Write data from io.Reader into hello.txt
	n, err := store.Write("hello.txt", r, length)
}
```

- See more examples in [go-storage-example](https://github.com/rgglez/go-storage-example).
- Read [more docs](https://beyondstorage.io/docs/go-storage/services/ipfs) about go-service-ipfs.
