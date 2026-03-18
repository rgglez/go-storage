# endpoint

Both human and machine readable endpoint format.

## Format

```
<protocol>:<value>+
```

For example:

- File: `file:/var/cache/data`
- HTTP: `http:example.com:80`
- HTTPS: `https:example.com:443`

## Quick Start

```go
ep, err := endpoint.Parse("https:example.com")
if err != nil {
	log.Fatal("parse: ", err)
}

switch ep.Protocol() {
case ProtocolHTTP:
    url, host, port := ep.HTTP()
    log.Println("url: ", url)
    log.Println("host: ", host)
    log.Println("port: ", port)
case ProtocolHTTPS:
    url, host, port := ep.HTTPS()
    log.Println("url: ", url)
    log.Println("host: ", host)
    log.Println("port: ", port)
case ProtocolFile:
    path := ep.File()
    log.Println("path: ", path)
default:
    panic("unsupported protocol")
}
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
