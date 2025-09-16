package client

import (
	openapi "github.com/rancher/muchang/darabonba-openapi/client"
	openapiutil "github.com/rancher/muchang/darabonba-openapi/utils"
	"github.com/rancher/muchang/utils/tea"
	util "github.com/rancher/muchang/utils/tea-utils/service"
)

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("resourcemanager"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := openapiutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

type ListResourceGroupsRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListResourceGroupsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupsRequestTag) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsRequestTag) SetKey(v string) *ListResourceGroupsRequestTag {
	s.Key = &v
	return s
}

func (s *ListResourceGroupsRequestTag) SetValue(v string) *ListResourceGroupsRequestTag {
	s.Value = &v
	return s
}

type ListResourceGroupsResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4B450CA1-36E8-4AA2-8461-86B42BF4CC4E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the resource groups.
	ResourceGroups *ListResourceGroupsResponseBodyResourceGroups `json:"ResourceGroups,omitempty" xml:"ResourceGroups,omitempty" type:"Struct"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourceGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponseBody) SetPageNumber(v int32) *ListResourceGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListResourceGroupsResponseBody) SetPageSize(v int32) *ListResourceGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListResourceGroupsResponseBody) SetRequestId(v string) *ListResourceGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceGroupsResponseBody) SetResourceGroups(v *ListResourceGroupsResponseBodyResourceGroups) *ListResourceGroupsResponseBody {
	s.ResourceGroups = v
	return s
}

func (s *ListResourceGroupsResponseBody) SetTotalCount(v int32) *ListResourceGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type ListResourceGroupsResponseBodyResourceGroups struct {
	ResourceGroup []*ListResourceGroupsResponseBodyResourceGroupsResourceGroup `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty" type:"Repeated"`
}

func (s ListResourceGroupsResponseBodyResourceGroups) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupsResponseBodyResourceGroups) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponseBodyResourceGroups) SetResourceGroup(v []*ListResourceGroupsResponseBodyResourceGroupsResourceGroup) *ListResourceGroupsResponseBodyResourceGroups {
	s.ResourceGroup = v
	return s
}

type ListResourceGroupsResponseBodyResourceGroupsResourceGroup struct {
	// The ID of the Alibaba Cloud account to which the resource group belongs.
	//
	// example:
	//
	// 123456789****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The time when the resource group was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The display name of the resource group.
	//
	// example:
	//
	// my-project
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-9gLOoK****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The identifier of the resource group.
	//
	// example:
	//
	// my-project
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the resource group. Valid values:
	//
	// 	- Creating: The resource group is being created.
	//
	// 	- OK: The resource group is created.
	//
	// 	- PendingDelete: The resource group is waiting to be deleted.
	//
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags that are added to the resource group.
	Tags *ListResourceGroupsResponseBodyResourceGroupsResourceGroupTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s ListResourceGroupsResponseBodyResourceGroupsResourceGroup) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupsResponseBodyResourceGroupsResourceGroup) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroup) SetAccountId(v string) *ListResourceGroupsResponseBodyResourceGroupsResourceGroup {
	s.AccountId = &v
	return s
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroup) SetCreateDate(v string) *ListResourceGroupsResponseBodyResourceGroupsResourceGroup {
	s.CreateDate = &v
	return s
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroup) SetDisplayName(v string) *ListResourceGroupsResponseBodyResourceGroupsResourceGroup {
	s.DisplayName = &v
	return s
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroup) SetId(v string) *ListResourceGroupsResponseBodyResourceGroupsResourceGroup {
	s.Id = &v
	return s
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroup) SetName(v string) *ListResourceGroupsResponseBodyResourceGroupsResourceGroup {
	s.Name = &v
	return s
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroup) SetStatus(v string) *ListResourceGroupsResponseBodyResourceGroupsResourceGroup {
	s.Status = &v
	return s
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroup) SetTags(v *ListResourceGroupsResponseBodyResourceGroupsResourceGroupTags) *ListResourceGroupsResponseBodyResourceGroupsResourceGroup {
	s.Tags = v
	return s
}

type ListResourceGroupsResponseBodyResourceGroupsResourceGroupTags struct {
	Tag []*ListResourceGroupsResponseBodyResourceGroupsResourceGroupTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListResourceGroupsResponseBodyResourceGroupsResourceGroupTags) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupsResponseBodyResourceGroupsResourceGroupTags) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroupTags) SetTag(v []*ListResourceGroupsResponseBodyResourceGroupsResourceGroupTagsTag) *ListResourceGroupsResponseBodyResourceGroupsResourceGroupTags {
	s.Tag = v
	return s
}

type ListResourceGroupsResponseBodyResourceGroupsResourceGroupTagsTag struct {
	// The tag key.
	//
	// example:
	//
	// k1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// v1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListResourceGroupsResponseBodyResourceGroupsResourceGroupTagsTag) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupsResponseBodyResourceGroupsResourceGroupTagsTag) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroupTagsTag) SetTagKey(v string) *ListResourceGroupsResponseBodyResourceGroupsResourceGroupTagsTag {
	s.TagKey = &v
	return s
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroupTagsTag) SetTagValue(v string) *ListResourceGroupsResponseBodyResourceGroupsResourceGroupTagsTag {
	s.TagValue = &v
	return s
}

type ListResourceGroupsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponse) SetHeaders(v map[string]*string) *ListResourceGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListResourceGroupsResponse) SetStatusCode(v int32) *ListResourceGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceGroupsResponse) SetBody(v *ListResourceGroupsResponseBody) *ListResourceGroupsResponse {
	s.Body = v
	return s
}

type ListResourceGroupsRequest struct {
	// The display name of the resource group. This parameter specifies a filter condition for the query. Fuzzy match is supported.
	//
	// The display name can be a maximum of 50 characters in length.
	//
	// example:
	//
	// my-project
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// Specifies whether to return the information of tags. Valid values:
	//
	// 	- false (default value)
	//
	// 	- true
	//
	// >  If you configure the Tag parameter, the system returns the information of tags regardless of the setting of the `IncludeTags` parameter.
	//
	// example:
	//
	// false
	IncludeTags *bool `json:"IncludeTags,omitempty" xml:"IncludeTags,omitempty"`
	// The identifier of the resource group. This parameter specifies a filter condition for the query. Fuzzy match is supported.
	//
	// The identifier can be a maximum of 50 characters in length and can contain letters, digits, and hyphens (-).
	//
	// example:
	//
	// my-project
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource group. This parameter specifies a filter condition for the query.
	//
	// The ID can be a maximum of 18 characters in length and must start with `rg-`.
	//
	// >  This parameter is incorporated into the `ResourceGroupIds` parameter. If you configure both the `ResourceGroupId` and `ResourceGroupIds` parameters, the value of the `ResourceGroupIds` parameter prevails.
	//
	// example:
	//
	// rg-9gLOoK****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The IDs of the resource groups. This parameter specifies a filter condition for the query.
	//
	// You can specify a maximum of 100 resource group IDs.
	//
	// >  If you configure both the `ResourceGroupId` and `ResourceGroupIds` parameters, the value of the `ResourceGroupIds` parameter prevails.
	ResourceGroupIds []*string `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty" type:"Repeated"`
	// The status of the resource group. This parameter specifies a filter condition for the query. Valid values:
	//
	// 	- Creating: The resource group is being created.
	//
	// 	- OK: The resource group is created.
	//
	// 	- PendingDelete: The resource group is waiting to be deleted.
	//
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag. This parameter specifies a filter condition for the query.
	Tag []*ListResourceGroupsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListResourceGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsRequest) SetDisplayName(v string) *ListResourceGroupsRequest {
	s.DisplayName = &v
	return s
}

func (s *ListResourceGroupsRequest) SetIncludeTags(v bool) *ListResourceGroupsRequest {
	s.IncludeTags = &v
	return s
}

func (s *ListResourceGroupsRequest) SetName(v string) *ListResourceGroupsRequest {
	s.Name = &v
	return s
}

func (s *ListResourceGroupsRequest) SetPageNumber(v int32) *ListResourceGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourceGroupsRequest) SetPageSize(v int32) *ListResourceGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourceGroupsRequest) SetResourceGroupId(v string) *ListResourceGroupsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListResourceGroupsRequest) SetResourceGroupIds(v []*string) *ListResourceGroupsRequest {
	s.ResourceGroupIds = v
	return s
}

func (s *ListResourceGroupsRequest) SetStatus(v string) *ListResourceGroupsRequest {
	s.Status = &v
	return s
}

func (s *ListResourceGroupsRequest) SetTag(v []*ListResourceGroupsRequestTag) *ListResourceGroupsRequest {
	s.Tag = v
	return s
}

// Description:
//
// You can call this API operation to query all resource groups within the current account. You can also call this API operation to query a specific resource group based on the status, ID, identifier, or display name of the resource group.
//
// This topic provides an example on how to call the API operation to query the basic information about the resource groups `rg-1hSBH2****` and `rg-9gLOoK****` within the current account.
//
// @param request - ListResourceGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceGroupsResponse
func (client *Client) ListResourceGroupsWithOptions(request *ListResourceGroupsRequest, runtime *util.RuntimeOptions) (_result *ListResourceGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		query["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeTags)) {
		query["IncludeTags"] = request.IncludeTags
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupIds)) {
		query["ResourceGroupIds"] = request.ResourceGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourceGroups"),
		Version:     tea.String("2020-03-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListResourceGroupsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListResourceGroupsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Description:
//
// You can call this API operation to query all resource groups within the current account. You can also call this API operation to query a specific resource group based on the status, ID, identifier, or display name of the resource group.
//
// This topic provides an example on how to call the API operation to query the basic information about the resource groups `rg-1hSBH2****` and `rg-9gLOoK****` within the current account.
//
// @param request - ListResourceGroupsRequest
//
// @return ListResourceGroupsResponse
func (client *Client) ListResourceGroups(request *ListResourceGroupsRequest) (_result *ListResourceGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListResourceGroupsResponse{}
	_body, _err := client.ListResourceGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
