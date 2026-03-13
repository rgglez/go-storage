// Copyright The go-storage Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package httpclient

import (
	"net"
	"time"
)

// Conn is a generic stream-oriented network connection.
type Conn struct {
	net.Conn
	readTimeout  time.Duration
	writeTimeout time.Duration
}

// Read will read from the conn.
func (c *Conn) Read(buf []byte) (n int, err error) {
	err = c.SetDeadline(time.Now().Add(c.readTimeout))
	if err != nil {
		return
	}
	defer func() {
		// Clean read timeout so that this will not affect further read
		// It's safe to ignore the returning error: even if it don’t return now, it will return via next read.
		_ = c.SetDeadline(time.Time{})
	}()

	return c.Conn.Read(buf)
}

// Write will write into the conn.
func (c *Conn) Write(buf []byte) (n int, err error) {
	err = c.SetDeadline(time.Now().Add(c.writeTimeout))
	if err != nil {
		return
	}
	defer func() {
		// Clean read timeout so that this will not affect further write
		// It's safe to ignore the returning error: even if it don’t return now, it will return via next write.
		_ = c.SetDeadline(time.Time{})
	}()

	return c.Conn.Write(buf)
}
