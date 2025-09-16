package client

import (
	"context"

	openapiutil "github.com/rancher/muchang/darabonba-openapi/utils"
	"github.com/rancher/muchang/utils/tea/dara"
)

// Summary:
//
// Creates a Container Service for Kubernetes (ACK) cluster. For example, you can create an ACK managed cluster, ACK Serverless cluster, ACK Edge cluster, or registered cluster. When you create an ACK cluster, you need to configure the cluster information, components, and cloud resources used by ACK.
//
// Description:
//
// ### [](#-openapi-)Generate API request parameters through the ACK console
//
// When calling the CreateCluster operation to create a cluster, if the API call fails due to invalid parameter settings, you can generate valid request parameters through the ACK console. Follow these steps:
//
// 1.  Log on to the [ACK console](https://csnew.console.aliyun.com). In the left-side navigation pane, click **Clusters**.
//
// 2.  On the **Clusters*	- page, click **Cluster Templates**.
//
// 3.  In the Select Cluster Template dialog box, select the type of cluster you want to create and click Create. Then, configure the cluster parameters.
//
// 4.  In the **Confirm*	- step, click **Generate API Request Parameters**.
//
//	The API request parameters are displayed in the API Request Parameters dialog box.
//
// @param request - CreateClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClusterResponse
func (client *Client) CreateClusterWithContext(ctx context.Context, request *CreateClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccessControlList) {
		body["access_control_list"] = request.AccessControlList
	}

	if !dara.IsNil(request.Addons) {
		body["addons"] = request.Addons
	}

	if !dara.IsNil(request.ApiAudiences) {
		body["api_audiences"] = request.ApiAudiences
	}

	if !dara.IsNil(request.AuditLogConfig) {
		body["audit_log_config"] = request.AuditLogConfig
	}

	if !dara.IsNil(request.AutoMode) {
		body["auto_mode"] = request.AutoMode
	}

	if !dara.IsNil(request.AutoRenew) {
		body["auto_renew"] = request.AutoRenew
	}

	if !dara.IsNil(request.AutoRenewPeriod) {
		body["auto_renew_period"] = request.AutoRenewPeriod
	}

	if !dara.IsNil(request.ChargeType) {
		body["charge_type"] = request.ChargeType
	}

	if !dara.IsNil(request.CisEnabled) {
		body["cis_enabled"] = request.CisEnabled
	}

	if !dara.IsNil(request.CloudMonitorFlags) {
		body["cloud_monitor_flags"] = request.CloudMonitorFlags
	}

	if !dara.IsNil(request.ClusterDomain) {
		body["cluster_domain"] = request.ClusterDomain
	}

	if !dara.IsNil(request.ClusterSpec) {
		body["cluster_spec"] = request.ClusterSpec
	}

	if !dara.IsNil(request.ClusterType) {
		body["cluster_type"] = request.ClusterType
	}

	if !dara.IsNil(request.ContainerCidr) {
		body["container_cidr"] = request.ContainerCidr
	}

	if !dara.IsNil(request.ControlPlaneConfig) {
		body["control_plane_config"] = request.ControlPlaneConfig
	}

	if !dara.IsNil(request.ControlplaneLogComponents) {
		body["controlplane_log_components"] = request.ControlplaneLogComponents
	}

	if !dara.IsNil(request.ControlplaneLogProject) {
		body["controlplane_log_project"] = request.ControlplaneLogProject
	}

	if !dara.IsNil(request.ControlplaneLogTtl) {
		body["controlplane_log_ttl"] = request.ControlplaneLogTtl
	}

	if !dara.IsNil(request.CpuPolicy) {
		body["cpu_policy"] = request.CpuPolicy
	}

	if !dara.IsNil(request.CustomSan) {
		body["custom_san"] = request.CustomSan
	}

	if !dara.IsNil(request.DeletionProtection) {
		body["deletion_protection"] = request.DeletionProtection
	}

	if !dara.IsNil(request.DisableRollback) {
		body["disable_rollback"] = request.DisableRollback
	}

	if !dara.IsNil(request.EnableRrsa) {
		body["enable_rrsa"] = request.EnableRrsa
	}

	if !dara.IsNil(request.EncryptionProviderKey) {
		body["encryption_provider_key"] = request.EncryptionProviderKey
	}

	if !dara.IsNil(request.EndpointPublicAccess) {
		body["endpoint_public_access"] = request.EndpointPublicAccess
	}

	if !dara.IsNil(request.ExtraSans) {
		body["extra_sans"] = request.ExtraSans
	}

	if !dara.IsNil(request.FormatDisk) {
		body["format_disk"] = request.FormatDisk
	}

	if !dara.IsNil(request.ImageId) {
		body["image_id"] = request.ImageId
	}

	if !dara.IsNil(request.ImageType) {
		body["image_type"] = request.ImageType
	}

	if !dara.IsNil(request.Instances) {
		body["instances"] = request.Instances
	}

	if !dara.IsNil(request.IpStack) {
		body["ip_stack"] = request.IpStack
	}

	if !dara.IsNil(request.IsEnterpriseSecurityGroup) {
		body["is_enterprise_security_group"] = request.IsEnterpriseSecurityGroup
	}

	if !dara.IsNil(request.KeepInstanceName) {
		body["keep_instance_name"] = request.KeepInstanceName
	}

	if !dara.IsNil(request.KeyPair) {
		body["key_pair"] = request.KeyPair
	}

	if !dara.IsNil(request.KubernetesVersion) {
		body["kubernetes_version"] = request.KubernetesVersion
	}

	if !dara.IsNil(request.LoadBalancerId) {
		body["load_balancer_id"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.LoadBalancerSpec) {
		body["load_balancer_spec"] = request.LoadBalancerSpec
	}

	if !dara.IsNil(request.LoggingType) {
		body["logging_type"] = request.LoggingType
	}

	if !dara.IsNil(request.LoginPassword) {
		body["login_password"] = request.LoginPassword
	}

	if !dara.IsNil(request.MaintenanceWindow) {
		body["maintenance_window"] = request.MaintenanceWindow
	}

	if !dara.IsNil(request.MasterAutoRenew) {
		body["master_auto_renew"] = request.MasterAutoRenew
	}

	if !dara.IsNil(request.MasterAutoRenewPeriod) {
		body["master_auto_renew_period"] = request.MasterAutoRenewPeriod
	}

	if !dara.IsNil(request.MasterCount) {
		body["master_count"] = request.MasterCount
	}

	if !dara.IsNil(request.MasterInstanceChargeType) {
		body["master_instance_charge_type"] = request.MasterInstanceChargeType
	}

	if !dara.IsNil(request.MasterInstanceTypes) {
		body["master_instance_types"] = request.MasterInstanceTypes
	}

	if !dara.IsNil(request.MasterPeriod) {
		body["master_period"] = request.MasterPeriod
	}

	if !dara.IsNil(request.MasterPeriodUnit) {
		body["master_period_unit"] = request.MasterPeriodUnit
	}

	if !dara.IsNil(request.MasterSystemDiskCategory) {
		body["master_system_disk_category"] = request.MasterSystemDiskCategory
	}

	if !dara.IsNil(request.MasterSystemDiskPerformanceLevel) {
		body["master_system_disk_performance_level"] = request.MasterSystemDiskPerformanceLevel
	}

	if !dara.IsNil(request.MasterSystemDiskSize) {
		body["master_system_disk_size"] = request.MasterSystemDiskSize
	}

	if !dara.IsNil(request.MasterSystemDiskSnapshotPolicyId) {
		body["master_system_disk_snapshot_policy_id"] = request.MasterSystemDiskSnapshotPolicyId
	}

	if !dara.IsNil(request.MasterVswitchIds) {
		body["master_vswitch_ids"] = request.MasterVswitchIds
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.NatGateway) {
		body["nat_gateway"] = request.NatGateway
	}

	if !dara.IsNil(request.NodeCidrMask) {
		body["node_cidr_mask"] = request.NodeCidrMask
	}

	if !dara.IsNil(request.NodeNameMode) {
		body["node_name_mode"] = request.NodeNameMode
	}

	if !dara.IsNil(request.NodePortRange) {
		body["node_port_range"] = request.NodePortRange
	}

	if !dara.IsNil(request.Nodepools) {
		body["nodepools"] = request.Nodepools
	}

	if !dara.IsNil(request.NumOfNodes) {
		body["num_of_nodes"] = request.NumOfNodes
	}

	if !dara.IsNil(request.OperationPolicy) {
		body["operation_policy"] = request.OperationPolicy
	}

	if !dara.IsNil(request.OsType) {
		body["os_type"] = request.OsType
	}

	if !dara.IsNil(request.Period) {
		body["period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		body["period_unit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.Platform) {
		body["platform"] = request.Platform
	}

	if !dara.IsNil(request.PodVswitchIds) {
		body["pod_vswitch_ids"] = request.PodVswitchIds
	}

	if !dara.IsNil(request.Profile) {
		body["profile"] = request.Profile
	}

	if !dara.IsNil(request.ProxyMode) {
		body["proxy_mode"] = request.ProxyMode
	}

	if !dara.IsNil(request.RdsInstances) {
		body["rds_instances"] = request.RdsInstances
	}

	if !dara.IsNil(request.RegionId) {
		body["region_id"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["resource_group_id"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.RrsaConfig) {
		body["rrsa_config"] = request.RrsaConfig
	}

	if !dara.IsNil(request.Runtime) {
		body["runtime"] = request.Runtime
	}

	if !dara.IsNil(request.SecurityGroupId) {
		body["security_group_id"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.SecurityHardeningOs) {
		body["security_hardening_os"] = request.SecurityHardeningOs
	}

	if !dara.IsNil(request.ServiceAccountIssuer) {
		body["service_account_issuer"] = request.ServiceAccountIssuer
	}

	if !dara.IsNil(request.ServiceCidr) {
		body["service_cidr"] = request.ServiceCidr
	}

	if !dara.IsNil(request.ServiceDiscoveryTypes) {
		body["service_discovery_types"] = request.ServiceDiscoveryTypes
	}

	if !dara.IsNil(request.SnatEntry) {
		body["snat_entry"] = request.SnatEntry
	}

	if !dara.IsNil(request.SocEnabled) {
		body["soc_enabled"] = request.SocEnabled
	}

	if !dara.IsNil(request.SshFlags) {
		body["ssh_flags"] = request.SshFlags
	}

	if !dara.IsNil(request.Tags) {
		body["tags"] = request.Tags
	}

	if !dara.IsNil(request.Taints) {
		body["taints"] = request.Taints
	}

	if !dara.IsNil(request.TimeoutMins) {
		body["timeout_mins"] = request.TimeoutMins
	}

	if !dara.IsNil(request.Timezone) {
		body["timezone"] = request.Timezone
	}

	if !dara.IsNil(request.UserCa) {
		body["user_ca"] = request.UserCa
	}

	if !dara.IsNil(request.UserData) {
		body["user_data"] = request.UserData
	}

	if !dara.IsNil(request.Vpcid) {
		body["vpcid"] = request.Vpcid
	}

	if !dara.IsNil(request.VswitchIds) {
		body["vswitch_ids"] = request.VswitchIds
	}

	if !dara.IsNil(request.WorkerAutoRenew) {
		body["worker_auto_renew"] = request.WorkerAutoRenew
	}

	if !dara.IsNil(request.WorkerAutoRenewPeriod) {
		body["worker_auto_renew_period"] = request.WorkerAutoRenewPeriod
	}

	if !dara.IsNil(request.WorkerDataDisks) {
		body["worker_data_disks"] = request.WorkerDataDisks
	}

	if !dara.IsNil(request.WorkerInstanceChargeType) {
		body["worker_instance_charge_type"] = request.WorkerInstanceChargeType
	}

	if !dara.IsNil(request.WorkerInstanceTypes) {
		body["worker_instance_types"] = request.WorkerInstanceTypes
	}

	if !dara.IsNil(request.WorkerPeriod) {
		body["worker_period"] = request.WorkerPeriod
	}

	if !dara.IsNil(request.WorkerPeriodUnit) {
		body["worker_period_unit"] = request.WorkerPeriodUnit
	}

	if !dara.IsNil(request.WorkerSystemDiskCategory) {
		body["worker_system_disk_category"] = request.WorkerSystemDiskCategory
	}

	if !dara.IsNil(request.WorkerSystemDiskPerformanceLevel) {
		body["worker_system_disk_performance_level"] = request.WorkerSystemDiskPerformanceLevel
	}

	if !dara.IsNil(request.WorkerSystemDiskSize) {
		body["worker_system_disk_size"] = request.WorkerSystemDiskSize
	}

	if !dara.IsNil(request.WorkerSystemDiskSnapshotPolicyId) {
		body["worker_system_disk_snapshot_policy_id"] = request.WorkerSystemDiskSnapshotPolicyId
	}

	if !dara.IsNil(request.WorkerVswitchIds) {
		body["worker_vswitch_ids"] = request.WorkerVswitchIds
	}

	if !dara.IsNil(request.ZoneId) {
		body["zone_id"] = request.ZoneId
	}

	if !dara.IsNil(request.ZoneIds) {
		body["zone_ids"] = request.ZoneIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCluster"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateClusterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a node pool for a Container Service for Kubernetes (ACK) cluster. You can use node pools to facilitate node management. For example, you can schedule, configure, or maintain nodes by node pool, and enable auto scaling for a node pool. We recommend that you use a managed node pool, which can help automate specific O\\\\\\&M tasks for nodes, such as Common Vulnerabilities and Exposures (CVE) patching and node repair. This reduces your O\\\\\\&M workload.
//
// @param request - CreateClusterNodePoolRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClusterNodePoolResponse
func (client *Client) CreateClusterNodePoolWithContext(ctx context.Context, ClusterId *string, request *CreateClusterNodePoolRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateClusterNodePoolResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoMode) {
		body["auto_mode"] = request.AutoMode
	}

	if !dara.IsNil(request.AutoScaling) {
		body["auto_scaling"] = request.AutoScaling
	}

	if !dara.IsNil(request.Count) {
		body["count"] = request.Count
	}

	if !dara.IsNil(request.EfloNodeGroup) {
		body["eflo_node_group"] = request.EfloNodeGroup
	}

	if !dara.IsNil(request.HostNetwork) {
		body["host_network"] = request.HostNetwork
	}

	if !dara.IsNil(request.InterconnectConfig) {
		body["interconnect_config"] = request.InterconnectConfig
	}

	if !dara.IsNil(request.InterconnectMode) {
		body["interconnect_mode"] = request.InterconnectMode
	}

	if !dara.IsNil(request.Intranet) {
		body["intranet"] = request.Intranet
	}

	if !dara.IsNil(request.KubernetesConfig) {
		body["kubernetes_config"] = request.KubernetesConfig
	}

	if !dara.IsNil(request.Management) {
		body["management"] = request.Management
	}

	if !dara.IsNil(request.MaxNodes) {
		body["max_nodes"] = request.MaxNodes
	}

	if !dara.IsNil(request.NodeConfig) {
		body["node_config"] = request.NodeConfig
	}

	if !dara.IsNil(request.NodepoolInfo) {
		body["nodepool_info"] = request.NodepoolInfo
	}

	if !dara.IsNil(request.ScalingGroup) {
		body["scaling_group"] = request.ScalingGroup
	}

	if !dara.IsNil(request.TeeConfig) {
		body["tee_config"] = request.TeeConfig
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateClusterNodePool"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/nodepools"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateClusterNodePoolResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the DeleteCluster operation to delete a cluster and specify whether to delete or retain the relevant cluster resources. Before you delete a cluster, you must manually delete workloads in the cluster, such as Deployments, StatefulSets, Jobs, and CronJobs. Otherwise, you may fail to delete the cluster.
//
// Description:
//
// Warning:
//
//   - Subscription ECS instances and Lingjun nodes in a cluster cannot be automatically released. To avoid unnecessary costs, we recommend that you manually release the resources. For more information, see \\<a href="{0}" target="_blank">Rules for deleting clusters and releasing nodes\\</a>.
//
//   - If the SLB instance of the API server uses the subscription billing method, it cannot be automatically released. To avoid unnecessary costs, we recommend that you manually release it.
//
//   - By default, virtual private clouds (VPCs), vSwitches, security groups, and RAM roles are retained if they are used by other resources. To avoid unnecessary costs, we recommend that you manually release the resources.
//
//   - Elastic container instances created on virtual nodes are automatically released.
//
//   - Some resources created together with a cluster are not automatically released when the cluster is deleted. After the cluster is deleted, you are still charged for the resources. Release or retain the resources based on your actual needs. The resources include Simple Log Service projects automatically created by the cluster and dynamically provisioned disks.
//
// @param tmpReq - DeleteClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteClusterResponse
func (client *Client) DeleteClusterWithContext(ctx context.Context, ClusterId *string, tmpReq *DeleteClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteClusterResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DeleteClusterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeleteOptions) {
		request.DeleteOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeleteOptions, dara.String("delete_options"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RetainResources) {
		request.RetainResourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RetainResources, dara.String("retain_resources"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeleteOptionsShrink) {
		query["delete_options"] = request.DeleteOptionsShrink
	}

	if !dara.IsNil(request.KeepSlb) {
		query["keep_slb"] = request.KeepSlb
	}

	if !dara.IsNil(request.RetainAllResources) {
		query["retain_all_resources"] = request.RetainAllResources
	}

	if !dara.IsNil(request.RetainResourcesShrink) {
		query["retain_resources"] = request.RetainResourcesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCluster"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteClusterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// null
//
// @param request - DeleteClusterNodepoolRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteClusterNodepoolResponse
func (client *Client) DeleteClusterNodepoolWithContext(ctx context.Context, ClusterId *string, NodepoolId *string, request *DeleteClusterNodepoolRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteClusterNodepoolResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Force) {
		query["force"] = request.Force
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteClusterNodepool"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/nodepools/" + dara.PercentEncode(dara.StringValue(NodepoolId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteClusterNodepoolResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the DescribeClusterDetail operation to query the details of a Container Service for Kubernetes (ACK) cluster by cluster ID.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterDetailResponse
func (client *Client) DescribeClusterDetailWithContext(ctx context.Context, ClusterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClusterDetailResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClusterDetail"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClusterDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about all node pools in a cluster.
//
// @param request - DescribeClusterNodePoolsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterNodePoolsResponse
func (client *Client) DescribeClusterNodePoolsWithContext(ctx context.Context, ClusterId *string, request *DescribeClusterNodePoolsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClusterNodePoolsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NodepoolName) {
		query["NodepoolName"] = request.NodepoolName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClusterNodePools"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/nodepools"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClusterNodePoolsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// null
//
// @param request - DescribeClusterNodesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterNodesResponse
func (client *Client) DescribeClusterNodesWithContext(ctx context.Context, ClusterId *string, request *DescribeClusterNodesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClusterNodesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIds) {
		query["instanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.NodepoolId) {
		query["nodepool_id"] = request.NodepoolId
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.State) {
		query["state"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClusterNodes"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/nodes"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClusterNodesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Kubeconfig files store identity and authentication information that is used by clients to access Container Service for Kubernetes (ACK) clusters. To use a kubectl client to manage an ACK cluster, you need to use the corresponding kubeconfig file to connect to the ACK cluster. We recommend that you keep kubeconfig files confidential and revoke kubeconfig files that are not in use. This helps prevent data leaks caused by the disclosure of kubeconfig files.
//
// Description:
//
//	  The default validity period of a kubeconfig file is 3 years. 180 days before a kubeconfig file expires, you can renew it in the Container Service for Kubernetes (ACK) console or by calling API operations. After a kubeconfig file is renewed, the kubeconfig file is valid for 3 years. The previous kubeconfig file still remains valid until expiration. We recommend that you renew your kubeconfig file at the earliest opportunity.
//
//		- We recommend that you keep kubeconfig files confidential and revoke kubeconfig files that are not in use. This helps prevent data leaks caused by the disclosure of kubeconfig files.
//
// @param request - DescribeClusterUserKubeconfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterUserKubeconfigResponse
func (client *Client) DescribeClusterUserKubeconfigWithContext(ctx context.Context, ClusterId *string, request *DescribeClusterUserKubeconfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClusterUserKubeconfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PrivateIpAddress) {
		query["PrivateIpAddress"] = request.PrivateIpAddress
	}

	if !dara.IsNil(request.TemporaryDurationMinutes) {
		query["TemporaryDurationMinutes"] = request.TemporaryDurationMinutes
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClusterUserKubeconfig"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/k8s/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/user_config"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClusterUserKubeconfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all clusters in a specified region.
//
// @param request - DescribeClustersForRegionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClustersForRegionResponse
func (client *Client) DescribeClustersForRegionWithContext(ctx context.Context, regionId *string, request *DescribeClustersForRegionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeClustersForRegionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["cluster_id"] = request.ClusterId
	}

	if !dara.IsNil(request.ClusterSpec) {
		query["cluster_spec"] = request.ClusterSpec
	}

	if !dara.IsNil(request.ClusterType) {
		query["cluster_type"] = request.ClusterType
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["page_number"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	if !dara.IsNil(request.Profile) {
		query["profile"] = request.Profile
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClustersForRegion"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/regions/" + dara.PercentEncode(dara.StringValue(regionId)) + "/clusters"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClustersForRegionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the detailed information about Kubernetes versions, including the version number, release date, expiration date, compatible OSs, and runtime.
//
// @param request - DescribeKubernetesVersionMetadataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeKubernetesVersionMetadataResponse
func (client *Client) DescribeKubernetesVersionMetadataWithContext(ctx context.Context, request *DescribeKubernetesVersionMetadataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeKubernetesVersionMetadataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterType) {
		query["ClusterType"] = request.ClusterType
	}

	if !dara.IsNil(request.KubernetesVersion) {
		query["KubernetesVersion"] = request.KubernetesVersion
	}

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	if !dara.IsNil(request.Profile) {
		query["Profile"] = request.Profile
	}

	if !dara.IsNil(request.QueryUpgradableVersion) {
		query["QueryUpgradableVersion"] = request.QueryUpgradableVersion
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.Runtime) {
		query["runtime"] = request.Runtime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeKubernetesVersionMetadata"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/metadata/versions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("array"),
	}
	_result = &DescribeKubernetesVersionMetadataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries detailed information about a task, such as the task type, status, and progress.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTaskInfoResponse
func (client *Client) DescribeTaskInfoWithContext(ctx context.Context, taskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeTaskInfoResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTaskInfo"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tasks/" + dara.PercentEncode(dara.StringValue(taskId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTaskInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the ModifyClusterNodePool operation to modify the configuration of a node pool with the specified node pool ID.
//
// @param request - ModifyClusterNodePoolRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyClusterNodePoolResponse
func (client *Client) ModifyClusterNodePoolWithContext(ctx context.Context, ClusterId *string, NodepoolId *string, request *ModifyClusterNodePoolRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyClusterNodePoolResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoScaling) {
		body["auto_scaling"] = request.AutoScaling
	}

	if !dara.IsNil(request.Concurrency) {
		body["concurrency"] = request.Concurrency
	}

	if !dara.IsNil(request.KubernetesConfig) {
		body["kubernetes_config"] = request.KubernetesConfig
	}

	if !dara.IsNil(request.Management) {
		body["management"] = request.Management
	}

	if !dara.IsNil(request.NodepoolInfo) {
		body["nodepool_info"] = request.NodepoolInfo
	}

	if !dara.IsNil(request.ScalingGroup) {
		body["scaling_group"] = request.ScalingGroup
	}

	if !dara.IsNil(request.TeeConfig) {
		body["tee_config"] = request.TeeConfig
	}

	if !dara.IsNil(request.UpdateNodes) {
		body["update_nodes"] = request.UpdateNodes
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyClusterNodePool"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/nodepools/" + dara.PercentEncode(dara.StringValue(NodepoolId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyClusterNodePoolResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes nodes from a node pool.
//
// Description:
//
//	  When you remove a node, the pods on the node are migrated to other nodes. This may cause service interruptions. We recommend that you remove nodes during off-peak hours.
//
//		- The operation may have unexpected risks. Back up the data before you perform this operation.
//
//		- Nodes remain in the Unschedulable state when they are being removed.
//
//		- The system removes only worker nodes. It does not remove master nodes.
//
//		- Even if you set the `release_node` parameter to `true`, subscription nodes are not released. You must release the subscription nodes in the [ECS console](https://ecs.console.aliyun.com/) after you remove the nodes.
//
// @param tmpReq - RemoveNodePoolNodesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveNodePoolNodesResponse
func (client *Client) RemoveNodePoolNodesWithContext(ctx context.Context, ClusterId *string, NodepoolId *string, tmpReq *RemoveNodePoolNodesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RemoveNodePoolNodesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &RemoveNodePoolNodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("instance_ids"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Nodes) {
		request.NodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Nodes, dara.String("nodes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Concurrency) {
		query["concurrency"] = request.Concurrency
	}

	if !dara.IsNil(request.DrainNode) {
		query["drain_node"] = request.DrainNode
	}

	if !dara.IsNil(request.InstanceIdsShrink) {
		query["instance_ids"] = request.InstanceIdsShrink
	}

	if !dara.IsNil(request.NodesShrink) {
		query["nodes"] = request.NodesShrink
	}

	if !dara.IsNil(request.ReleaseNode) {
		query["release_node"] = request.ReleaseNode
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveNodePoolNodes"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/nodepools/" + dara.PercentEncode(dara.StringValue(NodepoolId)) + "/nodes"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveNodePoolNodesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Outdated Kubernetes versions may have security and stability issues. We recommend that you update the Kubernetes version of your cluster at the earliest opportunity to enjoy the new features of the new Kubernetes version. You can call the UpgradeCluster operation to manually upgrade a cluster.
//
// Description:
//
// After successfully calling the UpgradeCluster interface, this API returns the `task_id` of the upgrade task. You can manage this operation task by calling the following task APIs:
//
// - [Call DescribeTaskInfo to query task details](https://help.aliyun.com/document_detail/2667985.html)
//
// - [Call PauseTask to pause a running task](https://help.aliyun.com/document_detail/2667986.html)
//
// - [Call ResumeTask to resume a task that has been paused](https://help.aliyun.com/document_detail/2667987.html)
//
// - [Call CancelTask to cancel a running task](https://help.aliyun.com/document_detail/2667988.html)
//
// @param request - UpgradeClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeClusterResponse
func (client *Client) UpgradeClusterWithContext(ctx context.Context, ClusterId *string, request *UpgradeClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpgradeClusterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ComponentName) {
		body["component_name"] = request.ComponentName
	}

	if !dara.IsNil(request.MasterOnly) {
		body["master_only"] = request.MasterOnly
	}

	if !dara.IsNil(request.NextVersion) {
		body["next_version"] = request.NextVersion
	}

	if !dara.IsNil(request.RollingPolicy) {
		body["rolling_policy"] = request.RollingPolicy
	}

	if !dara.IsNil(request.Version) {
		body["version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeCluster"),
		Version:     dara.String("2015-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/clusters/" + dara.PercentEncode(dara.StringValue(ClusterId)) + "/upgrade"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeClusterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
