[![Services Test Fs](https://github.com/rgglez/go-storage/actions/workflows/services-test-fs.yml/badge.svg)](https://github.com/rgglez/go-storage/actions/workflows/services-test-fs.yml)

# fs

Local file system service support for [go-storage](https://github.com/rgglez/go-storage).

## Install

```go
go get github.com/rgglez/go-storage/services/fs/v4
```

## Usage

```go
import (
	"log"

	_ "github.com/rgglez/go-storage/services/fs/v4"
	"github.com/rgglez/go-storage/v5/services"
)

func main() {
	store, err := services.NewStoragerFromString("fs:///path/to/workdir")
	if err != nil {
		log.Fatal(err)
	}
	
	// Write data from io.Reader into hello.txt
	n, err := store.Write("hello.txt", r, length)
}
```

- See more examples in [go-storage-example](https://github.com/rgglez/go-storage-example).
- Read [more docs](https://beyondstorage.io/docs/go-storage/services/fs) about go-service-fs.
