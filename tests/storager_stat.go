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
package tests

import (
	"bytes"
	"io"
	"math/rand"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"

	"github.com/rgglez/go-storage/v5/pkg/randbytes"
)

type storageStatSuite struct {
	suite.Suite

	p *StorageSuite

	size    int64
	content []byte
	path    string
}

func (s *storageStatSuite) SetupTest() {
	var err error

	s.size = rand.Int63n(4 * 1024 * 1024) // Max file size is 4MB
	s.content, err = io.ReadAll(io.LimitReader(randbytes.NewRand(), s.size))
	s.NoError(err)

	s.path = uuid.New().String()

	_, err = s.p.store.Write(s.path, bytes.NewReader(s.content), s.size)
	s.NoError(err)
}

func (s *storageStatSuite) TearDownTest() {
	err := s.p.store.Delete(s.path)
	s.NoError(err)
}

func (s *storageStatSuite) TestStat() {
	o, err := s.p.store.Stat(s.path)
	s.NoError(err)
	s.NotNil(o)

	osize, ok := o.GetContentLength()
	s.True(ok)
	s.Equal(osize, s.size)
}
