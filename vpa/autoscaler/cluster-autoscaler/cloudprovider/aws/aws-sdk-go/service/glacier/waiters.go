// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glacier

import (
	"time"

	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
)

// WaitUntilVaultExists uses the Amazon Glacier API operation
// DescribeVault to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *Glacier) WaitUntilVaultExists(input *DescribeVaultInput) error {
	return c.WaitUntilVaultExistsWithContext(aws.BackgroundContext(), input)
}

// WaitUntilVaultExistsWithContext is an extended version of WaitUntilVaultExists.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Glacier) WaitUntilVaultExistsWithContext(ctx aws.Context, input *DescribeVaultInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilVaultExists",
		MaxAttempts: 15,
		Delay:       request.ConstantWaiterDelay(3 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 200,
			},
			{
				State:    request.RetryWaiterState,
				Matcher:  request.ErrorWaiterMatch,
				Expected: "ResourceNotFoundException",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeVaultInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeVaultRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilVaultNotExists uses the Amazon Glacier API operation
// DescribeVault to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *Glacier) WaitUntilVaultNotExists(input *DescribeVaultInput) error {
	return c.WaitUntilVaultNotExistsWithContext(aws.BackgroundContext(), input)
}

// WaitUntilVaultNotExistsWithContext is an extended version of WaitUntilVaultNotExists.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Glacier) WaitUntilVaultNotExistsWithContext(ctx aws.Context, input *DescribeVaultInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilVaultNotExists",
		MaxAttempts: 15,
		Delay:       request.ConstantWaiterDelay(3 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:    request.RetryWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 200,
			},
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.ErrorWaiterMatch,
				Expected: "ResourceNotFoundException",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeVaultInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeVaultRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}
