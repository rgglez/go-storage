[![Services Test Dropbox](https://github.com/rgglez/go-storage/actions/workflows/services-test-dropbox.yml/badge.svg)](https://github.com/rgglez/go-storage/actions/workflows/services-test-dropbox.yml)

# dropbox

[Dropbox](https://www.dropbox.com) service support for [go-storage](https://github.com/rgglez/go-storage).

## Install

```go
go get github.com/rgglez/go-storage/services/dropbox/v3
```

## Usage

```go
import (
	"log"

	_ "github.com/rgglez/go-storage/services/dropbox/v3"
	"github.com/rgglez/go-storage/v5/services"
)

func main() {
	store, err := services.NewStoragerFromString("dropbox:///path/to/workdir?credential=apikey:<apikey>")
	if err != nil {
		log.Fatal(err)
	}
	
	// Write data from io.Reader into hello.txt
	n, err := store.Write("hello.txt", r, length)
}
```

- See more examples in [go-storage-example](https://github.com/rgglez/go-storage-example).
- Read [more docs](https://beyondstorage.io/docs/go-storage/services/dropbox) about go-service-dropbox.
