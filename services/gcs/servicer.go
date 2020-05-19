package gcs

import (
	"fmt"

	gs "cloud.google.com/go/storage"
	"google.golang.org/api/iterator"

	"github.com/Xuanwo/storage"
	"github.com/Xuanwo/storage/services"
	"github.com/Xuanwo/storage/types"
	ps "github.com/Xuanwo/storage/types/pairs"
)

// Service is the gcs config.
type Service struct {
	service   *gs.Client
	projectID string
}

// String implements Servicer.String
func (s *Service) String() string {
	return fmt.Sprintf("Servicer gcs")
}

// List implements Servicer.List
func (s *Service) List(pairs ...*types.Pair) (err error) {
	defer func() {
		err = s.formatError(services.OpList, err, "")
	}()

	opt, err := s.parsePairList(pairs...)
	if err != nil {
		return err
	}

	it := s.service.Buckets(opt.Context, s.projectID)
	for {
		bucketAttr, err := it.Next()
		// Next will return iterator.Done if there is no more items.
		if err != nil && err == iterator.Done {
			return nil
		}
		if err != nil {
			return err
		}
		store, err := s.newStorage(ps.WithName(bucketAttr.Name))
		if err != nil {
			return err
		}
		opt.StoragerFunc(store)
	}
}

// Get implements Servicer.Get
func (s *Service) Get(name string, pairs ...*types.Pair) (st storage.Storager, err error) {
	defer func() {
		err = s.formatError(services.OpGet, err, name)
	}()

	store, err := s.newStorage(ps.WithName(name))
	if err != nil {
		return nil, err
	}
	return store, nil
}

// Create implements Servicer.Create
func (s *Service) Create(name string, pairs ...*types.Pair) (st storage.Storager, err error) {
	defer func() {
		err = s.formatError(services.OpCreate, err, name)
	}()

	opt, err := s.parsePairCreate(pairs...)
	if err != nil {
		return nil, err
	}

	store, err := s.newStorage(ps.WithName(name))
	if err != nil {
		return nil, err
	}
	err = store.bucket.Create(opt.Context, s.projectID, nil)
	if err != nil {
		return nil, err
	}
	return store, nil
}

// Delete implements Servicer.Delete
func (s *Service) Delete(name string, pairs ...*types.Pair) (err error) {
	defer func() {
		err = s.formatError(services.OpDelete, err, name)
	}()

	opt, err := s.parsePairDelete(pairs...)
	if err != nil {
		return err
	}

	store, err := s.newStorage(ps.WithName(name))
	if err != nil {
		return err
	}
	err = store.bucket.Delete(opt.Context)
	if err != nil {
		return err
	}
	return nil
}

// newStorage will create a new client.
func (s *Service) newStorage(pairs ...*types.Pair) (st *Storage, err error) {
	opt, err := parseStoragePairNew(pairs...)
	if err != nil {
		return nil, err
	}

	bucket := s.service.Bucket(opt.Name)

	store := &Storage{
		bucket: bucket,
		name:   opt.Name,

		workDir: "/",
	}

	if opt.HasWorkDir {
		store.workDir = opt.WorkDir
	}
	return store, nil
}

func (s *Service) formatError(op string, err error, name string) error {
	if err == nil {
		return nil
	}

	return &services.ServiceError{
		Op:       op,
		Err:      formatError(err),
		Servicer: s,
		Name:     name,
	}
}