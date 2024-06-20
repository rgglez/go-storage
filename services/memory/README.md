[![Services Test Memory](https://github.com/rgglez/go-storage/actions/workflows/services-test-memory.yml/badge.svg)](https://github.com/rgglez/go-storage/actions/workflows/services-test-memory.yml)

# memory

memory service support for [go-storage](https://github.com/rgglez/go-storage).

## Install

```go
go get github.com/rgglez/go-storage/services/memory
```

## Usage

```go
import (
	"log"

	_ "github.com/rgglez/go-storage/services/memory"
	"github.com/rgglez/go-storage/v5/services"
)

func main() {
	store, err := services.NewStoragerFromString("memory:///path/to/workdir")
	if err != nil {
		log.Fatal(err)
	}
	
	// Write data from io.Reader into hello.txt
	n, err := store.Write("hello.txt", r, length)
}
```

- See more examples in [go-storage-example](https://github.com/rgglez/go-storage-example).
