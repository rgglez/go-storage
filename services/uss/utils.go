package uss

import (
	"fmt"
	"strings"

	"github.com/upyun/go-sdk/upyun"

	"github.com/Xuanwo/storage"
	"github.com/Xuanwo/storage/pkg/credential"
	"github.com/Xuanwo/storage/pkg/httpclient"
	"github.com/Xuanwo/storage/services"
	"github.com/Xuanwo/storage/types"
	ps "github.com/Xuanwo/storage/types/pairs"
)

// NewStorager will create Storager only.
func NewStorager(pairs ...*types.Pair) (storage.Storager, error) {
	return newStorager(pairs...)
}

func newStorager(pairs ...*types.Pair) (store *Storage, err error) {
	defer func() {
		if err != nil {
			err = &services.InitError{Op: services.OpNewStorager, Type: Type, Err: err, Pairs: pairs}
		}
	}()

	store = &Storage{}

	opt, err := parseStoragePairNew(pairs...)
	if err != nil {
		return
	}

	credProtocol, cred := opt.Credential.Protocol(), opt.Credential.Value()
	if credProtocol != credential.ProtocolHmac {
		return nil, services.NewPairUnsupportedError(ps.WithCredential(opt.Credential))
	}

	cfg := &upyun.UpYunConfig{
		Bucket:   opt.Name,
		Operator: cred[0],
		Password: cred[1],
	}
	store.bucket = upyun.NewUpYun(cfg)
	// Set http client
	store.bucket.SetHTTPClient(httpclient.New(opt.HTTPClientOptions))
	store.name = opt.Name
	store.workDir = "/"
	if opt.HasWorkDir {
		store.workDir = opt.WorkDir
	}
	return
}

// ref: https://help.upyun.com/knowledge-base/errno/
func formatError(err error) error {
	fn := func(s string) bool {
		return strings.Contains(err.Error(), `"code": `+s)
	}

	switch {
	case !fn(""):
		// If body is empty
		switch {
		case strings.Contains(err.Error(), "404"):
			return fmt.Errorf("%w: %v", services.ErrObjectNotExist, err)
		default:
			return err
		}
	case fn("40400001"):
		// 40400001:	file or directory not found
		return fmt.Errorf("%w: %v", services.ErrObjectNotExist, err)
	case fn("40100017"), fn("40100019"), fn("40300011"):
		// 40100017: user need permission
		// 40100019: account forbidden
		// 40300011: has no permission to delete
		return fmt.Errorf("%w: %v", services.ErrPermissionDenied, err)
	default:
		return err
	}
}