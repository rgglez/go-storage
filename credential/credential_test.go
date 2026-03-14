// Copyright The go-storage Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package credential

import (
	"errors"
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestProvider(t *testing.T) {
	protocol := uuid.New().String()
	args := []string{uuid.New().String(), uuid.New().String()}

	p := Credential{protocol: protocol, args: args}

	assert.Equal(t, protocol, p.Protocol())
	assert.EqualValues(t, args, p.Value())
}

func TestParse(t *testing.T) {
	cases := []struct {
		name  string
		cfg   string
		value Credential
		err   error
	}{
		{
			"hmac",
			"hmac:ak:sk",
			Credential{protocol: ProtocolHmac, args: []string{"ak", "sk"}},
			nil,
		},
		{
			"api key",
			"apikey:key",
			Credential{protocol: ProtocolAPIKey, args: []string{"key"}},
			nil,
		},
		{
			"file",
			"file:/path/to/file",
			Credential{protocol: ProtocolFile, args: []string{"/path/to/file"}},
			nil,
		},
		{
			"env",
			"env",
			Credential{protocol: ProtocolEnv},
			nil,
		},
		{
			"base64",
			"base64:aGVsbG8sd29ybGQhCg==",
			Credential{protocol: ProtocolBase64, args: []string{"aGVsbG8sd29ybGQhCg=="}},
			nil,
		},
		{
			"basic",
			"basic:user:password",
			Credential{protocol: ProtocolBasic, args: []string{"user", "password"}},
			nil,
		},
		{
			"not supported protocol",
			"notsupported:ak:sk",
			Credential{},
			ErrUnsupportedProtocol,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			p, err := Parse(tt.cfg)
			if tt.err == nil {
				assert.Nil(t, err)
			} else {
				assert.True(t, errors.Is(err, tt.err))
			}
			assert.EqualValues(t, tt.value, p)
		})
	}
}

func ExampleParse() {
	cred, err := Parse("hmac:access_key:secret_key")
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
}
