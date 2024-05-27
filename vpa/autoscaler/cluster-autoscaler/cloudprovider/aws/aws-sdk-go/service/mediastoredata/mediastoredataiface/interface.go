// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package mediastoredataiface provides an interface to enable mocking the AWS Elemental MediaStore Data Plane service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package mediastoredataiface

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/service/mediastoredata"
)

// MediaStoreDataAPI provides an interface to enable mocking the
// mediastoredata.MediaStoreData service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// AWS Elemental MediaStore Data Plane.
//	func myFunc(svc mediastoredataiface.MediaStoreDataAPI) bool {
//	    // Make svc.DeleteObject request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := mediastoredata.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockMediaStoreDataClient struct {
//	    mediastoredataiface.MediaStoreDataAPI
//	}
//	func (m *mockMediaStoreDataClient) DeleteObject(input *mediastoredata.DeleteObjectInput) (*mediastoredata.DeleteObjectOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockMediaStoreDataClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type MediaStoreDataAPI interface {
	DeleteObject(*mediastoredata.DeleteObjectInput) (*mediastoredata.DeleteObjectOutput, error)
	DeleteObjectWithContext(aws.Context, *mediastoredata.DeleteObjectInput, ...request.Option) (*mediastoredata.DeleteObjectOutput, error)
	DeleteObjectRequest(*mediastoredata.DeleteObjectInput) (*request.Request, *mediastoredata.DeleteObjectOutput)

	DescribeObject(*mediastoredata.DescribeObjectInput) (*mediastoredata.DescribeObjectOutput, error)
	DescribeObjectWithContext(aws.Context, *mediastoredata.DescribeObjectInput, ...request.Option) (*mediastoredata.DescribeObjectOutput, error)
	DescribeObjectRequest(*mediastoredata.DescribeObjectInput) (*request.Request, *mediastoredata.DescribeObjectOutput)

	GetObject(*mediastoredata.GetObjectInput) (*mediastoredata.GetObjectOutput, error)
	GetObjectWithContext(aws.Context, *mediastoredata.GetObjectInput, ...request.Option) (*mediastoredata.GetObjectOutput, error)
	GetObjectRequest(*mediastoredata.GetObjectInput) (*request.Request, *mediastoredata.GetObjectOutput)

	ListItems(*mediastoredata.ListItemsInput) (*mediastoredata.ListItemsOutput, error)
	ListItemsWithContext(aws.Context, *mediastoredata.ListItemsInput, ...request.Option) (*mediastoredata.ListItemsOutput, error)
	ListItemsRequest(*mediastoredata.ListItemsInput) (*request.Request, *mediastoredata.ListItemsOutput)

	ListItemsPages(*mediastoredata.ListItemsInput, func(*mediastoredata.ListItemsOutput, bool) bool) error
	ListItemsPagesWithContext(aws.Context, *mediastoredata.ListItemsInput, func(*mediastoredata.ListItemsOutput, bool) bool, ...request.Option) error

	PutObject(*mediastoredata.PutObjectInput) (*mediastoredata.PutObjectOutput, error)
	PutObjectWithContext(aws.Context, *mediastoredata.PutObjectInput, ...request.Option) (*mediastoredata.PutObjectOutput, error)
	PutObjectRequest(*mediastoredata.PutObjectInput) (*request.Request, *mediastoredata.PutObjectOutput)
}

var _ MediaStoreDataAPI = (*mediastoredata.MediaStoreData)(nil)
