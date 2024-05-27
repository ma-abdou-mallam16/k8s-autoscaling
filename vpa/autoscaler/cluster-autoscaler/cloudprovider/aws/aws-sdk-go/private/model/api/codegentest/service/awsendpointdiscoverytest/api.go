// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package awsendpointdiscoverytest

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/awsutil"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/crr"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
)

const opDescribeEndpoints = "DescribeEndpoints"

// DescribeEndpointsRequest generates a "aws/request.Request" representing the
// client's request for the DescribeEndpoints operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See DescribeEndpoints for more information on using the DescribeEndpoints
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//	// Example sending a request using the DescribeEndpointsRequest method.
//	req, resp := client.DescribeEndpointsRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *AwsEndpointDiscoveryTest) DescribeEndpointsRequest(input *DescribeEndpointsInput) (req *request.Request, output *DescribeEndpointsOutput) {
	op := &request.Operation{
		Name:       opDescribeEndpoints,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeEndpointsInput{}
	}

	output = &DescribeEndpointsOutput{}
	req = c.newRequest(op, input, output)
	return
}

// DescribeEndpoints API operation for AwsEndpointDiscoveryTest.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AwsEndpointDiscoveryTest's
// API operation DescribeEndpoints for usage and error information.
func (c *AwsEndpointDiscoveryTest) DescribeEndpoints(input *DescribeEndpointsInput) (*DescribeEndpointsOutput, error) {
	req, out := c.DescribeEndpointsRequest(input)
	return out, req.Send()
}

// DescribeEndpointsWithContext is the same as DescribeEndpoints with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeEndpoints for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AwsEndpointDiscoveryTest) DescribeEndpointsWithContext(ctx aws.Context, input *DescribeEndpointsInput, opts ...request.Option) (*DescribeEndpointsOutput, error) {
	req, out := c.DescribeEndpointsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type discovererDescribeEndpoints struct {
	Client        *AwsEndpointDiscoveryTest
	Required      bool
	EndpointCache *crr.EndpointCache
	Params        map[string]*string
	Key           string
	req           *request.Request
}

func (d *discovererDescribeEndpoints) Discover() (crr.Endpoint, error) {
	input := &DescribeEndpointsInput{
		Operation: d.Params["op"],
	}

	resp, err := d.Client.DescribeEndpoints(input)
	if err != nil {
		return crr.Endpoint{}, err
	}

	endpoint := crr.Endpoint{
		Key: d.Key,
	}

	for _, e := range resp.Endpoints {
		if e.Address == nil {
			continue
		}

		address := *e.Address

		var scheme string
		if idx := strings.Index(address, "://"); idx != -1 {
			scheme = address[:idx]
		}

		if len(scheme) == 0 {
			address = fmt.Sprintf("%s://%s", d.req.HTTPRequest.URL.Scheme, address)
		}

		cachedInMinutes := aws.Int64Value(e.CachePeriodInMinutes)
		u, err := url.Parse(address)
		if err != nil {
			continue
		}

		addr := crr.WeightedAddress{
			URL:     u,
			Expired: time.Now().Add(time.Duration(cachedInMinutes) * time.Minute),
		}

		endpoint.Add(addr)
	}

	d.EndpointCache.Add(endpoint)

	return endpoint, nil
}

func (d *discovererDescribeEndpoints) Handler(r *request.Request) {
	endpointKey := crr.BuildEndpointKey(d.Params)
	d.Key = endpointKey
	d.req = r

	endpoint, err := d.EndpointCache.Get(d, endpointKey, d.Required)
	if err != nil {
		r.Error = err
		return
	}

	if endpoint.URL != nil && len(endpoint.URL.String()) > 0 {
		r.HTTPRequest.URL = endpoint.URL
	}
}

const opTestDiscoveryIdentifiersRequired = "TestDiscoveryIdentifiersRequired"

