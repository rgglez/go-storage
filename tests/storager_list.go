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
	"errors"
	"fmt"
	"io"
	"math/rand"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"

	ps "github.com/rgglez/go-storage/v5/pairs"
	"github.com/rgglez/go-storage/v5/pkg/randbytes"
	"github.com/rgglez/go-storage/v5/types"
)

type storageListSuite struct {
	suite.Suite

	p *StorageSuite

	base   string
	length int
	paths  []string
}

func (s *storageListSuite) SetupTest() {
	size := rand.Int63n(256)

	s.length = rand.Intn(16)
	s.base = uuid.NewString()
	s.paths = make([]string, s.length)

	for i := 0; i < s.length; i++ {
		s.paths[i] = fmt.Sprintf("%s/%s", s.base, uuid.NewString())

		_, err := s.p.store.Write(s.paths[i],
			io.LimitReader(randbytes.NewRand(), size), size)
		s.NoError(err)
	}
}

func (s *storageListSuite) TearDownTest() {
	for i := 0; i < s.length; i++ {
		err := s.p.store.Delete(s.paths[i])
		s.NoError(err)
	}
}

func (s *storageListSuite) TestList() {
	it, err := s.p.store.List(s.base, ps.WithListMode(types.ListModeDir))
	s.NoError(err)
	s.NotNil(it)

	paths := make([]string, 0)
	for {
		o, err := it.Next()
		if errors.Is(err, types.IterateDone) {
			break
		}
		s.NoError(err)

		paths = append(paths, o.Path)
	}
	s.ElementsMatch(s.paths, paths)
}

func (s *storageListSuite) TestListWithoutListMode() {
	it, err := s.p.store.List(s.base)
	s.NoError(err)
	s.NotNil(it)

	paths := make([]string, 0)
	for {
		o, err := it.Next()
		if errors.Is(err, types.IterateDone) {
			break
		}
		s.NoError(err)

		paths = append(paths, o.Path)
	}
	s.ElementsMatch(s.paths, paths)
}

func (s *storageListSuite) TestListEmptyDir() {
	if !s.p.store.Features().CreateDir {
		s.T().Skipf("store doesn't support CreateDir, skip TestListEmptyDir.")
	}

	path := uuid.New().String()

	virtualDir := s.p.store.Features().VirtualDir
	if !virtualDir {
		_, err := s.p.store.CreateDir(path)
		s.NoError(err)
	}

	defer func(isVirtualDir bool) {
		if !isVirtualDir {
			err := s.p.store.Delete(path, ps.WithObjectMode(types.ModeDir))
			s.NoError(err)
		}
	}(virtualDir)

	it, err := s.p.store.List(path, ps.WithListMode(types.ListModeDir))
	s.NoError(err)
	s.NotNil(it)

	o, err := it.Next()
	s.ErrorIs(err, types.IterateDone)
	s.Nil(o)
}
