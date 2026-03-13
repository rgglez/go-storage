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
package hdfs

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/rgglez/go-storage/v5/pkg/iowrap"
	"github.com/rgglez/go-storage/v5/services"
	"github.com/rgglez/go-storage/v5/types"
)

func (s *Storage) create(path string, opt pairStorageCreate) (o *types.Object) {
	rp := s.getAbsPath(path)
	if opt.ObjectMode.IsDir() && opt.HasObjectMode {
		o = s.newObject(false)
		o.Mode = types.ModeDir
	} else {
		o = s.newObject(false)
		o.Mode = types.ModeRead
	}

	o.ID = rp
	o.Path = path
	return o
}

func (s *Storage) createDir(ctx context.Context, path string, opt pairStorageCreateDir) (o *types.Object, err error) {
	rp := s.getAbsPath(path)

	//	If dirname is already a directory,
	// 	MkdirAll does nothing and returns nil.
	err = s.hdfs.MkdirAll(rp, 0755)
	//	If dirname is not exist ,it will create a Mkdir rpc communication
	//	So we just need to catch other errors
	if err != nil {
		return nil, err
	}

	o = s.newObject(true)
	o.ID = rp
	o.Path = path
	o.Mode |= types.ModeDir
	return o, err
}

func (s *Storage) delete(ctx context.Context, path string, opt pairStorageDelete) (err error) {
	rp := s.getAbsPath(path)

	err = s.hdfs.Remove(rp)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		// Omit `file not exist` error here
		// ref: [GSP-46](https://github.com/rgglez/specs/blob/master/rfcs/46-idempotent-delete.md)
		err = nil
	}
	return err
}

func (s *Storage) list(ctx context.Context, path string, opt pairStorageList) (oi *types.ObjectIterator, err error) {
	if !opt.HasListMode || opt.ListMode.IsDir() {
		input := &listDirInput{
			rp:                s.getAbsPath(path),
			continuationToken: opt.ContinuationToken,
		}
		return types.NewObjectIterator(ctx, s.listDirNext, input), nil
	} else {
		return nil, services.ListModeInvalidError{Actual: opt.ListMode}
	}
}

func (s *Storage) metadata(opt pairStorageMetadata) (meta *types.StorageMeta) {
	meta = types.NewStorageMeta()
	meta.WorkDir = s.workDir
	return meta
}

func (s *Storage) move(ctx context.Context, src string, dst string, opt pairStorageMove) (err error) {
	rs := s.getAbsPath(src)
	rd := s.getAbsPath(dst)

	stat, err := s.hdfs.Stat(rd)
	if err == nil {
		if stat.IsDir() {
			return services.ErrObjectModeInvalid
		}
	}

	return s.hdfs.Rename(rs, rd)
}

func (s *Storage) read(ctx context.Context, path string, w io.Writer, opt pairStorageRead) (n int64, err error) {
	rp := s.getAbsPath(path)
	f, err := s.hdfs.Open(rp)
	if err != nil {
		return n, err
	}

	defer func() {
		closeErr := f.Close()
		if err == nil {
			err = closeErr
		}
	}()

	if opt.HasOffset {
		_, err := f.Seek(opt.Offset, 0)
		if err != nil {
			return 0, err
		}
	}

	var rc io.ReadCloser
	rc = f

	if opt.HasSize {
		return io.CopyN(w, rc, opt.Size)
	}
	if opt.HasIoCallback {
		rc = iowrap.CallbackReadCloser(rc, opt.IoCallback)
	}

	return io.Copy(w, rc)
}

func (s *Storage) stat(ctx context.Context, path string, opt pairStorageStat) (o *types.Object, err error) {
	rp := s.getAbsPath(path)

	stat, err := s.hdfs.Stat(rp)
	if err != nil {
		return nil, err
	}

	o = s.newObject(true)
	o.ID = rp
	o.Path = path

	if stat.IsDir() {
		o.Mode |= types.ModeDir
		return
	}

	if stat.Mode().IsRegular() {
		o.Mode |= types.ModeRead
		o.SetContentLength(stat.Size())
		o.SetLastModified(stat.ModTime())
	}

	return o, nil
}

func (s *Storage) write(ctx context.Context, path string, r io.Reader, size int64, opt pairStorageWrite) (n int64, err error) {
	rp := s.getAbsPath(path)
	dir := filepath.Dir(rp)

	if size == 0 {
		r = bytes.NewReader([]byte{})
	}

	if r == nil {
		return 0, services.ErrObjectNotExist
	}

	//	If dirname is already a directory,
	// 	MkdirAll does nothing and returns nil.
	err = s.hdfs.MkdirAll(dir, 0755)
	//	If dirname is not exist ,it will create a Mkdir rpc communication
	//	So we just need to catch other errors
	if err != nil {
		return 0, err
	}

	_, err = s.hdfs.Stat(rp)
	if err == nil {
		//	If the error returned by Stat is nil, the path must exist.
		err = s.hdfs.Remove(rp)

		if err != nil && errors.Is(err, os.ErrNotExist) {
			// Omit `file not exist` error here
			// ref: [GSP-46](https://github.com/rgglez/specs/blob/master/rfcs/46-idempotent-delete.md)
			err = nil
		}
	}

	//	This ensures that err can only be os.ErrNotExist
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return 0, err
	}

	f, err := s.hdfs.Create(rp)
	if err != nil {
		return 0, err
	}

	defer func() {
		closeErr := f.Close()
		if err == nil {
			err = closeErr
		}
	}()

	if opt.HasIoCallback {
		r = iowrap.CallbackReader(r, opt.IoCallback)
	}

	return io.CopyN(f, r, size)
}
