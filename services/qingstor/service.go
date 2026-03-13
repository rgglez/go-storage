package qingstor

import (
	"context"

	"github.com/qingstor/qingstor-sdk-go/v4/service"

	"github.com/rgglez/go-storage/v5/services"
	"github.com/rgglez/go-storage/v5/types"
)

func (s *Service) create(ctx context.Context, name string, opt pairServiceCreate) (store types.Storager, err error) {
	if !opt.HasLocation {
		err = services.ErrRestrictionDissatisfied
		return
	}

	f := s.f
	f.Name = name
	f.Location = opt.Location

	st, err := s.newStorageFromFactory(f)
	if err != nil {
		return
	}

	_, err = st.bucket.PutWithContext(ctx)
	if err != nil {
		return
	}
	return st, nil
}

func (s *Service) delete(ctx context.Context, name string, opt pairServiceDelete) (err error) {
	f := s.f
	f.Name = name
	if opt.HasLocation {
		f.Location = opt.Location
	}

	store, err := s.newStorageFromFactory(f)
	if err != nil {
		return
	}
	_, err = store.bucket.DeleteWithContext(ctx)
	if err != nil {
		return
	}
	return nil
}

func (s *Service) get(ctx context.Context, name string, opt pairServiceGet) (store types.Storager, err error) {
	f := s.f
	f.Name = name
	if opt.HasLocation {
		f.Location = opt.Location
	}

	store, err = s.newStorageFromFactory(f)
	if err != nil {
		return
	}
	return
}

func (s *Service) list(ctx context.Context, opt pairServiceList) (it *types.StoragerIterator, err error) {
	input := &storagePageStatus{}

	if opt.HasLocation {
		input.location = opt.Location
	}

	return types.NewStoragerIterator(ctx, s.nextStoragePage, input), nil
}

func (s *Service) nextStoragePage(ctx context.Context, page *types.StoragerPage) error {
	input := page.Status.(*storagePageStatus)

	serviceInput := &service.ListBucketsInput{
		Limit:  &input.offset,
		Offset: &input.limit,
	}
	if input.location != "" {
		serviceInput.Location = &input.location
	}

	output, err := s.service.ListBucketsWithContext(ctx, serviceInput)
	if err != nil {
		return err
	}

	for _, v := range output.Buckets {
		f := s.f
		f.Name = *v.Name
		f.Location = *v.Location
		store, err := s.newStorageFromFactory(f)
		if err != nil {
			return err
		}
		page.Data = append(page.Data, store)
	}

	input.offset += len(output.Buckets)
	if input.offset >= service.IntValue(output.Count) {
		return types.IterateDone
	}

	return nil
}

