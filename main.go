package main

import (
	"context"
	"encoding/json"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/version"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"

	"k8s.io/client-go/rest"
)

const (
	k8sComponentNamespace = "kube-system"
)

func main() {

	cf := genericclioptions.NewConfigFlags(true)
	rest.SetDefaultWarningHandler(rest.NoWarnings{})
	clientConfig := cf.ToRawKubeConfigLoader()
	rc, err := clientConfig.ClientConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(rc)
	if err != nil {
		panic(err.Error())
	}
	components := make([]Component, 0)
	etcdPods, err := clientset.CoreV1().Pods(k8sComponentNamespace).List(context.Background(), metav1.ListOptions{LabelSelector: "component=etcd"})
	if err != nil {
		panic(err.Error())
	}
	for _, pod := range etcdPods.Items {
		components = append(components, Component{
			Name:      pod.Spec.Containers[0].Name,
			Container: pod.Spec.Containers[0].Image,
		})
	}

	cmPods, err := clientset.CoreV1().Pods(k8sComponentNamespace).List(context.Background(), metav1.ListOptions{LabelSelector: "component=kube-controller-manager"})
	if err != nil {
		panic(err.Error())
	}

	for _, pod := range cmPods.Items {
		components = append(components, Component{
			Name:      pod.Spec.Containers[0].Name,
			Container: pod.Spec.Containers[0].Image,
		})
	}
	apiPods, err := clientset.CoreV1().Pods(k8sComponentNamespace).List(context.Background(), metav1.ListOptions{LabelSelector: "component=kube-apiserver"})
	if err != nil {
		panic(err.Error())
	}

	for _, pod := range apiPods.Items {
		components = append(components, Component{
			Name:      pod.Spec.Containers[0].Name,
			Container: pod.Spec.Containers[0].Image,
		})
	}

	schedulerPod, err := clientset.CoreV1().Pods(k8sComponentNamespace).List(context.Background(), metav1.ListOptions{LabelSelector: "component=kube-scheduler"})
	if err != nil {
		panic(err.Error())
	}

	for _, pod := range schedulerPod.Items {
		components = append(components, Component{
			Name:      pod.Spec.Containers[0].Name,
			Container: pod.Spec.Containers[0].Image,
		})
	}

	nodes, err := clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	nodesInfo := make([]NodeInfo, 0)
	for _, node := range nodes.Items {
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
		})
	}
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
	serverVersion, err := clientset.ServerVersion()
	if err != nil {
		panic(err.Error())
	}
	rawCfg, err := clientConfig.RawConfig()
	if err != nil {
		panic(err.Error())
	}
	clusterName := rawCfg.Contexts[rawCfg.CurrentContext].Cluster
	k8sCluster := &K8sCluster{
		ClusterName:  clusterName,
		Version:      serverVersion,
		ControlPlane: ControlPlane{Components: components},
		NodesInfo:    nodesInfo,
		Addons:       addons,
	}
	b, err := json.Marshal(k8sCluster)
	fmt.Print(string(b))
}

type K8sCluster struct {
	ClusterName  string        `json:"cluster_name"`
	Version      *version.Info `json:"version"`
	ControlPlane ControlPlane  `json:"control_plane"`
	NodesInfo    []NodeInfo    `json:"node_info"`
	Addons       []Addon       `json:"addons"`
}

type NodeInfo struct {
	NodeName                string `json:"node_name"`
	KubeletVersion          string `json:"kubelet_version"`
	ContainerRuntimeVersion string `json:"container_runtime_version"`
	OsImage                 string `json:"os_image"`
	Hostname                string `json:"host_name"`
	KernelVersion           string `json:"kernel_version"`
	KubeProxyVersion        string `json:"kube_proxy_version"`
	OperatingSystem         string `json:"operating_system"`
	Architecture            string `json:"architecture"`
}

type Addon struct {
	Name      string `json:"name"`
	Container string `json:"container"`
}

type ControlPlane struct {
	Components []Component `json:"components"`
}

type Component struct {
	Name      string `json:"name"`
	Container string `json:"container"`
}
