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
package httpclient

import (
	"net/http"
	"time"
)

// Options is the httpclient supported options.
type Options struct {
	// Dialer related options
	DialConnectTimeout time.Duration

	// Underlying connection related options
	ConnReadTimeout  time.Duration
	ConnWriteTimeout time.Duration

	// HTTP client related options
}

// New will create new http client.
func New(o *Options) *http.Client {
	dialer := NewDialer()

	if o != nil {
		if o.DialConnectTimeout > 0 {
			dialer.WithConnectTimeout(o.DialConnectTimeout)
		}
		if o.ConnReadTimeout > 0 {
			dialer.WithReadTimeout(o.ConnReadTimeout)
		}
		if o.ConnWriteTimeout > 0 {
			dialer.WithWriteTimeout(o.ConnWriteTimeout)
		}
	}

	hc := &http.Client{
		Transport: &http.Transport{
			DialContext: dialer.DialContext,

			// Support http proxy from env.
			Proxy: http.ProxyFromEnvironment,
			// Specify timeout for tls handshake.
			TLSHandshakeTimeout: 10 * time.Second,
			// Specify max idle conns across all hosts.
			MaxIdleConns: 0,
			// Specify max idle conns across per host.
			MaxIdleConnsPerHost: 100,
			// Specify timeout for closing idle (keep-alive) connection.
			IdleConnTimeout: 90 * time.Second,
			// Specify timeout that waiting for server's approve before sending data.
			ExpectContinueTimeout: time.Second,
			// Gzip file should not be auto-decompressed
			DisableCompression: true,
		},
		// http client used in storage don't need to follow redirect, return directly.
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		// We will handle timeout by ourselves, and disable http.Client's timeout.
		Timeout: 0,
	}
	return hc
}
