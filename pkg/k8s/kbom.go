package k8s

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-containerregistry/pkg/name"
	containerimage "github.com/google/go-containerregistry/pkg/name"
	corev1 "k8s.io/api/core/v1"
	k8sapierror "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	k8sComponentNamespace = "kube-system"
)

type Cluster struct {
	clientSet *kubernetes.Clientset
	cConfig   clientcmd.ClientConfig
}

func NewCluster(clientSet *kubernetes.Clientset, clientConfig clientcmd.ClientConfig) *Cluster {
	return &Cluster{clientSet: clientSet, cConfig: clientConfig}
}

func (cluster *Cluster) CreateClusterSbom() (*ClusterBom, error) {
	nodesInfo := cluster.CollectNodes()
	// collect addons info
	var components []Component
	var err error
	labels := map[string]string{
		k8sComponentNamespace: "component",
	}
	if cluster.isOpenShift() {
		labels = map[string]string{
			"openshift-kube-apiserver":          "apiserver",
			"openshift-kube-controller-manager": "kube-controller-manager",
			"openshift-kube-scheduler":          "scheduler",
			"openshift-etcd":                    "etcd",
		}
	}
	components, err = cluster.collectComponents(labels)
	if err != nil {
		return nil, err
	}
	metadata, err := cluster.getBasicMetadata()
	if err != nil {
		return nil, err
	}
	addonLabels := map[string]string{
		k8sComponentNamespace: "k8s-app",
	}
	addons, err := cluster.collectComponents(addonLabels)
	if err != nil {
		return nil, err
	}
	return cluster.CreateClusterBom(metadata, components, addons, nodesInfo), nil
}

func (cluster *Cluster) CreateComResult() (*ClusterBom, error) {
	nodesInfo := cluster.CollectNodes()
	// collect addons info
	var components []Component
	var err error
	labels := map[string]string{
		k8sComponentNamespace: "component",
	}
	if cluster.isOpenShift() {
		labels = map[string]string{
			"openshift-kube-apiserver":          "apiserver",
			"openshift-kube-controller-manager": "kube-controller-manager",
			"openshift-kube-scheduler":          "scheduler",
			"openshift-etcd":                    "etcd",
		}
	}
	components, err = cluster.collectComponents(labels)
	if err != nil {
		return nil, err
	}
	metadata, err := cluster.getBasicMetadata()
	if err != nil {
		return nil, err
	}
	addonLabels := map[string]string{
		k8sComponentNamespace: "k8s-app",
	}
	addons, err := cluster.collectComponents(addonLabels)
	if err != nil {
		return nil, err
	}
	return cluster.CreateClusterBom(metadata, components, addons, nodesInfo), nil
}

func (cluster *Cluster) CreatePkgBom() (*BomResult, error) {
	nodesInfo := cluster.CollectNodes()
	// collect addons info
	var components []Component
	var err error
	labels := map[string]string{
		k8sComponentNamespace: "component",
	}
	if cluster.isOpenShift() {
		labels = map[string]string{
			"openshift-kube-apiserver":          "apiserver",
			"openshift-kube-controller-manager": "kube-controller-manager",
			"openshift-kube-scheduler":          "scheduler",
			"openshift-etcd":                    "etcd",
		}
	}
	components, err = cluster.collectComponents(labels)
	if err != nil {
		return nil, err
	}
	metadata, err := cluster.targetMetadata()
	if err != nil {
		return nil, err
	}
	addonLabels := map[string]string{
		k8sComponentNamespace: "k8s-app",
	}
	addons, err := cluster.collectComponents(addonLabels)
	if err != nil {
		return nil, err
	}
	return cluster.GetPackages(metadata, components, addons, nodesInfo), nil
}

