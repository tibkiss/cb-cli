package cli

import (
	"sort"
	"strings"
	"testing"

	"github.com/hortonworks/hdc-cli/models"
)

func clusterSkeleton(stackParams map[string]string, credParams map[string]interface{}, netParams map[string]interface{}) (ClusterSkeletonResult, *models.StackResponse, *models.CredentialResponse, *models.BlueprintResponse, *models.NetworkResponse) {
	skeleton := ClusterSkeletonResult{}
	sr := models.StackResponse{
		Name:         "stack-name",
		Status:       &(&stringWrapper{"status"}).s,
		StatusReason: &(&stringWrapper{"status-reason"}).s,
		HdpVersion:   &(&stringWrapper{"hdp-version"}).s,
		Parameters:   stackParams,
	}
	cr := models.CredentialResponse{
		Parameters: credParams,
	}
	br := models.BlueprintResponse{
		Name: "blueprint-name",
		AmbariBlueprint: models.AmbariBlueprint{
			Blueprint: models.Blueprint{
				Name: &(&stringWrapper{"ambari-blueprint-name"}).s,
			},
		},
	}
	nj := models.NetworkResponse{
		Parameters: netParams,
	}

	return skeleton, &sr, &cr, &br, &nj
}

func defaultNetworkParams() map[string]interface{} {
	np := make(map[string]interface{})
	np["internetGatewayId"] = ""
	return np
}

func TestFillMinimumSet(t *testing.T) {
	skeleton, sr, cr, br, nj := clusterSkeleton(nil, nil, defaultNetworkParams())

	skeleton.fill(sr, cr, br, nil, nil, nj, nil, nil, nil)

	if skeleton.ClusterName != sr.Name {
		t.Errorf("name not match %s == %s", sr.Name, skeleton.ClusterName)
	}
	if skeleton.InstanceRole != "" {
		t.Errorf("instance role not empty %s", skeleton.RemoteAccess)
	}
	if skeleton.Status != *sr.Status {
		t.Errorf("status not match %s == %s", *sr.Status, skeleton.Status)
	}
	if skeleton.StatusReason != *sr.StatusReason {
		t.Errorf("status reason not match %s == %s", *sr.StatusReason, skeleton.StatusReason)
	}
	if skeleton.ClusterInputs != nil {
		t.Errorf("cluster inputs not empty %s", skeleton.ClusterInputs)
	}
	if skeleton.HDPVersion != "hdp" {
		t.Errorf("HDP version not match hdp == %s", skeleton.HDPVersion)
	}
	if skeleton.ClusterType != br.Name {
		t.Errorf("cluster type not match %s == %s", br.Name, skeleton.ClusterType)
	}
	if skeleton.Network != nil {
		t.Errorf("network not empty %s", *skeleton.Network)
	}
	if skeleton.HiveMetastore != nil {
		t.Errorf("hive meta store not empty %s", *skeleton.HiveMetastore)
	}
	if skeleton.SSHKeyName != "" {
		t.Errorf("ssh key name not empty %s", skeleton.SSHKeyName)
	}
	if skeleton.RemoteAccess != "" {
		t.Errorf("remote access not empty %s", skeleton.RemoteAccess)
	}
	if skeleton.WebAccess != false {
		t.Error("web access must be false")
	}
}

func TestFillWithInstanceProfileStrategy(t *testing.T) {
	sp := make(map[string]string)
	sp["instanceProfileStrategy"] = "strategy"
	skeleton, sr, cr, br, nj := clusterSkeleton(sp, nil, defaultNetworkParams())

	skeleton.fill(sr, cr, br, nil, nil, nj, nil, nil, nil)

	if skeleton.InstanceRole != sp["instanceProfileStrategy"] {
		t.Errorf("instance role not match %s == %s", sp["instanceProfileStrategy"], skeleton.InstanceRole)
	}
}

func TestFillWithUseExistingInstanceProfileStrategy(t *testing.T) {
	sp := make(map[string]string)
	sp["instanceProfileStrategy"] = "USE_EXISTING"
	sp["instanceProfile"] = "s3-role"
	skeleton, sr, cr, br, nj := clusterSkeleton(sp, nil, defaultNetworkParams())

	skeleton.fill(sr, cr, br, nil, nil, nj, nil, nil, nil)

	if skeleton.InstanceRole != sp["instanceProfile"] {
		t.Errorf("instance role not match %s == %s", sp["instanceProfile"], skeleton.InstanceRole)
	}
}

