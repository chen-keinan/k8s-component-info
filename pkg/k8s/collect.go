package k8s

import (
	"context"

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

func CollectCoreComponents(clientset *kubernetes.Clientset) []Component {
	coreLabelSelector := "component"
	corePods, err := clientset.CoreV1().Pods(k8sComponentNamespace).List(context.TODO(), metav1.ListOptions{LabelSelector: coreLabelSelector})
	if err != nil {
		panic(err.Error())
	}

	components := make([]Component, 0)
	for _, pod := range corePods.Items {
		components = append(components, Component{
			Name:      pod.Spec.Containers[0].Name,
			Container: pod.Spec.Containers[0].Image,
		})
	}
	return components
}

func CollectAddons(clientset *kubernetes.Clientset) []Addon {
	labelSelector := "k8s-app"
	pods, err := clientset.CoreV1().Pods(k8sComponentNamespace).List(context.TODO(), metav1.ListOptions{LabelSelector: labelSelector})
	if err != nil {
		panic(err.Error())
	}
	addons := make([]Addon, 0)
	for _, pod := range pods.Items {
		addons = append(addons, Addon{
			Name:      pod.Name,
			Container: pod.Spec.Containers[0].Image,
		})
	}
	return addons
}
