package k8s

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-containerregistry/pkg/name"
	containerimage "github.com/google/go-containerregistry/pkg/name"
	corev1 "k8s.io/api/core/v1"
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

func (bom *Cluster) CreateClusterSbom(clusterType string) (*ClusterBom, error) {
	nodesInfo := bom.CollectNodes()
	// collect addons info
	var components []Component
	var err error
	var labels map[string]string
	if clusterType == "ocp" {
		labels = map[string]string{
			"openshift-kube-apiserver":          "apiserver",
			"openshift-kube-controller-manager": "kube-controller-manager",
			"openshift-kube-scheduler":          "scheduler",
			"openshift-etcd":                    "etcd",
		}
	} else {
		// collect core components
		labels = map[string]string{
			k8sComponentNamespace: "component",
		}
	}
	components, err = bom.collectComponents(labels)
	if err != nil {
		return nil, err
	}
	metadata, err := bom.getBasicMetadata()
	if err != nil {
		return nil, err
	}
	addonLabels := map[string]string{
		k8sComponentNamespace: "k8s-app",
	}
	addons, err := bom.collectComponents(addonLabels)
	if err != nil {
		return nil, err
	}
	return bom.CreateClusterBom(metadata, components, addons, nodesInfo), nil
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

func (bom *Cluster) getBasicMetadata() (Metadata, error) {
	rawCfg, err := bom.cConfig.RawConfig()
	if err != nil {
		return Metadata{}, err
	}
	clusterName := rawCfg.Contexts[rawCfg.CurrentContext].Cluster
	version, err := bom.clientSet.ServerVersion()
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

func (bom *Cluster) CollectNodes() []NodeInfo {
	nodes, err := bom.clientSet.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
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

func (bom *Cluster) collectComponents(labels map[string]string) ([]Component, error) {
	components := make([]Component, 0)
	for namespace, labelSelector := range labels {
		pods := getPodsInfo(bom.clientSet, labelSelector, namespace)
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
				c, err := bom.GetBaseComponent(imageRef, imageName)
				if err != nil {
					continue
				}
				components = append(components, c)
			}

		}
	}
	return components, nil
}
