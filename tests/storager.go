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
package tests

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/rgglez/go-storage/v5/types"
)

func TestStorager(t *testing.T, store types.Storager) {
	suite.Run(t, &StorageSuite{store: store})
}

type StorageSuite struct {
	suite.Suite
	store types.Storager
}

func (s *StorageSuite) TestString() {
	v := s.store.String()
	s.NotEmpty(v, "String() should not be empty.")
}

func (s *StorageSuite) TestMetadata() {
	m := s.store.Metadata()
	s.NotNil(m, "Metadata() should not return nil.")
}

func (s *StorageSuite) TestRead() {
	fe := s.store.Features()

	if !fe.Delete || !fe.Read {
		s.T().Skipf("store doesn't support Delete and Read, skip TestRead.")
	}

	suite.Run(s.T(), &storageReadSuite{p: s})
}

func (s *StorageSuite) TestWrite() {
	fe := s.store.Features()

	if !fe.Delete || !fe.Write {
		s.T().Skipf("store doesn't support Delete and Write, skip TestWrite.")
	}

	suite.Run(s.T(), &storageWriteSuite{p: s})
}

func (s *StorageSuite) TestStat() {
	fe := s.store.Features()

	if !fe.Delete || !fe.Write || !fe.Stat {
		s.T().Skipf("store doesn't support Delete, Write and Stat, skip TestStat.")
	}

	suite.Run(s.T(), &storageStatSuite{p: s})
}

func (s *StorageSuite) TestDelete() {
	fe := s.store.Features()

	if !fe.Delete || !fe.Write {
		s.T().Skipf("store doesn't support Delete, Write, skip TestDelete.")
	}

	suite.Run(s.T(), &storageDeleteSuite{p: s})
}

func (s *StorageSuite) TestList() {
	fe := s.store.Features()

	if !fe.Delete || !fe.Write || !fe.List {
		s.T().Skipf("store doesn't support Delete, Write and List, skip TestList.")
	}

	suite.Run(s.T(), &storageListSuite{p: s})
}

func (s *StorageSuite) TestPath() {
	fe := s.store.Features()

	if !fe.Delete || !fe.Write || !fe.Read {
		s.T().Skipf("store doesn't support Delete, Write and Read, skip TestPath.")
	}

	suite.Run(s.T(), &storagePathSuite{p: s})
}