func (bom *Cluster) GetBaseComponent(imageRef name.Reference, imageName name.Reference) (Component, error) {
	repoName := imageRef.Context().RepositoryStr()
	registryName := imageRef.Context().RegistryStr()
	if strings.HasPrefix(repoName, "library/sha256") {
		repoName = imageName.Context().RepositoryStr()
		registryName = imageName.Context().RegistryStr()
	}

	return Component{
		Repository: repoName,
		Registry:   registryName,
		Name:       fmt.Sprintf("%s:%s", repoName, imageName.Identifier()),
		Digest:     imageRef.Context().Digest(imageRef.Identifier()).DigestStr(),
		Version:    imageName.Identifier(),
	}, nil
}

func (cluster *Cluster) getBasicMetadata() (Metadata, error) {
	rawCfg, err := cluster.cConfig.RawConfig()
	if err != nil {
		return Metadata{}, err
	}
	clusterName := rawCfg.Contexts[rawCfg.CurrentContext].Cluster
	version, err := cluster.clientSet.ServerVersion()
	if err != nil {
		return Metadata{}, err
	}
	tools := []Tool{
		{
			Vendor:  "aquasecurity",
			Name:    "trivy",
			Version: "0.38.1",
		},
	}
	return Metadata{
		Timestamp: CurrentTimeStamp(),
		Tools:     tools,
		Component: Component{
			Name:    clusterName,
			Version: version.GitVersion,
		},
	}, nil
}

func (bom *Cluster) CreateClusterBom(metadata Metadata, components []Component, Addons []Component, nodeInfo []NodeInfo) *ClusterBom {
	cluster := &ClusterBom{
		Metadata: metadata,
		ControlPlane: ControlPlane{
			Components: components,
		},
		NodesInfo: nodeInfo,
		Addons:    Addons,
	}
	return cluster
}

func (cluster *Cluster) CollectNodes() []NodeInfo {
	nodes, err := cluster.clientSet.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	nodesInfo := make([]NodeInfo, 0)
	for _, node := range nodes.Items {
		nodeRole := "worker"
		if _, ok := node.Labels["node-role.kubernetes.io/control-plane"]; ok {
			nodeRole = "master"
		}
		if _, ok := node.Labels["node-role.kubernetes.io/master"]; ok {
			nodeRole = "master"
		}
		nodesInfo = append(nodesInfo, NodeInfo{
			NodeName:                node.Name,
			KubeletVersion:          node.Status.NodeInfo.KubeletVersion,
			ContainerRuntimeVersion: node.Status.NodeInfo.ContainerRuntimeVersion,
			OsImage:                 node.Status.NodeInfo.OSImage,
			Hostname:                node.ObjectMeta.Name,
			KernelVersion:           node.Status.NodeInfo.KernelVersion,
			KubeProxyVersion:        node.Status.NodeInfo.KernelVersion,
			OperatingSystem:         node.Status.NodeInfo.OperatingSystem,
			Architecture:            node.Status.NodeInfo.Architecture,
			NodeRole:                nodeRole,
		})
	}
	return nodesInfo
}

func getPodsInfo(clientset *kubernetes.Clientset, labelSelector string, namespace string) *corev1.PodList {
	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{LabelSelector: labelSelector})
	if err != nil {
		panic(err.Error())
	}
	return pods
}

func (cluster *Cluster) collectComponents(labels map[string]string) ([]Component, error) {
	components := make([]Component, 0)
	for namespace, labelSelector := range labels {
		pods := getPodsInfo(cluster.clientSet, labelSelector, namespace)
		for _, pod := range pods.Items {
			for _, s := range pod.Status.ContainerStatuses {
				imageRef, err := containerimage.ParseReference(s.ImageID)
				if err != nil {
					return nil, err
				}
				imageName, err := containerimage.ParseReference(s.Image)
				if err != nil {
					return nil, err
				}
				c, err := cluster.GetBaseComponent(imageRef, imageName)
				if err != nil {
					continue
				}
				components = append(components, c)
			}

		}
	}
	return components, nil
}

