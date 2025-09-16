package client

import (
	openapi "github.com/rancher/muchang/darabonba-openapi/client"
	openapiutil "github.com/rancher/muchang/darabonba-openapi/utils"
	"github.com/rancher/muchang/utils/tea/dara"
)

type Client struct {
	openapi.Client
	DisableSDKError *bool
}

func NewClient(config *openapiutil.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapiutil.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = dara.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-hangzhou":                 dara.String("vpc.aliyuncs.com"),
		"cn-shanghai-finance-1":       dara.String("vpc.aliyuncs.com"),
		"cn-shenzhen-finance-1":       dara.String("vpc.aliyuncs.com"),
		"cn-north-2-gov-1":            dara.String("vpc.aliyuncs.com"),
		"ap-northeast-2-pop":          dara.String("vpc.aliyuncs.com"),
		"cn-beijing-finance-pop":      dara.String("vpc.aliyuncs.com"),
		"cn-beijing-gov-1":            dara.String("vpc.aliyuncs.com"),
		"cn-beijing-nu16-b01":         dara.String("vpc.aliyuncs.com"),
		"cn-edge-1":                   dara.String("vpc-nebula.cn-qingdao-nebula.aliyuncs.com"),
		"cn-fujian":                   dara.String("vpc.aliyuncs.com"),
		"cn-haidian-cm12-c01":         dara.String("vpc.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          dara.String("vpc.aliyuncs.com"),
		"cn-hangzhou-finance":         dara.String("vpc.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": dara.String("vpc.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": dara.String("vpc-pre.cn-hangzhou.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": dara.String("vpc-inner-pre.cn-hangzhou.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": dara.String("vpc-pre.cn-hangzhou.aliyuncs.com"),
		"cn-hangzhou-test-306":        dara.String("vpc-pre.cn-hangzhou.aliyuncs.com"),
		"cn-hongkong-finance-pop":     dara.String("vpc.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       dara.String("vpc-nebula.cn-qingdao-nebula.aliyuncs.com"),
		"cn-qingdao-nebula":           dara.String("vpc-nebula.cn-qingdao-nebula.aliyuncs.com"),
		"cn-shanghai-et15-b01":        dara.String("vpc-pre.cn-hangzhou.aliyuncs.com"),
		"cn-shanghai-et2-b01":         dara.String("vpc.aliyuncs.com"),
		"cn-shanghai-inner":           dara.String("vpc.aliyuncs.com"),
		"cn-shanghai-internal-test-1": dara.String("vpc-pre.cn-hangzhou.aliyuncs.com"),
		"cn-shenzhen-inner":           dara.String("vpc.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         dara.String("vpc.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        dara.String("vpc.aliyuncs.com"),
		"cn-wuhan":                    dara.String("vpc.aliyuncs.com"),
		"cn-yushanfang":               dara.String("vpc.aliyuncs.com"),
		"cn-zhangbei":                 dara.String("vpc.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        dara.String("vpc.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     dara.String("vpc.cn-zhangjiakou.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       dara.String("vpc-nebula.cn-qingdao-nebula.aliyuncs.com"),
		"eu-west-1-oxs":               dara.String("vpc-nebula.cn-shenzhen-cloudstone.aliyuncs.com"),
		"rus-west-1-pop":              dara.String("vpc.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("vpc"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !dara.IsNil(endpoint) {
		_result = endpoint
		return _result, _err
	}

	if !dara.IsNil(endpointMap) && !dara.IsNil(endpointMap[dara.StringValue(regionId)]) {
		_result = endpointMap[dara.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := openapiutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
