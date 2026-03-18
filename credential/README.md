# credential

Both human and machine-readable credential format.

## Format

```
<protocol>:<value>+
```

For example:

- hmac: `hmac:access_key:secret_key`
- apikey: `apikey:apikey`
- file: `file:/path/to/config/file`
- basic: `basic:user:password`

## Quick Start

```go
cred, err := credential.Parse("hmac:access_key:secret_key)
if err != nil {
    log.Fatal("parse: ", err)
}

switch cred.Protocol() {
case ProtocolHmac:
    ak, sk := cred.Hmac()
    log.Println("access_key: ", ak)
    log.Println("secret_key: ", sk)
case ProtocolAPIKey:
    apikey := cred.APIKey()
    log.Println("apikey: ", apikey)
case ProtocolFile:
    path := cred.File()
    log.Println("path: ", path)
case ProtocolEnv:
    log.Println("use env value")
case ProtocolBase64:
    content := cred.Base64()
    log.Println("base64: ", content)
case ProtocolBasic:
    user, password := cred.Basic()
    log.Println("user: ", user)
    log.Println("password: ", password)
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