func (cluster *Cluster) isOpenShift() bool {
	ctx := context.Background()
	_, err := cluster.clientSet.CoreV1().Namespaces().Get(ctx, "openshift-kube-apiserver", metav1.GetOptions{})
	return !k8sapierror.IsNotFound(err)
}

func (cluster *Cluster) GetPackages(metadata TargetMetadata, components []Component, Addons []Component, nodeInfo []NodeInfo) *BomResult {
	packages := make([]Package, 0)
	components = append(components, Addons...)
	for _, c := range components {
		packages = append(packages, Package{
			ID:      fmt.Sprintf("%s@%s", c.Name, c.Version),
			Name:    c.Name,
			Version: c.Version,
			Digest:  c.Digest,
		})
	}
	nodePkgs := cluster.NodeInfoToPkg(nodeInfo, metadata.Version)
	packages = append(packages, nodePkgs...)
	br := &BomResult{
		Packages: packages,
		Target:   fmt.Sprintf("%s@%s", metadata.Name, metadata.Version),
		Type:     "Cluster",
		Class:    "Kubertnetes",
	}
	return br
}

func (cluster *Cluster) NodeInfoToPkg(nodesInfo []NodeInfo, version string) []Package {
	packages := make([]Package, 0)
	for _, n := range nodesInfo {
		kubeletPkg := Package{
			ID:      fmt.Sprintf("%s@%s", "kubelet", n.KubeletVersion),
			Name:    "kubelet",
			Version: n.KubeletVersion,
		}
		packages = append(packages, kubeletPkg)
		kubeProxyPkg := Package{
			ID:      fmt.Sprintf("%s@%s", "kube-proxy", n.KubeProxyVersion),
			Name:    "kube-proxy",
			Version: n.KubeProxyVersion,
		}
		packages = append(packages, kubeProxyPkg)
		runtimeParts := strings.Split(n.ContainerRuntimeVersion, "://")
		runtimeName := strings.TrimSpace(runtimeParts[0])
		runtimeVersion := strings.TrimSpace(runtimeParts[1])
		containerdPkg := Package{
			ID:      fmt.Sprintf("pkg:%s@v%s", runtimeName, runtimeVersion),
			Name:    runtimeName,
			Version: runtimeVersion,
		}
		packages = append(packages, containerdPkg)
		osParts := strings.Split(n.OsImage, " ")
		osName := strings.TrimSpace(osParts[0])
		osVersion := strings.TrimSpace(osParts[1])
		osPkg := Package{
			ID:        fmt.Sprintf("%s@%s", osName, osVersion),
			Name:      osName,
			Version:   osVersion,
			DependsOn: []string{kubeletPkg.Name, kubeProxyPkg.Name, containerdPkg.Name},
			Properties: []KeyValue{
				{
					Name:  "architecture",
					Value: n.Architecture,
				},
				{
					Name:  "kernel_version",
					Value: n.KernelVersion,
				},
				{
					Name:  "operating_system",
					Value: n.OperatingSystem,
				},
			},
		}
		packages = append(packages, osPkg)
		packages = append(packages, Package{
			ID:        fmt.Sprintf("%s@%s", n.NodeName, version),
			Name:      n.NodeName,
			Version:   version,
			DependsOn: []string{osPkg.Name},
			Properties: []KeyValue{
				{
					Name:  "node-role",
					Value: n.NodeRole,
				},
				{
					Name:  "host_name",
					Value: n.Hostname,
				},
			},
		})
	}
	return packages
}

func (cluster *Cluster) targetMetadata() (TargetMetadata, error) {
	rawCfg, err := cluster.cConfig.RawConfig()
	if err != nil {
		return TargetMetadata{}, err
	}
	clusterName := rawCfg.Contexts[rawCfg.CurrentContext].Cluster
	version, err := cluster.clientSet.ServerVersion()
	if err != nil {
		return TargetMetadata{}, err
	}
	return TargetMetadata{
		Name:    clusterName,
		Version: version.GitVersion,
	}, nil
}