// TestDiscoveryIdentifiersRequiredRequest generates a "aws/request.Request" representing the
// client's request for the TestDiscoveryIdentifiersRequired operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See TestDiscoveryIdentifiersRequired for more information on using the TestDiscoveryIdentifiersRequired
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//	// Example sending a request using the TestDiscoveryIdentifiersRequiredRequest method.
//	req, resp := client.TestDiscoveryIdentifiersRequiredRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *AwsEndpointDiscoveryTest) TestDiscoveryIdentifiersRequiredRequest(input *TestDiscoveryIdentifiersRequiredInput) (req *request.Request, output *TestDiscoveryIdentifiersRequiredOutput) {
	op := &request.Operation{
		Name:       opTestDiscoveryIdentifiersRequired,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &TestDiscoveryIdentifiersRequiredInput{}
	}

	output = &TestDiscoveryIdentifiersRequiredOutput{}
	req = c.newRequest(op, input, output)
	// if custom endpoint for the request is set to a non empty string,
	// we skip the endpoint discovery workflow.
	if req.Config.Endpoint == nil || *req.Config.Endpoint == "" {
		de := discovererDescribeEndpoints{
			Required:      true,
			EndpointCache: c.endpointCache,
			Params: map[string]*string{
				"op":  aws.String(req.Operation.Name),
				"Sdk": input.Sdk,
			},
			Client: c,
		}

		for k, v := range de.Params {
			if v == nil {
				delete(de.Params, k)
			}
		}

		req.Handlers.Build.PushFrontNamed(request.NamedHandler{
			Name: "crr.endpointdiscovery",
			Fn:   de.Handler,
		})
	}
	return
}

// TestDiscoveryIdentifiersRequired API operation for AwsEndpointDiscoveryTest.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AwsEndpointDiscoveryTest's
// API operation TestDiscoveryIdentifiersRequired for usage and error information.
func (c *AwsEndpointDiscoveryTest) TestDiscoveryIdentifiersRequired(input *TestDiscoveryIdentifiersRequiredInput) (*TestDiscoveryIdentifiersRequiredOutput, error) {
	req, out := c.TestDiscoveryIdentifiersRequiredRequest(input)
	return out, req.Send()
}

// TestDiscoveryIdentifiersRequiredWithContext is the same as TestDiscoveryIdentifiersRequired with the addition of
// the ability to pass a context and additional request options.
//
// See TestDiscoveryIdentifiersRequired for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AwsEndpointDiscoveryTest) TestDiscoveryIdentifiersRequiredWithContext(ctx aws.Context, input *TestDiscoveryIdentifiersRequiredInput, opts ...request.Option) (*TestDiscoveryIdentifiersRequiredOutput, error) {
	req, out := c.TestDiscoveryIdentifiersRequiredRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opTestDiscoveryOptional = "TestDiscoveryOptional"

// TestDiscoveryOptionalRequest generates a "aws/request.Request" representing the
// client's request for the TestDiscoveryOptional operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See TestDiscoveryOptional for more information on using the TestDiscoveryOptional
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//	// Example sending a request using the TestDiscoveryOptionalRequest method.
//	req, resp := client.TestDiscoveryOptionalRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *AwsEndpointDiscoveryTest) TestDiscoveryOptionalRequest(input *TestDiscoveryOptionalInput) (req *request.Request, output *TestDiscoveryOptionalOutput) {
	op := &request.Operation{
		Name:       opTestDiscoveryOptional,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &TestDiscoveryOptionalInput{}
	}

	output = &TestDiscoveryOptionalOutput{}
	req = c.newRequest(op, input, output)
	// if custom endpoint for the request is set to a non empty string,
	// we skip the endpoint discovery workflow.
	if req.Config.Endpoint == nil || *req.Config.Endpoint == "" {
		if aws.BoolValue(req.Config.EnableEndpointDiscovery) {
			de := discovererDescribeEndpoints{
				Required:      false,
				EndpointCache: c.endpointCache,
				Params: map[string]*string{
					"op": aws.String(req.Operation.Name),
				},
				Client: c,
			}

			for k, v := range de.Params {
				if v == nil {
					delete(de.Params, k)
				}
			}

			req.Handlers.Build.PushFrontNamed(request.NamedHandler{
				Name: "crr.endpointdiscovery",
				Fn:   de.Handler,
			})
		}
	}
	return
}

