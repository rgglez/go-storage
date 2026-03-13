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
package webdav

import (
	"fmt"

	"github.com/rgglez/go-storage/v5/types"
)

// Service is the webdav config.
// It is not usable, only for generate code
type Service struct {
	f Factory

	defaultPairs types.DefaultServicePairs
	features     types.ServiceFeatures

	types.UnimplementedServicer
}

// String implements Servicer.String
func (s *Service) String() string {
	return fmt.Sprintf("Servicer webdav")
}

// NewServicer is not usable, only for generate code
func NewServicer(pairs ...types.Pair) (types.Servicer, error) {
	f := Factory{}
	err := f.WithPairs(pairs...)
	if err != nil {
		return nil, err
	}
	return f.NewServicer()
}

// newService is not usable, only for generate code
func (f *Factory) newService() (srv *Service, err error) {
	srv = &Service{
		f:        *f,
		features: f.serviceFeatures(),
	}
	return
}

// Storage is the example client.
type Storage struct {
	f Factory

	defaultPairs types.DefaultStoragePairs
	features     types.StorageFeatures

	types.UnimplementedStorager
}

// String implements Storager.String
func (s *Storage) String() string {
	panic("implement me")
}

// NewStorager will create Storager only.
func NewStorager(pairs ...types.Pair) (types.Storager, error) {
	f := Factory{}
	err := f.WithPairs(pairs...)
	if err != nil {
		return nil, err
	}
	return f.newStorage()
}

// newStorager will create a storage client.
func (f *Factory) newStorage() (store *Storage, err error) {
	panic("implement me")
}

func (s *Storage) formatError(op string, err error, path ...string) error {
	panic("implement me")
}
