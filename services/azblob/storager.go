package azblob

import (
	"fmt"
	"io"
	"strings"

	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/Xuanwo/storage/pkg/iowrap"
	"github.com/Xuanwo/storage/services"
	"github.com/Xuanwo/storage/types"
	"github.com/Xuanwo/storage/types/metadata"
)

// Storage is the azblob service client.
type Storage struct {
	bucket azblob.ContainerURL

	name    string
	workDir string
}

// String implements Storager.String
func (s *Storage) String() string {
	return fmt.Sprintf(
		"Storager azblob {Name: %s, WorkDir: %s}",
		s.name, "/"+s.workDir,
	)
}

// Metadata implements Storager.Metadata
func (s *Storage) Metadata(pairs ...*types.Pair) (m metadata.StorageMeta, err error) {
	m = metadata.NewStorageMeta()
	m.Name = s.name
	m.WorkDir = s.workDir
	return m, nil
}

// List implements Storager.List
func (s *Storage) List(path string, pairs ...*types.Pair) (err error) {
	defer func() {
		err = s.formatError("list", err, path)
	}()

	opt, err := parseStoragePairList(pairs...)
	if err != nil {
		return err
	}

	rp := s.getAbsPath(path)

	marker := azblob.Marker{}
	var output *azblob.ListBlobsFlatSegmentResponse
	for {
		output, err = s.bucket.ListBlobsFlatSegment(opt.Context, marker, azblob.ListBlobsSegmentOptions{
			Prefix: rp,
		})
		if err != nil {
			return err
		}

		for _, v := range output.Segment.BlobItems {
			o := &types.Object{
				ID:         v.Name,
				Name:       s.getRelPath(v.Name),
				Type:       types.ObjectTypeDir,
				Size:       *v.Properties.ContentLength,
				UpdatedAt:  v.Properties.LastModified,
				ObjectMeta: metadata.NewObjectMeta(),
			}
			o.SetContentType(*v.Properties.ContentType)
			o.SetContentMD5(string(v.Properties.ContentMD5))

			storageClass, err := formatStorageClass(v.Properties.AccessTier)
			if err != nil {
				return err
			}
			o.SetStorageClass(storageClass)

			opt.FileFunc(o)
		}

		marker = output.NextMarker
		if !marker.NotDone() {
			break
		}
	}

	return
}

// Read implements Storager.Read
func (s *Storage) Read(path string, pairs ...*types.Pair) (r io.ReadCloser, err error) {
	defer func() {
		err = s.formatError("read", err, path)
	}()

	opt, err := parseStoragePairRead(pairs...)
	if err != nil {
		return nil, err
	}

	rp := s.getAbsPath(path)

	output, err := s.bucket.NewBlockBlobURL(rp).Download(opt.Context, 0, azblob.CountToEnd, azblob.BlobAccessConditions{}, false)
	if err != nil {
		return nil, err
	}

	r = output.Body(azblob.RetryReaderOptions{})
	if opt.HasReadCallbackFunc {
		r = iowrap.CallbackReadCloser(r, opt.ReadCallbackFunc)
	}
	return r, nil
}

// Write implements Storager.Write
func (s *Storage) Write(path string, r io.Reader, pairs ...*types.Pair) (err error) {
	defer func() {
		err = s.formatError("write", err, path)
	}()

	opt, err := parseStoragePairWrite(pairs...)
	if err != nil {
		return err
	}

	rp := s.getAbsPath(path)

	if opt.HasReadCallbackFunc {
		r = iowrap.CallbackReader(r, opt.ReadCallbackFunc)
	}

	// TODO: add checksum and storage class support.
	_, err = s.bucket.NewBlockBlobURL(rp).Upload(opt.Context, iowrap.ReadSeekCloser(r),
		azblob.BlobHTTPHeaders{}, azblob.Metadata{}, azblob.BlobAccessConditions{})
	if err != nil {
		return err
	}
	return nil
}

// Stat implements Storager.Stat
func (s *Storage) Stat(path string, pairs ...*types.Pair) (o *types.Object, err error) {
	defer func() {
		err = s.formatError("stat", err, path)
	}()

	opt, err := parseStoragePairStat(pairs...)
	if err != nil {
		return nil, err
	}

	rp := s.getAbsPath(path)

	output, err := s.bucket.NewBlockBlobURL(rp).GetProperties(opt.Context, azblob.BlobAccessConditions{})
	if err != nil {
		return nil, err
	}

	o = &types.Object{
		ID:         rp,
		Name:       path,
		Type:       types.ObjectTypeFile,
		Size:       output.ContentLength(),
		UpdatedAt:  output.LastModified(),
		ObjectMeta: metadata.NewObjectMeta(),
	}

	storageClass, err := formatStorageClass(azblob.AccessTierType(output.AccessTier()))
	if err != nil {
		return nil, err
	}
	o.SetStorageClass(storageClass)
	return o, nil
}

// Delete implements Storager.Delete
func (s *Storage) Delete(path string, pairs ...*types.Pair) (err error) {
	defer func() {
		err = s.formatError("delete", err, path)
	}()

	opt, err := parseStoragePairStat(pairs...)
	if err != nil {
		return err
	}

	rp := s.getAbsPath(path)

	_, err = s.bucket.NewBlockBlobURL(rp).Delete(opt.Context,
		azblob.DeleteSnapshotsOptionNone, azblob.BlobAccessConditions{})
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) getAbsPath(path string) string {
	return strings.TrimPrefix(s.workDir+"/"+path, "/")
}

func (s *Storage) getRelPath(path string) string {
	return strings.TrimPrefix(path, s.workDir+"/")
}

func (s *Storage) formatError(op string, err error, path ...string) error {
	if err == nil {
		return nil
	}

	// Handle errors returned by azblob.
	e, ok := err.(azblob.StorageError)
	if ok {
		err = formatAzblobError(e)
	}

	return &services.StorageError{
		Op:       op,
		Err:      err,
		Storager: s,
		Path:     path,
	}
}