// TestDiscoveryOptional API operation for AwsEndpointDiscoveryTest.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AwsEndpointDiscoveryTest's
// API operation TestDiscoveryOptional for usage and error information.
func (c *AwsEndpointDiscoveryTest) TestDiscoveryOptional(input *TestDiscoveryOptionalInput) (*TestDiscoveryOptionalOutput, error) {
	req, out := c.TestDiscoveryOptionalRequest(input)
	return out, req.Send()
}

// TestDiscoveryOptionalWithContext is the same as TestDiscoveryOptional with the addition of
// the ability to pass a context and additional request options.
//
// See TestDiscoveryOptional for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AwsEndpointDiscoveryTest) TestDiscoveryOptionalWithContext(ctx aws.Context, input *TestDiscoveryOptionalInput, opts ...request.Option) (*TestDiscoveryOptionalOutput, error) {
	req, out := c.TestDiscoveryOptionalRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opTestDiscoveryRequired = "TestDiscoveryRequired"

// TestDiscoveryRequiredRequest generates a "aws/request.Request" representing the
// client's request for the TestDiscoveryRequired operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See TestDiscoveryRequired for more information on using the TestDiscoveryRequired
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//	// Example sending a request using the TestDiscoveryRequiredRequest method.
//	req, resp := client.TestDiscoveryRequiredRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *AwsEndpointDiscoveryTest) TestDiscoveryRequiredRequest(input *TestDiscoveryRequiredInput) (req *request.Request, output *TestDiscoveryRequiredOutput) {
	op := &request.Operation{
		Name:       opTestDiscoveryRequired,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &TestDiscoveryRequiredInput{}
	}

	output = &TestDiscoveryRequiredOutput{}
	req = c.newRequest(op, input, output)
	// if custom endpoint for the request is set to a non empty string,
	// we skip the endpoint discovery workflow.
	if req.Config.Endpoint == nil || *req.Config.Endpoint == "" {
		if aws.BoolValue(req.Config.EnableEndpointDiscovery) {
			de := discovererDescribeEndpoints{
				Required:      false,
				EndpointCache: c.endpointCache,
				Params: map[string]*string{
					"op": aws.String(req.Operation.Name),
				},
				Client: c,
			}

			for k, v := range de.Params {
				if v == nil {
					delete(de.Params, k)
				}
			}

			req.Handlers.Build.PushFrontNamed(request.NamedHandler{
				Name: "crr.endpointdiscovery",
				Fn:   de.Handler,
			})
		}
	}
	return
}

// TestDiscoveryRequired API operation for AwsEndpointDiscoveryTest.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AwsEndpointDiscoveryTest's
// API operation TestDiscoveryRequired for usage and error information.
func (c *AwsEndpointDiscoveryTest) TestDiscoveryRequired(input *TestDiscoveryRequiredInput) (*TestDiscoveryRequiredOutput, error) {
	req, out := c.TestDiscoveryRequiredRequest(input)
	return out, req.Send()
}

// TestDiscoveryRequiredWithContext is the same as TestDiscoveryRequired with the addition of
// the ability to pass a context and additional request options.
//
// See TestDiscoveryRequired for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AwsEndpointDiscoveryTest) TestDiscoveryRequiredWithContext(ctx aws.Context, input *TestDiscoveryRequiredInput, opts ...request.Option) (*TestDiscoveryRequiredOutput, error) {
	req, out := c.TestDiscoveryRequiredRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeEndpointsInput struct {
	_ struct{} `type:"structure"`

	Operation *string `type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s DescribeEndpointsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s DescribeEndpointsInput) GoString() string {
	return s.String()
}

// SetOperation sets the Operation field's value.
func (s *DescribeEndpointsInput) SetOperation(v string) *DescribeEndpointsInput {
	s.Operation = &v
	return s
}

type DescribeEndpointsOutput struct {
	_ struct{} `type:"structure"`

	// Endpoints is a required field
	Endpoints []*Endpoint `type:"list" required:"true"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s DescribeEndpointsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s DescribeEndpointsOutput) GoString() string {
	return s.String()
}

