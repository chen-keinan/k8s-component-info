package k8s

import (
	"context"

	"github.com/google/go-containerregistry/pkg/name"
	containerimage "github.com/google/go-containerregistry/pkg/name"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
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

func CollectCoreComponents[T any](components []T, clientset *kubernetes.Clientset, getComponent func(name.Reference, name.Reference) (T, error)) ([]T, error) {
	labelSelector := "component"
	pods := GetPodsInfo(clientset, labelSelector, k8sComponentNamespace)

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
			c, err := getComponent(imageRef, imageName)
			if err != nil {
				continue
			}
			components = append(components, c)
		}

	}
	return components, nil
}

func GetPodsInfo(clientset *kubernetes.Clientset, labelSelector string, namespace string) *corev1.PodList {
	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{LabelSelector: labelSelector})
	if err != nil {
		panic(err.Error())
	}
	return pods
}

func CollectAddons[T any](addons []T, clientset *kubernetes.Clientset, getComponent func(name.Reference, name.Reference) (T, error)) ([]T, error) {
	labelSelector := "k8s-app"
	pods := GetPodsInfo(clientset, labelSelector, k8sComponentNamespace)
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
			c, err := getComponent(imageRef, imageName)
			if err != nil {
				continue
			}
			addons = append(addons, c)
		}
	}
	return addons, nil
}

func CollectOpenShiftComponents[T any](components []T, clientset *kubernetes.Clientset, getComponent func(name.Reference, name.Reference) (T, error)) ([]T, error) {
	namespaceLabel := map[string]string{
		"openshift-kube-apiserver":          "apiserver",
		"openshift-kube-controller-manager": "kube-controller-manager",
		"openshift-kube-scheduler":          "scheduler",
		"openshift-etcd":                    "etcd",
	}
	for namespace, labelSelector := range namespaceLabel {
		pods := GetPodsInfo(clientset, labelSelector, namespace)

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
				c, err := getComponent(imageRef, imageName)
				if err != nil {
					continue
				}
				components = append(components, c)
			}

		}
	}
	return components, nil
}
