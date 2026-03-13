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
package qingstor

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"testing"

	"bou.ke/monkey"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/qingstor/qingstor-sdk-go/v4/config"
	"github.com/qingstor/qingstor-sdk-go/v4/service"
	"github.com/stretchr/testify/assert"

	"github.com/rgglez/go-storage/credential"
	"github.com/rgglez/go-storage/v5/pairs"
	"github.com/rgglez/go-storage/v5/services"
)

func TestService_String(t *testing.T) {
	accessKey := uuid.New().String()
	secretKey := uuid.New().String()

	srv, err := NewServicer(pairs.WithCredential(credential.NewHmac(accessKey, secretKey).String()))
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, srv.String())
}

func TestService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("with location", func(t *testing.T) {
		mockService := NewMockService(ctrl)

		srv := Service{
			service: mockService,
		}

		name := uuid.New().String()
		location := uuid.New().String()

		mockService.EXPECT().Bucket(gomock.Any(), gomock.Any()).DoAndReturn(func(bucketName, inputLocation string) (*service.Bucket, error) {
			assert.Equal(t, name, bucketName)
			assert.Equal(t, location, inputLocation)
			return &service.Bucket{
				Properties: &service.Properties{
					BucketName: &name,
					Zone:       &location,
				},
			}, nil
		})

		s, err := srv.Get(name, pairs.WithLocation(location))
		assert.NoError(t, err)
		assert.NotNil(t, s)
	})

	t.Run("without location", func(t *testing.T) {
		mockService := NewMockService(ctrl)

		srv := &Service{}
		srv.service = mockService
		srv.config = &config.Config{
			AccessKeyID:     uuid.New().String(),
			SecretAccessKey: uuid.New().String(),
			Host:            uuid.New().String(),
			Port:            1234,
			Protocol:        "https",
		}

		name := uuid.New().String()
		location := uuid.New().String()

		// Patch http Head.
		fn := func(client *http.Client, url string) (*http.Response, error) {
			header := http.Header{}
			header.Set(
				"location",
				fmt.Sprintf("%s://%s.%s/%s",
					srv.config.Protocol, location, srv.config.Host, name),
			)
			return &http.Response{
				StatusCode: http.StatusMovedPermanently,
				Header:     header,
			}, nil
		}
		monkey.PatchInstanceMethod(reflect.TypeOf(srv.client), "Head", fn)

		// Mock Bucket.
		mockService.EXPECT().Bucket(gomock.Any(), gomock.Any()).DoAndReturn(func(bucketName, inputLocation string) (*service.Bucket, error) {
			assert.Equal(t, name, bucketName)
			assert.Equal(t, location, inputLocation)
			return &service.Bucket{
				Properties: &service.Properties{
					BucketName: &name,
					Zone:       &location,
				},
			}, nil
		})

		s, err := srv.Get(name)
		assert.NoError(t, err)
		assert.NotNil(t, s)
	})

	t.Run("invalid bucket name", func(t *testing.T) {
		mockService := NewMockService(ctrl)

		srv := &Service{}
		srv.service = mockService
		srv.config = &config.Config{
			AccessKeyID:     uuid.New().String(),
			SecretAccessKey: uuid.New().String(),
			Host:            uuid.New().String(),
			Port:            1234,
			Protocol:        uuid.New().String(),
		}

		s, err := srv.Get("1234")
		assert.Error(t, err)
		assert.Nil(t, s)
	})
}

func TestService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := NewMockService(ctrl)

	srv := Service{
		service: mockService,
	}

	// Test case1: without location
	path := uuid.New().String()
	_, err := srv.Create(path)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, services.ErrRestrictionDissatisfied))

	// Test case2: with valid location.
	path = uuid.New().String()
	location := uuid.New().String()

	// Monkey the bucket's Put method
	bucket := &service.Bucket{}
	fn := func(*service.Bucket, context.Context) (*service.PutBucketOutput, error) {
		t.Log("Bucket put has been called")
		return &service.PutBucketOutput{}, nil
	}
	monkey.PatchInstanceMethod(reflect.TypeOf(bucket), "PutWithContext", fn)

	mockService.EXPECT().Bucket(gomock.Any(), gomock.Any()).Do(func(inputPath, inputLocation string) {
		assert.Equal(t, path, inputPath)
		assert.Equal(t, location, inputLocation)
	}).Return(bucket, nil)

	_, err = srv.Create(path, pairs.WithLocation(location))
	assert.NoError(t, err)
}

func TestService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := NewMockService(ctrl)

	srv := Service{
		service: mockService,
	}

	{
		name := uuid.New().String()
		location := uuid.New().String()

		// Patch bucket.Delete
		bucket := &service.Bucket{}
		fn := func(*service.Bucket, context.Context) (*service.DeleteBucketOutput, error) {
			t.Log("Bucket delete has been called")
			return nil, nil
		}
		monkey.PatchInstanceMethod(reflect.TypeOf(bucket), "DeleteWithContext", fn)

		mockService.EXPECT().Bucket(gomock.Any(), gomock.Any()).DoAndReturn(func(bucketName, inputLocation string) (*service.Bucket, error) {
			assert.Equal(t, name, bucketName)
			assert.Equal(t, location, inputLocation)
			return bucket, nil
		})

		err := srv.Delete(name, pairs.WithLocation(location))
		assert.NoError(t, err)
	}
}

