package k8s

import (
	"k8s.io/apimachinery/pkg/version"
)

type Cluster struct {
	ClusterName  string        `json:"cluster_name,omitempty"`
	Version      *version.Info `json:"version,omitempty"`
	ControlPlane ControlPlane  `json:"control_plane,omitempty"`
	NodesInfo    []NodeInfo    `json:"nodes,omitempty"`
	Addons       []*Component  `json:"addons,omitempty"`
}

type NodeInfo struct {
	NodeRole                string `json:"node_role,omitempty"`
	NodeName                string `json:"node_name,omitempty"`
	KubeletVersion          string `json:"kubelet_version,omitempty"`
	ContainerRuntimeVersion string `json:"container_runtime_version,omitempty"`
	OsImage                 string `json:"os_image,omitempty"`
	Hostname                string `json:"host_name,omitempty"`
	KernelVersion           string `json:"kernel_version,omitempty"`
	KubeProxyVersion        string `json:"kube_proxy_version,omitempty"`
	OperatingSystem         string `json:"operating_system,omitempty"`
	Architecture            string `json:"architecture,omitempty"`
}

type ControlPlane struct {
	Components []*Component `json:"components,omitempty"`
}

type Component struct {
	BomRef   string    `json:"bom-ref,omitempty"`
	Type     string    `json:"type,omitempty"`
	Name     string    `json:"name,omitempty"`
	Purl     string    `json:"purl,omitempty"`
	Version  string    `json:"version,omitempty"`
	Licenses []License `json:"licenses,omitempty"`
}

type License struct {
	Expression string `json:"expression,omitempty"`
}
