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
package tests

import (
	"testing"

	"github.com/google/uuid"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/rgglez/go-storage/v5/pairs"
	"github.com/rgglez/go-storage/v5/types"
)

func TestDirer(t *testing.T, store types.Storager) {
	Convey("Given a basic Storager", t, func() {

		Convey("When CreateDir", func() {
			path := uuid.New().String()
			_, err := store.CreateDir(path)

			defer func() {
				err := store.Delete(path, pairs.WithObjectMode(types.ModeDir))
				if err != nil {
					t.Error(err)
				}
			}()

			Convey("The first returned error should be nil", func() {
				So(err, ShouldBeNil)
			})

			o, err := store.CreateDir(path)
			Convey("The second returned error also should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("The Object Path should equal to the input path", func() {
				So(o.Path, ShouldEqual, path)
			})

			Convey("The Object Mode should be dir", func() {
				// Dir object's mode must be Dir.
				So(o.Mode.IsDir(), ShouldBeTrue)
			})
		})

		Convey("When Create with ModeDir", func() {
			path := uuid.New().String()
			o := store.Create(path, pairs.WithObjectMode(types.ModeDir))

			defer func() {
				err := store.Delete(path, pairs.WithObjectMode(types.ModeDir))
				if err != nil {
					t.Error(err)
				}
			}()

			Convey("The Object Path should equal to the input path", func() {
				So(o.Path, ShouldEqual, path)
			})

			Convey("The Object Mode should be dir", func() {
				// Dir object's mode must be Dir.
				So(o.Mode.IsDir(), ShouldBeTrue)
			})
		})

		Convey("When Stat with ModeDir", func() {
			path := uuid.New().String()
			_, err := store.CreateDir(path)
			if err != nil {
				t.Error(err)
			}

			defer func() {
				err := store.Delete(path, pairs.WithObjectMode(types.ModeDir))
				if err != nil {
					t.Error(err)
				}
			}()

			o, err := store.Stat(path, pairs.WithObjectMode(types.ModeDir))

			Convey("The error should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("The Object Path should equal to the input path", func() {
				So(o.Path, ShouldEqual, path)
			})

			Convey("The Object Mode should be dir", func() {
				// Dir object's mode must be Dir.
				So(o.Mode.IsDir(), ShouldBeTrue)
			})
		})

		Convey("When Delete with ModeDir", func() {
			path := uuid.New().String()
			_, err := store.CreateDir(path)
			if err != nil {
				t.Error(err)
			}

			err = store.Delete(path, pairs.WithObjectMode(types.ModeDir))
			Convey("The first returned error should be nil", func() {
				So(err, ShouldBeNil)
			})

			err = store.Delete(path, pairs.WithObjectMode(types.ModeDir))
			Convey("The second returned error also should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}