func TestFillWithExistingNetwork(t *testing.T) {
	np := make(map[string]interface{})
	np["vpcId"] = "vpcId"
	np["subnetId"] = "subnetId"
	skeleton, sr, cr, br, nj := clusterSkeleton(nil, nil, np)

	skeleton.fill(sr, cr, br, nil, nil, nj, nil, nil, nil)

	expected := Network{VpcId: "vpcId", SubnetId: "subnetId"}
	if *skeleton.Network != expected {
		t.Errorf("network not match %s == %s", expected, *skeleton.Network)
	}
}

func TestFillWithRDSConfig(t *testing.T) {
	skeleton, sr, cr, br, nj := clusterSkeleton(nil, nil, defaultNetworkParams())

	rdsName := "rds-name"
	rcr := []*models.RDSConfigResponse{{
		Name: rdsName,
		Type: &(&stringWrapper{HIVE_RDS}).s,
	},
	}

	skeleton.fill(sr, cr, br, nil, nil, nj, rcr, nil, nil)

	if skeleton.HiveMetastore == nil {
		t.Error("meta store is empty")
	} else if skeleton.HiveMetastore.Name != rdsName {
		t.Errorf("meta store name not match %s == %s", rdsName, skeleton.HiveMetastore.Name)
	}
}

func TestFillWithRDSConfigsdrgwsr(t *testing.T) {
	skeleton, sr, cr, br, nj := clusterSkeleton(nil, nil, defaultNetworkParams())

	groups := make([]*models.InstanceGroupResponse, 0)
	groups = append(groups, &models.InstanceGroupResponse{Group: MASTER, NodeCount: 1})
	groups = append(groups, &models.InstanceGroupResponse{Group: WORKER, NodeCount: 2})
	sr.InstanceGroups = groups

	n := int32(10)

	tm := make(map[string]*models.TemplateResponse)
	tm[MASTER] = &models.TemplateResponse{
		InstanceType: "master",
		VolumeType:   &(&stringWrapper{"type"}).s,
		VolumeSize:   n,
		VolumeCount:  n,
	}
	tm[WORKER] = &models.TemplateResponse{
		InstanceType: "worker",
		VolumeType:   &(&stringWrapper{"type"}).s,
		VolumeSize:   n,
		VolumeCount:  n,
	}

	sm := make(map[string][]*models.SecurityRuleResponse)
	rules := make([]*models.SecurityRuleResponse, 0)
	rules = append(rules, &models.SecurityRuleResponse{Ports: "ports"})
	sm["master"] = rules

	skeleton.fill(sr, cr, br, tm, sm, nj, nil, nil, nil)

	if skeleton.Master.InstanceCount != 1 {
		t.Errorf("master instance count not match 1 == %d", skeleton.Master.InstanceCount)
	}
	if skeleton.Master.InstanceType != tm[MASTER].InstanceType {
		t.Errorf("master instance type not match %s == %s", tm[MASTER].InstanceType, skeleton.Master.InstanceType)
	}
	if skeleton.Master.VolumeType != *tm[MASTER].VolumeType {
		t.Errorf("master volume type not match %s == %s", *tm[MASTER].VolumeType, skeleton.Master.VolumeType)
	}
	if *skeleton.Master.VolumeSize != tm[MASTER].VolumeSize {
		t.Errorf("master volume size not match %d == %d", tm[MASTER].VolumeSize, skeleton.Master.VolumeSize)
	}
	if *skeleton.Master.VolumeCount != tm[MASTER].VolumeCount {
		t.Errorf("master volume count not match %d == %d", tm[MASTER].VolumeCount, skeleton.Master.VolumeCount)
	}
	if skeleton.Worker.InstanceCount != 2 {
		t.Errorf("worker instance count not match 2 == %d", skeleton.Worker.InstanceCount)
	}
	if skeleton.Worker.InstanceType != tm[WORKER].InstanceType {
		t.Errorf("worker instance type not match %s == %s", tm[WORKER].InstanceType, skeleton.Worker.InstanceType)
	}
	if skeleton.Worker.VolumeType != *tm[WORKER].VolumeType {
		t.Errorf("worker volume type not match %s == %s", *tm[WORKER].VolumeType, skeleton.Worker.VolumeType)
	}
	if *skeleton.Worker.VolumeSize != tm[WORKER].VolumeSize {
		t.Errorf("worker volume size not match %d == %d", tm[WORKER].VolumeSize, skeleton.Worker.VolumeSize)
	}
	if *skeleton.Worker.VolumeCount != tm[WORKER].VolumeCount {
		t.Errorf("worker volume count not match %d == %d", tm[WORKER].VolumeCount, skeleton.Worker.VolumeCount)
	}
}