func TestService_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := NewMockService(ctrl)

	srv := &Service{
		service: mockService,
	}
	mockService.EXPECT().Bucket(gomock.Any(), gomock.Any()).DoAndReturn(func(inputName, inputLocation string) (*service.Bucket, error) {
		return &service.Bucket{
			Config: &config.Config{},
		}, nil
	}).AnyTimes()

	{
		// Test request with location.
		name := uuid.New().String()
		location := uuid.New().String()

		mockService.EXPECT().ListBucketsWithContext(gomock.Eq(context.Background()), gomock.Any()).DoAndReturn(func(ctx context.Context, input *service.ListBucketsInput) (*service.ListBucketsOutput, error) {
			assert.Equal(t, location, *input.Location)
			return &service.ListBucketsOutput{
				Buckets: []*service.BucketType{
					{Name: &name, Location: &location},
				},
			}, nil
		})

		it, err := srv.List(pairs.WithLocation(location))
		assert.NoError(t, err)
		assert.NotNil(t, it)
		st, err := it.Next()
		if err != nil {
			t.Error(err)
		}
		assert.NotNil(t, st)
	}

	{
		// Test request without location.
		name := uuid.New().String()
		location := uuid.New().String()

		mockService.EXPECT().ListBucketsWithContext(gomock.Eq(context.Background()), gomock.Any()).DoAndReturn(func(ctx context.Context, input *service.ListBucketsInput) (*service.ListBucketsOutput, error) {
			assert.Nil(t, input.Location)
			return &service.ListBucketsOutput{
				Buckets: []*service.BucketType{
					{Name: &name, Location: &location},
				},
			}, nil
		})

		it, err := srv.List()
		assert.NoError(t, err)
		assert.NotNil(t, it)
		// assert.Equal(t, 1, len(s))
		_, err = it.Next()
		assert.NoError(t, err)
	}
}

func ExampleNew() {
	_, _, err := New(
		pairs.WithCredential(
			credential.NewHmac("test_access_key", "test_secret_key").String(),
		),
	)
	if err != nil {
		log.Printf("service init failed: %v", err)
	}
}

func ExampleService_Get() {
	srv, _, err := New(
		pairs.WithCredential(
			credential.NewHmac("test_access_key", "test_secret_key").String(),
		),
	)
	if err != nil {
		log.Printf("service init failed: %v", err)
	}

	store, err := srv.Get("bucket_name", pairs.WithLocation("location"))
	if err != nil {
		log.Printf("service get bucket failed: %v", err)
	}
	log.Printf("%v", store)
}

func Test_isWorkDirValid(t *testing.T) {
	type args struct {
		wd string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "not a dir",
			args: args{wd: "/path/to/file"},
			want: false,
		},
		{
			name: "not a abs dir",
			args: args{wd: "path/to/file/"},
			want: false,
		},
		{
			name: "start with multi-slash",
			args: args{wd: "///multi-slash/to/file/"},
			want: false,
		},
		{
			name: "end with multi-slash",
			args: args{wd: "/path/to/multi-slash///"},
			want: false,
		},
		{
			name: "root path",
			args: args{wd: "/"},
			want: true,
		},
		{
			name: "normal abs dir",
			args: args{wd: "/path/to/dir/"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isWorkDirValid(tt.args.wd)
			if got != tt.want {
				t.Errorf("isWorkDirValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_newStorage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	blankBucket := service.Bucket{
		Config: &config.Config{},
	}
	validWorkDir, invalidWorkDir := "/valid/dir/", "invalid/dir"
	type args struct {
		name     string
		location string
		workDir  string
	}
	tests := []struct {
		name       string
		wd         string
		args       args
		wantBucket *service.Bucket
		targetErr  error
		wantErr    bool
	}{
		{
			name: "normal case",
			wd:   validWorkDir,
			args: args{
				name:     uuid.New().String(),
				location: uuid.New().String(),
				workDir:  validWorkDir,
			},
			wantBucket: &blankBucket,
			targetErr:  nil,
			wantErr:    false,
		},
		{
			name: "invalid work dir",
			wd:   invalidWorkDir,
			args: args{
				name:     uuid.New().String(),
				location: uuid.New().String(),
				workDir:  invalidWorkDir,
			},
			wantBucket: nil,
			targetErr:  ErrWorkDirInvalid,
			wantErr:    true,
		},
		{
			name: "blank work dir",
			wd:   "/",
			args: args{
				name:     uuid.New().String(),
				location: uuid.New().String(),
			},
			wantBucket: &blankBucket,
			targetErr:  nil,
			wantErr:    false,
		},
		{
			name: "invalid bucket name",
			wd:   validWorkDir,
			args: args{
				name:     "invalid bucket name",
				location: uuid.New().String(),
				workDir:  validWorkDir,
			},
			wantBucket: nil,
			targetErr:  ErrBucketNameInvalid,
			wantErr:    true,
		},
		{
			name:       "no name, fail with invalid bucket name",
			args:       args{},
			wantBucket: nil,
			targetErr:  ErrBucketNameInvalid,
			wantErr:    true,
		},
	}

	mockService := NewMockService(ctrl)
	mockService.EXPECT().Bucket(gomock.Any(), gomock.Any()).DoAndReturn(func(inputName, inputLocation string) (*service.Bucket, error) {
		return &blankBucket, nil
	}).AnyTimes()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Factory{
				Name:     tt.args.name,
				Location: tt.args.location,
				WorkDir:  tt.args.workDir,
			}
			s := &Service{
				f:       f,
				service: mockService,
				config:  &config.Config{},
			}
			gotStore, err := s.newStorageFromFactory(f)
			if (err != nil) != tt.wantErr {
				t.Errorf("newStorage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr || err != nil {
				assert.Nil(t, gotStore, tt.name)
				assert.True(t, errors.Is(err, tt.targetErr), tt.name)
			} else {
				assert.Equal(t, tt.wantBucket, gotStore.bucket, tt.name)
				assert.Equal(t, tt.wd, gotStore.workDir, tt.name)
			}
		})
	}
}
