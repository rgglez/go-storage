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
package tar

import (
	"archive/tar"
	"bytes"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rgglez/go-storage/endpoint"
	"github.com/rgglez/go-storage/v5/pairs"
	"github.com/rgglez/go-storage/v5/types"
)

func setupTest(t *testing.T) (filename string) {
	f, err := os.CreateTemp(os.TempDir(), "")
	if err != nil {
		t.Fatal("create temp", err)
	}
	defer f.Close()

	filename = f.Name()

	tw := tar.NewWriter(f)
	defer tw.Close()

	files := []struct {
		Name    string
		Mode    int64
		Content string
	}{
		{"dir/", 0700 | int64(os.ModeDir), ""},
		{"world.txt", 0600, "world!"},
		{"dir/hello.txt", 0600, "hello,"},
	}

	for _, v := range files {
		err = tw.WriteHeader(&tar.Header{
			Name: v.Name,
			Mode: v.Mode,
			Size: int64(len(v.Content)),
		})
		if err != nil {
			t.Fatal(err)
		}
		if len(v.Content) > 0 {
			if _, err := tw.Write([]byte(v.Content)); err != nil {
				t.Fatal(err)
			}
		}
	}

	return filename
}

func TestStorage_List(t *testing.T) {
	filename := setupTest(t)

	s, err := NewStorager(
		pairs.WithEndpoint(endpoint.NewFile(filename).String()),
	)
	if err != nil {
		t.Fatal("new storage failed", err)
	}

	it, err := s.List("")
	if err != nil {
		t.Fatal("list", err)
	}

	files := make([]string, 0)

	for {
		o, err := it.Next()
		if err != nil && errors.Is(err, types.IterateDone) {
			break
		}
		if err != nil {
			t.Fatal("next", err)
		}

		files = append(files, o.Path)
	}

	assert.EqualValues(t, []string{
		"dir/",
		"world.txt",
		"dir/hello.txt",
	}, files)
}

func TestStorage_Read(t *testing.T) {
	filename := setupTest(t)

	s, err := NewStorager(
		pairs.WithEndpoint(endpoint.NewFile(filename).String()),
	)
	if err != nil {
		t.Fatal("new storage failed", err)
	}

	buf := &bytes.Buffer{}

	_, err = s.Read("world.txt", buf)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "world!", buf.String())
}

func TestStorage_Stat(t *testing.T) {
	filename := setupTest(t)

	s, err := NewStorager(
		pairs.WithEndpoint(endpoint.NewFile(filename).String()),
	)
	if err != nil {
		t.Fatal("new storage failed", err)
	}

	o, err := s.Stat("world.txt")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "world.txt", o.Path)
	assert.Equal(t, int64(6), o.MustGetContentLength())
}