func TestFillWithSshKey(t *testing.T) {
	cp := make(map[string]interface{})
	cp["existingKeyPairName"] = "ssh-key-name"
	skeleton, sr, cr, br, nj := clusterSkeleton(nil, cp, defaultNetworkParams())

	skeleton.fill(sr, cr, br, nil, nil, nj, nil, nil, nil)

	if skeleton.SSHKeyName != cp["existingKeyPairName"] {
		t.Errorf("ssh key name not match %s == %s", cp["existingKeyPairName"], skeleton.SSHKeyName)
	}
}

func TestFillWithSecurityMap(t *testing.T) {
	skeleton, sr, cr, br, nj := clusterSkeleton(nil, nil, defaultNetworkParams())

	cl := models.ClusterResponse{
		Gateway: &models.GatewayJSON{
			ExposedServices: []string{"AMBARI"},
		},
	}
	sr.Cluster = &cl

	sm := make(map[string][]*models.SecurityRuleResponse)
	rules := make([]*models.SecurityRuleResponse, 0)
	rules = append(rules, &models.SecurityRuleResponse{Ports: "22,9443,8443"})
	sm["master"] = rules
	sm["worker"] = rules

	skeleton.fill(sr, cr, br, nil, sm, nj, nil, nil, nil)

	expected := []string{"master", "worker"}
	sort.Strings(expected)
	actual := strings.Split(skeleton.RemoteAccess, ",")
	sort.Strings(actual)

	if strings.Join(actual, "") != strings.Join(expected, "") {
		t.Errorf("remote access not match %s == %s", expected, actual)
	}
	if skeleton.WebAccess != true {
		t.Error("web access must be true")
	}
}

func TestFillWithSecurityMapDefaultPorts(t *testing.T) {
	skeleton, sr, cr, br, nj := clusterSkeleton(nil, nil, defaultNetworkParams())

	sm := make(map[string][]*models.SecurityRuleResponse)
	rules := make([]*models.SecurityRuleResponse, 0)
	rules = append(rules, &models.SecurityRuleResponse{Ports: "22,9443"})
	sm["master"] = rules

	skeleton.fill(sr, cr, br, nil, sm, nj, nil, nil, nil)

	if skeleton.WebAccess != false {
		t.Error("web access must be false")
	}
}

func TestFillWithCluster(t *testing.T) {
	skeleton, sr, cr, br, nj := clusterSkeleton(nil, nil, defaultNetworkParams())

	inputs := make([]*models.BlueprintInput, 0)
	inputs = append(inputs, &models.BlueprintInput{Name: &(&stringWrapper{"property"}).s, PropertyValue: &(&stringWrapper{"value"}).s})
	sr.Cluster = &models.ClusterResponse{
		Status:          &(&stringWrapper{"cluster-status"}).s,
		StatusReason:    &(&stringWrapper{"cluster-reason"}).s,
		BlueprintInputs: inputs,
	}

	skeleton.fill(sr, cr, br, nil, nil, nj, nil, nil, nil)

	if skeleton.Status != *sr.Cluster.Status {
		t.Errorf("status not match %s == %s", *sr.Cluster.Status, skeleton.Status)
	}
	if skeleton.StatusReason != *sr.Cluster.StatusReason {
		t.Errorf("status reason not match %s == %s", *sr.Cluster.StatusReason, skeleton.StatusReason)
	}

	if len(skeleton.ClusterInputs) != 1 || skeleton.ClusterInputs["property"] != "value" {
		t.Errorf("cluster inputs not match map[property:value] == %s", skeleton.ClusterInputs)
	}
}

func TestFillWithNoInstances(t *testing.T) {
	skeleton, sr, _, _, _ := clusterSkeleton(nil, nil, nil)

	skeleton.fill(sr, nil, nil, nil, nil, nil, nil, nil, nil)

	if skeleton.Nodes != UNKNOWN {
		t.Errorf("nodes status not match %s == %s", skeleton.Nodes, UNKNOWN)
	}
}

