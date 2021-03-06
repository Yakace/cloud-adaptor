package tke

import (
	"context"
	"fmt"
	"time"

	rainbondv1alpha1 "github.com/goodrain/rainbond-operator/api/v1alpha1"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tke "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tke/v20180525"
	"goodrain.com/cloud-adaptor/adaptor"
	"goodrain.com/cloud-adaptor/adaptor/v1alpha1"
)

type tkeAdaptor struct {
	accessKeyID     string
	accessKeySecret string
	tkeclient       *tke.Client
}

//Create create ack adaptor
func Create(accessKeyID, accessKeySecret string) (adaptor.CloudAdaptor, error) {
	credential := common.NewCredential(accessKeyID, accessKeySecret)
	client, err := tke.NewClient(credential, "ap-guangzhou", profile.NewClientProfile())
	if err != nil {
		return nil, err
	}
	return &tkeAdaptor{
		accessKeyID:     accessKeyID,
		accessKeySecret: accessKeySecret,
		tkeclient:       client,
	}, nil
}

func toString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func (t *tkeAdaptor) ClusterList() ([]*v1alpha1.Cluster, error) {
	req := tke.NewDescribeClustersRequest()
	res, err := t.tkeclient.DescribeClusters(req)
	if err != nil {
		return nil, fmt.Errorf("Query cluster list from tencent api failure %s", err.Error())
	}
	var clusters []*v1alpha1.Cluster
	fmt.Printf("%+v", *res.Response.Clusters[0])
	for _, c := range res.Response.Clusters {
		createTime, _ := time.Parse(time.RFC3339, toString(c.CreatedTime))
		clusters = append(clusters, &v1alpha1.Cluster{
			Name:           toString(c.ClusterName),
			ClusterID:      toString(c.ClusterId),
			Created:        v1alpha1.NewTime(createTime),
			State:          toString(c.ClusterStatus),
			ClusterType:    toString(c.ClusterType),
			CurrentVersion: toString(c.ClusterVersion),
			VPCID:          toString(c.ClusterNetworkSettings.VpcId),
			RegionID:       "",
			NetworkMode:    "",
			SubnetCIDR:     "",
			PodCIDR:        toString(c.ClusterNetworkSettings.ClusterCIDR),
		})
	}
	return clusters, nil
}

func (t *tkeAdaptor) DescribeCluster(clusterID string) (*v1alpha1.Cluster, error) {
	return nil, nil
}

func (t *tkeAdaptor) CreateCluster(v1alpha1.CreateClusterConfig) (*v1alpha1.Cluster, error) {
	return nil, nil
}

//DeleteCluster delete cluster
func (t *tkeAdaptor) DeleteCluster(clusterID string) error {
	return nil
}

func (t *tkeAdaptor) GetKubeConfig(clusterID string) (*v1alpha1.KubeConfig, error) {
	return nil, nil
}

func (t *tkeAdaptor) VPCList(regionID string) ([]*v1alpha1.VPC, error) {
	return nil, nil
}

func (t *tkeAdaptor) CreateVPC(v *v1alpha1.VPC) error {
	return nil
}

func (t *tkeAdaptor) DeleteVPC(regionID, vpcID string) error {
	return nil
}

func (t *tkeAdaptor) DescribeVPC(regionID, vpcID string) (*v1alpha1.VPC, error) {
	return nil, nil
}

func (t *tkeAdaptor) CreateVSwitch(v *v1alpha1.VSwitch) error {
	return nil
}

func (t *tkeAdaptor) DescribeVSwitch(regionID, vswitchID string) (*v1alpha1.VSwitch, error) {
	return nil, nil
}

func (t *tkeAdaptor) DeleteVSwitch(regionID, vswitchID string) error {
	return nil
}

func (t *tkeAdaptor) ListZones(regionID string) ([]*v1alpha1.Zone, error) {
	return nil, nil
}

func (t *tkeAdaptor) ListInstanceType(regionID string) ([]*v1alpha1.InstanceType, error) {
	return nil, nil
}

func (t *tkeAdaptor) CreateDB(*v1alpha1.Database) error {
	return nil
}
func (t *tkeAdaptor) CreateNAS(regionID, zoneID string) (string, error) {
	return "", nil
}

func (t *tkeAdaptor) GetNasZone(regionID string) (string, error) {
	return "", nil
}

func (t *tkeAdaptor) GetNASInfo(regionID, fileSystemID string) (*v1alpha1.NasStorageInfo, error) {
	return nil, nil
}

func (t *tkeAdaptor) CreateNASMountTarget(regionID, fileSystemID, VpcID, VSwitchID string) (string, error) {
	return "", nil
}

func (t *tkeAdaptor) CreateLoadBalancer(regionID string) (*v1alpha1.LoadBalancer, error) {
	return nil, nil
}

func (t *tkeAdaptor) BoundLoadBalancerToCluster(regionID, VpcID, loadBalancerID string, endpoints []string) error {
	return nil
}

// init rainbond region
func (t *tkeAdaptor) InitRainbondRegion(initConfig *v1alpha1.RainbondInitConfig) error {
	return nil
}

func (t *tkeAdaptor) SetSecurityGroup(regionID, securityGroupID string) error {
	return nil
}

func (t *tkeAdaptor) DescribeAvailableResourceZones(regionID, InstanceType string) ([]*v1alpha1.AvailableResourceZone, error) {
	return nil, nil
}

func (t *tkeAdaptor) CreateRainbondKubernetes(ctx context.Context, config *v1alpha1.KubernetesClusterConfig, rollback func(step, message, status string)) *v1alpha1.Cluster {
	return nil
}
func (t *tkeAdaptor) GetRainbondInitConfig(cluster *v1alpha1.Cluster, gateway, chaos []*rainbondv1alpha1.K8sNode, rollback func(step, message, status string)) *v1alpha1.RainbondInitConfig {
	return nil
}

func (t *tkeAdaptor) ExpansionNode(ctx context.Context, en *v1alpha1.ExpansionNode, rollback func(step, message, status string)) *v1alpha1.Cluster {
	return nil
}
