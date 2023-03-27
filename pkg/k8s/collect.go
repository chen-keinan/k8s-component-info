package k8s

import (
	"context"
	"strings"

	"fmt"

	containerimage "github.com/google/go-containerregistry/pkg/name"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const (
	k8sComponentNamespace = "kube-system"
)

func CollectNodes(clientset *kubernetes.Clientset) []NodeInfo {
	nodes, err := clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
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

func CollectCoreComponents(clientset *kubernetes.Clientset) []*Component {
	labelSelector := "component"
	pods := GetPodsInfo(clientset, labelSelector, k8sComponentNamespace)
	components := make([]*Component, 0)

	for _, pod := range pods.Items {
		for _, s := range pod.Status.ContainerStatuses {
			c, err := GetComponent(s)
			if err != nil {
				continue
			}
			components = append(components, c)
		}

	}
	return components
}

func CollectAddons(clientset *kubernetes.Clientset) []*Component {
	labelSelector := "k8s-app"
	pods := GetPodsInfo(clientset, labelSelector, k8sComponentNamespace)
	addons := make([]*Component, 0)
	for _, pod := range pods.Items {
		for _, s := range pod.Status.ContainerStatuses {
			c, err := GetComponent(s)
			if err != nil {
				continue
			}
			addons = append(addons, c)
		}
	}
	return addons
}

func GetPodsInfo(clientset *kubernetes.Clientset, labelSelector string, namespace string) *corev1.PodList {
	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{LabelSelector: labelSelector})
	if err != nil {
		panic(err.Error())
	}
	return pods
}

const containerType = "container"

func GetComponent(containerStatus corev1.ContainerStatus) (*Component, error) {
	imageRef, err := containerimage.ParseReference(containerStatus.ImageID)
	if err != nil {
		return nil, err
	}
	imageName, err := containerimage.ParseReference(containerStatus.Image)
	if err != nil {
		return nil, err
	}
	repoName := imageRef.Context().RepositoryStr()
	registryName := imageRef.Context().RegistryStr()
	if strings.HasPrefix(repoName, "library/sha256") {
		repoName = imageName.Context().RepositoryStr()
		registryName = imageName.Context().RegistryStr()
	}

	bomRef := fmt.Sprintf("pkg:oci/%s@%s?repository_url=%s/library/%s",
		repoName,
		imageRef.Context().Digest(imageRef.Identifier()).DigestStr(),
		registryName,
		repoName)
	return &Component{
		Type:        containerType,
		NameVersion: containerStatus.Image,
		Name:        containerStatus.Name,
		BomRef:      bomRef,
		Purl:        bomRef,
	}, nil
}

func GetDependencies(ref string, components []*Component) []Dependency {
	dependencies := make([]Dependency, 0)
	dependsOn := make([]string, 0)
	for _, c := range components {
		dependsOn = append(dependsOn, fmt.Sprintf("pkg:%s/%s", c.Type, strings.Replace(c.NameVersion, ":", "@", -1)))
	}
	dependencies = append(dependencies, Dependency{
		Ref:       ref,
		DependsOn: dependsOn,
	})
	return dependencies
}

func CollectOpenShiftComponents(clientset *kubernetes.Clientset) []*Component {
	namespaceLabel := map[string]string{
		"openshift-kube-apiserver":          "apiserver",
		"openshift-kube-controller-manager": "kube-controller-manager",
		"openshift-kube-scheduler":          "scheduler",
		"openshift-etcd":                    "etcd",
	}
	components := make([]*Component, 0)
	for namespace, labelSelector := range namespaceLabel {
		pods := GetPodsInfo(clientset, labelSelector, namespace)

		for _, pod := range pods.Items {
			for _, s := range pod.Status.ContainerStatuses {
				c, err := GetComponent(s)
				if err != nil {
					continue
				}
				components = append(components, c)
			}

		}
	}
	return components
}
