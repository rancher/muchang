package client

import (
	"context"

	openapiutil "github.com/rancher/muchang/darabonba-openapi/utils"
	"github.com/rancher/muchang/utils/tea/dara"
)

// Summary:
//
// Queries Alibaba Cloud regions. When you call this operation, you can specify parameters, such as InstanceChargeType and ResourceType, in the request.
//
// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithContext(ctx context.Context, request *DescribeRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.InstanceChargeType) {
		query["InstanceChargeType"] = request.InstanceChargeType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2014-05-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of all instance types or a specific instance type provided by Elastic Compute Service (ECS). You can understand the specifications and performance of instance types based on the response and select an instance type that meets your business requirements.
//
// Description:
//
//	  **Paged query**: You can set MaxResults to specify the maximum number of entries to return in a single call. If the number of entries to return exceeds the specified MaxResults value, the response includes a NextToken value. You can set NextToken to the return value and specify MaxResults in your next request to DescribeInstanceTypes to retrieve the next page of results.
//
//		- When you call this operation, if you do not set NextToken to paginate the results, only the first page of results is returned by default and includes a maximum of 100 entries. To retrieve further pages of results, set NextToken or pass filter conditions in your requests to DescribeInstanceTypes.
//
// >  MaxResults specifies the maximum number of entries per page. The maximum value of this parameter is changed from 1600 to 100 for all users as of November 15, 2023. If you called the DescribeInstanceTypes operation in 2022, you can use 1600 as the maximum value before November 15, 2023.
//
//   - The DescribeInstanceTypes operation is used to query only the specifications and performance information of instance types. To query instance types that are available in a specific region, call the [DescribeAvailableResource](https://help.aliyun.com/document_detail/66186.html) operation.
//
//   - To use special instance types such as instance types that are unavailable for purchase, [submit a ticket](https://smartservice.console.aliyun.com/service/create-ticket-intl).
//
// @param request - DescribeInstanceTypesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceTypesResponse
func (client *Client) DescribeInstanceTypesWithContext(ctx context.Context, request *DescribeInstanceTypesRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceTypesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdditionalAttributes) {
		query["AdditionalAttributes"] = request.AdditionalAttributes
	}

	if !dara.IsNil(request.CpuArchitecture) {
		query["CpuArchitecture"] = request.CpuArchitecture
	}

	if !dara.IsNil(request.CpuArchitectures) {
		query["CpuArchitectures"] = request.CpuArchitectures
	}

	if !dara.IsNil(request.GPUSpec) {
		query["GPUSpec"] = request.GPUSpec
	}

	if !dara.IsNil(request.GpuSpecs) {
		query["GpuSpecs"] = request.GpuSpecs
	}

	if !dara.IsNil(request.InstanceCategories) {
		query["InstanceCategories"] = request.InstanceCategories
	}

	if !dara.IsNil(request.InstanceCategory) {
		query["InstanceCategory"] = request.InstanceCategory
	}

	if !dara.IsNil(request.InstanceFamilyLevel) {
		query["InstanceFamilyLevel"] = request.InstanceFamilyLevel
	}

	if !dara.IsNil(request.InstanceTypeFamilies) {
		query["InstanceTypeFamilies"] = request.InstanceTypeFamilies
	}

	if !dara.IsNil(request.InstanceTypeFamily) {
		query["InstanceTypeFamily"] = request.InstanceTypeFamily
	}

	if !dara.IsNil(request.InstanceTypes) {
		query["InstanceTypes"] = request.InstanceTypes
	}

	if !dara.IsNil(request.LocalStorageCategories) {
		query["LocalStorageCategories"] = request.LocalStorageCategories
	}

	if !dara.IsNil(request.LocalStorageCategory) {
		query["LocalStorageCategory"] = request.LocalStorageCategory
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.MaximumCpuCoreCount) {
		query["MaximumCpuCoreCount"] = request.MaximumCpuCoreCount
	}

	if !dara.IsNil(request.MaximumCpuSpeedFrequency) {
		query["MaximumCpuSpeedFrequency"] = request.MaximumCpuSpeedFrequency
	}

	if !dara.IsNil(request.MaximumCpuTurboFrequency) {
		query["MaximumCpuTurboFrequency"] = request.MaximumCpuTurboFrequency
	}

	if !dara.IsNil(request.MaximumGPUAmount) {
		query["MaximumGPUAmount"] = request.MaximumGPUAmount
	}

	if !dara.IsNil(request.MaximumMemorySize) {
		query["MaximumMemorySize"] = request.MaximumMemorySize
	}

	if !dara.IsNil(request.MinimumBaselineCredit) {
		query["MinimumBaselineCredit"] = request.MinimumBaselineCredit
	}

	if !dara.IsNil(request.MinimumCpuCoreCount) {
		query["MinimumCpuCoreCount"] = request.MinimumCpuCoreCount
	}

	if !dara.IsNil(request.MinimumCpuSpeedFrequency) {
		query["MinimumCpuSpeedFrequency"] = request.MinimumCpuSpeedFrequency
	}

	if !dara.IsNil(request.MinimumCpuTurboFrequency) {
		query["MinimumCpuTurboFrequency"] = request.MinimumCpuTurboFrequency
	}

	if !dara.IsNil(request.MinimumDiskQuantity) {
		query["MinimumDiskQuantity"] = request.MinimumDiskQuantity
	}

	if !dara.IsNil(request.MinimumEniIpv6AddressQuantity) {
		query["MinimumEniIpv6AddressQuantity"] = request.MinimumEniIpv6AddressQuantity
	}

	if !dara.IsNil(request.MinimumEniPrivateIpAddressQuantity) {
		query["MinimumEniPrivateIpAddressQuantity"] = request.MinimumEniPrivateIpAddressQuantity
	}

	if !dara.IsNil(request.MinimumEniQuantity) {
		query["MinimumEniQuantity"] = request.MinimumEniQuantity
	}

	if !dara.IsNil(request.MinimumEriQuantity) {
		query["MinimumEriQuantity"] = request.MinimumEriQuantity
	}

	if !dara.IsNil(request.MinimumGPUAmount) {
		query["MinimumGPUAmount"] = request.MinimumGPUAmount
	}

	if !dara.IsNil(request.MinimumInitialCredit) {
		query["MinimumInitialCredit"] = request.MinimumInitialCredit
	}

	if !dara.IsNil(request.MinimumInstanceBandwidthRx) {
		query["MinimumInstanceBandwidthRx"] = request.MinimumInstanceBandwidthRx
	}

	if !dara.IsNil(request.MinimumInstanceBandwidthTx) {
		query["MinimumInstanceBandwidthTx"] = request.MinimumInstanceBandwidthTx
	}

	if !dara.IsNil(request.MinimumInstancePpsRx) {
		query["MinimumInstancePpsRx"] = request.MinimumInstancePpsRx
	}

	if !dara.IsNil(request.MinimumInstancePpsTx) {
		query["MinimumInstancePpsTx"] = request.MinimumInstancePpsTx
	}

	if !dara.IsNil(request.MinimumLocalStorageAmount) {
		query["MinimumLocalStorageAmount"] = request.MinimumLocalStorageAmount
	}

	if !dara.IsNil(request.MinimumLocalStorageCapacity) {
		query["MinimumLocalStorageCapacity"] = request.MinimumLocalStorageCapacity
	}

	if !dara.IsNil(request.MinimumMemorySize) {
		query["MinimumMemorySize"] = request.MinimumMemorySize
	}

	if !dara.IsNil(request.MinimumPrimaryEniQueueNumber) {
		query["MinimumPrimaryEniQueueNumber"] = request.MinimumPrimaryEniQueueNumber
	}

	if !dara.IsNil(request.MinimumQueuePairNumber) {
		query["MinimumQueuePairNumber"] = request.MinimumQueuePairNumber
	}

	if !dara.IsNil(request.MinimumSecondaryEniQueueNumber) {
		query["MinimumSecondaryEniQueueNumber"] = request.MinimumSecondaryEniQueueNumber
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.NvmeSupport) {
		query["NvmeSupport"] = request.NvmeSupport
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhysicalProcessorModel) {
		query["PhysicalProcessorModel"] = request.PhysicalProcessorModel
	}

	if !dara.IsNil(request.PhysicalProcessorModels) {
		query["PhysicalProcessorModels"] = request.PhysicalProcessorModels
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceTypes"),
		Version:     dara.String("2014-05-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceTypesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries one or more key pairs.
//
// @param request - DescribeKeyPairsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeKeyPairsResponse
func (client *Client) DescribeKeyPairsWithContext(ctx context.Context, request *DescribeKeyPairsRequest, runtime *dara.RuntimeOptions) (_result *DescribeKeyPairsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IncludePublicKey) {
		query["IncludePublicKey"] = request.IncludePublicKey
	}

	if !dara.IsNil(request.KeyPairFingerPrint) {
		query["KeyPairFingerPrint"] = request.KeyPairFingerPrint
	}

	if !dara.IsNil(request.KeyPairName) {
		query["KeyPairName"] = request.KeyPairName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeKeyPairs"),
		Version:     dara.String("2014-05-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeKeyPairsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries resources in a zone. You can query the resources available in a zone before you create Elastic Compute Service (ECS) instances by calling the RunInstances operation or before you change instance types by calling the ModifyInstanceSpec operation.
//
// Description:
//
// The value of `DestinationResource` determines whether you need to specify additional parameters. When you select a value in the following chain for DestinationResource, the more to the right the selected value is ordered, the more parameters you must specify.
//
//   - Sequence: `Zone > IoOptimized > InstanceType = Network = ddh > SystemDisk > DataDisk`
//
//   - Examples:
//
//   - If you set `DestinationResource` to `DataDisk`, take note of the following items:
//
//   - If you set `ResourceType` to `disk` to query the categories of data disks regardless of whether the disks are attached to ECS instances, you can leave `InstanceType` empty.
//
//   - If you set `ResourceType` to `instance` to query the categories of data disks that are attached to ECS instances, you must specify `InstanceType` and `DataDiskCategory` due to instance type-specific limits on data disks.
//
//   - If you set `DestinationResource` to `SystemDisk` and `ResourceType` to `instance`, you must specify `InstanceType` due to instance type-specific limits on system disks.
//
//   - If you set `DestinationResource` to `InstanceType`, we recommend that you specify `IoOptimized` and `InstanceType`.
//
//   - To query the ecs.g5.large instance type in all zones of the China (Hangzhou) region, set `RegionId to cn-hangzhou, DestinationResource to InstanceType, IoOptimized to optimized, and InstanceType to ecs.g5.large`.
//
//   - To query the zones in which the ecs.g5.large instance type is available in the China (Hangzhou) region, set `RegionId to cn-hangzhou, DestinationResource to Zone, IoOptimized to optimized, and InstanceType to ecs.g5.large`.
//
// **To query the zones in which the ecs.g5.large instance type is available in the China (Hangzhou) region, specify parameters as follows:**
//
//	"RegionId": "cn-hangzhou",
//
//	"DestinationResource": "Zone",
//
//	"InstanceType": "ecs.g5.large"
//
// **To query the ecs.g5.large instance type in all zones of the China (Hangzhou) region, specify parameters as follows:**
//
//	"RegionId": "cn-hangzhou",
//
//	"DestinationResource": "InstanceType""InstanceType": "ecs.g5.large"
//
// **To query data disks of the ultra disk category in Hangzhou Zone B regardless of whether the disks are attached to ECS instances, specify parameters as follows:**
//
//	"RegionId": "cn-hangzhou",
//
//	"ZoneId": "cn-hangzhou-b",
//
//	"ResourceType": "disk",
//
//	"DestinationResource": "DataDisk"
//
// **To query data disks purchased together with ecs.g7.large instances that reside in Hangzhou Zone B and use Enterprise SSDs (ESSDs) as system disks, specify parameters as follows:**
//
//	"RegionId": "cn-hangzhou",
//
//	"ZoneId": "cn-hangzhou-b",
//
//	"ResourceType": "instance",
//
//	"InstanceType": "ecs.g7.large",
//
//	"DestinationResource": "SystemDisk",
//
//	"SystemDiskCategory": "cloud_essd"
//
// @param request - DescribeAvailableResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAvailableResourceResponse
func (client *Client) DescribeAvailableResourceWithContext(ctx context.Context, request *DescribeAvailableResourceRequest, runtime *dara.RuntimeOptions) (_result *DescribeAvailableResourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cores) {
		query["Cores"] = request.Cores
	}

	if !dara.IsNil(request.DataDiskCategory) {
		query["DataDiskCategory"] = request.DataDiskCategory
	}

	if !dara.IsNil(request.DedicatedHostId) {
		query["DedicatedHostId"] = request.DedicatedHostId
	}

	if !dara.IsNil(request.DestinationResource) {
		query["DestinationResource"] = request.DestinationResource
	}

	if !dara.IsNil(request.InstanceChargeType) {
		query["InstanceChargeType"] = request.InstanceChargeType
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.IoOptimized) {
		query["IoOptimized"] = request.IoOptimized
	}

	if !dara.IsNil(request.Memory) {
		query["Memory"] = request.Memory
	}

	if !dara.IsNil(request.NetworkCategory) {
		query["NetworkCategory"] = request.NetworkCategory
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	if !dara.IsNil(request.SpotDuration) {
		query["SpotDuration"] = request.SpotDuration
	}

	if !dara.IsNil(request.SpotStrategy) {
		query["SpotStrategy"] = request.SpotStrategy
	}

	if !dara.IsNil(request.SystemDiskCategory) {
		query["SystemDiskCategory"] = request.SystemDiskCategory
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAvailableResource"),
		Version:     dara.String("2014-05-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAvailableResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Alibaba Cloud regions. You can specify parameters, such as InstanceChargeType and ResourceType, in the request.
//
// Description:
//
// ## [](#)Usage notes
//
// When you call this operation, only a list of zones and some resource information of each zone are returned. If you want to query instance types and disk categories that are available for purchase in a specific zone, we recommend that you call the [DescribeAvailableResource](https://help.aliyun.com/document_detail/66186.html) operation.
//
// @param request - DescribeZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeZonesResponse
func (client *Client) DescribeZonesWithContext(ctx context.Context, request *DescribeZonesRequest, runtime *dara.RuntimeOptions) (_result *DescribeZonesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.InstanceChargeType) {
		query["InstanceChargeType"] = request.InstanceChargeType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SpotStrategy) {
		query["SpotStrategy"] = request.SpotStrategy
	}

	if !dara.IsNil(request.Verbose) {
		query["Verbose"] = request.Verbose
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeZones"),
		Version:     dara.String("2014-05-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeZonesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Elastic Compute Service (ECS) instance types supported by an image.
//
// @param request - DescribeImageSupportInstanceTypesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeImageSupportInstanceTypesResponse
func (client *Client) DescribeImageSupportInstanceTypesWithContext(ctx context.Context, request *DescribeImageSupportInstanceTypesRequest, runtime *dara.RuntimeOptions) (_result *DescribeImageSupportInstanceTypesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionType) {
		query["ActionType"] = request.ActionType
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeImageSupportInstanceTypes"),
		Version:     dara.String("2014-05-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeImageSupportInstanceTypesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
