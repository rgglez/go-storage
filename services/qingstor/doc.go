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
/*
Package qingstor provided support for qingstor object storage (https://www.qingcloud.com/products/qingstor/)
*/
package qingstor

//go:generate go run github.com/golang/mock/mockgen -package qingstor -destination mock_test.go github.com/qingstor/qingstor-sdk-go/v4/interface Service,Bucket
//go:generate go run ./internal/cmd
