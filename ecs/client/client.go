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

func (client *Client) Init(config *openapiutil.Config) (err error) {
	err = client.Client.Init(config)
	if err != nil {
		return err
	}
	client.EndpointRule = dara.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-hangzhou":                 dara.String("ecs-cn-hangzhou.aliyuncs.com"),
		"cn-shanghai-finance-1":       dara.String("ecs-cn-hangzhou.aliyuncs.com"),
		"cn-shenzhen-finance-1":       dara.String("ecs-cn-hangzhou.aliyuncs.com"),
		"cn-north-2-gov-1":            dara.String("ecs.aliyuncs.com"),
		"ap-northeast-2-pop":          dara.String("ecs.aliyuncs.com"),
		"cn-beijing-finance-pop":      dara.String("ecs.aliyuncs.com"),
		"cn-beijing-gov-1":            dara.String("ecs.aliyuncs.com"),
		"cn-beijing-nu16-b01":         dara.String("ecs-cn-hangzhou.aliyuncs.com"),
		"cn-edge-1":                   dara.String("ecs.cn-qingdao-nebula.aliyuncs.com"),
		"cn-fujian":                   dara.String("ecs-cn-hangzhou.aliyuncs.com"),
		"cn-haidian-cm12-c01":         dara.String("ecs-cn-hangzhou.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          dara.String("ecs-cn-hangzhou.aliyuncs.com"),
		"cn-hangzhou-finance":         dara.String("ecs.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": dara.String("ecs-cn-hangzhou.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": dara.String("ecs-cn-hangzhou.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": dara.String("ecs-cn-hangzhou.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": dara.String("ecs-cn-hangzhou.aliyuncs.com"),
		"cn-hangzhou-test-306":        dara.String("ecs-cn-hangzhou.aliyuncs.com"),
		"cn-hongkong-finance-pop":     dara.String("ecs.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       dara.String("ecs.cn-qingdao-nebula.aliyuncs.com"),
		"cn-shanghai-et15-b01":        dara.String("ecs-cn-hangzhou.aliyuncs.com"),
		"cn-shanghai-et2-b01":         dara.String("ecs-cn-hangzhou.aliyuncs.com"),
		"cn-shanghai-inner":           dara.String("ecs.aliyuncs.com"),
		"cn-shanghai-internal-test-1": dara.String("ecs-cn-hangzhou.aliyuncs.com"),
		"cn-shenzhen-inner":           dara.String("ecs.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         dara.String("ecs-cn-hangzhou.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        dara.String("ecs-cn-hangzhou.aliyuncs.com"),
		"cn-wuhan":                    dara.String("ecs.aliyuncs.com"),
		"cn-yushanfang":               dara.String("ecs.aliyuncs.com"),
		"cn-zhangbei":                 dara.String("ecs.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        dara.String("ecs-cn-hangzhou.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     dara.String("ecs.cn-zhangjiakou.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       dara.String("ecs.cn-qingdao-nebula.aliyuncs.com"),
		"eu-west-1-oxs":               dara.String("ecs.cn-shenzhen-cloudstone.aliyuncs.com"),
		"rus-west-1-pop":              dara.String("ecs.aliyuncs.com"),
	}
	err = client.CheckConfig(config)
	if err != nil {
		return err
	}
	client.Endpoint, err = client.GetEndpoint(dara.String("ecs"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (result *string, err error) {
	if !dara.IsNil(endpoint) {
		result = endpoint
		return result, err
	}

	if !dara.IsNil(endpointMap) && !dara.IsNil(endpointMap[dara.StringValue(regionId)]) {
		result = endpointMap[dara.StringValue(regionId)]
		return result, err
	}

	body, err := openapiutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if err != nil {
		return result, err
	}
	result = body
	return result, err
}