func TestFillWitMixedHostStatuses(t *testing.T) {
	skeleton, _, _, _, _ := clusterSkeleton(nil, nil, nil)
	host1 := "host1.example.com"
	host2 := "host2.example.com"
	metadata := []*models.InstanceMetaData{
		{DiscoveryFQDN: &host1,
			InstanceGroup: &(&stringWrapper{MASTER}).s},
		{DiscoveryFQDN: &host2,
			InstanceGroup: &(&stringWrapper{WORKER}).s},
	}
	ig := []*models.InstanceGroupResponse{{Metadata: metadata}}
	hm := []*models.HostMetadata{
		{Name: host1,
			State: &(&stringWrapper{UNHEALTHY}).s},
		{Name: host2,
			State: &(&stringWrapper{HEALTHY}).s},
	}
	hg := []*models.HostGroupResponse{
		{Name: MASTER,
			Metadata: hm},
	}
	cr := models.ClusterResponse{
		HostGroups: hg,
	}
	sr := models.StackResponse{
		Name:           "stack-name",
		InstanceGroups: ig,
		Cluster:        &cr,
	}

	skeleton.fill(&sr, nil, nil, nil, nil, nil, nil, nil, nil)

	if skeleton.Nodes != UNHEALTHY {
		t.Errorf("nodes status not match %s == %s", skeleton.Nodes, UNHEALTHY)
	}
}

func TestFillWithHealthyHostStatuses(t *testing.T) {
	skeleton, _, _, _, _ := clusterSkeleton(nil, nil, nil)
	host1 := "host1.example.com"
	host2 := "host2.example.com"
	metadata := []*models.InstanceMetaData{
		{DiscoveryFQDN: &host1,
			InstanceGroup: &(&stringWrapper{MASTER}).s},
		{DiscoveryFQDN: &host2,
			InstanceGroup: &(&stringWrapper{WORKER}).s},
	}
	ig := []*models.InstanceGroupResponse{{Metadata: metadata}}
	hm := []*models.HostMetadata{
		{Name: host1,
			State: &(&stringWrapper{HEALTHY}).s},
		{Name: host2,
			State: &(&stringWrapper{HEALTHY}).s},
	}
	hg := []*models.HostGroupResponse{
		{Name: MASTER,
			Metadata: hm},
	}
	cr := models.ClusterResponse{
		HostGroups: hg,
	}
	sr := models.StackResponse{
		Name:           "stack-name",
		InstanceGroups: ig,
		Cluster:        &cr,
	}

	skeleton.fill(&sr, nil, nil, nil, nil, nil, nil, nil, nil)

	if skeleton.Nodes != HEALTHY {
		t.Errorf("nodes status not match %s == %s", skeleton.Nodes, HEALTHY)
	}
}

func TestFillWithUnknownHostStatuses(t *testing.T) {
	skeleton, _, _, _, _ := clusterSkeleton(nil, nil, nil)
	host1 := "host1.example.com"
	host2 := "host2.example.com"
	metadata := []*models.InstanceMetaData{
		{DiscoveryFQDN: &host1,
			InstanceGroup: &(&stringWrapper{MASTER}).s},
		{DiscoveryFQDN: &host2,
			InstanceGroup: &(&stringWrapper{WORKER}).s},
	}
	ig := []*models.InstanceGroupResponse{{Metadata: metadata}}
	cr := models.ClusterResponse{}
	sr := models.StackResponse{
		Name:           "stack-name",
		InstanceGroups: ig,
		Cluster:        &cr,
	}

	skeleton.fill(&sr, nil, nil, nil, nil, nil, nil, nil, nil)

	if skeleton.Nodes != UNKNOWN {
		t.Errorf("nodes status not match %s == %s", skeleton.Nodes, UNKNOWN)
	}
}

func TestFillWithClusterAvailable(t *testing.T) {
	skeleton, sr, cr, br, nj := clusterSkeleton(nil, nil, defaultNetworkParams())

	sr.Cluster = &models.ClusterResponse{
		Status: &(&stringWrapper{"AVAILABLE"}).s,
	}

	skeleton.fill(sr, cr, br, nil, nil, nj, nil, nil, nil)

	if skeleton.Status != *sr.Status {
		t.Errorf("status not match %s == %s", *sr.Status, skeleton.Status)
	}
}