// SetEndpoints sets the Endpoints field's value.
func (s *DescribeEndpointsOutput) SetEndpoints(v []*Endpoint) *DescribeEndpointsOutput {
	s.Endpoints = v
	return s
}

type Endpoint struct {
	_ struct{} `type:"structure"`

	// Address is a required field
	Address *string `type:"string" required:"true"`

	// CachePeriodInMinutes is a required field
	CachePeriodInMinutes *int64 `type:"long" required:"true"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s Endpoint) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s Endpoint) GoString() string {
	return s.String()
}

// SetAddress sets the Address field's value.
func (s *Endpoint) SetAddress(v string) *Endpoint {
	s.Address = &v
	return s
}

// SetCachePeriodInMinutes sets the CachePeriodInMinutes field's value.
func (s *Endpoint) SetCachePeriodInMinutes(v int64) *Endpoint {
	s.CachePeriodInMinutes = &v
	return s
}

type TestDiscoveryIdentifiersRequiredInput struct {
	_ struct{} `type:"structure"`

	// Sdk is a required field
	Sdk *string `type:"string" required:"true"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s TestDiscoveryIdentifiersRequiredInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s TestDiscoveryIdentifiersRequiredInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TestDiscoveryIdentifiersRequiredInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "TestDiscoveryIdentifiersRequiredInput"}
	if s.Sdk == nil {
		invalidParams.Add(request.NewErrParamRequired("Sdk"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetSdk sets the Sdk field's value.
func (s *TestDiscoveryIdentifiersRequiredInput) SetSdk(v string) *TestDiscoveryIdentifiersRequiredInput {
	s.Sdk = &v
	return s
}

type TestDiscoveryIdentifiersRequiredOutput struct {
	_ struct{} `type:"structure"`

	RequestSuccessful *bool `type:"boolean"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s TestDiscoveryIdentifiersRequiredOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s TestDiscoveryIdentifiersRequiredOutput) GoString() string {
	return s.String()
}

// SetRequestSuccessful sets the RequestSuccessful field's value.
func (s *TestDiscoveryIdentifiersRequiredOutput) SetRequestSuccessful(v bool) *TestDiscoveryIdentifiersRequiredOutput {
	s.RequestSuccessful = &v
	return s
}

type TestDiscoveryOptionalInput struct {
	_ struct{} `type:"structure"`

	Sdk *string `type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s TestDiscoveryOptionalInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s TestDiscoveryOptionalInput) GoString() string {
	return s.String()
}

// SetSdk sets the Sdk field's value.
func (s *TestDiscoveryOptionalInput) SetSdk(v string) *TestDiscoveryOptionalInput {
	s.Sdk = &v
	return s
}

type TestDiscoveryOptionalOutput struct {
	_ struct{} `type:"structure"`

	RequestSuccessful *bool `type:"boolean"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s TestDiscoveryOptionalOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s TestDiscoveryOptionalOutput) GoString() string {
	return s.String()
}

// SetRequestSuccessful sets the RequestSuccessful field's value.
func (s *TestDiscoveryOptionalOutput) SetRequestSuccessful(v bool) *TestDiscoveryOptionalOutput {
	s.RequestSuccessful = &v
	return s
}

type TestDiscoveryRequiredInput struct {
	_ struct{} `type:"structure"`

	Sdk *string `type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s TestDiscoveryRequiredInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s TestDiscoveryRequiredInput) GoString() string {
	return s.String()
}

// SetSdk sets the Sdk field's value.
func (s *TestDiscoveryRequiredInput) SetSdk(v string) *TestDiscoveryRequiredInput {
	s.Sdk = &v
	return s
}

type TestDiscoveryRequiredOutput struct {
	_ struct{} `type:"structure"`

	RequestSuccessful *bool `type:"boolean"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s TestDiscoveryRequiredOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s TestDiscoveryRequiredOutput) GoString() string {
	return s.String()
}

// SetRequestSuccessful sets the RequestSuccessful field's value.
func (s *TestDiscoveryRequiredOutput) SetRequestSuccessful(v bool) *TestDiscoveryRequiredOutput {
	s.RequestSuccessful = &v
	return s
}
