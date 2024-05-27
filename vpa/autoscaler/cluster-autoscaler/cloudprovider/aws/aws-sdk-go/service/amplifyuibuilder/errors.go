// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package amplifyuibuilder

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// An internal error has occurred. Please retry your request.
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodeInvalidParameterException for service response error code
	// "InvalidParameterException".
	//
	// An invalid or out-of-range value was supplied for the input parameter.
	ErrCodeInvalidParameterException = "InvalidParameterException"

	// ErrCodeResourceConflictException for service response error code
	// "ResourceConflictException".
	//
	// The resource specified in the request conflicts with an existing resource.
	ErrCodeResourceConflictException = "ResourceConflictException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The requested resource does not exist, or access was denied.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServiceQuotaExceededException for service response error code
	// "ServiceQuotaExceededException".
	//
	// You exceeded your service quota. Service quotas, also referred to as limits,
	// are the maximum number of service resources or operations for your Amazon
	// Web Services account.
	ErrCodeServiceQuotaExceededException = "ServiceQuotaExceededException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// The request was denied due to request throttling.
	ErrCodeThrottlingException = "ThrottlingException"

	// ErrCodeUnauthorizedException for service response error code
	// "UnauthorizedException".
	//
	// You don't have permission to perform this operation.
	ErrCodeUnauthorizedException = "UnauthorizedException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"InternalServerException":       newErrorInternalServerException,
	"InvalidParameterException":     newErrorInvalidParameterException,
	"ResourceConflictException":     newErrorResourceConflictException,
	"ResourceNotFoundException":     newErrorResourceNotFoundException,
	"ServiceQuotaExceededException": newErrorServiceQuotaExceededException,
	"ThrottlingException":           newErrorThrottlingException,
	"UnauthorizedException":         newErrorUnauthorizedException,
}
