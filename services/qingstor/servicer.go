package qingstor

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/qingstor/qingstor-sdk-go/v4/config"
	iface "github.com/qingstor/qingstor-sdk-go/v4/interface"
	"github.com/qingstor/qingstor-sdk-go/v4/service"

	"github.com/Xuanwo/storage"
	"github.com/Xuanwo/storage/pkg/headers"
	"github.com/Xuanwo/storage/services"
	"github.com/Xuanwo/storage/types"
	ps "github.com/Xuanwo/storage/types/pairs"
)

// Service is the qingstor service config.
type Service struct {
	config  *config.Config
	service iface.Service

	client *http.Client
}

// String implements Service.String
func (s *Service) String() string {
	if s.config == nil {
		return fmt.Sprintf("Servicer qingstor")
	}
	return fmt.Sprintf("Servicer qingstor {AccessKey: %s}", s.config.AccessKeyID)
}

// List implements Servicer.List
func (s *Service) List(pairs ...*types.Pair) (err error) {
	defer func() {
		err = s.formatError(services.OpList, err, "")
	}()

	opt, err := s.parsePairList(pairs...)
	if err != nil {
		return
	}

	input := &service.ListBucketsInput{}
	if opt.HasLocation {
		input.Location = &opt.Location
	}

	offset := 0
	var output *service.ListBucketsOutput
	for {
		input.Offset = service.Int(offset)

		output, err = s.service.ListBuckets(input)
		if err != nil {
			return
		}

		for _, v := range output.Buckets {
			store, err := s.newStorage(ps.WithName(*v.Name), ps.WithLocation(*v.Location))
			if err != nil {
				return err
			}
			opt.StoragerFunc(store)
		}

		offset += len(output.Buckets)
		if offset >= service.IntValue(output.Count) {
			break
		}
	}
	return nil
}

// Get implements Servicer.Get
func (s *Service) Get(name string, pairs ...*types.Pair) (store storage.Storager, err error) {
	defer func() {
		err = s.formatError(services.OpGet, err, name)
	}()

	_, err = s.parsePairGet(pairs...)
	if err != nil {
		return
	}

	pairs = append(pairs, ps.WithName(name))

	store, err = s.newStorage(pairs...)
	if err != nil {
		return
	}
	return
}

// Create implements Servicer.Create
func (s *Service) Create(name string, pairs ...*types.Pair) (store storage.Storager, err error) {
	defer func() {
		err = s.formatError(services.OpCreate, err, name)
	}()

	_, err = s.parsePairCreate(pairs...)
	if err != nil {
		return
	}

	// ServicePairCreate requires location, so we don't need to add location into pairs
	pairs = append(pairs, ps.WithName(name))

	st, err := s.newStorage(pairs...)
	if err != nil {
		return
	}

	_, err = st.bucket.Put()
	if err != nil {
		return
	}
	return st, nil
}

// Delete implements Servicer.Delete
func (s *Service) Delete(name string, pairs ...*types.Pair) (err error) {
	defer func() {
		err = s.formatError(services.OpDelete, err, name)
	}()

	_, err = s.parsePairDelete(pairs...)
	if err != nil {
		return
	}

	pairs = append(pairs, ps.WithName(name))

	store, err := s.newStorage(pairs...)
	if err != nil {
		return
	}
	_, err = store.bucket.Delete()
	if err != nil {
		return
	}
	return nil
}

func (s *Service) newStorage(pairs ...*types.Pair) (store *Storage, err error) {
	opt, err := parseStoragePairNew(pairs...)
	if err != nil {
		return
	}

	// WorkDir should be an abs path, start and ends with "/"
	if opt.HasWorkDir && !isWorkDirValid(opt.WorkDir) {
		err = ErrInvalidWorkDir
		return
	}
	// set work dir into root path if no work dir passed
	if !opt.HasWorkDir {
		opt.WorkDir = "/"
	}

	if !IsBucketNameValid(opt.Name) {
		err = ErrInvalidBucketName
		return
	}

	// Detect location automatically
	if !opt.HasLocation {
		opt.Location, err = s.detectLocation(opt.Name)
		if err != nil {
			return
		}
	}

	bucket, err := s.service.Bucket(opt.Name, opt.Location)
	if err != nil {
		return
	}

	st := &Storage{
		bucket:     bucket,
		config:     bucket.Config,
		properties: bucket.Properties,

		workDir: "/",
	}

	if opt.HasWorkDir {
		st.workDir = opt.WorkDir
	}
	if opt.HasDisableURICleaning {
		st.config.DisableURICleaning = opt.DisableURICleaning
	}
	return st, nil
}

func (s *Service) detectLocation(name string) (location string, err error) {
	defer func() {
		err = s.formatError("detect_location", err, "")
	}()

	url := fmt.Sprintf("%s://%s.%s:%d", s.config.Protocol, name, s.config.Host, s.config.Port)

	r, err := s.client.Head(url)
	if err != nil {
		return
	}
	if r.StatusCode != http.StatusTemporaryRedirect {
		err = fmt.Errorf("head status is %d instead of %d", r.StatusCode, http.StatusTemporaryRedirect)
		return
	}

	// Example URL: https://bucket.zone.qingstor.com
	location = strings.Split(r.Header.Get(headers.Location), ".")[1]
	return
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

// isWorkDirValid check qingstor work dir
// work dir must start with only one "/" (abs path), and end with only one "/" (a dir).
// If work dir is the root path, set it to "/".
func isWorkDirValid(wd string) bool {
	return strings.HasPrefix(wd, "/") && // must start with "/"
		strings.HasSuffix(wd, "/") && // must end with "/"
		!strings.HasPrefix(wd, "//") && // not start with more than one "/"
		!strings.HasSuffix(wd, "//") // not end with more than one "/"
}