/*
Package qingstor provided support for qingstor object storage (https://www.qingcloud.com/products/qingstor/)
*/
package qingstor

//go:generate ../../internal/bin/service
//go:generate mockgen -package qingstor -destination mock_test.go github.com/yunify/qingstor-sdk-go/v3/interface Service,Bucket