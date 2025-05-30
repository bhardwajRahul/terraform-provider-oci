// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package networkfirewall

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListNatRulesRequest wrapper for the ListNatRules operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/ListNatRules.go.html to see an example of how to use ListNatRulesRequest.
type ListNatRulesRequest struct {

	// Unique Network Firewall Policy identifier
	NetworkFirewallPolicyId *string `mandatory:"true" contributesTo:"path" name:"networkFirewallPolicyId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` or `opc-prev-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListNatRulesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListNatRulesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique priority order for NAT Rules in the network firewall policy.
	NatRulePriorityOrder *int `mandatory:"false" contributesTo:"query" name:"natRulePriorityOrder"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListNatRulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListNatRulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListNatRulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListNatRulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListNatRulesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListNatRulesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListNatRulesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListNatRulesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListNatRulesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListNatRulesResponse wrapper for the ListNatRules operation
type ListNatRulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of NatRuleCollection instances
	NatRuleCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response, previous pages of results exist. For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. This is to get the page counts overall.
	OpcPageCount *string `presentIn:"header" name:"opc-page-count"`

	// For pagination of a list of items. This provides the count of total items across pages.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListNatRulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListNatRulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListNatRulesSortOrderEnum Enum with underlying type: string
type ListNatRulesSortOrderEnum string

// Set of constants representing the allowable values for ListNatRulesSortOrderEnum
const (
	ListNatRulesSortOrderAsc  ListNatRulesSortOrderEnum = "ASC"
	ListNatRulesSortOrderDesc ListNatRulesSortOrderEnum = "DESC"
)

var mappingListNatRulesSortOrderEnum = map[string]ListNatRulesSortOrderEnum{
	"ASC":  ListNatRulesSortOrderAsc,
	"DESC": ListNatRulesSortOrderDesc,
}

var mappingListNatRulesSortOrderEnumLowerCase = map[string]ListNatRulesSortOrderEnum{
	"asc":  ListNatRulesSortOrderAsc,
	"desc": ListNatRulesSortOrderDesc,
}

// GetListNatRulesSortOrderEnumValues Enumerates the set of values for ListNatRulesSortOrderEnum
func GetListNatRulesSortOrderEnumValues() []ListNatRulesSortOrderEnum {
	values := make([]ListNatRulesSortOrderEnum, 0)
	for _, v := range mappingListNatRulesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListNatRulesSortOrderEnumStringValues Enumerates the set of values in String for ListNatRulesSortOrderEnum
func GetListNatRulesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListNatRulesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNatRulesSortOrderEnum(val string) (ListNatRulesSortOrderEnum, bool) {
	enum, ok := mappingListNatRulesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListNatRulesSortByEnum Enum with underlying type: string
type ListNatRulesSortByEnum string

// Set of constants representing the allowable values for ListNatRulesSortByEnum
const (
	ListNatRulesSortByTimecreated ListNatRulesSortByEnum = "timeCreated"
	ListNatRulesSortByDisplayname ListNatRulesSortByEnum = "displayName"
)

var mappingListNatRulesSortByEnum = map[string]ListNatRulesSortByEnum{
	"timeCreated": ListNatRulesSortByTimecreated,
	"displayName": ListNatRulesSortByDisplayname,
}

var mappingListNatRulesSortByEnumLowerCase = map[string]ListNatRulesSortByEnum{
	"timecreated": ListNatRulesSortByTimecreated,
	"displayname": ListNatRulesSortByDisplayname,
}

// GetListNatRulesSortByEnumValues Enumerates the set of values for ListNatRulesSortByEnum
func GetListNatRulesSortByEnumValues() []ListNatRulesSortByEnum {
	values := make([]ListNatRulesSortByEnum, 0)
	for _, v := range mappingListNatRulesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListNatRulesSortByEnumStringValues Enumerates the set of values in String for ListNatRulesSortByEnum
func GetListNatRulesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListNatRulesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNatRulesSortByEnum(val string) (ListNatRulesSortByEnum, bool) {
	enum, ok := mappingListNatRulesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
