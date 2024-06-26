[![Services Test Hdfs](https://github.com/rgglez/go-storage/actions/workflows/services-test-hdfs.yml/badge.svg)](https://github.com/rgglez/go-storage/actions/workflows/services-test-hdfs.yml)

# hdfs 

Hadoop Distributed File System (HDFS) support for [go-storage](https://github.com/rgglez/go-storage).

## Install

```go
go get go.beyondstorage.io/services/hdfs
```

## Usage

```go
import (
	"log"
	
	_ "go.beyondstorage.io/services/hdfs"
	"go.beyondstorage.io/v5/services"
)

func main() {
	store, err := services.NewStoragerFromString("hdfs:///path/to/workdir?endpoint=tcp:<host>:<port>")
	if err != nil {
		log.Fatal(err)
	}
	
	// Write data from io.Reader into hello.txt
	n, err := store.Write("hello.txt", r, length)
}
```

- See more examples in [go-storage-example](https://github.com/rgglez/go-storage-example).
- Read [more docs](https://beyondstorage.io/docs/go-storage/services/hdfs) about go-service-hdfs.