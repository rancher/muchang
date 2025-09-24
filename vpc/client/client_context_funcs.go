package client

import (
	"context"

	openapiutil "github.com/rancher/muchang/darabonba-openapi/utils"
	"github.com/rancher/muchang/utils/tea/dara"
)

// DescribeVpcsWithContext queries virtual private clouds (VPCs).
func (client *Client) DescribeVpcsWithContext(ctx context.Context, request *DescribeVpcsRequest, runtime *dara.RuntimeOptions) (result *DescribeVpcsResponse, err error) {
	err = request.Validate()
	if err != nil {
		return result, err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DhcpOptionsSetId) {
		query["DhcpOptionsSetId"] = request.DhcpOptionsSetId
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EnableIpv6) {
		query["EnableIpv6"] = request.EnableIpv6
	}

	if !dara.IsNil(request.IsDefault) {
		query["IsDefault"] = request.IsDefault
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
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

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.VpcName) {
		query["VpcName"] = request.VpcName
	}

	if !dara.IsNil(request.VpcOwnerId) {
		query["VpcOwnerId"] = request.VpcOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVpcs"),
		Version:     dara.String("2016-04-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	result = &DescribeVpcsResponse{}
	body, err := client.CallApiWithCtx(ctx, params, req, runtime)
	if err != nil {
		return result, err
	}
	err = dara.Convert(body, &result)
	return result, err
}

// DescribeVSwitchesWithContext queries the information about available vSwitches that are used for an internal network.
func (client *Client) DescribeVSwitchesWithContext(ctx context.Context, request *DescribeVSwitchesRequest, runtime *dara.RuntimeOptions) (result *DescribeVSwitchesResponse, err error) {
	err = request.Validate()
	if err != nil {
		return result, err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EnableIpv6) {
		query["EnableIpv6"] = request.EnableIpv6
	}

	if !dara.IsNil(request.IsDefault) {
		query["IsDefault"] = request.IsDefault
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
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

	if !dara.IsNil(request.RouteTableId) {
		query["RouteTableId"] = request.RouteTableId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VSwitchName) {
		query["VSwitchName"] = request.VSwitchName
	}

	if !dara.IsNil(request.VSwitchOwnerId) {
		query["VSwitchOwnerId"] = request.VSwitchOwnerId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVSwitches"),
		Version:     dara.String("2016-04-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	result = &DescribeVSwitchesResponse{}
	body, err := client.CallApiWithCtx(ctx, params, req, runtime)
	if err != nil {
		return result, err
	}
	err = dara.Convert(body, &result)
	return result, err
}
