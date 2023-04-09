package k8s

type Cluster struct {
	Metadata     Metadata     `json:"metadata,omitempty"`
	Version      string       `json:"version,omitempty"`
	ControlPlane ControlPlane `json:"control_plane,omitempty"`
	NodesInfo    []NodeInfo   `json:"nodes,omitempty"`
	Addons       []Component `json:"addons,omitempty"`
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
	Components []Component `json:"components,omitempty"`
}

type Component struct {
	Type       string `json:"type,omitempty"`
	Name       string `json:"name,omitempty"`
	Version    string `json:"version,omitempty"`
	Repository string `json:"repository,omitempty"`
	Registry   string `json:"registry,omitempty"`
	Digest     string `json:"digest,omitempty"`
}

type Metadata struct {
	Timestamp string    `json:"timestamp,omitempty"`
	Component Component `json:"component,omitempty"`
	Tools     []Tool    `json:"tools,omitempty"`
}

type Tool struct {
	Vendor  string `json:"vendor,omitempty"`
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}
