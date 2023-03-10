package k8s

import (
	"k8s.io/apimachinery/pkg/version"
)

type Cluster struct {
	ClusterName  string        `json:"cluster_name"`
	Version      *version.Info `json:"version"`
	ControlPlane ControlPlane  `json:"control_plane"`
	NodesInfo    []NodeInfo    `json:"node_info"`
	Addons       []Addon       `json:"addons"`
}

type NodeInfo struct {
	NodeRole                string `json:"node_role"`
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
