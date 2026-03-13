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
	"fmt"
	"io"
	"math/rand"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"

	"github.com/rgglez/go-storage/v5/pkg/randbytes"
)

type storagePathSuite struct {
	suite.Suite

	p *StorageSuite

	base string
	path string
}

func (s *storagePathSuite) SetupTest() {
	s.base = uuid.NewString()
	s.path = uuid.NewString()
}

func (s *storagePathSuite) TearDownTest() {
	path := fmt.Sprintf("%s/%s", s.base, s.path)

	err := s.p.store.Delete(path)
	s.NoError(err)
}

func (s *storagePathSuite) TestAbsPath() {
	m := s.p.store.Metadata()

	path := fmt.Sprintf("%s%s/%s", m.WorkDir, s.base, s.path)

	size := rand.Int63n(4 * 1024 * 1024)
	content, err := io.ReadAll(io.LimitReader(randbytes.NewRand(), size))
	s.NoError(err)

	_, err = s.p.store.Write(path, bytes.NewReader(content), size)
	s.NoError(err)

	var buf bytes.Buffer

	n, err := s.p.store.Read(path, &buf)
	s.NoError(err)
	s.Equal(size, n)
}

func (s *storagePathSuite) TestBackslash() {
	path := s.base + "\\" + s.path

	size := rand.Int63n(4 * 1024 * 1024)
	content, err := io.ReadAll(io.LimitReader(randbytes.NewRand(), size))
	s.NoError(err)

	_, err = s.p.store.Write(path, bytes.NewReader(content), size)
	s.NoError(err)

	var buf bytes.Buffer

	n, err := s.p.store.Read(path, &buf)
	s.NoError(err)
	s.Equal(size, n)
}
