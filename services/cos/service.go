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
package cos

import (
	"context"

	typ "github.com/rgglez/go-storage/v5/types"
)

func (s *Service) create(ctx context.Context, name string, opt pairServiceCreate) (store typ.Storager, err error) {
	f := s.f
	f.Name = name
	f.Location = opt.Location
	st, err := f.newStorage()
	if err != nil {
		return nil, err
	}
	_, err = st.bucket.Put(ctx, nil)
	if err != nil {
		return nil, err
	}
	return st, nil
}

func (s *Service) delete(ctx context.Context, name string, opt pairServiceDelete) (err error) {
	f := s.f
	f.Name = name
	f.Location = opt.Location
	st, err := f.newStorage()
	if err != nil {
		return err
	}
	_, err = st.bucket.Delete(ctx)
	if err != nil {
		return err
	}
	return
}

func (s *Service) get(ctx context.Context, name string, opt pairServiceGet) (store typ.Storager, err error) {
	f := s.f
	f.Name = name
	f.Location = opt.Location
	st, err := f.newStorage()
	if err != nil {
		return nil, err
	}
	return st, nil
}

func (s *Service) list(ctx context.Context, opt pairServiceList) (it *typ.StoragerIterator, err error) {
	return typ.NewStoragerIterator(ctx, s.nextStoragePage, nil), nil
}

func (s *Service) nextStoragePage(ctx context.Context, page *typ.StoragerPage) error {
	output, _, err := s.service.Service.Get(ctx)
	if err != nil {
		return err
	}

	for _, v := range output.Buckets {
		f := s.f
		f.Name = v.Name
		f.Location = v.Region
		store, err := f.newStorage()
		if err != nil {
			return err
		}

		page.Data = append(page.Data, store)
	}

	return typ.IterateDone
}
