## How run integration tests

### Run tests locally

Copy example files and update corresponding values.

```shell
cp Makefile.env.exmaple Makefile.env
```

Run tests

```shell
make integration_test
```

### Run tests in CI

Set following environment variables:

```shell
export STORAGE_KODO_INTEGRATION_TEST=on
export STORAGE_KODO_CREDENTIAL=hmac:access_key:secret_key
export STORAGE_KODO_NAME=bucketname
export STORAGE_KODO_ENDPOINT=endpoint
```

Run tests

```shell
make integration_test
```


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
