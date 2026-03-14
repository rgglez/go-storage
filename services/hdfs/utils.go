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
package hdfs

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/colinmarc/hdfs/v2"

	"github.com/rgglez/go-storage/endpoint"
	ps "github.com/rgglez/go-storage/v5/pairs"
	"github.com/rgglez/go-storage/v5/services"
	"github.com/rgglez/go-storage/v5/types"
)

// Service is not usable by hdfs, only required for code generation.
type Service struct {
	f Factory

	defaultPairs types.DefaultServicePairs
	features     types.ServiceFeatures

	types.UnimplementedServicer
}

// String implements Servicer.String
func (s *Service) String() string {
	return fmt.Sprintf("Servicer hdfs")
}

// Storage is the example client.
type Storage struct {
	hdfs *hdfs.Client

	defaultPairs DefaultStoragePairs
	features     StorageFeatures

	workDir string

	types.UnimplementedStorager
	types.UnimplementedDirer
	types.UnimplementedMover
	types.UnimplementedAppender
}

// String implements Storager.String
func (s *Storage) String() string {
	return fmt.Sprintf("Storager hdfs {WorkDir: %s}", s.workDir)
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

// newService is not usable by hdfs, only required for code generation.
func (f *Factory) newService() (srv *Service, err error) {
	srv = &Service{}
	return
}

// newStorage creates an hdfs Storage.
func (f *Factory) newStorage() (store *Storage, err error) {
	defer func() {
		if err != nil {
			err = services.InitError{Op: "new_storager", Type: Type, Err: formatError(err)}
		}
	}()

	store = &Storage{
		workDir:  "/",
		features: f.storageFeatures(),
	}
	if f.WorkDir != "" {
		store.workDir = f.WorkDir
	}

	ep, err := endpoint.Parse(f.Endpoint)
	if err != nil {
		return nil, err
	}

	var addr string

	switch ep.Protocol() {
	case endpoint.ProtocolTCP:
		addr, _, _ = ep.TCP()
	default:
		return nil, services.PairUnsupportedError{Pair: ps.WithEndpoint(f.Endpoint)}
	}
	store.hdfs, err = hdfs.New(addr)
	if err != nil {
		return nil, errors.New("hdfs address is not exist")
	}

	return store, nil
}

func formatError(err error) error {
	if _, ok := err.(services.InternalError); ok {
		return err
	}

	switch {
	case errors.Is(err, os.ErrNotExist):
		return fmt.Errorf("%w: %v", services.ErrObjectNotExist, err)
	case errors.Is(err, os.ErrPermission):
		return fmt.Errorf("%w: %v", services.ErrPermissionDenied, err)
	default:
		return fmt.Errorf("%w: %v", services.ErrUnexpected, err)
	}
}

func (s *Storage) getAbsPath(fp string) string {
	if filepath.IsAbs(fp) {
		return fp
	}
	return path.Join(s.workDir, fp)
}

func (s *Storage) formatError(op string, err error, path ...string) error {
	if err == nil {
		return nil
	}
	return services.StorageError{
		Op:       op,
		Err:      formatError(err),
		Storager: s,
		Path:     path,
	}
}

func (s *Storage) newObject(done bool) *types.Object {
	return types.NewObject(s, done)
}
