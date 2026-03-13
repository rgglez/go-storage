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
package types

import (
	"context"
)

type Interceptor func(ctx context.Context, method string) func(error)

func ChainInterceptor(interceptors ...Interceptor) Interceptor {
	n := len(interceptors)
	fns := make([]func(error), n)

	return func(ctx context.Context, method string) func(err error) {
		for k, v := range interceptors {
			fn := v(ctx, method)
			fns[n-k-1] = fn
		}
		return func(err error) {
			for _, v := range fns {
				v(err)
			}
		}
	}
}
