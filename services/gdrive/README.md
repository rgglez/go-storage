[![Services Test Gdrive](https://github.com/rgglez/go-storage/actions/workflows/services-test-gdrive.yml/badge.svg)](https://github.com/rgglez/go-storage/actions/workflows/services-test-gdrive.yml)

# gdrive

Google Drive service support for [go-storage](https://github.com/rgglez/go-storage).

## Install

```go
go get github.com/rgglez/go-storage/services/gdrive
```

## Usage

```go
import (
	"log"

	_ "github.com/rgglez/go-storage/services/gdrive"
	"github.com/rgglez/go-storage/v5/services"
)

func main() {
	store, err := services.NewStoragerFromString("gdrive://path/to/work_dir?name=<a_meaningful_name>?credential=file:<absolute_path_to_credentials>")
	if err != nil {
		log.Fatal(err)
	}

	// Write data from io.Reader into hello.txt
	n, err := store.Write("hello.txt", r, length)
}
```

- See more examples in [go-storage-example](https://github.com/rgglez/go-storage-example).
- Read [more docs](https://beyondstorage.io/docs/go-storage/services/gdrive) about go-service-gdrive.
