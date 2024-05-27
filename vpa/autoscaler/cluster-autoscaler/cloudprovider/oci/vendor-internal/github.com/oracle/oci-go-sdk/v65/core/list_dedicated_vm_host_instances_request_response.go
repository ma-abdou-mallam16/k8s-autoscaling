// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package core

import (
	"fmt"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/oci/vendor-internal/github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDedicatedVmHostInstancesRequest wrapper for the ListDedicatedVmHostInstances operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/ListDedicatedVmHostInstances.go.html to see an example of how to use ListDedicatedVmHostInstancesRequest.
type ListDedicatedVmHostInstancesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The OCID of the dedicated VM host.
	DedicatedVmHostId *string `mandatory:"true" contributesTo:"path" name:"dedicatedVmHostId"`

	// The name of the availability domain.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" contributesTo:"query" name:"availabilityDomain"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by availability domain if the scope of the resource type is within a
	// single availability domain. If you call one of these "List" operations without specifying
	// an availability domain, the resources are grouped by availability domain, then sorted.
	SortBy ListDedicatedVmHostInstancesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListDedicatedVmHostInstancesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDedicatedVmHostInstancesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDedicatedVmHostInstancesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDedicatedVmHostInstancesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDedicatedVmHostInstancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDedicatedVmHostInstancesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDedicatedVmHostInstancesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDedicatedVmHostInstancesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDedicatedVmHostInstancesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDedicatedVmHostInstancesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDedicatedVmHostInstancesResponse wrapper for the ListDedicatedVmHostInstances operation
type ListDedicatedVmHostInstancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []DedicatedVmHostInstanceSummary instances
	Items []DedicatedVmHostInstanceSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListDedicatedVmHostInstancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDedicatedVmHostInstancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDedicatedVmHostInstancesSortByEnum Enum with underlying type: string
type ListDedicatedVmHostInstancesSortByEnum string

// Set of constants representing the allowable values for ListDedicatedVmHostInstancesSortByEnum
const (
	ListDedicatedVmHostInstancesSortByTimecreated ListDedicatedVmHostInstancesSortByEnum = "TIMECREATED"
	ListDedicatedVmHostInstancesSortByDisplayname ListDedicatedVmHostInstancesSortByEnum = "DISPLAYNAME"
)

var mappingListDedicatedVmHostInstancesSortByEnum = map[string]ListDedicatedVmHostInstancesSortByEnum{
	"TIMECREATED": ListDedicatedVmHostInstancesSortByTimecreated,
	"DISPLAYNAME": ListDedicatedVmHostInstancesSortByDisplayname,
}

var mappingListDedicatedVmHostInstancesSortByEnumLowerCase = map[string]ListDedicatedVmHostInstancesSortByEnum{
	"timecreated": ListDedicatedVmHostInstancesSortByTimecreated,
	"displayname": ListDedicatedVmHostInstancesSortByDisplayname,
}

// GetListDedicatedVmHostInstancesSortByEnumValues Enumerates the set of values for ListDedicatedVmHostInstancesSortByEnum
func GetListDedicatedVmHostInstancesSortByEnumValues() []ListDedicatedVmHostInstancesSortByEnum {
	values := make([]ListDedicatedVmHostInstancesSortByEnum, 0)
	for _, v := range mappingListDedicatedVmHostInstancesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDedicatedVmHostInstancesSortByEnumStringValues Enumerates the set of values in String for ListDedicatedVmHostInstancesSortByEnum
func GetListDedicatedVmHostInstancesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListDedicatedVmHostInstancesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDedicatedVmHostInstancesSortByEnum(val string) (ListDedicatedVmHostInstancesSortByEnum, bool) {
	enum, ok := mappingListDedicatedVmHostInstancesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDedicatedVmHostInstancesSortOrderEnum Enum with underlying type: string
type ListDedicatedVmHostInstancesSortOrderEnum string

// Set of constants representing the allowable values for ListDedicatedVmHostInstancesSortOrderEnum
const (
	ListDedicatedVmHostInstancesSortOrderAsc  ListDedicatedVmHostInstancesSortOrderEnum = "ASC"
	ListDedicatedVmHostInstancesSortOrderDesc ListDedicatedVmHostInstancesSortOrderEnum = "DESC"
)

var mappingListDedicatedVmHostInstancesSortOrderEnum = map[string]ListDedicatedVmHostInstancesSortOrderEnum{
	"ASC":  ListDedicatedVmHostInstancesSortOrderAsc,
	"DESC": ListDedicatedVmHostInstancesSortOrderDesc,
}

var mappingListDedicatedVmHostInstancesSortOrderEnumLowerCase = map[string]ListDedicatedVmHostInstancesSortOrderEnum{
	"asc":  ListDedicatedVmHostInstancesSortOrderAsc,
	"desc": ListDedicatedVmHostInstancesSortOrderDesc,
}

// GetListDedicatedVmHostInstancesSortOrderEnumValues Enumerates the set of values for ListDedicatedVmHostInstancesSortOrderEnum
func GetListDedicatedVmHostInstancesSortOrderEnumValues() []ListDedicatedVmHostInstancesSortOrderEnum {
	values := make([]ListDedicatedVmHostInstancesSortOrderEnum, 0)
	for _, v := range mappingListDedicatedVmHostInstancesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDedicatedVmHostInstancesSortOrderEnumStringValues Enumerates the set of values in String for ListDedicatedVmHostInstancesSortOrderEnum
func GetListDedicatedVmHostInstancesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDedicatedVmHostInstancesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDedicatedVmHostInstancesSortOrderEnum(val string) (ListDedicatedVmHostInstancesSortOrderEnum, bool) {
	enum, ok := mappingListDedicatedVmHostInstancesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
