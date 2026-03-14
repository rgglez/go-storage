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
	"context"
	"net"
	"time"
)

// Dialer is the dialer the storage used for stream-oriented network connection.
type Dialer struct {
	*net.Dialer
	readTimeout  time.Duration
	writeTimeout time.Duration
}

// WithConnectTimeout will configure dialer's timeout
func (d *Dialer) WithConnectTimeout(t time.Duration) *Dialer {
	d.Dialer.Timeout = t
	return d
}

// WithReadTimeout will configure underlying conn's read timeout
func (d *Dialer) WithReadTimeout(t time.Duration) *Dialer {
	d.readTimeout = t
	return d
}

// WithWriteTimeout will configure underlying conn's write timeout
func (d *Dialer) WithWriteTimeout(t time.Duration) *Dialer {
	d.writeTimeout = t
	return d
}

// DialContext connects to the address on the named network using
// the provided context.
func (d *Dialer) DialContext(ctx context.Context, network, addr string) (net.Conn, error) {
	c, err := d.Dialer.DialContext(ctx, network, addr)
	if err != nil {
		return nil, err
	}

	conn := &Conn{
		Conn:         c,
		readTimeout:  d.readTimeout,
		writeTimeout: d.writeTimeout,
	}
	return conn, nil
}

// NewDialer will create a new dialer with preset default value:
//   - dialer connect timeout will be 60s
//   - underlying conn read timeout will be 30s
//   - underlying conn write timeout will be 30s
func NewDialer() *Dialer {
	d := &net.Dialer{
		Timeout: time.Minute,
		Resolver: &net.Resolver{
			// Use go builtin dns resolver instead of cgo
			PreferGo: true,
		},
	}

	return &Dialer{d, 30 * time.Second, 30 * time.Second}
}
